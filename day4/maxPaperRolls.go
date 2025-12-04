package day4

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)

func MaxPaperRolls() bool {
	layoutFile, err := GatherInput()
	if err != nil {
		fmt.Println("Error gathering input:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	defer layoutFile.Close()

	layoutInfo, err := IngestLayout(layoutFile)
	if err != nil {
		fmt.Println("Error ingesting layout:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	numRemoved, err := RemoveAllPossiblePaperRolls(layoutInfo)
	if err != nil {
		fmt.Println("Error removing paper rolls:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	fmt.Println("Maximum number of paper rolls removed: ", numRemoved)
	return utils.RunAnotherPuzzlePrompt()
}

func RemoveAllPossiblePaperRolls(layout *LayoutInfo) (int, error) {
	numRemoved := 0
	updatedLayoutInfo := &LayoutInfo{
		Rows: layout.Rows,
		Cols: layout.Cols,
		Grid: make([][]rune, layout.Rows),
	}
	copy(updatedLayoutInfo.Grid, layout.Grid)
	accessiblePositions := GetAccessiblePositions(updatedLayoutInfo)
	for len(accessiblePositions) > 0 {
		RemovePaperRolls(updatedLayoutInfo, accessiblePositions)
		numRemoved += len(accessiblePositions)
		accessiblePositions = GetAccessiblePositions(updatedLayoutInfo)
	}
	return numRemoved, nil
}

func RemovePaperRolls(layout *LayoutInfo, positions []GridPosition) {
	for _, pos := range positions {
		layout.Grid[pos.Row][pos.Col] = unoccupiedSymbol
	}
}