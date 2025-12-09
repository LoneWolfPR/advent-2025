package day7

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part1 () bool {
	diagramFile, err := GatherInput()
	if err != nil {
		fmt.Printf("error gathering input: %v", err)
	}

	diagram, err := ingestDiagramFile(diagramFile)
	if err != nil {
		fmt.Printf("error ingesting diagram file: %v", err)
	}

	splitCount, err := countSplits(diagram)
	if err != nil {
		fmt.Printf("error counting splits: %v", err)
	}
	fmt.Printf("The total number of tachyon beam splits is: %d", splitCount)
	return utils.RunAnotherPuzzlePrompt()
}

func countSplits(diagram [][]rune) (int, error) {
	splitCount := 0
	for rowIndex, row := range diagram {
		// No need to process the last row
		if rowIndex == len(diagram) - 2 {
			continue
		}
		for colIndex, gridSpace := range row {
			switch gridSpace {
			case startRune:
				// Found the start. Add beam directly below
				diagram[rowIndex + 1][colIndex] = beamRune
			case emptyRune:
				continue
			case beamRune:
				belowSpace := diagram[rowIndex + 1][colIndex]
				if belowSpace == emptyRune {
					// No split. Beam carries on
					diagram[rowIndex + 1][colIndex] = beamRune
				} else if belowSpace == splitterRune {
					// Split found. Increase count and divide beam
					splitCount++
					diagram[rowIndex + 1][colIndex - 1] = beamRune
					diagram[rowIndex + 1][colIndex + 1] = beamRune
				}
			case splitterRune:
				continue
			default:
				return 0, fmt.Errorf("error processing diagram at row %d, col %d. Invalid rune: %c", rowIndex, colIndex, gridSpace)
			}
		}
	}
	return splitCount, nil
}