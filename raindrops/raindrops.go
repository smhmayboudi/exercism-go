package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	result := ""
	conditions := [...]map[int]string{{3: "Pling"}, {5: "Plang"}, {7: "Plong"}}
	for _, value := range conditions {
		for key, val := range value {
			if number%key == 0 {
				result += val
			}
		}
	}
	if result == "" {
		result += strconv.Itoa(number)
	}
	return result
}
