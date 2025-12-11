package day9

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part2 () bool {
	squaresFile, err := GatherInput()
	if err != nil {
		fmt.Printf("error gathering input: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	squares, err := ingestSquaresFile(squaresFile, false)
	if err != nil {
		fmt.Printf("error ingesting squares file: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	area, err := findLargestRectAreaV2(squares)
	if err != nil {
		fmt.Printf("error calculating area %v", err)
	}
	fmt.Printf("The largest area possible is: %d", area)
	return utils.RunAnotherPuzzlePrompt()
}

func findLargestRectAreaV2(squares [][2]int) (int, error) {
	largestArea := 0
	squareCache := map[[2]int]bool{}
	for _, square1 := range squares {
		for _, square2 := range squares {
			if isRectLegal(squares, square1, square2, squareCache) {
				area := computeArea(square1, square2)
				if area > largestArea {
					largestArea = area
				}
			}
		}
	}
	return largestArea, nil
}

// Adds 1 to furthest corner to add padding around shape
func getFurthestCorner(squares [][2]int) [2]int {
	largestCol := 0
	largestRow := 0
	for _, square := range squares {
		if square[0] > largestRow {
			largestRow = square[0]
		}
		if square[1] > largestCol {
			largestCol = square[1]
		}
	}

	return [2]int{largestRow + 1, largestCol + 1}
}

func isRectLegal(squares [][2]int, square1 [2]int, square2 [2]int, cache map[[2]int]bool) bool {
	// Figure out direction of coordinates
	rowDirection := ""
	if square1[0] > square2[0] {
		rowDirection = "up"
	} else if square1[0] < square2[0] {
		rowDirection = "down"
	} else {
		rowDirection = "none"
	}

	colDirection := ""
	if square1[1] > square2[1] {
		rowDirection = "left"
	} else if square1[0] < square2[0] {
		rowDirection = "right"
	} else {
		rowDirection = "none"
	}


	endOfRows := false;
	currSquare := square1
	for !endOfRows {
		endOfCols := false;
		for !endOfCols {
			if !isSquareLegal(squares, currSquare, cache) {
				return false
			}
			if currSquare[1] == square2[1] {
				endOfCols = true
				continue
			}
			switch colDirection {
			case "left":
				currSquare[1]--
			case "right":
				currSquare[1]++
			case "none":
				endOfCols = true
			}
		}

		// Check if at end of rows. If not go to next row
		// depending on direction moving
		if currSquare[0] == square2[0] {
			endOfRows = true
			continue
		}
		switch rowDirection {
		case "up":
			currSquare[0]--
		case "down":
			currSquare[0]++
		case "none":
			endOfRows = true
		}
	}
	return true
}

func isSquareLegal(squares [][2]int, square [2]int, cache map[[2]int]bool) bool {
	if val, found := (cache)[square]; found {
		return val
	}
	isLegal := false
	// Cast rays in each direction to determine isLegal
	directions := []string{"up", "down", "left", "right"}
	boundariesCrossed := 0
	for _, direction := range directions {
		if castRay(squares, square, direction) {
			boundariesCrossed++
		}
	}
	if boundariesCrossed == 4 {
		isLegal = true
	}
	cache[square] = isLegal
	return isLegal
}

func castRay(squares [][2]int, start [2]int, direction string) bool {
	crossesBoundary := false
	// TODO: Implement logic to trace a line in the supplied direction from
	// the start until is crosses a boundary or encounters the edge of the grid.
	// For boundary use 0 to the left and up. Use index 0 from furthest corner for down.
	// Use index 1 from furthest corner for right
	// Loop by adding or subtracting from the relevant coordinate in start given the direction
	// For each step loop through squares and compare the current progress point of the ray to
	// the current square. If both coordinates in the progress point match the current square
	// the points coincide and we need to return true.
	// If one of the coordinates match then we have to test the progress point against the line
	// drawn between the current square and the next square. Do that by checking if the non-matching
	// coordinate lies between (inclusive) the corresponding coordinates of the two squares.
	// If so, we've crossed a boundary and return true
	// furthestCorner := getFurthestCorner(squares)

	return crossesBoundary
}