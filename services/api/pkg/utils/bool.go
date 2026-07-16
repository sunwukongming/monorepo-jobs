package utils

// BoolInt int转bool
func BoolInt(i int) bool {
	if i != 0 {
		return true
	}
	return false
}

// BoolString string转bool
func BoolString(s string) bool {
	if s != "" {
		return true
	}
	return false
}

// BoolFloat64 float转bool
func BoolFloat64(a float64) bool {
	if a != 0 {
		return true
	}
	return false
}
