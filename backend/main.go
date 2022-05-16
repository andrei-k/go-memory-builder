package main

import (
	"app/memory"

	"github.com/fatih/color"
)

func main() {
	playAgain := true

	for playAgain {
		memory.Play()
		playAgain = memory.AskToPlayAgain()
	}

	color.Green("Goodbye")
}
