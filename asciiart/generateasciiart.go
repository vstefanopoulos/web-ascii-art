package asciiart

import (
	"fmt"
	"strings"
)

func GenerateAsciiArt(inputString, banner, align string) (string, error) {
	var padStr string
	var result string
	var lastPad string

	fileLines, err := ReadFile(banner)
	if err != nil {
		return "", err
	}

	inputString = strings.ReplaceAll(inputString, "\r\n", "\n")
	lines := strings.Split(inputString, "\n")
	maxLen := findBiggerString(lines, fileLines)
	terminalWidth, err := lineLenInTerminal(lines[maxLen], fileLines)
	if err != nil {
		return "", err
	}

	for _, line := range lines {
		// Creating the padding string
		if align == "center" || align == "right" {
			lineWidth, err := lineLenInTerminal(line, fileLines)
			if err != nil {
				return "", err
			}
			padStr = paddingCR(terminalWidth, lineWidth, align)
		}
		if align == "justify" {
			temp, spaceCount := removeSpaces(line)
			lineWidth, err := lineLenInTerminal(temp, fileLines)
			if err != nil {
				return "", err
			}
			padStr, lastPad = paddingJustify(terminalWidth, lineWidth, spaceCount, align)
		}

		switch {
		case line == "":
			result += fmt.Sprint("\n")
		default:
			for currentLine := 0; currentLine < 8; currentLine++ {
				if align == "center" || align == "right" {
					result += fmt.Sprint(padStr)
				}
				for charPosition, char := range line {
					if align == "justify" && char == ' ' {
						result += printPad(line, lastPad, padStr, charPosition)
						continue
					}
					letterIndxInFile := (int(char) - 32) * 8
					result += fmt.Sprint(fileLines[letterIndxInFile+currentLine])
				}
				result += fmt.Sprint("\n")
			}
		}
	}
	return result, nil
}

// checks at which space char we are in the line and returns the pad unless its the last one
func printPad(line, lastPad, padStr string, charPosition int) string {
	spaceIndexes := findIndexes(line, " ")
	if len(lastPad) != 0 && charPosition == spaceIndexes[len(spaceIndexes)-1] {
		return fmt.Sprint(lastPad)
	}
	return fmt.Sprint(padStr)
}
