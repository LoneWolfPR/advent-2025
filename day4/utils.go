package day4

import (
	"bufio"
	"fmt"
	"os"
)

const occupiedSymbol   = '@'
const unoccupiedSymbol = '.'

type LayoutInfo struct {
	Rows		   int
	Cols           int
	Grid		   [][]rune
}

type GridPosition struct {
	Row int
	Col int
}

func GatherInput() (*os.File, error) {
	var filePath string
	fmt.Print("Enter path to department layout file: ")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file path: %w", err)
	}

	layoutFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	return layoutFile, nil
}

func IngestLayout(layoutFile *os.File) (*LayoutInfo, error) {
	layoutInfo := &LayoutInfo{}
	grid := [][]rune{}
	rows := 0
	cols := 0

	scanner := bufio.NewScanner(layoutFile)
	for scanner.Scan() {
		line := scanner.Text()
		if cols == 0 {
			cols = len(line)
		} else if len(line) != cols {
			return nil, fmt.Errorf("Inconsistent row lengths in layout file")
		}
		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		grid = append(grid, row)
		rows++
	}
	layoutInfo.Rows = rows
	layoutInfo.Cols = cols
	layoutInfo.Grid = grid
	return layoutInfo, nil
}

func CountAccessibleRolls(layout *LayoutInfo) int {
	accessibleCount := 0
	for r := 0; r < layout.Rows; r++ {
		for c := 0; c < layout.Cols; c++ {
			space := layout.Grid[r][c]
			if space != occupiedSymbol {
				continue
			}
			if IsAccessible(layout, r, c) {
				accessibleCount++
			}
		}
	}

	return accessibleCount
}

func GetAccessiblePositions(layout *LayoutInfo) []GridPosition {
	accessiblePositions := []GridPosition{}
	for r := 0; r < layout.Rows; r++ {
		for c := 0; c < layout.Cols; c++ {
			space := layout.Grid[r][c]
			if space != occupiedSymbol {
				continue
			}
			if IsAccessible(layout, r, c) {
				accessiblePositions = append(accessiblePositions, GridPosition{Row: r, Col: c})
			}
		}
	}
	return accessiblePositions
}

func IsAccessible(layout *LayoutInfo, row, col int) bool {
	numOccupiedNeighbors := 0
	maxOccupiedNeighbors := 3
	testablePositions := GetTestablePositions(layout, row, col)
	for _, pos := range testablePositions {
		if layout.Grid[pos.Row][pos.Col] == occupiedSymbol {
			numOccupiedNeighbors++
		}
	}
	return numOccupiedNeighbors <= maxOccupiedNeighbors
}

func GetTestablePositions(layout *LayoutInfo, row, col int) []GridPosition {
	positions := []GridPosition{
		{Row: row - 1, Col: col - 1}, // Up and Left
		{Row: row - 1, Col: col},     // Up
		{Row: row - 1, Col: col + 1}, // Up and Right
		{Row: row, Col: col - 1},     // Left
		{Row: row, Col: col + 1},     // Right
		{Row: row + 1, Col: col - 1}, // Down and Left
		{Row: row + 1, Col: col},     // Down
		{Row: row + 1, Col: col + 1}, // Down and Right
	}
	testablePositions := []GridPosition{}
	var (
		minRow = 0
		minCol = 0
		maxRow = layout.Rows - 1
		maxCol = layout.Cols - 1
	)
	for _, pos := range positions {
		if pos.Row >= minRow && pos.Row <= maxRow && pos.Col >= minCol && pos.Col <= maxCol {
			testablePositions = append(testablePositions, pos)
		}
	}
	return testablePositions
}