package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Type string

const (
	Discord Type = "discord"
	Slack   Type = "slack"
	All     Type = "all"
)

type ServerItem struct {
	Name    string
	Id      string
	Token   string
	Channel string
	Type    Type
}

type DiscordWebhookPayload struct {
	Content string `json:"content"`
}

type SlackWebhookPayload struct {
	Text string `json:"text"`
}

func (p ServerItem) ToString() {
	fmt.Printf("\nName: %s\nId: %s\nToken: %s\nChannel: %s\nType: %s\n", p.Name, p.Id, p.Token, p.Channel, p.Type)
}

func (p ServerItem) Send(msg string) {
	const DiscordBaseURL = "https://discord.com/api/webhooks/"
	const SlackBaseURL = "https://hooks.slack.com/services/"

	var url string
	var payload interface{}

	if p.Type == Discord {
		url = fmt.Sprintf("%s%s/%s", DiscordBaseURL, p.Id, p.Token)

		payload = DiscordWebhookPayload{
			Content: msg,
		}
	} else if p.Type == Slack {
		url = fmt.Sprintf("%s%s/%s/%s", SlackBaseURL, p.Id, p.Token, p.Channel)

		payload = SlackWebhookPayload{
			Text: msg,
		}
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON payload:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error sending webhook request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Sended notification. Status code:", resp.StatusCode)

	return
}

type Server struct {
	Items []ServerItem
}
