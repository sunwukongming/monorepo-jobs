package utils

import "strconv"

func IntBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IntVal(i interface{}) int {
	var k int
	if v, ok := i.(string); ok {
		k, _ = strconv.Atoi(v)
	} else if v, ok := (i.(bool)); ok {
		if v {
			k = 1
		}
	} else if v, ok := (i.(int)); ok {
		k = v
	} else if v, ok := (i.(int8)); ok {
		k = int(v)
	} else if v, ok := (i.(int16)); ok {
		k = int(v)
	} else if v, ok := (i.(int32)); ok {
		k = int(v)
	} else if v, ok := (i.(int64)); ok {
		k = int(v)
	} else if v, ok := (i.(uint)); ok {
		k = int(v)
	} else if v, ok := (i.(uint8)); ok {
		k = int(v)
	} else if v, ok := (i.(uint16)); ok {
		k = int(v)
	} else if v, ok := (i.(uint32)); ok {
		k = int(v)
	} else if v, ok := (i.(uint64)); ok {
		k = int(v)
	} else if v, ok := (i.(float32)); ok {
		k = int(v)
	} else if v, ok := (i.(float64)); ok {
		k = int(v)
	} else {
		//TODO
	}
	return k
}

func IntABS(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
