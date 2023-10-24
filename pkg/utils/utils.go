package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) []string {
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
