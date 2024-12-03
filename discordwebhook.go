package discordwebhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Thumbnail struct {
	Url string `json:"url"`
}

type Footer struct {
	Text     string `json:"text"`
	Icon_url string `json:"icon_url"`
}

type Embed struct {
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	Color       int       `json:"color"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Footer      Footer    `json:"footer"`
	Fields      []Field   `json:"fields"`
	Timestamp   time.Time `json:"timestamp"`
	Author      Author    `json:"author"`
}

type Author struct {
	Name     string `json:"name"`
	Icon_URL string `json:"icon_url"`
	Url      string `json:"url"`
}

type Attachment struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
}

type Hook struct {
	Username    string       `json:"username"`
	Avatar_url  string       `json:"avatar_url"`
	Content     string       `json:"content"`
	Embeds      []Embed      `json:"embeds"`
	Attachments []Attachment `json:"attachments"`
}

// WebhookRequest represents a single webhook payload.
type WebhookRequest struct {
	Link string
	Data []byte
}

// WebhookQueue manages the queue of rate-limited requests.
type WebhookQueue struct {
	queue chan WebhookRequest
	wg    sync.WaitGroup
}

var rateLimitQueue *WebhookQueue

// NewWebhookQueue initializes the global queue.
func NewWebhookQueue(bufferSize int) *WebhookQueue {
	return &WebhookQueue{
		queue: make(chan WebhookRequest, bufferSize),
	}
}

// Start begins processing the rate-limited requests in the queue.
func (wq *WebhookQueue) Start() {
	go func() {
		for req := range wq.queue {
			err := executeWithDelay(req.Link, req.Data)
			if err != nil {
				fmt.Printf("Failed to process webhook from queue: %v\n", err)
			}
			wq.wg.Done()
		}
	}()
}

// Add adds a request to the queue.
func (wq *WebhookQueue) Add(request WebhookRequest) {
	wq.wg.Add(1)
	wq.queue <- request
}

// Stop waits for all tasks to complete.
func (wq *WebhookQueue) Stop() {
	wq.wg.Wait()
	close(wq.queue)
}

// executeWithDelay retries a webhook request after a delay.
func executeWithDelay(link string, data []byte) error {
	// Wait for 10 seconds before retrying the request.
	time.Sleep(2 * time.Second)
	return ExecuteWebhook(link, data)
}

// ExecuteWebhook sends a webhook request with rate-limit handling.
func ExecuteWebhook(link string, data []byte) error {
	// Create the HTTP request.
	req, err := http.NewRequest("POST", link, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body.
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 || resp.StatusCode == 204 {
		// Success: No further processing needed.
		return nil
	}

	if resp.StatusCode == 429 {
		// Rate limit reached: Add to the rate limit queue.
		fmt.Println("Rate limit reached. Adding to the queue for retry.")
		rateLimitQueue.Add(WebhookRequest{Link: link, Data: data})
		return nil
	}

	// Handle unexpected status codes.
	return errors.New(fmt.Sprintf("Unexpected status code %d: %s", resp.StatusCode, bodyText))
}

// SendEmbed sends an embed to a Discord webhook.
func SendEmbed(link string, embeds Embed) error {
	hook := Hook{
		Embeds: []Embed{embeds},
	}
	payload, err := json.Marshal(hook)
	if err != nil {
		return err
	}
	return ExecuteWebhook(link, payload)
}

func init() {
	// Initialize the global rate-limit queue with a buffer size of 10.
	rateLimitQueue = NewWebhookQueue(10)
	rateLimitQueue.Start()
}
