package webhook

import (
	"strings"
	"testing"

	"github.com/ewilan-riviere/notifier/pkg/utils"
)

func TestWebhook(t *testing.T) {
	items := utils.ReadFile("../../.env.testing")

	webhook := Webhook{
		DiscordWebhook: strings.Split(items[0], "=")[1],
		SlackWebhook:   strings.Split(items[1], "=")[1],
	}

	discord := webhook.SendDiscord("Hello from Discord")
	slack := webhook.SendSlack("Hello from Slack")

	if !discord {
		t.Errorf("Discord webhook failed")
	}

	if !slack {
		t.Errorf("Slack webhook failed")
	}
}
