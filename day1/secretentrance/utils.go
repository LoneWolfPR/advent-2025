package secretentrance
import (
	"fmt"
	"os"
)

type InputData struct {
	SequenceFile *os.File
	StartingPoint int
}

func GatherInput() (*InputData, error) {
	var filePath string
	var startingPoint int
	fmt.Print("Enter path to rotation sequence file: ")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file path: %w", err)
	}

	fmt.Print("Enter starting point: ")
	_, err = fmt.Scanln(&startingPoint)
	if err != nil {
		return nil, fmt.Errorf("Error reading starting point: %w", err)
	}

	rotSeqFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}

	return &InputData{
		SequenceFile: rotSeqFile,
		StartingPoint: startingPoint,
	}, nil
}