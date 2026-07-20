package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestCamelCase(t *testing.T) {
	if s := CamelCase("good_bye"); s != "goodBye" {
		t.Errorf("CamelCase(good_bye) = %q, want goodBye", s)
	}
	if s := CamelCase("Good_bye"); s != "GoodBye" {
		t.Errorf("CamelCase(Good_bye) = %q, want GoodBye", s)
	}
}

func TestSnakeCase(t *testing.T) {
	s := SnakeCase("realName")
	if s != "real_name" {
		t.Fail()
	}
}

func TestUniqueId(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 1000; i++ {
		id := UniqueID()
		if seen[id] {
			t.Fatalf("UniqueID 第 %d 次碰撞: %s", i, id)
		}
		seen[id] = true
	}
}

func TestStringSub(t *testing.T) {
	s := "abcdef"
	if StringSub(s, 0, 3) != "abc" {
		t.Fail()
	}
	s = "你好，我是成功人士"
	if StringSub(s, 2, 3) != "，我是" {
		t.Fail()
	}
	if StringSub(s, -1, 3) != "士" {
		t.Fail()
	}
	if StringSub(s, 9, 3) != "" {
		t.Fail()
	}
	if StringSub(s, 0, 100) != s {
		t.Fail()
	}
}

func TestStringRandom(t *testing.T) {
	fmt.Println(StringRandom(StringAlnumPassword, 100))
}

func TestStringConcat(t *testing.T) {
	s := StringConcat("/", "path/", "to/", "/name")
	if s != "path/to/name" {
		t.Log(s)
		t.Fail()
	}
}

func TestReplace(t *testing.T) {
	replacer := strings.NewReplacer(
		"#name#", "名字",
		"#候选人姓名#", "名字",
		"#position#", "职位",
		"#职位名称#", "职位",
	)

	emailContent := replacer.Replace("<p>候选人姓名:#name#</p>↵↵<p>职位名:#position##position##position#</p>↵↵<p>公司名:#company#</p>↵↵<p>面试时间:")
	fmt.Println(emailContent)
}
