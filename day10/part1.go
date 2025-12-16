package day10

import (
	"fmt"
	"os"
	"slices"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part1() bool {
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
	buttonPresses := calcFewestButtonPresses(machines)

	fmt.Printf("The fewest button presses required to start all machines is: %d", buttonPresses)
	return utils.RunAnotherPuzzlePrompt()
}

func calcFewestButtonPresses(machines []MachineInfo) int {
	buttonPresses := 0
	logFile, err := os.Create("./day10/log.txt")
	if err != nil {
		panic("can't open log file")
	}

	for index, machine := range machines {
		machinePresses := calcFewestButtonPressesForMachine(machine)
		fmt.Fprintf(logFile, "Presses for line %d: %d\n", index + 1, machinePresses)
		buttonPresses += machinePresses
	}
	defer logFile.Close()
	return buttonPresses
}

func calcFewestButtonPressesForMachine(machine MachineInfo) int {
	buttonPresses := 0
	//checkedStates := map[string]bool{} // convert list of clicked button indexes to a string for the key
	// set a slice to keep track of number of times a light is toggled
	// create initial slice of all button indexes
	allButtons := []int{}
	for index := range machine.buttons {
		allButtons = append(allButtons, index)
	}
	allButtonsStr := fmt.Sprintf("%v", allButtons)
	cache := []string{allButtonsStr}
	allCombinations := buildButtonCombinations(allButtons, &cache)
	// the combination of all buttons needs to be added as well
	allCombinations = append(allCombinations, allButtons)
	
	for _, combination := range allCombinations {
		currState := make([]int, len(machine.requiredPattern))
		for _, button := range combination {
			for _, light := range machine.buttons[button] {
				currState[light]++
			}
		}
		valid := checkState(machine.requiredPattern, currState)
		if valid && (buttonPresses == 0 || len(combination) < buttonPresses) {
			buttonPresses = len(combination)
		}
	}
	return buttonPresses
}

func checkState(requiredPattern []int, state []int) bool {
	for index, expected := range requiredPattern {
		lightState := state[index] % 2
		if expected != lightState{
			return false
		}
	}
	return true
}

func buildButtonCombinations(allButtons []int, cache *[]string) [][]int {
	combinations := [][]int{}
	for index := range allButtons {
		delIndex := len(allButtons) - index - 1
		combination := removeButton(allButtons, delIndex)
		combinationStr := fmt.Sprintf("%v",combination)
		if slices.Contains(*cache, combinationStr) {
			continue
		}
		combinations = append(combinations, combination)
		*cache = append(*cache, combinationStr)
		if len(combination) > 1 {
			childCombinations := buildButtonCombinations(combination, cache)
			combinations = append(combinations, childCombinations...)
		} else {
			continue
		}
	}
	return combinations
}

func removeButton(buttons []int, indexToRemove int) []int {
	subset := []int{}
	for i := 0; i < len(buttons); i++ {
		if i == indexToRemove {
			continue
		}
		subset = append(subset, buttons[i])
	}
	return subset
}