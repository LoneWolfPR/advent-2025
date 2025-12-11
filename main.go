package main

import (
	"fmt"

	"github.com/lonewolfpr/advent-2025/day1"
	"github.com/lonewolfpr/advent-2025/day2"
	"github.com/lonewolfpr/advent-2025/day3"
	"github.com/lonewolfpr/advent-2025/day4"
	"github.com/lonewolfpr/advent-2025/day5"
	"github.com/lonewolfpr/advent-2025/day6"
	"github.com/lonewolfpr/advent-2025/day7"
	"github.com/lonewolfpr/advent-2025/day8"
	"github.com/lonewolfpr/advent-2025/day9"
	"github.com/lonewolfpr/advent-2025/day10"
	"github.com/lonewolfpr/advent-2025/day11"
	"github.com/lonewolfpr/advent-2025/day12"
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
	// Day 2
	2: {
		1 : PuzzleEntry{Title: "Gift Shop Part 1", Function: day2.Part1},
		2 : PuzzleEntry{Title: "Gift Shop Part 2", Function: day2.Part2},
	},
	// Day 3
	3: {
		1 : PuzzleEntry{Title: "Lobby (works for both parts)", Function: day3.Part1},
	},
	// Day 4
	4: {
		1 : PuzzleEntry{Title: "Printing Department Part 1", Function: day4.PaperRolls},
		2 : PuzzleEntry{Title: "Printing Department Part 2", Function: day4.MaxPaperRolls},
	},
	// Day 5
	5: {
		1 : PuzzleEntry{Title: "Cafeteria Part 1", Function: day5.IngredientCheck },
		2 : PuzzleEntry{Title: "Cafeteria Part 2", Function: day5.AllFreshIngredients },
	},
	// Day 6
	6: {
		1 : PuzzleEntry{Title: "Trash Compactor Part 1", Function: day6.Part1},
		2 : PuzzleEntry{Title: "Trash Compactor Part 2", Function: day6.Part2},
	},
	// Day 7
	7: {
		1 : PuzzleEntry{Title: "Laboratories Part 1", Function: day7.Part1},
		2 : PuzzleEntry{Title: "Laboratories Part 2", Function: day7.Part2},
	},
	// Day 8
	8: {
		1 : PuzzleEntry{Title: "Playground Part 1", Function: day8.Part1},
		2 : PuzzleEntry{Title: "Playground Part 2", Function: day8.Part2},
	},
	// Day 9
	9: {
		1 : PuzzleEntry{Title: "Movie Theater Part 1", Function: day9.Part1},
		2 : PuzzleEntry{Title: "Movie Theater Part 2", Function: day9.Part2},
	},
	// Day 10
	10: {
		1 : PuzzleEntry{Title: "Factory Part 1", Function: day10.Part1},
		//2 : PuzzleEntry{Title: "Factory Part 2", Function: day10.Part2},
	},
	// Day 11
	11: {
		1 : PuzzleEntry{Title: "Reactor Part 1", Function: day11.Part1},
		//2 : PuzzleEntry{Title: "Reactor Part 2", Function: day11.Part2},
	},
	// Day 12
	12: {
		1 : PuzzleEntry{Title: "Placeholder Part 1", Function: day12.Part1},
		//2 : PuzzleEntry{Title: "Placeholder Part 2", Function: day12.Part2},
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
