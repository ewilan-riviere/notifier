package main

import (
	"fmt"
	"strings"

	"github.com/ewilan-riviere/notifier/pkg/dotenv"
	"github.com/ewilan-riviere/notifier/pkg/webhook"
	"github.com/spf13/cobra"
)

func main() {
	var cmdSend = &cobra.Command{
		Use:   "send [service to send] [message to send]",
		Short: "Send a message to a service (Discord or Slack) from .env",
		Long:  `Send a message to a service (Discord or Slack) from .env. You can use 'discord' or 'slack' as argument.`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			service := args[0]
			message := args[1]

			if service != "discord" && service != "slack" {
				fmt.Println("Error: service not found, please use 'discord' or 'slack'")
				return
			}

			dotenv := dotenv.Make()
			webhook := webhook.Make()
			webhook.SetDiscord(dotenv.DiscordWebhook)
			webhook.SetSlack(dotenv.SlackWebhook)

			if service == "discord" {
				webhook.SendDiscord(message)
			}
			if service == "slack" {
				webhook.SendSlack(message)
			}
		},
	}

	var cmdSendWebhook = &cobra.Command{
		Use:   "webhook [webhook] [message to send]",
		Short: "Send a message to a service (Discord or Slack) from webhook URL",
		Long:  `Send a message to a service (Discord or Slack) from webhook URL.`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			url := args[0]
			message := args[1]

			webhook := webhook.Make()
			if strings.Contains(url, "discord") {
				webhook := webhook.SetDiscord(url)
				webhook.SendDiscord(message)
			}
			if strings.Contains(url, "slack") {
				webhook := webhook.SetSlack(url)
				webhook.SendSlack(message)
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdSend)
	rootCmd.AddCommand(cmdSendWebhook)
	rootCmd.Execute()
}
