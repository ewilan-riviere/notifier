package dotenv

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

type Dotenv struct {
	DiscordWebhook string
	SlackWebhook   string
}

func Make(name string) Dotenv {
	basePath := dotenvPath() // `/Users/ewilan/Workspace/notifier/notifier`
	split := strings.Split(basePath, "/")
	path := strings.Join(split[:len(split)-1], "/") // `/Users/ewilan/Workspace/notifier`

	dotenvPath := path + "/" + name

	err := godotenv.Load(dotenvPath)
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

	realPath, err := filepath.EvalSymlinks(exPath + "/notifier")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return realPath
}
