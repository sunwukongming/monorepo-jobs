package utils

import (
	"bytes"
	"math/rand"
	"strconv"
	"strings"
	"time"
	_ "time/tzdata"
	"unicode"
)

// StringAlpha 字母
var StringAlpha string

// StringNumber 数字
var StringNumber string = "1234567890"

// StringAlnum 字母数字
var StringAlnum string

// StringAlnumPassword 字母数字密码
var StringAlnumPassword string

func init() {
	bf := bytes.NewBuffer(nil)
	var i byte
	for i = 'a'; i <= 'z'; i++ {
		bf.WriteByte(i)
	}
	for i = 'A'; i <= 'Z'; i++ {
		bf.WriteByte(i)
	}
	StringAlpha = bf.String()
	StringAlnum = StringAlpha + StringNumber
	StringAlnumPassword = StringAlnum
	StringAlnumPassword = strings.ReplaceAll(StringAlnumPassword, "0", "")
	StringAlnumPassword = strings.ReplaceAll(StringAlnumPassword, "O", "")
	StringAlnumPassword = strings.ReplaceAll(StringAlnumPassword, "l", "")
}

// SnakeCase 下划线化
func SnakeCase(s string) string {
	b := &strings.Builder{}
	length := len(s)
	for i := 0; i < length; i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			b.WriteByte('_')
			b.WriteByte(s[i] + 32)
		} else {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

// CamelCase 驼峰化
func CamelCase(s string) string {
	b := &strings.Builder{}
	length := len(s)
	for i := 0; i < length; i++ {
		if s[i] == '_' {
			if i == length-1 {

			} else {
				i += 1
				c := s[i]
				if 'a' <= s[i] && s[i] <= 'z' {
					c = s[i] - 32
				}
				b.WriteByte(c)
			}
		} else {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

// FirstToLower 首字母消息
func FirstToLower(s string) string {
	if unicode.IsUpper([]rune(s)[0]) {
		sr := []rune(s)
		sr[0] = unicode.ToLower(sr[0])
		return string(sr)
	}
	return s
}

// UniqueID 生成唯一id
func UniqueID() string {
	t := time.Now().UnixNano()
	return strconv.FormatInt(t, 36)
}

// StringSub 字符串截取
func StringSub(s string, start int, length int) string {
	var a = []rune(s)
	n := len(a)
	if n == 0 {
		return ""
	}
	if start >= n {
		return ""
	} else if start < 0 {
		start = n + start
		if start < 0 {
			start = 0
		}
	}
	if start+length > n {
		length = n - start
	}
	return string(a[start : start+length])
}

// StringRandom 生成随机字符串
func StringRandom(s string, n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bf := bytes.NewBuffer(nil)
	for i := 0; i < n; i++ {
		idx := r.Intn(len(s))
		bf.WriteByte(s[idx])
	}
	return bf.String()
}

// StringConcat 字符串拼接
func StringConcat(sep string, a ...string) string {
	bf := bytes.NewBuffer(nil)
	for i, s := range a {
		if i == 0 {
			bf.WriteString(strings.TrimRight(s, sep))
		} else if i == len(a)-1 {
			bf.WriteString(strings.TrimLeft(s, sep))
		} else {
			bf.WriteString(strings.Trim(s, sep))
		}
		if i != len(a)-1 {
			bf.WriteString(sep)
		}
	}
	return bf.String()
}

func StringToTime(s string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	timestamp, err := time.ParseInLocation(timeLayout, s, TimeLocation)
	if err != nil {
		return 0
	}
	return timestamp.Unix()
}

func StringTrim(s string) string {
	return strings.Trim(s, " ")
}

// 数字转成excel的列标识字母
func Div(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if Num > 26 { //数据大于26需要进行拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}

	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}

func StringStartsWith(s string, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return s[:len(prefix)] == prefix
}
