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
	}

	diagram, err := ingestDiagramFile(diagramFile)
	if err != nil {
		fmt.Printf("error ingesting diagram file: %v", err)
	}
	startingIndex, err := findInitialStart(diagram[0])
	if err != nil {
		fmt.Printf("error finding start point: %v", err)
	}
	start := &Coordinates{
		row: 0,
		col: startingIndex,
	}
	timelineCount := countTimelines(diagram, *start)
	fmt.Printf("The total number of tachyon particle timelines is: %d", timelineCount)
	return utils.RunAnotherPuzzlePrompt()
}

func findInitialStart(line []rune) (int, error) {
	for index, space := range line {
		if space == startRune {
			return index, nil
		}
	}
	return 0, fmt.Errorf("starting rune not found")
}

func countTimelines(diagram [][]rune, start Coordinates) int {
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

/* func countTimelines(diagram [][]rune, start Coordinates) int {
	numTimelines := 1
	startingRow := start.row
	for i := startingRow; i < len(diagram) - 1; i++ {
		spaceBelow := diagram[i+1][start.col]
		breakLoop := false
		switch spaceBelow {
		case emptyRune:
			diagram[i+1][start.col] = beamRune
		case splitterRune:
			diagram[i+1][start.col-1] = beamRune
			leftStart := &Coordinates{
				row: i + 1,
				col: start.col - 1,
			}
			numTimelines = countTimelines(diagram, *leftStart)

			diagram[i+1][start.col+1] = beamRune
			rightStart := &Coordinates{
				row: i + 1,
				col: start.col + 1,
			}
			numTimelines += countTimelines(diagram, *rightStart)
			breakLoop = true
		}
		if breakLoop {
			break
		}
	}

	return numTimelines
} */

func printGrid (diagram [][]rune) {
	for _, row := range diagram {
		line := string(row)
		fmt.Println(line)
	}
}