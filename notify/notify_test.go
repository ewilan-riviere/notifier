package notify

import (
	"testing"

	"github.com/ewilan-riviere/notifier/pkg/utils"
)

func TestWebhook(t *testing.T) {
	path := "../.env.testing"
	discordWebhook := utils.ReadDotenv(path, "DISCORD_WEBHOOK")
	slackWebhook := utils.ReadDotenv(path, "SLACK_WEBHOOK")

	Notifier("test", discordWebhook)
	Notifier("test", slackWebhook)

	discordWebhookToken := utils.ReadDotenv(path, "DISCORD_WEBHOOK_TOKEN")
	slackWebhookToken := utils.ReadDotenv(path, "SLACK_WEBHOOK_TOKEN")

	Notifier("test", discordWebhookToken)
	Notifier("test", slackWebhookToken)
}
