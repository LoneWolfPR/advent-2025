package day10

import (
	"fmt"
	"os"
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

func ingestDiagramFile(file *os.File) []MachineInfo {
	machines := []MachineInfo{}

	return machines
}