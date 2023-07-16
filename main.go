package main

import (
	"flag"
	"fmt"
	"notifier/src/dotenv"
	"os"
)

func main() {
	fmt.Println("Welcome to Notifier!")
	fmt.Println("Example: `notifier -name=default -service=discord 'hello !'`")

	service := flag.String("service", "discord", "the service to send the message to: 'discord'|'slack'|'all'")
	name := flag.String("name", "default", "the name of the server")

	flag.Parse()

	help()

	if len(flag.Args()) < 1 {
		fmt.Println("Please provide a message.")
		os.Exit(1)
	}

	content := flag.Args()[0]
	fmt.Println("Message:", content)
	send(*name, *service, content)
}

func help() {
	fmt.Println()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] MESSAGE\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Send a message to services: Discord, Slack")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}
}

func send(name string, service string, content string) {
	server := dotenv.Handle()

	var sended bool = false
	for i := 0; i < len(server.Items); i++ {
		el := server.Items[i]

		if string(el.Type) == service {
			if el.Name == name {
				el.Send(content)
				sended = true
			}
		} else if service == "all" {
			if el.Name == name {
				el.Send(content)
				sended = true
			}
		}
	}

	if !sended {
		fmt.Println("Not found.")
		fmt.Println("Name:", name)
		fmt.Println("Service:", service)
	}
}
