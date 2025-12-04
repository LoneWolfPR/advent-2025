package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/lonewolfpr/advent-2025/utils"
)

func Part1() bool {
	inputData , err := GatherInput()
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer inputData.BatteryFile.Close()

	result, err := processBatteryBanks(inputData.BatteryFile, inputData.NumDigits)
	if err != nil {
		fmt.Println("Error processing battery banks:", err)
		return false
	}
	fmt.Println("Total joltage from battery banks: ", result)
	return utils.RunAnotherPuzzlePrompt()
}

func processBatteryBanks(file *os.File, numDigits int) (int, error) {
	totalJoltage := 0
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		joltageValueStr, err := GetJoltageValue(line, numDigits)
		if err != nil {
			return 0, fmt.Errorf("Error getting highest digits from line %s: %w", line, err)
		}
		bankJoltage, err := strconv.Atoi(joltageValueStr)
		if err != nil {
			return 0, fmt.Errorf("Error converting joltage string to int %s: %w", joltageValueStr, err)
		}
		totalJoltage += bankJoltage
	}

	return totalJoltage, nil
}

type InputData struct {
	BatteryFile *os.File
	NumDigits int
}

func GatherInput() (*InputData, error) {
	var filePath string
	var numDigits int
	fmt.Print("Enter path to battery banks file: ")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file path: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}

	fmt.Print("Enter number of digits to consider for joltage calculation: ")
	_, err = fmt.Scanln(&numDigits)
	if err != nil {
		return nil, fmt.Errorf("Error reading number of digits: %w", err)
	}

	return &InputData{
		BatteryFile: file,
		NumDigits: numDigits,
	}, nil
}

type HighestDigitInfo struct {
	Index int
	Value string
}

func GetJoltageValue(batteryStr string, numDigitsToGet int) (string, error) {
	highestValueIndex := 0
	joltageStr := ""
	for i := 0; i < numDigitsToGet; i++ {
		digitInfo, err := GetHighestDigitWithinRange(batteryStr, highestValueIndex, len(batteryStr)-(numDigitsToGet-i))
		if err != nil {
			return "", fmt.Errorf("Error getting highest digit within range: %w", err)
		}
		joltageStr += digitInfo.Value
		highestValueIndex = digitInfo.Index + 1
	}
	return joltageStr, nil
}

func GetHighestDigitWithinRange(batteryStr string, startIndex int, endIndex int) (*HighestDigitInfo, error) {
	highestValue := 0
	highestValueIndex := 0
	for i := startIndex; i <= endIndex; i++ {
		digitValue, err := strconv.Atoi(string(batteryStr[i]))
		if err != nil {
			return nil, fmt.Errorf("Error converting character to int %s: %w", string(batteryStr[i]), err)
		}
		if digitValue > highestValue {
			highestValue = digitValue
			highestValueIndex = i
		}
	}
	return &HighestDigitInfo{
		Index: highestValueIndex,
		Value: strconv.Itoa(highestValue),
	}, nil
}