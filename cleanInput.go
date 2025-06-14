package main

import (
	"regexp"
	"strings"
)

func cleanInput(text string) []string {
	// clean up trailing spaces
	text = strings.TrimSpace(text)

	if text == "" {
		return []string{}
	}

	// separate the words
	words := strings.Fields(text)

	// identify each punctuation character
	var cleanedWords []string
	punctuationRegex := regexp.MustCompile(`[^\w\s]`)

	// clean up punctuation characters and make Title case each
	for _, word := range words {
		cleaned := punctuationRegex.ReplaceAllString(word, "")
		cleaned = strings.Title(strings.ToLower(cleaned))
		cleanedWords = append(cleanedWords, cleaned)
	}

	return cleanedWords
}
