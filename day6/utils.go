package day6

import (
	"fmt"
	"os"
)

type HomeworkProblem struct {
	numbers []int
	operator string
}

func GatherInput() (*os.File, error) {
	var filePath string
	fmt.Print("Enter path to homework file: ")
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

func solveProblems(problems []HomeworkProblem) (int, error) {
	solution := 0
	for _, problem := range problems {
		problemAnswer := 0
		for index, num := range problem.numbers {
			switch problem.operator {
			case "+":
				problemAnswer += num
			case "*":
				if (index == 0) {
					problemAnswer = num
				} else {
					problemAnswer *= num
				}
			}
		}
		solution += problemAnswer
	}
	return solution, nil
}