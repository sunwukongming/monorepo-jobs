package utils

import (
	"fmt"
	"strconv"
)

func FloatDecimal(value float64, radix int) float64 {
	s := "%." + strconv.Itoa(radix) + "f"
	value, _ = strconv.ParseFloat(fmt.Sprintf(s, value), 64)
	return value
}
