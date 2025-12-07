package day5

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/lonewolfpr/advent-2025/utils"
)

func AllFreshIngredients() bool {
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
	idRanges := inventoryInfo.FreshIngredientRanges
	slices.SortFunc(idRanges, func(a, b FreshIngredientRange) int {
		if a.start < b.start {
			return -1
		}
		if a.start > b.start {
			return 1
		}
		return 0
	})
	combinedRanges := CombineRangesRecursive(
		idRanges,
		[]FreshIngredientRange{},
		[]FreshIngredientRange{},
	)
	ingredientCount := CountAllPossibleFreshIngredients(combinedRanges)
	fmt.Println("Total number of possible fresh ingredients:", ingredientCount)

	return utils.RunAnotherPuzzlePrompt()
}

func CountAllPossibleFreshIngredients(ranges []FreshIngredientRange) int {
	freshCount := 0
	for _, freshRange := range ranges {
		freshCount += freshRange.end - freshRange.start + 1
	}

	return freshCount
}

func CombineRangesRecursive(
	ranges []FreshIngredientRange, 
	modifiedRanges []FreshIngredientRange,
	staticRanges []FreshIngredientRange,
) ([]FreshIngredientRange) {
	remainingRanges := []FreshIngredientRange{}
	//modified := false

	mergedRange := FreshIngredientRange{
		start: 0,
		end: 0,
	}
	for i := 0; i < len(ranges); i++ {
		currRange := ranges[i]
		// On first iteration set merged to first range
		if i == 0 {
			mergedRange.start = currRange.start
			mergedRange.end = currRange.end
			//modified = true
			continue
		}

		// Check if current range and merged range overlap
		// and adjust merged range if so
		if mergedRange.start > currRange.start &&
		mergedRange.start <= currRange.end &&
		mergedRange.end > currRange.end {
			// First, check if merged range overlaps current range to the right,
			// and if so set mergedRange start equal to currRange start
			mergedRange.start = currRange.start
			//modified = true
		} else if mergedRange.start < currRange.start &&
		mergedRange.end >= currRange.start &&
		mergedRange.end < currRange.end {
			// If it doesn't overlap to the right check if it overlaps to the left,
			// and if so set mergedRange end equal to currRange end
			mergedRange.end = currRange.end
			//modified = true
		} else if (mergedRange.start >= currRange.start && mergedRange.end <= currRange.end) {
			// If it doesn't overlap to either side check if the currRange
			// completely surrounds the mergedRange, and if so set both the start and
			// end of the mergedRange to that of the currRange
			mergedRange.start = currRange.start
			mergedRange.end = currRange.end
		} else if (mergedRange.start < currRange.start && mergedRange.end > currRange.end) {
			// The merged range already completely surrounds the curr range, so the curr
			// range should be ignored and not carried over
			continue
		}else if mergedRange.start != currRange.start &&
		mergedRange.end != currRange.end {
			// This range does not overlap the mergedRange
			// Add it to remaining Ranges for recursive call
			remainingRanges = append(remainingRanges, currRange)
		}
	}

	modifiedRanges = append(modifiedRanges, mergedRange)
	// If there are no remaining ranges then need to now iterate through the modified
	// ranges to see if they overlap at all
	if len(remainingRanges) == 0 {
		if  reflect.DeepEqual(modifiedRanges, staticRanges){
			return modifiedRanges
		}
		return CombineRangesRecursive(modifiedRanges, []FreshIngredientRange{}, modifiedRanges)
	}
	return CombineRangesRecursive(remainingRanges, modifiedRanges, staticRanges)
}