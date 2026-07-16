package utils

import "strconv"

func UintVal(s string) uint {
	a, _ := strconv.ParseUint(s, 10, 32)
	return uint(a)
}
