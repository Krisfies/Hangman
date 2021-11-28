package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"fmt"
	"os"
)

type Hangman struct {
	word string
	hiddenWord []string
	attempt int
	position string
}

func main() {
	var h1 Hangman
	h1.attempt = 10
	h1.ReadFile()
	h1.CreateHidden()
	h1.Reveal()
	for i:= 0; i < 9999; i++ {
		h1.PlayerTurn()
		h1.HangmanPositions()
		h1.Win()
		h1.Loose()
	}
}

func (h *Hangman) ReadFile() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(37)
	data, err := ioutil.ReadFile("words.txt")
	content := string(data)
	words := strings.Split(content, "\n")
	if err != nil {
		fmt.Println(err)

	} else {
		for i := 0; i < n-1; i++ {
			h.word = words[i]
		}
	}
	fmt.Println(h.word)
}

func (h *Hangman) CreateHidden() {
	for i :=0; i < len(h.word); i++ {
		h.hiddenWord = append(h.hiddenWord, "_")
	}
}

func (h *Hangman) Reveal() {
	var randomLetter int
	n := len(h.word) / 2 - 1
	if n > 0 {
		for i := 1 ; i <= n ; i ++ {
			randomLetter = rand.Intn(len(h.word))
			for i, letter := range h.word {
				if i == randomLetter {
					h.hiddenWord[i] = string(letter)
					fmt.Println(h.hiddenWord)
				}
			}
		}
	}
	fmt.Println("Voici le mot à deviner :",h.hiddenWord)
}

func (h *Hangman) HangmanPositions() {
	data, err := ioutil.ReadFile("hangman.txt")
	content := string(data)
	positions := strings.Split(content, "\n")
	var n int = 10-h.attempt

	if err != nil {
		fmt.Println(err)

	} else {
		for i := 0 ; i < n ; i++ {
			h.position = positions[i]
		}
		fmt.Println(positions[n])
	}

}

func (h *Hangman) PlayerTurn() {	
	var s string
	found := false
	fmt.Println("Entrez un mot ou une lettre.")
	fmt.Scanln(&s)
	if len(s) > 0 {
		for _, letter := range s {
			if (letter < 'a' && letter > 'z') || (letter < 'A' && letter > 'Z') {
				fmt.Println("Vous devez entrez une lettre")
			}
		}
	} else {
		fmt.Println("Vous ne devez entrez qu'une seule lettre")
	}
	for i, letter := range h.word {
		if s == string(letter) {
			h.hiddenWord[i] = string(letter) 
			found = true
		}
	}
	if !found {
		fmt.Println("Votre lettre n'est pas présente dans le mot.")
		h.attempt -= 1
		fmt.Println("Il vous reste", h.attempt,"tentatives")
		h.HangmanPositions()
		fmt.Println(h.hiddenWord)
	} else {
		fmt.Println("Bravo, vous avez deviné une lettre !")
		fmt.Println(h.hiddenWord)
	}
}

func (h *Hangman) Win() {
	match := false
	Loop:
	for i, letter := range h.word {
		if h.hiddenWord[i] == string(letter) {
			match = true
		} else {
			match = false
			break Loop
		}
	}
	if match {
		fmt.Println("Bravo, tu as gagnés ! le mot était bel est bien :", h.word)
		os.Exit(0)
	}
}

func (h *Hangman) Loose() {
	if h.attempt == 0 {
		fmt.Println("Tu as perdu, tes chances sont épuisés et José est mort.")
		os.Exit(0)	
	}
}


// Le hangmanposition ne fonctionne pas.