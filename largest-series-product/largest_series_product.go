package lsproduct

import "sort"

func LargestSeriesProduct(digits string, span int) (int64, error) {
	runes := []rune(digits)
	mem := map[int]string{}
	for i := 0; i < len(runes)/3; i++ {
		key := int(digits[i+0]) * int(digits[i+1]) * int(digits[i+2])
		value := digits[i : i+span]
		mem[key] = value
	}
	keys := make([]int, 0, len(mem))
	for k := range mem {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return int64(keys[0]), nil
}
