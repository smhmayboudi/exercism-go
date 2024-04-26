package wordsearch

import (
	"errors"
)

var errWordNotFound = errors.New("Word Not Found")

// Solve finds the first instance of each word in a square grid.
// Returns errWordNotFound as only error.
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	m := make(map[string][2][2]int)
	for _, w := range words {
		if loc, ok := find(puzzle, w); ok {
			m[w] = loc
			continue
		}
		return nil, errWordNotFound
	}
	return m, nil
}

func find(puzzle []string, w string) (loc [2][2]int, ok bool) {
	for j, s := range puzzle {
		for i, r := range s {
			if byte(r) != w[0] {
				continue
			}
			if loc, ok = test(i, j, puzzle, w); ok {
				return loc, true
			}
		}
	}
	return
}

type dir struct {
	x, y int
}

var dirs = []dir{
	{0, -1},  // north
	{1, -1},  // northeast
	{1, 0},   // east
	{1, 1},   // southeast
	{0, 1},   // south
	{-1, 1},  // southwest
	{-1, 0},  // west
	{-1, -1}, // northwest
}

// test assumes that the first character already matches
func test(i, j int, puzzle []string, word string) (loc [2][2]int, ok bool) {
	l := len(word) - 1
	for _, d := range dirs {
		loc = [2][2]int{{i, j}, {i + d.x*l, j + d.y*l}}
		if ok = match(i, j, d, puzzle, word); ok {
			return
		}
	}
	return
}

func match(i, j int, d dir, puzzle []string, word string) bool {
	for t := range word {
		x, y := i+d.x*t, j+d.y*t
		if x < 0 || y < 0 || y >= len(puzzle) || x >= len(puzzle[y]) || word[t] != puzzle[y][x] {
			return false
		}
	}
	return true
}
