package main

import (
	"app/memory"

	"github.com/fatih/color"
)

func main() {
	// playAgain := true

	// for playAgain {
	for i := 0; i < 1; i++ {
		memory.Play()
		// playAgain = memory.GetYesOrNo("Would you like to play again (y/n)?")
		// fmt.Println()
	}

	color.Green("Goodbye")
}
