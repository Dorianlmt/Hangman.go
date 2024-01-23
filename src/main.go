package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	// Initialisation et exécution du jeu Hangman
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Bienvenue dans le HANGMAN")
	fmt.Println("1. Jouer au jeu")
	fmt.Println("2. Quitter")

	for {
		fmt.Print("Sélectionnez une option (1 jouer ou 2 quitter) : ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			clearTerminal()
			text := `                                                       ___
		            	|     |  |""""| |\  | |      |\  /| |""""| |\  | 
				|-----|  |----| | \ | |   _  | \/ | |----| | \ |
			        |     |  |    | |  \| |____| |    | |    | |  \|
					 `
			fmt.Println(text)
			time.Sleep(1 * time.Second)

			jeu()

		case "2":
			clearTerminal()
			text2 := `
			 ___    ____   ____   ____    ____          ___
			|      |    | |    | |    \  |    | \   /  |
			|   _  |    | |    | |     ) |    )  \ /   |---
			|____| |____| |____| |____/  |____)   |    |___
			`
			fmt.Println(text2)
			time.Sleep(5 * time.Second)
			return
		default:
			clearTerminal()
			fmt.Println("Option invalide. Veuillez choisir 1 pour jouer ou 2 pour quitter.")
		}
	}
}

func clearTerminal() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") // Utilise la commande "cls" pour effacer l'écran dans Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J") // Utilise une séquence ANSI pour effacer l'écran (Linux/macOS)
	}
}
