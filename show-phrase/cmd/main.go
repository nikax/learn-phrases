package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cwd, _ := os.Getwd()
	phrasesFilePath := cwd + "/../../save-phrase/data/phrases.txt"

	bytes, err := os.ReadFile(phrasesFilePath)
	if err != nil {
		fmt.Printf("can not read phrases from file %s", phrasesFilePath)
	}
	data := string(bytes)
	rows := strings.Split(strings.Trim(data, "\n"), "\n")

	words := make([][]string, 0, len(rows))

	for _, row := range rows {
		words = append(words, strings.Split(row, "|"))
	}

	answer := ""
	currentIndex := 0

	for {
		fmt.Println(words[currentIndex][0])
		fmt.Scanln(&answer)

		switch answer {
		case "y":
		case "n":
			fmt.Println(words[currentIndex][1])
			fmt.Printf(
				"https://translate.google.com/?sl=en&tl=ru&text=%s&op=translate\n",
				strings.ReplaceAll(words[currentIndex][0], " ", "+"),
			)
		default:
			printUsage()
		}

		currentIndex++
		if currentIndex > len(words)-1 {
			currentIndex = 0
		}
	}
}

func printUsage() {
	fmt.Println("Type \"y\" if you know the translation of the phrase or \"n\" if you don't")
}
