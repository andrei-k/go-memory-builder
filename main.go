package main

import (
	"app/memory"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	playAgain := true

	for playAgain {
		memory.Play()
		playAgain = memory.GetYesOrNo("Would you like to play again (y/n)?")
		fmt.Println()
	}

	color.Green("Goodbye")
}
