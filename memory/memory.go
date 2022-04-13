package memory

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

var words []string

func Play() {
	generateWords()
	displayWords()
}

// Generate randome words
func generateWords() {
	words = nil
	// This is a placeholder words for now
	// TODO: The app should grab 20 random words
	// through an API in JSON format
	words = append(words, "dog")
	words = append(words, "cat")
	words = append(words, "fish")
	words = append(words, "duck")
	words = append(words, "rabbit")
}

func displayWords() {
	fmt.Println()
	color.Green("Get ready to train your memory!")
	color.Green("Press ENTER to see the next word and ESC to quit")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for i, x := range words {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc {
			break
		}

		if key == keyboard.KeyEnter {
			color.Yellow("%d: %s\n", i+1, x)
		}
		i++
	}

	revealWords()
}

// The app should ask the user to re-display the words so they
// can see if they remembered all the words
func revealWords() {
	clearScreen()
	color.Green("How many words can you remember?")
	color.Green("Press ENTER to reveal the words to see how you did.")

	_, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}

	if key == keyboard.KeyEnter {
		for i, x := range words {
			color.Blue("%d: %s\n", i+1, x)
		}
	}
}

func GetYesOrNo(q string) bool {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println("")
		color.Green(q)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'y' || char == 'Y' {
			return true
		}
		return false
	}
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
