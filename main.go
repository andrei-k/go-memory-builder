package main

import (
	"app/memory"

	"github.com/fatih/color"
)

func main() {
	playAgain := true

	for playAgain {
		memory.Play()
		playAgain = memory.GetYesOrNo("Would you like to play again (y/n)?")
	}

	color.Green("Goodbye")
}
