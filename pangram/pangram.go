package pangram

func IsPangram(input string) bool {
	// const alphamask int = 0b11111111111111111111111111
	// result := 0
	// length := len(input)
	// for i := 0; i < length; i++ {
	// 	value := input[i]
	// 	if value > 96 && value < 123 {
	// 		result |= 1 << (value - 97)
	// 	} else if value > 64 && value < 91 {
	// 		result |= 1 << (value - 65)
	// 	}
	// }
	// return alphamask&result == alphamask

	const alphamask int = 0b11111111111111111111111111
	result := 0
	for _, value := range input {
		if value > 96 && value < 123 {
			result |= 1 << (value - 97)
		} else if value > 64 && value < 91 {
			result |= 1 << (value - 65)
		}
	}
	return alphamask&result == alphamask
}
