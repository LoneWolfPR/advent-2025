package day7

import (
	"bufio"
	"fmt"
	"os"
)

const (
	startRune = 'S'
	emptyRune = '.'
	beamRune = '|'
	splitterRune = '^'
)

func GatherInput() (*os.File, error) {
	var filePath string
	fmt.Print("Enter path to Tachyon Manifold Diagram file: ")
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

func ingestDiagramFile(file *os.File) ([][]rune, error) {
	diagram := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		diagram = append(diagram, line)
	}
	return diagram, nil
}