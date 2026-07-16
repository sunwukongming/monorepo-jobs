/**

Filename: 		int64.go
Author: 		gengming - gengming.zb@ccbft.com
Description:	loongrpc int64 tool logic
Create:			2022-07-03 14:21:23
Last Modified:	2022-07-03 15:03:11

*/

package utils

import "strconv"

func Int64ABS(a int64) int64 {
	if a < 0 {
		a = -1
	}
	return a
}

func Int64Val(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

// IntArrayContain 判断数是否在数组中
func Int64ArrayContain(a []int64, i int64) bool {
	for _, item := range a {
		if i == item {
			return true
		}
	}
	return false
}
