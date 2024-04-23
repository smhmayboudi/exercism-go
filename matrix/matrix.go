package matrix

import (
	"errors"
	"strconv"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	runes := []rune(s)
	le := len(runes)

	if runes[le-1] == '\n' {
		return nil, errors.New("error \\n at end")
	}

	numMR := [][]int{}
	numMC := []int{}
	for i := 0; i < le; {
		numS := ""

		for j := i; j < le; j++ {
			numC := runes[j]
			if '0' <= numC && numC <= '9' {
				numS = numS + string(numC)
				if j+1 == le {
					if numS != "" {
						num, err := strconv.Atoi(numS)
						if err != nil {
							return nil, errors.New("error strconv")
						}
						numMC = append(numMC, num)
					}

					numMR = append(numMR, numMC)
					numMC = []int{}
					i = i + len(numS) + 1
					break
				}
			} else if numC == '\n' {
				if numS != "" {
					num, err := strconv.Atoi(numS)
					if err != nil {
						return nil, errors.New("error strconv")
					}
					numMC = append(numMC, num)
				}

				numMR = append(numMR, numMC)
				numMC = []int{}
				i = i + len(numS) + 1
				break
			} else if numC == ' ' {
				if numS != "" {
					num, err := strconv.Atoi(numS)
					if err != nil {
						return nil, errors.New("error strconv")
					}
					numMC = append(numMC, num)
				}

				i = i + len(numS) + 1
				break
			} else {
				return nil, errors.New("error")
			}
		}
	}

	invalid := false
	invalid = invalid || len(numMR) == 1 && len(numMR[0]) == 1 && numMR[0][0] != 0
	for r := 0; r < len(numMR)-1; r++ {
		invalid = invalid || (len(numMR[r]) != len(numMR[r+1]))
	}
	if invalid {
		return nil, errors.New("error")
	}

	return numMR, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	mN := Matrix{}
	// for r := 0; r < len(m); r++ {
	for c := 0; c < len(m[0]); c++ {
		col := []int{}
		for r := 0; r < len(m); r++ {
			col = append(col, m[r][c])
		}
		mN = append(mN, col)
	}
	return mN
}

func (m Matrix) Rows() [][]int {
	mN := Matrix{}
	for r := 0; r < len(m); r++ {
		row := []int{}
		for c := 0; c < len(m[r]); c++ {
			row = append(row, m[r][c])
		}
		mN = append(mN, row)
	}
	return mN
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || len(m) <= row || col < 0 || len(m[row]) <= col {
		return false
	}
	m[row][col] = val
	return true
}
