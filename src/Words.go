package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"unicode"
)

func readWords(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	words := strings.Split(string(data), "\n")
	return words, nil
}

func chooseRandomWord(words []string) string {
	if len(words) == 1 {
		fmt.Println("Erreur : Il te faut des mots.")
	}
	randIndex := rand.Intn(len(words))

	word := words[randIndex]

	return word
}

func isAlpha(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
