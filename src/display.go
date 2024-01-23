package main

import (
	"fmt"
	"time"
)

func displayWord(word string, guessedLetters []bool) {
	time.Sleep(2 * time.Second)
	clearTerminal()
	fmt.Print("Mot : ")

	for i, char := range word {
		if guessedLetters[i] {
			fmt.Print(string(char))
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func allLettersGuessed(guessedLetters []bool) bool {
	for _, guessed := range guessedLetters {
		if !guessed {
			return false
		}
	}
	return true
}

func displayHangman(attempts int) {
	hangmanFigures := []string{

		`
              +---+
              |   |
              o   |
             /|\  |
             / \  |
                  |
            =========
                    `,

		`
              +---+
              |   |
              o   |
             /|\  |
             /    |
                  |
            =========
                    `,

		`
              +---+
              |   |
              o   |
             /|\  |
                  |
                  |
            =========
                    `,

		`
              +---+
              |   |
              o   |
             /|   |
                  |
                  |
            =========
                    `,

		`
              +---+
              |   |
              o   |
              |   |
                  |
                  |
            =========
                    `,

		`
                +---+
                |   |
                o   |                          
                    |
                    |
             =========
                        `,

		`
             +---+
             |   |                            
                 |
                 |
                 |
          =========
                          `,

		`
             +---+
                 |
                 |
                 |
                 |
            =======
                              `,

		`	
                |
                |
                |
                |
                |
          =========
		  `,
		`
            ========
            `,
	}

	if attempts < 0 {
		attempts = 0
	}
	fmt.Println(hangmanFigures[attempts])
}
