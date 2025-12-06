package day5

import (
	"fmt"
	"os"
)

func GatherInput() (*os.File, error) {
	var filePath string
	fmt.Print("Enter path to kitchen inventory file: ")
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