package secret

func Handshake(code uint) []string {
	var result []string
	if code&1 == 1 {
		result = append(result, "wink")
	}
	if code&2 == 2 {
		result = append(result, "double blink")
	}
	if code&4 == 4 {
		result = append(result, "close your eyes")
	}
	if code&8 == 8 {
		result = append(result, "jump")
	}
	if code&16 == 16 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}
	return result
}
