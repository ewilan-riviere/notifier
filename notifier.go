package main

import (
	"strings"

	"github.com/ewilan-riviere/notifier/pkg/webhook"
)

func notifier(msg string, url string) {
	webhook := webhook.Make()

	if strings.Contains(url, "discord") {
		webhook.SetDiscord(url)
		webhook.SendDiscord(msg)
	}

	if strings.Contains(url, "slack") {
		webhook.SetSlack(url)
		webhook.SendSlack(msg)
	}
}
