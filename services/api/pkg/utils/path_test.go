package utils

import (
	"testing"
)

func TestPathJoin(t *testing.T) {
	s := PathJoin("http://a.a.com", "/path/", "/to/")
	if s != "http://a.a.com/path/to/" {
		t.Fail()
	}
}

func TestPathBase(t *testing.T) {
	s := PathBase("abc.xlsx")
	if s != "abc" {
		t.Fail()
	}
}