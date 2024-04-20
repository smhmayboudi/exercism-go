package prime

func Factors(i int64) []int64 {
	var o = []int64{}
	for divisor := int64(2); i > 1; divisor++ {
		for ; i%divisor == 0; i /= divisor {
			o = append(o, divisor)
		}
	}
	return o
}
