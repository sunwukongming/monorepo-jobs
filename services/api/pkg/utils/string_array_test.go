package utils

import (
	"testing"
)

func TestStringArrayShuffle(t *testing.T) {
	a := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	StringArrayShuffle(a)
}
