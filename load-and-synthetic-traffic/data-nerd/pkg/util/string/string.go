package string

import "strings"

// InSlice check if an array is contained in a slice
func InSlice(a string, list []string) int {
	for index, b := range list {
		if strings.EqualFold(b, a) {
			return index
		}
	}
	return -1
}
