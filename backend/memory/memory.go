package memory

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"github.com/sethvargo/go-diceware/diceware"
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
	// words = append(words, "dog")
	// words = append(words, "cat")
	// words = append(words, "fish")
	// words = append(words, "duck")
	// words = append(words, "rabbit")

	// TODO: The app should grab 20 random words (nouns)
	// through an API in JSON format
	// This is a temporay setup for now...

	// Maybe use this: https://random-word-form.herokuapp.com/
	tempWords, err := diceware.Generate(10)
	if err != nil {
		log.Fatal(err)
	}

	tempJSON := `{
		"words": "` + strings.Join(tempWords, ",") + `"
	}`
	var myJSON map[string]string
	err = json.Unmarshal([]byte(tempJSON), &myJSON)
	if err != nil {
		fmt.Println("JSON decode error: ", err)
		return
	}
	words = strings.Split(myJSON["words"], ",")
}

func displayWords() {
	fmt.Println()
	color.Green("Get ready to train your memory!")
	color.Green("Press ENTER to see the next word or ESC to quit")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for index, word := range words {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc {
			color.Green("Goodbye")
			os.Exit(0)
		}

		if key == keyboard.KeyEnter {
			color.Yellow("%d: %s\n", index+1, word)
		}
		index++
	}

	revealWords()
}

// Ask to re-display the words so the user can check their memory
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

	if key == keyboard.KeyEsc {
		color.Green("Goodbye")
		os.Exit(0)
	}
}

func AskToPlayAgain() bool {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("")
	color.Green("Press ENTER to play again or ESC to quit")

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyEnter {
			return true
		}

		if key == keyboard.KeyEsc {
			return false
		}
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
