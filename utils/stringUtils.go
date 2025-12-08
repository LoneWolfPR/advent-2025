package utils

func ReverseString (input string) string {
	// convert input to slice of runes
	chars := []rune(input)

	// iterate through the slice, swapping elements from the beginning and end
	for i, j := 0, len(chars) - 1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	// convert slice back to a string
	return string(chars)
}