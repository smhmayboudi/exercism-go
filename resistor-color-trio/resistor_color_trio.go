package resistorcolortrio

import "strconv"

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	value0 := Ribbon(colors[0])
	value1 := Ribbon(colors[1])
	res := 10*value0 + value1
	for i := 2; i < len(colors) && i < 3; i++ {
		value := colors[i]
		ab := Ribbon(value)
		ten := 1
		for j := 0; j < ab; j++ {
			ten *= 10
		}
		res *= ten
	}
	if res/1_000_000_000 != 0 && res%1_000_000_000 == 0 {
		return strconv.Itoa(res/1_000_000_000) + " gigaohms"
	} else if res/1_000_000 != 0 && res%1_000_000 == 0 {
		return strconv.Itoa(res/1_000_000) + " megaohms"
	} else if res/1_000 != 0 && res%1_000 == 0 {
		return strconv.Itoa(res/1_000) + " kiloohms"
	}
	return strconv.Itoa(res) + " ohms"
}

func Ribbon(color string) int {
	switch color {
	case "black":
		return 0
	case "brown":
		return 1
	case "red":
		return 2
	case "orange":
		return 3
	case "yellow":
		return 4
	case "green":
		return 5
	case "blue":
		return 6
	case "violet":
		return 7
	case "grey":
		return 8
	case "white":
		return 9
	default:
		return 0
	}
}
