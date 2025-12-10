package day8

import (
	"fmt"
	"slices"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part1() bool {
	input, err := GatherInput(true)
	answer := 0
	if err != nil {
		fmt.Printf("error gathering input: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	boxes, err := ingestBoxFile(input.file)
	if err != nil {
		fmt.Printf("error processing box file: %v", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	distances := findAndSortDistances(boxes)
	circuits := createCircuits(distances, input.numConnections)
	answer = len(circuits[0]) * len(circuits[1]) * len(circuits[2])
	fmt.Println("The product of the size of the three largest circuits:",answer)
	return utils.RunAnotherPuzzlePrompt()
}



func createCircuits(distances []BoxDistance, connectionsToMake int) ([][]int) {
	circuits := [][]int{}
	connections := 0
	for i := 0; i < len(distances); i++ {
		currDistance := distances[i]
		foundCircuits := []int{}
		boxesToAdd := []int{}
		doNothing := false
		for index, circuit := range circuits {
			box1Found := slices.Contains(circuit, currDistance.box1Index)
			box2Found := slices.Contains(circuit, currDistance.box2Index)
			if box1Found && box2Found {
				doNothing = true
				continue
			} else if box1Found && !box2Found {
				foundCircuits = append(foundCircuits, index)
				boxesToAdd = append(boxesToAdd, currDistance.box2Index)
			} else if !box1Found && box2Found {
				foundCircuits = append(foundCircuits, index)
				boxesToAdd = append(boxesToAdd, currDistance.box1Index)
			}
		}
		if !doNothing {
			switch len(foundCircuits) {
			case 0:
				// If no match is found we create a new circuit
				newCircuit := []int{currDistance.box1Index, currDistance.box2Index}
				circuits = append(circuits, newCircuit)
			case 1:
				// If only one circuit is matched we just need to append the unmatched box to the circuit
				circuits[foundCircuits[0]] = append(circuits[foundCircuits[0]], boxesToAdd[0])
			case 2:
				// If both boxes are found in separate circuits we must merge the circuits instead of adding the boxes
				// to each circuit
				circuits[foundCircuits[0]] = append(circuits[foundCircuits[0]], circuits[foundCircuits[1]]...)
				circuits = slices.Delete(circuits, foundCircuits[1], foundCircuits[1] + 1)
			}
		}
		connections++
		if connections == connectionsToMake {
			break
		}
	}
	// Sort circuits by size
	slices.SortFunc(circuits, func(a, b []int) int {
		if len(a) < len(b) {
			return 1
		} else if len(a) > len(b) {
			return -1
		} else {
			return 0
		}
	})
	return circuits
}