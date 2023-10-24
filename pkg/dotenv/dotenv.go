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

func Make(name string) Dotenv {
	path := dotenvPath()
	dotenvPath := path + "/" + name
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading " + dotenvPath + " file")
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
