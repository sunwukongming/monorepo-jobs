package utils

import (
	"testing"
)

func TestTimeSecondHumanize(t *testing.T) {
	if TimeSecondHumanize(60) != "00:01:00" {
		t.Log("60 不正确")
		t.Fail()
	}
	if TimeSecondHumanize(59) != "00:00:59" {
		t.Log("59 不正确")
		t.Fail()
	}
	if TimeSecondHumanize(110) != "00:01:50" {
		t.Log("110 不正确")
		t.Fail()
	}
	if TimeSecondHumanize(3690) != "01:01:30" {
		t.Log("3690 不正确")
		t.Fail()
	}
}