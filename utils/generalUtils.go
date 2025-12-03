package utils

import "fmt"

func RunAnotherPuzzlePrompt() bool {
	var response string
	fmt.Println("\n---Puzzle Complete---")
	fmt.Print("Would you like to run another puzzle? (y/n): ")
	_, err := fmt.Scanln(&response)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return false
	}
	if response == "y" || response == "Y" {
		return true
	}
	return false
}