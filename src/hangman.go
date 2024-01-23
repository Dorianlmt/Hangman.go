package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func jeu() {
	words, err := readWords("words.txt")
	if err != nil {
		fmt.Println("Il te faut un fichier words.txt pour jouer")
		return
	}
	word := chooseRandomWord(words)
	attempts := 10
	guessedLetters := make([]bool, len(word))
	usedLetters := make(map[rune]bool) // lettres déjà proposées
	var guessedLettersString string    // stockage lettres déjà proposées

	for attempts > 0 {
		displayWord(word, guessedLetters)
		fmt.Println("Lettres déjà proposées : " + guessedLettersString)
		fmt.Print("\npropose :")
		reader := bufio.NewReader(os.Stdin)
		guess, _ := reader.ReadString('\n')
		guess = strings.ToLower(strings.TrimSpace(guess))

		if len(guess) != 1 && guess != word {
			fmt.Println("Merci de mettre une seule lettre ou de mettre le mot en entier .")
			continue
		}
		if guess == word {
			fmt.Println("Bravo ! Tu as trouvé le mot, c'est pas trop tôt :", word)
			fmt.Println("1.Rejouer, 2.Quitter")
			break
		}
		if !isAlpha(guess) {
			fmt.Println("Tu ne peux utiliser que des lettres... c'est logique c'est un mot :/.")
			continue
		}

		if usedLetters[rune(guess[0])] {
			fmt.Println("Tu as déjà proposé cette lettre. Essaye une autre lettre.")
			continue
		}

		guessedLettersString += guess + " " // ajout aux lettres déjà utilisées

		found := false
		for i, char := range word {
			if guessedLetters[i] {
				continue
			}
			if guess == string(char) {
				guessedLetters[i] = true
				found = true
			}
		}

		usedLetters[rune(guess[0])] = true // ajouter la lettre à la liste des lettres utilisées

		if !found {
			fmt.Println("Lettre incorrecte. Vous avez plus que ", attempts-1, "essais... Tu veux le tuer en fait??")
			attempts--
			displayHangman(attempts)
		} else if allLettersGuessed(guessedLetters) {
			displayWord(word, guessedLetters)
			fmt.Println("Félicitations ! Vous avez trouvé le mot :", word)
			break
		}
	}

	if attempts == 0 {
		fmt.Println("Perdu gros looser le mot était: ", word)
		fmt.Println("1.Rejouer, 2.Quitter")
	}
}
