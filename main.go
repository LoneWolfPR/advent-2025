package main

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/day1/secretentrance"
)

var dayPuzzleMap = map[int]map[string]func(){
	// Day 1
	1: {
		"Secret Entrance": secretentrance.Enter,
	},
}

func main() {
	var numDays = len(dayPuzzleMap)
	var selectedDay int
	var selectedPuzzle string

	fmt.Printf("Select a day (1-%d): ", numDays)
	_, err := fmt.Scanln(&selectedDay)
	if err != nil || selectedDay < 1 || selectedDay > numDays {
		fmt.Println("Invalid day selection.")
		return
	}

	puzzles := dayPuzzleMap[selectedDay]
	fmt.Printf("Available puzzles for Day %d:\n", selectedDay)
	i := 1
	puzzleKeys := make([]string, 0, len(puzzles))
	for key := range puzzles {
		fmt.Printf("%d. %s\n", i, key)
		puzzleKeys = append(puzzleKeys, key)
		i++
	}

	fmt.Print("Select a puzzle by number: ")
	var puzzleChoice int
	_, err = fmt.Scanln(&puzzleChoice)
	if err != nil || puzzleChoice < 1 || puzzleChoice > len(puzzles) {
		fmt.Println("Invalid puzzle selection.")
		return
	}

	selectedPuzzle = puzzleKeys[puzzleChoice-1]
	fmt.Printf("Running Day %d - %s:\n", selectedDay, selectedPuzzle)
	puzzles[selectedPuzzle]()
}