package utils

import (
	"strings"
)

// IntArrayJoin int连结
func IntArrayJoin(a []int, sep string) string {
	return strings.Join(StringArrayFromInt(a), sep)
}

// IntArrayContain 判断数是否在数组中
func IntArrayContain(a []int, i int) bool {
	for _, item := range a {
		if i == item {
			return true
		}
	}
	return false
}
