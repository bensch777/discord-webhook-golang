
# discord-webhook-golang
 Allows for easy webhook sending through discord's webhook API.

# Installation

```go
go get github.com/bensch777/discord-webhook-golang
```

# Code Example
```go
    embed := discordwebhook.Embed{
        Title:     "Example Webhook",
        Color:     15277667,
        Url:       "https://www....",
        Timestamp: time.Now(),
        Thumbnail: discordwebhook.Thumbnail{
            Url: "https://www....",
        },
        Author: discordwebhook.Author{
            Name:     "Autho Name",
            Icon_URL: "https://www....",
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
            Icon_url: "https://www....",
        },
    }
```