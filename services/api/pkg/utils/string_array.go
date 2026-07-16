/**

Filename: 		string_array.go
Author: 		gengming - gengming.zb@ccbft.com
Description:	loongrpc string array tool logic
Create:			2022-07-02 15:33:10
Last Modified:	2022-07-03 10:55:10

*/

package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// StringArrayMap 对字符串数组进行遍历操作
func StringArrayMap(a []string, f func(s string) string) []string {
	for i := range a {
		a[i] = f(a[i])
	}
	return a
}

// StringArrayContain 字符串数字包含
func StringArrayContain(a []string, s string) bool {
	for _, item := range a {
		if s == item {
			return true
		}
	}
	return false
}

// StringArrayRemoveAll 移除元素
func StringArrayRemoveAll(a []string, s string) []string {
	b := make([]string, 0)
	for _, item := range a {
		if item != s {
			b = append(b, item)
		}
	}
	return b
}

// StringArrayShuffle 随机打乱数组
func StringArrayShuffle(a []string) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := len(a)
	for i := 0; i < n; i++ {
		j := r.Intn(n)
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// StringArraySlice 取数组的一部分
func StringArraySlice(a []string, start, length int) []string {
	n := len(a)
	if n == 0 {
		return a
	}
	if start >= n {
		return []string{}
	} else if start < 0 {
		start = n + start
		if start < 0 {
			start = 0
		}
	}
	if start+length > n {
		length = n - start
	}
	return a[start : start+length]
}

// StringArrayFromInt 数字数组转字符串数组
func StringArrayFromInt(a []int) []string {
	b := make([]string, 0, len(a))
	for _, item := range a {
		b = append(b, strconv.Itoa(item))
	}
	return b
}

// StringArrayFromInt64 数字数组转字符串数组
func StringArrayFromInt64(a []int64) []string {
	b := make([]string, 0, len(a))
	for _, item := range a {
		b = append(b, strconv.FormatInt(item, 10))
	}
	return b
}

func StringArrayValue(a []string, i int) string {
	if i >= len(a) {
		return ""
	}
	return a[i]
}
