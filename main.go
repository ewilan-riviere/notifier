// Package notifier is a simple notifier CLI for Discord and Slack with Webhook.
//
// Examples/readme can be found on the GitHub page at https://github.com/ewilan-riviere/notifier
//
// If you want to use it as CLI, you can install it with:
//
//	go install github.com/ewilan-riviere/notifier
//
// Then you can use it like this:
//
//	notifier discord "Hello World" https://discord.com/api/webhooks/XXXXXXXXX/XXXXXXXXX
//	notifier slack "Hello World" https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXX
//
// If you want to use it as library, you can install it with:
//
//	go get github.com/ewilan-riviere/notifier
//
// Then you can use it like this:
//
//	package main
//
//	import "github.com/ewilan-riviere/notifier/notify"
//
//	func main() {
//		notify.Notifier("Hello World", "https://discord.com/api/webhooks/XXXXXXXXX/XXXXXXXXX")
//		notify.Notifier("Hello World", "https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXX")
//	}
package main

import (
	"github.com/ewilan-riviere/notifier/pkg/dotenv"
	"github.com/ewilan-riviere/notifier/pkg/webhook"
	"github.com/spf13/cobra"
)

func main() {
	var cmdDiscord = &cobra.Command{
		Use:   "discord [message to send] [optional: webhook URL]",
		Short: "Send a message to a Discord from .env or webhook URL",
		Long: `Send a message to a Discord from .env or webhook URL.
Webhook URL is optional if you have a .env file with DISCORD_WEBHOOK.
For Discord, webhook have this formate: https://discord.com/api/webhooks/XXXXXXXXX/XXXXXXXXX`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			msg := args[0]

			url := ""
			if len(args) > 1 {
				url = args[1]
			} else {
				dotenv := dotenv.Make(".env")
				url = dotenv.DiscordWebhook
			}

			webhook := webhook.Make()
			webhook = webhook.SetDiscord(url)
			webhook.SendDiscord(msg)
		},
	}

	var cmdSlack = &cobra.Command{
		Use:   "slack [message to send] [optional: webhook URL]",
		Short: "Send a message to a Slack from .env or webhook URL",
		Long: `Send a message to a Slack from .env or webhook URL.
Webhook URL is optional if you have a .env file with SLACK_WEBHOOK.
For Slack, webhook have this formate: https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXX`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			msg := args[0]

			url := ""
			if len(args) > 1 {
				url = args[1]
			} else {
				dotenv := dotenv.Make(".env")
				url = dotenv.SlackWebhook
			}

			webhook := webhook.Make()
			webhook = webhook.SetSlack(url)
			webhook.SendSlack(msg)
		},
	}

	var rootCmd = &cobra.Command{Use: "notifier"}
	rootCmd.AddCommand(cmdDiscord)
	rootCmd.AddCommand(cmdSlack)
	rootCmd.Execute()
}
