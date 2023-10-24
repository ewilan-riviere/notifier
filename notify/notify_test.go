package notify

import (
	"testing"

	"github.com/ewilan-riviere/notifier/pkg/utils"
)

func TestWebhook(t *testing.T) {
	path := "../.env.testing"
	discordWebhook := utils.ReadDotenv(path, "DISCORD_WEBHOOK")
	slackWebhook := utils.ReadDotenv(path, "SLACK_WEBHOOK")

	dw := Notifier("test", discordWebhook)
	sw := Notifier("test", slackWebhook)

	if !dw {
		t.Errorf("Discord webhook failed")
	}

	if !sw {
		t.Errorf("Slack webhook failed")
	}

	discordWebhookToken := utils.ReadDotenv(path, "DISCORD_WEBHOOK_TOKEN")
	slackWebhookToken := utils.ReadDotenv(path, "SLACK_WEBHOOK_TOKEN")

	dwt := Notifier("test", discordWebhookToken)
	swt := Notifier("test", slackWebhookToken)

	if !dwt {
		t.Errorf("Discord webhook token failed")
	}

	if !swt {
		t.Errorf("Slack webhook token failed")
	}
}
