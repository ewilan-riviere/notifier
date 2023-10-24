package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Webhook struct {
	DiscordWebhook string
	SlackWebhook   string
}

type DiscordPayload struct {
	Content string `json:"content"`
}

type SlackPayload struct {
	Text string `json:"text"`
}

func Make() Webhook {
	return Webhook{
		DiscordWebhook: "",
		SlackWebhook:   "",
	}
}

func (w Webhook) SetDiscord(webhook string) Webhook {
	if !strings.HasPrefix(webhook, "http") {
		webhook = "https://discord.com/api/webhooks/" + webhook
	}

	w.DiscordWebhook = webhook

	return w
}

func (w Webhook) SetSlack(webhook string) Webhook {
	if !strings.HasPrefix(webhook, "http") {
		webhook = "https://hooks.slack.com/services/" + webhook
	}

	w.SlackWebhook = webhook

	return w
}

func (w Webhook) SendDiscord(msg string) bool {
	return send(w.DiscordWebhook, DiscordPayload{Content: msg}, "discord")
}

func (w Webhook) SendSlack(msg string) bool {
	return send(w.SlackWebhook, SlackPayload{Text: msg}, "slack")
}

func send(url string, payload interface{}, service string) bool {
	if url == "" {
		fmt.Println("Error: please check your " + service + " webhook URL " + url + " and try again")
		return false
	}

	success := false
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON payload:", err)
		return success
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error sending webhook request:", err)
		return success
	}
	defer resp.Body.Close()

	fmt.Println("Notification sended to " + service + " (status: " + strconv.Itoa(resp.StatusCode) + ")")

	if resp.StatusCode == 204 || resp.StatusCode == 200 {
		success = true
	}

	if !success {
		fmt.Print("Error: please check your " + service + " webhook URL " + url + " and try again\n")
	}

	return success
}
