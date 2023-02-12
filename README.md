
# discord-webhook-golang
 Allows for easy webhook sending through discord's webhook API.

# Installation

```go
go get github.com/bensch777/discord-webhook-golang
```

# Code Example
```go
    embed := discordhooks.Embed{
        Title:     "Example Webhook",
        Color:     15277667,
        Url:       "https://www....",
        Timestamp: time.Now(),
        Thumbnail: discordhooks.Thumbnail{
            Url: "https://www....",
        },
        Author: discordhooks.Author{
            Name:     "Autho Name",
            Icon_URL: "https://www....",
        },
        Fields: []discordhooks.Field{
            discordhooks.Field{
                Name:   "Feld 1",
                Value:  "Feld Value 1",
                Inline: true,
            },
            discordhooks.Field{
                Name:   "Feld 2",
                Value:  "Feld Value 2",
                Inline: true,
            },
            discordhooks.Field{
                Name:   "Feld 3",
                Value:  "Feld Value 3",
                Inline: false,
            },
        },
        Footer: discordhooks.Footer{
            Text:     "Footer Text",
            Icon_url: "https://www....",
        },
    }
```