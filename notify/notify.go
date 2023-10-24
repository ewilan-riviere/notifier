package notify

import (
	"fmt"
	"strings"

	"github.com/ewilan-riviere/notifier/pkg/webhook"
)

func Notifier(msg string, url string) bool {
	success := false
	webhook := webhook.Make()

	if strings.Contains(url, "discord") {
		webhook := webhook.SetDiscord(url)
		webhook.SendDiscord(msg)
		success = true
	}

	if strings.Contains(url, "slack") {
		webhook := webhook.SetSlack(url)
		webhook.SendSlack(msg)
		success = true
	}

	if !success {
		split := strings.Split(url, "/")

		if len(split) == 2 {
			webhook := webhook.SetDiscord(url)
			webhook.SendDiscord(msg)
			success = true
		}

		if len(split) == 3 {
			webhook := webhook.SetSlack(url)
			webhook.SendSlack(msg)
			success = true
		}
	}

	if !success {
		fmt.Println("Error: Invalid webhook URL")
	}

	return success
}
