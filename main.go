package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if any command-line arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a subcommand.")
		os.Exit(1)
	}

	// Get the subcommand from the first command-line argument
	subcommand := os.Args[1]

	// Handle the subcommand
	switch subcommand {
	case "hello":
		fmt.Println("Hello, World!")
	default:
		fmt.Printf("Unknown subcommand: %s\n", subcommand)
		os.Exit(1)
	}
}
