package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func main() {
	// Placeholder words for now
	// TODO: The app should grab 20 randome words
	// through an API in JSON format
	var words []string
	words = append(words, "dog")
	words = append(words, "cat")
	words = append(words, "fish")
	words = append(words, "duck")
	words = append(words, "rabbit")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ENTER to see next word and ESC to quit")

	i := 0

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc {
			fmt.Println()
			break
		}

		if key == keyboard.KeyEnter {
			if i+1 == len(words) {
				fmt.Print("\033[H\033[2J")
				fmt.Println("The end")
				break
				// TODO: Instead of quitting, the app should ask the user
				// to re-display the words so they can see if they remembered
				// all the words
			} else {
				fmt.Printf("%d: %s\n", i+1, words[i])
			}
		}
		i++
	}
}

func prompt() {
	fmt.Print("-> ")
}
