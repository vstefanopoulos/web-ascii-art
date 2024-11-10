package asciiart

import (
	"fmt"
	"strings"
)

// returns a string with space characters to be printed to the left or instead of space characters
func paddingCR(terminalWidth, lineWidth int, align string) string {
	var padding int
	var div int
	switch align {
	case "center":
		div = 2
	case "right":
		div = 1
	}
	padding = (terminalWidth - lineWidth) / div

	if padding > 0 {
		return fmt.Sprintf(strings.Repeat(" ", padding))
	}
	return ""
}

func paddingJustify(terminalWidth, lineWidth, spaceCount int, align string) (string, string) {
	var padding int
	var lastPad int
	if spaceCount == 0 {
		return "", ""
	}
	padding = (terminalWidth - lineWidth) / spaceCount

	if extraSpaces := (terminalWidth - lineWidth) % spaceCount; extraSpaces != 0 {
		padding += extraSpaces / spaceCount
		lastPad = padding + extraSpaces%spaceCount
	}

	if padding > 0 {
		return fmt.Sprintf(strings.Repeat(" ", padding)), fmt.Sprintf(strings.Repeat(" ", lastPad))
	}
	return "", ""
}

// returns the width of the given string in characters when printed with ascii-art
func lineLenInTerminal(s string, fileLines []string) (int, error) {
	var lenInputString int
	for _, char := range s {
		if int(char) > 126 || int(char) < 32 {
			return 0, fmt.Errorf("Bad request")
		}
		lenInputString += len(fileLines[(int(char)-32)*8])
	}
	return lenInputString, nil
}

// removes spaces from the string to calculate the padding for justify
func removeSpaces(s string) (string, int) {
	var count int
	var result string
	for _, char := range s {
		if char == ' ' {
			count++
			continue
		}
		result += string(char)
	}
	return result, count
}

func findBiggerString(sl, fileLines []string) int {
	var result int = 0
	var bigger int
	for i, s := range sl {
		lenS, err := lineLenInTerminal(s, fileLines)
		if err != nil {
			return 0
		}
		if lenS > bigger {
			bigger = lenS
			result = i
		}
	}
	return result
}
