package alphametics

import "strings"

func Solve(puzzle string) (map[string]int, error) {
	strs := strings.Split(puzzle, " == ")
	o1 := [2]map[rune]int{}
	for i1, v1 := range strs {
		o1[i1] = map[rune]int{}
		for _, v2 := range v1 {
			if v2 < 'A' || 'Z' < v2 {
				continue
			}
			o1[i1][v2]++
		}
	}
}
