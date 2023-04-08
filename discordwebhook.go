package discordwebhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

func ExecuteWebhook(link string, data []byte) error {

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
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("%s\n", bodyText))
	}
	if resp.StatusCode == 429 {
		fmt.Println("Rate limit reached")
		time.Sleep(time.Second * 5)
		return ExecuteWebhook(link, data)
	}
	return err
}

func SendEmbed(link string, embeds Embed) error {
	hook := Hook{
		Embeds: []Embed{embeds},
	}
	payload, err := json.Marshal(hook)
	if err != nil {
		return err
	}
	err = ExecuteWebhook(link, payload)
	return err

}
