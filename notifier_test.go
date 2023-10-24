package main

import (
	"strings"
	"testing"

	"github.com/ewilan-riviere/notifier/pkg/utils"
)

func TestWebhook(t *testing.T) {
	items := utils.ReadFile(".env.testing")
	discordWebhook := strings.Split(items[0], "=")[1]
	slackWebhook := strings.Split(items[1], "=")[1]

	Notifier("test", discordWebhook)
	Notifier("test", slackWebhook)
}
