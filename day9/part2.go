package day9

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)

type Line struct {
	pt1 [2]int
	pt2 [2]int
}
type TheaterInfo struct {
	squares [][2]int
	verticalEdges []Line
	horizontalEdges []Line
	farCorner [2]int
	cache map[[2]int]bool 
}
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
	theaterInfo := &TheaterInfo{
		squares: squares,
		farCorner: getFurthestCorner(squares),
		cache: map[[2]int]bool{},
	}
	findEdges(theaterInfo)

	area, err := findLargestRectAreaV2(theaterInfo)
	if err != nil {
		fmt.Printf("error calculating area %v", err)
	}
	fmt.Printf("The largest area possible is: %d", area)
	return utils.RunAnotherPuzzlePrompt()
}

func findEdges(theaterInfo *TheaterInfo) {
	squares := theaterInfo.squares
	for index, square := range squares {
		// Squares wrap. So if we're at the end of the list the next square in
		// the shape would be the first in the list
		var nextSquare [2]int
		if index < (len(squares) - 1) {
			nextSquare = squares[index + 1]
		} else {
			nextSquare = squares[0]
		}
		edge := Line{
			pt1: square,
			pt2: nextSquare,
		}
		if square[0] == nextSquare[0] {
			// horizontal
			theaterInfo.horizontalEdges = append(theaterInfo.horizontalEdges, edge)
		} else {
			// they always share a row or column. so if not horizontal it's vertical
			theaterInfo.verticalEdges = append(theaterInfo.verticalEdges, edge)
		}
	}
}

func findLargestRectAreaV2(theaterInfo *TheaterInfo) (int, error) {
	largestArea := 0
	squares := theaterInfo.squares
	for i, square1 := range squares {
		for j := i + 1; j < len(squares); j++ {
			square2 := squares[j]
			if isRectLegal(theaterInfo, square1, square2) {
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

func isRectLegal(theaterInfo *TheaterInfo, square1 [2]int, square2 [2]int) bool {
	// First, check if 4 corners are legal
	// first corner
	corners := [][2]int {
		square1,
		{square1[0], square2[1]},
		{square2[0], square1[1]},
		square2,
	}
	for _, corner := range corners {
		if !isSquareLegal(theaterInfo, corner) {
			return false
		}
	}

	// If all 4 corners are equal we must check each edge
	edges := []Line {
		{
			pt1: corners[0],
			pt2: corners[1],
		},
		{
			pt1: corners[1],
			pt2: corners[2],
		},
		{
			pt1: corners[2],
			pt2: corners[3],
		},
		{
			pt1: corners[3],
			pt2: corners[1],
		},
	}
	for _, edge := range edges {
		edgeOrientation := "horizontal"
		if edge.pt1[0] == edge.pt2[0] {
			edgeOrientation = "vertical"
		}
		if intersectsEdge(theaterInfo, edge, edgeOrientation, false) {
			return false
		}
	}
	return true
}

func isSquareLegal(theaterInfo *TheaterInfo, square [2]int) bool {
	if val, found := theaterInfo.cache[square]; found {
		return val
	}
	isLegal := false
	// Cast rays in each direction to determine isLegal
	directions := []string{"up", "down", "left", "right"}
	boundariesCrossed := 0
	for _, direction := range directions {
		if castRay(theaterInfo, square, direction) {
			boundariesCrossed++
		}
	}
	if boundariesCrossed == 4 {
		isLegal = true
	}
	theaterInfo.cache[square] = isLegal
	return isLegal
}

func castRay(theaterInfo *TheaterInfo, start [2]int, direction string) bool {
	furthestCorner := theaterInfo.farCorner
	// a ray is a line from the start to the edge of the grid in the direction
	// supplied.
	endPt := [2]int{}
	edgeOrientation := "vertical"

	switch direction {
	case "up":
		endPt[0] = furthestCorner[0]
		endPt[1] = start[1]
		edgeOrientation = "horizontal"
	case "down":
		endPt[0] = 0
		endPt[1] = start[1]
		edgeOrientation = "horizontal"
	case "left":
		endPt[0] = start[0]
		endPt[1] = 0
	case "right":
		endPt[0] = start[0]
		endPt[1] = furthestCorner[1]
	}
	ray := Line{
		pt1: start,
		pt2: endPt,
	}

	// Test if a ray goes through one of the edges. If so return true
	return intersectsEdge(theaterInfo, ray, edgeOrientation, true)
}

func intersectsEdge(theaterInfo *TheaterInfo, subject Line, edgeOrientation string, edgeInclusive bool) bool {

	betweenFunc := isBetweenExclusive
	if edgeInclusive {
		betweenFunc = isBetweenInclusive
	}
	if edgeOrientation == "vertical" {
		for _, edge := range theaterInfo.verticalEdges {
			if betweenFunc(edge.pt1[0], edge.pt2[0], subject.pt1[0]) &&
			betweenFunc(subject.pt1[1], subject.pt2[1], edge.pt1[1]) {
				return true
			}
		}
	} else {
		for _, edge := range theaterInfo.horizontalEdges {
			if betweenFunc(edge.pt1[1], edge.pt2[1], subject.pt1[1]) &&
			betweenFunc(subject.pt1[0], subject.pt2[0], edge.pt1[0]) {
				return true
			}
		}
	}
	return false
}

func isBetweenInclusive(bound1, bound2, subject int) bool {
	if bound1 >= bound2 {
		return subject >= bound2 && subject <= bound1
	}
	return subject >= bound1 && subject <= bound2
}

func isBetweenExclusive(bound1, bound2, subject int) bool {
	if bound1 > bound2 {
		return subject > bound2 && subject < bound1
	}
	return subject > bound1 && subject < bound2
}