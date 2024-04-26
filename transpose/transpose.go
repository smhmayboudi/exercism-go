package transpose

func Transpose(input []string) []string {
	rows := len(input)
	if rows == 0 {
		return []string{}
	}
	cols := len(input[0])
	for _, v := range input {
		if cols < len(v) {
			cols = len(v)
		}
	}
	out := make([]string, cols)
	for _, value := range input {
		for col, val := range value {
			out[col] += string(val)
		}
	}
	return out
}
