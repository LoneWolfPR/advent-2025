package secretentrance

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Enter() {
	var filePath string
	var startingPoint int
	fmt.Print("Enter path to rotation sequence file: ")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		fmt.Println("Error reading file path:", err)
		return
	}

	fmt.Print("Enter starting point: ")
	_, err = fmt.Scanln(&startingPoint)
	if err != nil {
		fmt.Println("Error reading starting point:", err)
		return
	}

	rotSeqFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer rotSeqFile.Close()

	result, err := processRotationSequence(rotSeqFile, startingPoint)
	if err != nil {
		fmt.Println("Error processing rotation sequence:", err)
		return
	}
	fmt.Println("Passkey result: ", result)
}

func processRotationSequence(file *os.File, startingPoint int) (int, error) {
	var currPosition = startingPoint
	var zeroCount = 0
	var splitIndex = 1
	var maxStepValue = 99

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var newPosition int
		line := scanner.Text()
		direction := strings.ToUpper(line[:splitIndex])
		steps, err := strconv.Atoi(line[splitIndex:])
		if err != nil {
			return 0, fmt.Errorf("Invalid number of steps in line: %s", line)
		}
		moddedSteps := steps % (maxStepValue + 1)

		switch direction {
		case "L":
			newPosition = currPosition - moddedSteps
			if newPosition < 0 {
				newPosition = maxStepValue + newPosition + 1
			}
		case "R":
			newPosition = currPosition + moddedSteps
			if newPosition > maxStepValue {
				newPosition = newPosition - maxStepValue - 1
			}
		default:
			return 0, fmt.Errorf("Invalid direction in line: %s", line)	
		}
		currPosition = newPosition
		if currPosition == 0 {
			zeroCount++
		}
	}
	return zeroCount, nil
}