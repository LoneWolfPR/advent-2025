package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part2 () bool {
	homeworkFile, err := GatherInput()
	if err != nil {
		fmt.Println("Error gathering input:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	problems, err := IngestInput2(homeworkFile)
	if err != nil {
		fmt.Println("Error processing file:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	solution, err := solveProblems(problems)
	if err != nil {
		fmt.Println("Error solving problems:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	fmt.Println("Answer:", solution)
	return utils.RunAnotherPuzzlePrompt()
}

func IngestInput2 (file *os.File) ([]HomeworkProblem, error) {
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lineStr := scanner.Text()
		trimmedLine := strings.Trim(lineStr," ")
		if len(trimmedLine) == 0 {
			break;
		}
		reversedLine := utils.ReverseString(lineStr)
		lines = append(lines, reversedLine)
	}

	problems, err := generateProblems2(lines)
	if err != nil {
		return nil, fmt.Errorf("Problem generating problems: %v", err)
	}
	return problems, nil
}

func generateProblems2 (lines []string) ([]HomeworkProblem, error) {
	problems := []HomeworkProblem{}
	numLines := len(lines)

	problem := HomeworkProblem{
		numbers: []int{},
		operator: "",
	}
	for index, _ := range lines[0] {
		numString := ""
		spaceCount := 0
		for i := 0; i < numLines; i++ {
			currChar := string(lines[i][index])
			// Empty space, count and move on. Might be in column between problems
			if currChar == " " {
				spaceCount++
				continue
			}
			// At this point it's either a number or operator
			_, err := strconv.Atoi(currChar)
			if err != nil {
				problem.operator = currChar
			} else {
				numString += currChar
			}
		}

		if spaceCount == numLines {
			problems = append(problems, problem)
			problem.numbers = []int{}
			problem.operator = ""
		} else {
			number, err := strconv.Atoi(numString)
			if err != nil {
				return nil, fmt.Errorf("problem converting string to number: %s", numString)
			}
			problem.numbers = append(problem.numbers, number)
		}
	}
	// there should be one remaining problem because we don't end on a gap of all spaces
	problems = append(problems, problem)
	return problems, nil
}


/* Implementation Notes
Step 1: Read In lines as rows
Step 2: All lines should be same length so loop through the first line just for index purposes.
For each index initialize a new string for holding the int value. Then pull the character from 
each row and if the character is not a space append it to the string. If the last character is
an operator save it to the operator property of the problem. After reading the last character of
the last line add convert the value string to an int and add it to the numbers array of the problem.
If the value string is all spaces we're done with that problem and on the next index start a new problem.
*/