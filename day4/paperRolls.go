package day4

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/utils"
)

func PaperRolls() bool {
	layoutFile, err := GatherInput()
	if err != nil {
		fmt.Println("Error gathering input:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	defer layoutFile.Close()

	layoutInfo, err := IngestLayout(layoutFile)
	if err != nil {
		fmt.Println("Error ingesting layout:", err)
		return utils.RunAnotherPuzzlePrompt()
	}
	accessibleRolls := CountAccessibleRolls(layoutInfo)
	fmt.Println("Total accessible paper rolls: ", accessibleRolls)
	return utils.RunAnotherPuzzlePrompt()
}


