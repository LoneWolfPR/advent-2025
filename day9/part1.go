package day9

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part1 () bool {
	squaresFile, err := GatherInput()
	if err != nil {
		fmt.Printf("error gathering input: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	squares, err := ingestSquaresFile(squaresFile, true)
	if err != nil {
		fmt.Printf("error ingesting squares file: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	area, err := findLargestRectArea(squares)
	if err != nil {
		fmt.Printf("error calculating area %v", err)
	}
	fmt.Printf("The area of the largest possible red rectangle is: %d", area)
	return utils.RunAnotherPuzzlePrompt()
}

func findLargestRectArea(squares [][2]int) (int, error) {
	largestArea := 0
	for _, square1 := range squares {
		for _, square2 := range squares {
			if square1 == square2 {
				continue
			}
			area := computeArea(square1, square2)
			if area > largestArea {
				largestArea = area
			}
		}
	}

	return largestArea, nil
}
