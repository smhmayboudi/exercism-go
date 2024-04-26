package matrix

import (
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix [][]int

type Pair struct {
	x int
	y int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	matrix := make(Matrix, len(rows))
	for i := 0; i < len(rows); i++ {
		nums := strings.Split(rows[i], " ")
		matrix[i] = make([]int, len(nums))
		for j := 0; j < len(nums); j++ {
			num, err := strconv.Atoi(nums[j])
			if err != nil {
				return nil, err
			}
			matrix[i] = append(matrix[i], num)
		}
	}
	return &matrix, nil
}

func (m *Matrix) Saddle() []Pair {
	panic("Please implement the Saddle function")
}
