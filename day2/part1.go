package day2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part1() bool {
	idRangesFile, err := GatherInput()
	if err != nil {
		fmt.Println("Error gathering input:", err)
	}
	defer idRangesFile.Close()

	reader := bufio.NewReader(idRangesFile)
	contentBytes, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading file content:", err)
	}
	allRanges := string(contentBytes)
	result, err := processIDRanges(allRanges)
	if err != nil {
		fmt.Println("Error processing ID ranges:", err)
	}
	fmt.Println("Sum of all invalid product IDs: ", result)
	return utils.RunAnotherPuzzlePrompt()
}

func processIDRanges(rangesContent string) (int, error) {
	invalidIDSum := 0
	const delimiter = ","
	for idRange := range strings.SplitSeq(rangesContent, delimiter) {
		rangeSum, err := proccessRange(idRange)
		if err != nil {
			return 0, fmt.Errorf("Error processing range %s: %w", idRange, err)
		}
		invalidIDSum += rangeSum
	}
	return invalidIDSum, nil
}

func proccessRange(rangeStr string) (int, error) {
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
		if !validateProductID(productID) {
			invalidIDSum += productID
		}
	}
	return invalidIDSum, nil
}

func validateProductID(productID int) bool {
	idStr := fmt.Sprintf("%d", productID)
	idLen := len(idStr)
	// if the product ID has an odd number of digits, it's valid
	if idLen%2 != 0 {
		return true
	}
	midIndex := idLen / 2
	firstHalf := idStr[:midIndex]
	secondHalf := idStr[midIndex:]
	// valid if the two halves are different
	return firstHalf != secondHalf
}

func GatherInput() (*os.File, error) {
	var filePath string
	fmt.Print("Enter path to product ID ranges file: ")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file path: %w", err)
	}

	idRangesFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}

	return idRangesFile, nil
}