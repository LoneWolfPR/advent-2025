package day7

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)

type Coordinates struct {
	row int
	col int
}

func Part2 () bool {
	diagramFile, err := GatherInput()
	if err != nil {
		fmt.Printf("error gathering input: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	diagram, err := ingestDiagramFile(diagramFile)
	if err != nil {
		fmt.Printf("error ingesting diagram file: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	timelineCount := countTimelines(diagram)
	fmt.Printf("The total number of tachyon particle timelines is: %d", timelineCount)
	return utils.RunAnotherPuzzlePrompt()
}

func countTimelines(diagram [][]rune) int {
	lastCellValue := 0
	cache := make(map[Coordinates]int)
	for i := len(diagram) - 1; i >= 0; i-- {
		row := diagram[i]
		for colIndex, cell := range row {
			switch cell {
			case emptyRune:
				continue
			case splitterRune:
				cellCoords := &Coordinates{
					row: i,
					col: colIndex,
				}
				lastCellValue = getSplitterValue(diagram, *cellCoords, &cache)
			case startRune:
				return lastCellValue
			}
		}
	}
	return lastCellValue
}

func getSplitterValue(diagram [][]rune, start Coordinates, cache *map[Coordinates]int) int {
	leftStart := &Coordinates{
		row: start.row + 1,
		col: start.col - 1,
	}
	rightStart := &Coordinates{
		row: start.row + 1,
		col: start.col + 1,
	}
	timelineCount := getPathValueFromSplitter(diagram, *leftStart, cache) + getPathValueFromSplitter(diagram, *rightStart, cache)
	(*cache)[start] = timelineCount
	return timelineCount
}

func getPathValueFromSplitter(diagram [][]rune, currCoords Coordinates, cache *map[Coordinates]int) int {
	resultLeftFound := false
	for  !resultLeftFound {
		if currCoords.row == len(diagram) - 1 {
			break
		}
		nextCell := diagram[currCoords.row][currCoords.col]
		switch nextCell {
		case emptyRune:
			currCoords.row++
		case splitterRune:
			if val, found := (*cache)[currCoords]; found {
				return val
			}
			cellValue := getSplitterValue(diagram, currCoords, cache)
			(*cache)[currCoords] = cellValue
			return cellValue
		}
	}
	return 1
}