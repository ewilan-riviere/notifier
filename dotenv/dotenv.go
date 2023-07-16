package dotenv

import (
	"log"
	"os"
	"strings"

	server "notifier/server"

	"github.com/joho/godotenv"
)

func Handle(path string) server.Server {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DISCORD_SERVERS := os.Getenv("DISCORD_SERVERS")
	SLACK_SERVERS := os.Getenv("SLACK_SERVERS")

	items := []server.ServerItem{}

	discordItems := parse(DISCORD_SERVERS, server.Discord)
	slackItems := parse(SLACK_SERVERS, server.Slack)

	items = append(discordItems, slackItems...)

	listing := server.Server{
		Items: items,
	}

	return listing
}

func extract(array []string, index int) string {
	if len(array) > index {
		return array[index]
	}

	return ""
}

func parse(servers string, serverType server.Type) []server.ServerItem {
	var listing []server.ServerItem
	items := strings.Split(servers, ",")

	for i := 0; i < len(items); i++ {
		list := strings.Split(items[i], ":")
		item := server.ServerItem{
			Name:    extract(list, 0),
			Id:      extract(list, 1),
			Token:   extract(list, 2),
			Channel: extract(list, 3),
			Type:    serverType,
		}

		listing = append(listing, item)
	}

	return listing
}
