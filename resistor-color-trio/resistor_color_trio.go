package resistorcolortrio

import (
	"fmt"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	ribbon := map[string]int{
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
	value0 := ribbon[colors[0]]
	value1 := ribbon[colors[1]]
	res := 10*value0 + value1
	for i := 2; i < len(colors) && i < 3; i++ {
		value := colors[i]
		ab := ribbon[value]
		ten := 1
		for j := 0; j < ab; j++ {
			ten *= 10
		}
		res *= ten
	}
	if res/1_000_000_000 != 0 && res%1_000_000_000 == 0 {
		return fmt.Sprintf("%d gigaohms", res/1_000_000_000)
	} else if res/1_000_000 != 0 && res%1_000_000 == 0 {
		return fmt.Sprintf("%d megaohms", res/1_000_000)
	} else if res/1_000 != 0 && res%1_000 == 0 {
		return fmt.Sprintf("%d kiloohms", res/1_000)
	}
	return fmt.Sprintf("%d ohms", res)
}
