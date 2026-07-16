/**

Filename: 		path.go
Author: 		gengming - gengming.zb@ccbft.com
Description:	loongrpc path tool logic
Create:			2022-07-01 16:11:12
Last Modified:	2022-07-01 17:23:32

*/

package utils

import (
	"bytes"
	"path"
	"strings"
)

func PathJoin(a ...string) string {
	bs := bytes.NewBuffer(nil)
	for i, s := range a {
		if i == 0 {
			s = strings.TrimRight(s, "/")
			bs.WriteString(s)
		} else if i == len(a)-1 {
			s = strings.TrimLeft(s, "/")
			bs.WriteString("/")
			bs.WriteString(s)
		} else {
			s = strings.Trim(s, "/")
			bs.WriteString("/")
			bs.WriteString(s)
		}
	}
	return bs.String()
}

func PathBase(s string) string {
	return s[:len(s)-len(path.Ext(s))]
}
