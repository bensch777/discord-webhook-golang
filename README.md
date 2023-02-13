
# discord-webhook-golang
 Allows for easy webhook sending through discord's webhook API.

# Installation

```go
go get github.com/bensch777/discord-webhook-golang
```

# Code Example
```go
package main

import (
    "encoding/json"
    "log"
    "time"
    "github.com/bensch777/discord-webhook-golang"
)

func main() {

    var webhookurl = "https://discord.com/api/webhooks/1069721907429122218/AXcbveVUfztv5Xh5y5uOp....."

    embed := discordwebhook.Embed{
        Title:     "Example Webhook",
        Color:     15277667,
        Url:       "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
        Timestamp: time.Now(),
        Thumbnail: discordwebhook.Thumbnail{
            Url: "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
        },
        Author: discordwebhook.Author{
            Name:     "Autho Name",
            Icon_URL: "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
        },
        Fields: []discordwebhook.Field{
            discordwebhook.Field{
                Name:   "Feld 1",
                Value:  "Feld Value 1",
                Inline: true,
            },
            discordwebhook.Field{
                Name:   "Feld 2",
                Value:  "Feld Value 2",
                Inline: true,
            },
            discordwebhook.Field{
                Name:   "Feld 3",
                Value:  "Feld Value 3",
                Inline: false,
            },
        },
        Footer: discordwebhook.Footer{
            Text:     "Footer Text",
            Icon_url: "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
        },
    }

    SendEmbed(webhookurl, embed)

}


func SendEmbed(link string, embeds discordwebhook.Embed) error {

    hook := discordwebhook.Hook{
        Username:   "Captain Hook",
        Avatar_url: "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
        Content:    "Message",
        Embeds:     []discordwebhook.Embed{embeds},
    }

    payload, err := json.Marshal(hook)
    if err != nil {
        log.Fatal(err)
    }
    err = discordwebhook.ExecuteWebhook(link, payload)
    return err

}
```