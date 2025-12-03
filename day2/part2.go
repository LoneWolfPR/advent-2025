package day2

import (
	"fmt"
	"strings"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part2() bool {
	allRanges, err := GatherInput()
	if err != nil {
		fmt.Println(err)
		return false
	}
	result, err := ProcessIDRanges(allRanges, validateProductIDNew)
	if err != nil {
		fmt.Println("Error processing ID ranges:", err)
	}
	fmt.Println("Sum of all invalid product IDs: ", result)
	return utils.RunAnotherPuzzlePrompt()
}

func validateProductIDNew(productID int) bool {
	idStr := fmt.Sprintf("%d", productID)
	idLen := len(idStr)
	midIndex := idLen / 2
	for i := 0; i < midIndex; i++ {
		pattern := idStr[:i+1]
		if isRepeatingSequence(idStr, pattern) {
			return false
		}
		
	}
	return true
}

func isRepeatingSequence(fullString string, pattern string) bool {
	subStrs := []string{}
	for subStr := range strings.SplitSeq(fullString, pattern) {
		if subStr != "" {
			subStrs = append(subStrs, subStr)
		}
	}

	return len(subStrs) == 0
}
