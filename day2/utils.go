package day2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func GatherInput() (string, error) {
	var filePath string
	fmt.Print("Enter path to product ID ranges file: ")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		return "", fmt.Errorf("Error reading file path: %w", err)
	}

	idRangesFile, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("Error opening file: %w", err)
	}
	defer idRangesFile.Close()

	reader := bufio.NewReader(idRangesFile)
	contentBytes, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading file content:", err)
	}
	allRanges := string(contentBytes)

	return allRanges, nil
}

func ProcessIDRanges(rangesContent string, validationFunction func(int) bool) (int, error) {
	invalidIDSum := 0
	const delimiter = ","
	for idRange := range strings.SplitSeq(rangesContent, delimiter) {
		rangeSum, err := proccessRange(idRange, validationFunction)
		if err != nil {
			return 0, fmt.Errorf("Error processing range %s: %w", idRange, err)
		}
		invalidIDSum += rangeSum
	}
	return invalidIDSum, nil
}

func proccessRange(rangeStr string, validationFunction func(int) bool) (int, error) {
	invalidIDSum := 0
	bounds := strings.Split(rangeStr, "-")
	if len(bounds) != 2 {
		return 0, fmt.Errorf("Invalid range format: %s", rangeStr)
	}
	var lowerBound, upperBound int
	_, err := fmt.Sscanf(bounds[0], "%d", &lowerBound)
	if err != nil {
		return 0, fmt.Errorf("Invalid lower bound in range %s: %w", rangeStr, err)
	}
	_, err = fmt.Sscanf(bounds[1], "%d", &upperBound)
	if err != nil {
		return 0, fmt.Errorf("Invalid upper bound in range %s: %w", rangeStr, err)
	}
	for productID := lowerBound; productID <= upperBound; productID++ {
		if !validationFunction(productID) {
			invalidIDSum += productID
		}
	}
	return invalidIDSum, nil
}
