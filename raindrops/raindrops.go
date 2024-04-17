package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	result := ""
	conditions := map[int]string{3: "Pling", 5: "Plang", 7: "Plong"}
	for key, value := range conditions {
		if number%key == 0 {
			result += value
		}
	}
	if result == "" {
		result += strconv.Itoa(number)
	}
	return result
}
