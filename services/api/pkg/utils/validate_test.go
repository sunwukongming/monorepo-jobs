package utils

import (
	"testing"
)

func TestReg(t *testing.T) {
	if !(ValidateIsEmail("mingming@gmail.com")) {
		t.Fail()
	}
}
