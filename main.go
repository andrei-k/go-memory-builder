package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
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

	color.Green("Press ENTER to see next word and ESC to quit")

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
				clearScreen()
				color.Green("The end")
				break
				// TODO: Instead of quitting, the app should ask the user
				// to re-display the words so they can see if they remembered
				// all the words
			} else {
				color.Yellow("%d: %s\n", i+1, words[i])
			}
		}
		i++
	}
}

func prompt() {
	fmt.Print("-> ")
}

// Clears the Terminal screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// Linux and Mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
