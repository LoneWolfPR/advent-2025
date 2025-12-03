package day2

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part1() bool {
	allRanges, err := GatherInput()
	if err != nil {
		fmt.Println(err)
		return false
	}
	result, err := ProcessIDRanges(allRanges, validateProductID)
	if err != nil {
		fmt.Println("Error processing ID ranges:", err)
	}
	fmt.Println("Sum of all invalid product IDs: ", result)
	return utils.RunAnotherPuzzlePrompt()
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