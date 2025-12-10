package day8

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type BoxLocation struct {
	x float64
	y float64
	z float64
}

type BoxDistance struct {
	distance float64
	box1Index int
	box2Index int
}

type Day8Input struct {
	file *os.File
	numConnections int
}

func CalcDistance (box1 BoxLocation, box2 BoxLocation) float64 {
	xDif := box1.x - box2.x
	yDif := box1.y - box2.y
	zDif := box1.z - box2.z

	dist := math.Sqrt((xDif * xDif) + (yDif * yDif) + (zDif * zDif))
	return dist
}

func GatherInput(gatherConnections bool) (*Day8Input, error) {
	var filePath string
	var strConnections string
	input := &Day8Input{}
	fmt.Print("Enter path to Junction Box position file: ")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file path: %w", err)
	}

	if gatherConnections {
		fmt.Print("How many connections would you like to find? ")
		_, err = fmt.Scanln(&strConnections)
		if err != nil {
			return nil, fmt.Errorf("Error reading connections input")
		}
		input.numConnections, err = strconv.Atoi(strConnections)
		if err != nil {
			return nil, fmt.Errorf("That was not a number: %s", strConnections)
		}
	}

	input.file, err = os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	return input, nil
}

func ingestBoxFile(file *os.File) ([]BoxLocation, error) {
	boxes := []BoxLocation{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		coordinateStrings := strings.Split(line, ",")
		if len(coordinateStrings) != 3 {
			return nil, fmt.Errorf("invalid number of coordinates for location: %v", coordinateStrings)
		}
		xVal, err := strconv.ParseFloat(coordinateStrings[0], 64)
		if err != nil {
			return nil, fmt.Errorf("x coordinate is not an int: %s", coordinateStrings[0])
		}
		yVal, err := strconv.ParseFloat(coordinateStrings[1], 64)
		if err != nil {
			return nil, fmt.Errorf("y coordinate is not an int: %s", coordinateStrings[1])
		}
		zVal, err := strconv.ParseFloat(coordinateStrings[2], 64)
		if err != nil {
			return nil, fmt.Errorf("z coordinate is not an int: %s", coordinateStrings[2])
		}
		box := BoxLocation{
			x: xVal,
			y: yVal,
			z: zVal,
		}
		boxes = append(boxes, box)
	}

	return boxes, nil
}

func findAndSortDistances(boxes []BoxLocation) []BoxDistance {
	distances := []BoxDistance{}
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			boxDistance := BoxDistance{
				distance: CalcDistance(boxes[i], boxes[j]),
				box1Index: i,
				box2Index: j,
			}
			distances = append(distances, boxDistance)
		}
	}

	// Sort distances
	slices.SortFunc(distances, func(a, b BoxDistance) int {
		if a.distance < b.distance {
			return -1
		} else if a.distance > b.distance {
			return 1
		} else {
			return 0
		}
	})

	return distances
}