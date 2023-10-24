package notify

import (
	"strings"

	"github.com/ewilan-riviere/notifier/pkg/webhook"
)

func Notifier(msg string, url string) {
	webhook := webhook.Make()

	if strings.Contains(url, "discord") {
		webhook := webhook.SetDiscord(url)
		webhook.SendDiscord(msg)
	}

	if strings.Contains(url, "slack") {
		webhook := webhook.SetSlack(url)
		webhook.SendSlack(msg)
	}
}
