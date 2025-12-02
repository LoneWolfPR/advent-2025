package main

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/day1/secretentrance"
)
type PuzzleEntry struct {
	Title string
	Function func()
}

var dayPuzzleMap = map[int]map[int]PuzzleEntry{
	// Day 1
	1: {
		1 : PuzzleEntry{Title: "Secret Entrance Part 1", Function: secretentrance.Part1},
		2 : PuzzleEntry{Title: "Secret Entrance Part 2", Function: secretentrance.Part2},
	},
}

func main() {
	var numDays = len(dayPuzzleMap)
	var selectedDay int
	var selectedPuzzle PuzzleEntry

	fmt.Printf("Select a day (1-%d): ", numDays)
	_, err := fmt.Scanln(&selectedDay)
	if err != nil || selectedDay < 1 || selectedDay > numDays {
		fmt.Println("Invalid day selection.")
		return
	}

	puzzles := dayPuzzleMap[selectedDay]
	fmt.Printf("Available puzzles for Day %d:\n", selectedDay)
	
	for i := 1; i <= len(puzzles); i++{
		fmt.Printf("%d. %s\n", i, puzzles[i].Title)
	}

	fmt.Print("Select a puzzle by number: ")
	var puzzleChoice int
	_, err = fmt.Scanln(&puzzleChoice)
	if err != nil || puzzleChoice < 1 || puzzleChoice > len(puzzles) {
		fmt.Println("Invalid puzzle selection.")
		return
	}

	selectedPuzzle = puzzles[puzzleChoice]
	fmt.Printf("Running Day %d - %s:\n", selectedDay, selectedPuzzle.Title)
	selectedPuzzle.Function()
}