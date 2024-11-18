package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	taboos := tabooWords()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		sentence := scanner.Text()

		if strings.EqualFold("exit", sentence) {
			break
		} else {
			sentence = fixSentence(sentence, taboos)
			fmt.Println(sentence)
		}
	}

	fmt.Println("Bye!")
}

func tabooWords() []string {
	var filename string
	fmt.Scan(&filename)

	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var taboos []string

	for scanner.Scan() {
		taboos = append(taboos, scanner.Text())
	}

	return taboos
}

func fixSentence(sentence string, taboos []string) string {
	words := strings.Split(sentence, " ")

	var result []string

	for _, word := range words {
		if isTaboo(word, taboos) {
			result = append(result, bluredWord(word))
		} else {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}

func isTaboo(word string, taboos []string) bool {
	for _, taboo := range taboos {
		if strings.EqualFold(taboo, word) {
			return true
		}
	}
	return false
}

func bluredWord(word string) string {
	count := utf8.RuneCountInString(word)

	return strings.Repeat("*", count)
}
