package main

import (
	"testing"

	"github.com/ewilan-riviere/notifier/notify"
	"github.com/ewilan-riviere/notifier/pkg/utils"
)

func TestWebhook(t *testing.T) {
	discordWebhook := utils.ReadDotenv(".env.testing", "DISCORD_WEBHOOK")
	slackWebhook := utils.ReadDotenv(".env.testing", "SLACK_WEBHOOK")

	notify.Notifier("test", discordWebhook)
	notify.Notifier("test", slackWebhook)
}
