package etl

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for key, value := range in {
		for i := 0; i < len(value); i++ {
			newKey := []rune(value[i])[0]
			out[string(newKey-'A'+'a')] = key
		}
	}
	return out
}
