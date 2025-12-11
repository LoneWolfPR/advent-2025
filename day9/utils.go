package day9

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func GatherInput() (*os.File, error) {
	var filePath string
	fmt.Print("Enter path to list of red square coordinates: ")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file path: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	return file, nil
}

func ingestSquaresFile(file *os.File, sort bool) ([][2]int, error) {
	squares := [][2]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coordStrings := strings.Split(line, ",")
		if len(coordStrings) != 2 {
			return nil, fmt.Errorf("invalid number of coordinates for location: %v", coordStrings)
		}
		rowVal, err := strconv.Atoi(coordStrings[0])
		if err != nil {
			return nil, fmt.Errorf("error converting x coordinate to integer:  %v", coordStrings[0])
		}
		colVal, err := strconv.Atoi(coordStrings[1])
		if err != nil {
			return nil, fmt.Errorf("error converting y coordinate to integer:  %v", coordStrings[1])
		}
		square := [2]int{ rowVal, colVal }
		squares = append(squares, square)
	}
	// Sort the squares by row and column
	if sort {
		slices.SortFunc(squares, func(a [2]int, b [2]int) int {
			if a[0] < b[0] {
				return -1
			} else if a[0] > b[0] {
				return 1
			} else {
				// Same row, check column
				if a[1] < b[1] {
					return -1
				} else if a[1] > b[1] {
					return 1
				} else {
					return 0
				}
			}
		})
	}
	return squares, nil
}

func computeArea(square1 [2]int, square2 [2]int) int {
	length := intAbs((square1[0]-square2[0])) + 1 
	width := intAbs((square1[1] - square2[1])) + 1
	return length * width
}

func intAbs(intVal int) int {
	if intVal < 0 {
		return -intVal
	}
	return intVal
}