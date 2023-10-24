package dotenv

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Dotenv struct {
	DiscordWebhook string
	SlackWebhook   string
}

func Make() Dotenv {
	path := dotenvPath()
	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Dotenv{
		DiscordWebhook: os.Getenv("DISCORD_WEBHOOK"),
		SlackWebhook:   os.Getenv("SLACK_WEBHOOK"),
	}
}

func (d Dotenv) ToString() {
	fmt.Printf("\nDiscordWebhook: %s\nSlackWebhook: %s\n", d.DiscordWebhook, d.SlackWebhook)
}

func dotenvPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	return exPath
}
