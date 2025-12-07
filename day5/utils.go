package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FreshIngredientRange struct {
	start int
	end int
}

type InventoryInfo struct {
	FreshIngredientRanges []FreshIngredientRange
	Ingredients []int
}

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

func IngestInventory(inventoryFile *os.File) (*InventoryInfo, error) {
	const (
		rangesMode = "ranges"
		ingredientsMode = "ingredients"
	)
	inventoryInfo := &InventoryInfo{
		FreshIngredientRanges: []FreshIngredientRange{},
		Ingredients: []int{},
	}

	scanner := bufio.NewScanner(inventoryFile)
	currMode := rangesMode
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currMode = ingredientsMode
			continue
		}

		switch currMode {
		case rangesMode:
			newRangeSlice := strings.Split(line, "-")
			if len(newRangeSlice) != 2 {
				return nil, fmt.Errorf("Range was malformed: %s", line)
			}
			newStart, err := strconv.Atoi(newRangeSlice[0])
			if err != nil {
				return nil, fmt.Errorf("Error converting range to int: %s", line)
			}
			newEnd, err := strconv.Atoi(newRangeSlice[1])
			if err != nil {
				return nil, fmt.Errorf("Error converting range to int: %s", line)
			}
			newRange := &FreshIngredientRange{
				start: newStart,
				end: newEnd,
			}
			inventoryInfo.FreshIngredientRanges = append(inventoryInfo.FreshIngredientRanges, *newRange)
		case ingredientsMode:
			ingredientId, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("Error converting ingredient ID to int: %s", line)
			}
			inventoryInfo.Ingredients = append(inventoryInfo.Ingredients, ingredientId)
		}
	}
	return inventoryInfo, nil
}