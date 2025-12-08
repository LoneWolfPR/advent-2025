package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part1 () bool {
	homeworkFile, err := GatherInput()
	if err != nil {
		fmt.Println("Error gathering input:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	problems, err := IngestInput(homeworkFile)
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

func IngestInput(file *os.File) ([]HomeworkProblem, error) {
	scanner := bufio.NewScanner(file)
	lines := [][]string{}
	for scanner.Scan() {
		lineStr := scanner.Text()
		line := strings.Split(lineStr, " ")
		cleanedLine := []string{}
		for _, entry := range line {
			if entry != "" {
				cleanedLine = append(cleanedLine, entry)
			}
		}
		lines = append(lines, cleanedLine)
	}

	problems := generateProblems(lines)
	return problems, nil
}

func generateProblems(lines [][]string) []HomeworkProblem {
	problems := []HomeworkProblem{}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			var currProblem *HomeworkProblem
			if i == 0 {
				problems = append(problems, HomeworkProblem{
					numbers: []int{},
					operator: "",
				})
			}
			currProblem = &problems[j]
			entry := lines[i][j]
			val, err := strconv.Atoi(entry)
			if err != nil {
				// value is the operator because it's not an int
				currProblem.operator = entry
			} else {
				currProblem.numbers = append(currProblem.numbers, val)
			}
		}
	}

	return problems
}