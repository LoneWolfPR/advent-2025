package day5

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)

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

	freshCount, err := CountFreshIngredients(inventoryInfo)
	if err != nil {
		fmt.Println("Error counting fresh ingredients:", err)
		return utils.RunAnotherPuzzlePrompt()
	}

	fmt.Println("The number of still fresh ingredients is:", freshCount)
	return utils.RunAnotherPuzzlePrompt()
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