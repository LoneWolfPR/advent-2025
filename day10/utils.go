package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MachineInfo struct {
	requiredPattern []int
	buttons [][]int
	joltageRequirements []int
}

func GatherInput() (*os.File, error) {
	var filePath string
	fmt.Print("Enter path to Indicator Light Diagram file: ")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file path: %w", err)
	}

	diagramFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	return diagramFile, nil
}

func ingestDiagramFile(file *os.File) ([]MachineInfo, error) {
	machines := []MachineInfo{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		machine := MachineInfo{}
		line := strings.TrimSpace(scanner.Text())
		pieces := strings.Split(line, " ")
		for _, piece := range pieces {
			switch string(piece[0]) {
			case "[":
				pattern, err := getRequiredPattern(piece)
				if err != nil {
					return nil, fmt.Errorf("Error reading pattern: %v", err)
				}
				machine.requiredPattern = pattern
			case "(":
				button, err := buildButton(piece)
				if err != nil {
					return nil, fmt.Errorf("Error building button: %v", err)
				}
				machine.buttons = append(machine.buttons, button)
			case "{":
				joltages, err := gatherJoltages(piece)
				if err != nil {
					return nil, fmt.Errorf("Error gathering joltages: %v", err)
				}
				machine.joltageRequirements = joltages
			}
		}
		machines = append(machines, machine)
	}

	return machines, nil
}

func getRequiredPattern(patternStr string) ([]int, error) {
	requiredPattern := []int{}
	cleanPattern := strings.Trim(patternStr, "[]")
	for _, char := range cleanPattern {
		switch string(char){
		case ".":
			requiredPattern = append(requiredPattern, 0)
		case "#":
			requiredPattern = append(requiredPattern, 1)
		default:
			return nil, fmt.Errorf("Invalid character in pattern: %s", string(char))
		}
	}

	return requiredPattern, nil
}

func buildButton(btnConfigStr string) ([]int, error) {
	btnConfig := []int{}
	cleanConfig := strings.Trim(btnConfigStr, "()")
	lights := strings.Split(cleanConfig, ",")
	for _, light := range lights {
		lightIndex, err := strconv.Atoi(light)
		if err != nil {
			return nil, fmt.Errorf("light index is not valid: %s", light)
		}
		btnConfig = append(btnConfig, lightIndex)
	}

	return btnConfig, nil
}

func gatherJoltages(joltagesStr string) ([]int, error) {
	joltages := []int{}
	cleanJoltagesStr := strings.Trim(joltagesStr, "{}")
	joltageStrs := strings.Split(cleanJoltagesStr, ",")
	for _, joltageStr := range joltageStrs {
		joltage, err := strconv.Atoi(joltageStr)
		if err != nil {
			return nil, fmt.Errorf("Joltage is not valid: %d", joltage)
		}
		joltages = append(joltages, joltage)
	}

	return joltages, nil
}