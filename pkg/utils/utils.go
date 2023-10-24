package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	return items
}

func ReadDotenv(path string, k string) string {
	items := ReadFile(path)
	dotenvItems := map[string]string{}

	for _, item := range items {
		key := strings.Split(item, "=")[0]
		value := strings.Split(item, "=")[1]
		dotenvItems[key] = value
	}

	return dotenvItems[k]
}
