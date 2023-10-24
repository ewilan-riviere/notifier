package webhook

import (
	"testing"
)

func TestWebhook(t *testing.T) {
	webhook := Webhook{
		DiscordWebhook: "https://discord.com/api/webhooks/1100099112159957112/TCxcWwGibz_DMtK1M4OOMOHaV62mP9fbKYkotr9x1ZicJOv9JC2NtO9yN4jpNDs3Od3D",
		SlackWebhook:   "https://hooks.slack.com/services/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",
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
