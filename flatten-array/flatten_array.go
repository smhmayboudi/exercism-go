package flatten

func Flatten(nested any) []any {
	numsIN := Flat(nested)
	out := []any{}
	for i := 0; i < len(numsIN); i++ {
		n, _ := numsIN[i].(int)
		out = append(out, n)
	}
	return out
}

func Flat(nested any) []any {
	nums := []any{}
	ok := false

	nums, ok = nested.([]any)
	for ok == false {
		return Flat(nums)
	}

	return nums
}
