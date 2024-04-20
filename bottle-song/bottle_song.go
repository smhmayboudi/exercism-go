package bottlesong

import (
	"fmt"
	"strings"
)

func Recite(startBottles, takeDown int) []string {
	// "Ten green bottles hanging on the wall,"
	// "Ten green bottles hanging on the wall,"
	// "And if one green bottle should accidentally fall,"
	// "There'll be nine green bottles hanging on the wall."
	num := map[int][2]string{
		10: {"Ten", "bottles"},
		9:  {"Nine", "bottles"},
		8:  {"Eight", "bottles"},
		7:  {"Seven", "bottles"},
		6:  {"Six", "bottles"},
		5:  {"Five", "bottles"},
		4:  {"Four", "bottles"},
		3:  {"Three", "bottles"},
		2:  {"Two", "bottles"},
		1:  {"One", "bottle"},
		0:  {"No", "bottles"},
	}
	res := []string{}
	for i, j := startBottles, 0; 0 < i && j < takeDown; i, j = i-1, j+1 {
		res = append(res, fmt.Sprintf("%s green %s hanging on the wall,", num[i][0], num[i][1]))
		res = append(res, fmt.Sprintf("%s green %s hanging on the wall,", num[i][0], num[i][1]))
		res = append(res, "And if one green bottle should accidentally fall,")
		res = append(res, fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(num[i-1][0]), num[i-1][1]))
		if j+1 < takeDown {
			res = append(res, "")
		}
	}
	return res
}
