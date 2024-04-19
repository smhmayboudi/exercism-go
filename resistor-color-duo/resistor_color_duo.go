package resistorcolorduo

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	sum := 0
	bandColors := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}
	for index, value := range bandColors {
		if colors[0] == index {
			sum += 10 * value
		}
		if colors[1] == index {
			sum += value
		}
	}
	return sum
}
