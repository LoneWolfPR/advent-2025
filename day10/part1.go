package day10

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)
func Part1() bool {
	diagramFile, err := GatherInput()
	if err != nil {
		fmt.Printf("Error gathering input: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	machines := ingestDiagramFile(diagramFile)
	buttonPresses := calcFewestButtonPresses(machines)

	fmt.Printf("The fewest button presses required to start all machines is: %d", buttonPresses)
	return utils.RunAnotherPuzzlePrompt()
}

func calcFewestButtonPresses([]MachineInfo) int {
	buttonPresses := 0

	return buttonPresses
}