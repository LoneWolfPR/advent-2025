package day10

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part2() bool {
	diagramFile, err := GatherInput()
	if err != nil {
		fmt.Printf("Error gathering input: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	machines, err := ingestDiagramFile(diagramFile)
	if err != nil {
		fmt.Printf("Error ingesting diagram: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	buttonPresses := calcJoltageButtonPresses(machines)

	fmt.Printf("The fewest button presses required to configure joltages for all machines is: %d", buttonPresses)
	return utils.RunAnotherPuzzlePrompt()
}

func calcJoltageButtonPresses(machines []MachineInfo) int {
	buttonPresses := 0

	return buttonPresses
}