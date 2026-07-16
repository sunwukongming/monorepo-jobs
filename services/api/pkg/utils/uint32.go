package utils

func Uint32ArrayFromStringArray(a []string) []uint32 {
	b := make([]uint32, len(a))
	for _, item := range a {
		b = append(b, uint32(IntVal(item)))
	}
	return b
}
