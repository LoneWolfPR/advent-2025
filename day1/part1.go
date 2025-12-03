package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	InputData, err := GatherInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer InputData.SequenceFile.Close()
	
	result, err := processRotationSequence(InputData.SequenceFile, InputData.StartingPoint)
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