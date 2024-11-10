package asciiart

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	var lineCount int = 8

	for scanner.Scan() {
		// every 8 lines there is one that only contains "\n" including the first
		if lineCount == 8 {
			lineCount = 0
			continue
		}
		lines = append(lines, scanner.Text())
		lineCount++
	}
	return lines, nil
}

func writeToFile(filename, content string) error {
	// Create or truncate the file (overwrite if exists)
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write the content to the file
	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
