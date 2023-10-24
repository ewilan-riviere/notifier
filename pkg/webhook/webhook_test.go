package webhook

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestWebhook(t *testing.T) {
	items := readDotenv("../../.env.testing")

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

func readDotenv(path string) []string {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	items := []string{}
	for fileScanner.Scan() {
		items = append(items, fileScanner.Text())
	}

	readFile.Close()

	for _, eachline := range items {
		fmt.Println(eachline)
	}

	return items
}
