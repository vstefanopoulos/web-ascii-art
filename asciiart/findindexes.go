package asciiart

import (
	"strings"
)

// FindColorIndexes returns all the indexes that the substring occupies in the
// line in a slice of ints. If the substring is not present or not given it returns nil
func findIndexes(line, subString string) []int {
	var indexes []int
	var index int
	var prev int
	end := len(subString)
	if !strings.Contains(line, subString) || subString == "" {
		return nil
	}
	for {
		index = strings.Index(line, subString)
		if index != -1 {
			for i := 0; i < end; i++ {
				indexes = append(indexes, index+prev+i)
			}
			line = line[index+end:]
			prev += index + end
		} else {
			break
		}
	}
	return indexes
}
