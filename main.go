package main

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/day1"
	"github.com/lonewolfpr/advent-2025/day2"
)
type PuzzleEntry struct {
	Title string
	Function func() bool
}

var dayPuzzleMap = map[int]map[int]PuzzleEntry{
	// Day 1
	1: {
		1 : PuzzleEntry{Title: "Secret Entrance Part 1", Function: day1.Part1},
		2 : PuzzleEntry{Title: "Secret Entrance Part 2", Function: day1.Part2},
	},
	2: {
		1 : PuzzleEntry{Title: "Product ID Validation Part 1", Function: day2.Part1},
		2 : PuzzleEntry{Title: "Product ID Validation Part 2", Function: day2.Part2},
	},
}

func main() {
	fmt.Println("Welcome to Advent 2025!")
	keepRunning := true
	for keepRunning {
		keepRunning = mainMenu()
	}
}

func mainMenu() bool {
	var numDays = len(dayPuzzleMap)
	var selectedDay int
	var selectedPuzzle PuzzleEntry
	fmt.Println("\n---Main Menu---")
	fmt.Printf("Select a day (1-%d) or press 0 to quit: ", numDays)
	_, err := fmt.Scanln(&selectedDay)
	if err != nil || selectedDay < 0 || selectedDay > numDays {
		fmt.Println("Invalid day selection.")
		return true
	}
	if selectedDay == 0 {
		fmt.Println("Exiting Advent 2025. Goodbye!")
		return false
	}

	puzzles := dayPuzzleMap[selectedDay]
	fmt.Printf("Available puzzles for Day %d:\n", selectedDay)
	
	// Start with Go Back Option
	fmt.Println("0. Go Back")
	for i := 1; i <= len(puzzles); i++{
		fmt.Printf("%d. %s\n", i, puzzles[i].Title)
	}

	fmt.Print("Select a puzzle by number: ")
	var puzzleChoice int
	_, err = fmt.Scanln(&puzzleChoice)
	if err != nil || puzzleChoice < 0 || puzzleChoice > len(puzzles) {
		fmt.Println("Invalid puzzle selection.")
		return true
	}
	if puzzleChoice == 0 {
		return true
	}

	selectedPuzzle = puzzles[puzzleChoice]
	fmt.Printf("Running Day %d - %s:\n", selectedDay, selectedPuzzle.Title)
	return selectedPuzzle.Function()
}
