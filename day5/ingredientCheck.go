package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lonewolfpr/advent-2025/utils"
)

type FreshIngredientRange struct {
	start int
	end int
}

type InventoryInfo struct {
	FreshIngredientRanges []FreshIngredientRange
	Ingredients []int
}


func IngredientCheck() bool {
	inventoryFile, err := GatherInput()
	if err != nil {
		fmt.Println("Error gathering input:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	defer inventoryFile.Close()

	inventoryInfo, err := IngestInventory(inventoryFile)
	if err != nil {
		fmt.Println("Error ingesting inventory:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	fmt.Println("Number of Ranges:", len(inventoryInfo.FreshIngredientRanges))
	fmt.Println("Number of Ingredients:", len(inventoryInfo.Ingredients))

	freshCount, err := CountFreshIngredients(inventoryInfo)
	if err != nil {
		fmt.Println("Error counting fresh ingredients:", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	fmt.Println("The number of still fresh ingredients is:", freshCount)
	return utils.RunAnotherPuzzlePrompt()
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

func CountFreshIngredients(inventoryInfo *InventoryInfo) (int, error) {
	freshCount := 0

	for _, ingredient := range inventoryInfo.Ingredients {
		for _, freshRange := range inventoryInfo.FreshIngredientRanges {
			if ingredient >= freshRange.start && ingredient <= freshRange.end {
				freshCount++
				break
			}
		}
	}
	return freshCount, nil
}