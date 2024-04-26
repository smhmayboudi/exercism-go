package matrix

import (
	"math"
	"strconv"
	"strings"
)

type Matrix [][]int
type Pair [2]int

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	matrix := make(Matrix, len(rows))
	for i, row := range rows {
		cells := strings.Fields(row)
		matrix[i] = make([]int, len(cells))
		for j, cell := range cells {
			val, _ := strconv.Atoi(cell)
			matrix[i][j] = val
		}
	}
	return &matrix, nil
}

func (m *Matrix) Saddle() (pairs []Pair) {

	// Initialise rowMaxs and colMins
	rowMaxs, colMins := make([]int, len(*m)), make([]int, len((*m)[0]))
	for i := range colMins {
		colMins[i] = math.MaxInt
	}

	// Update the max for each row and min for each column
	for i, row := range *m {
		for j, cell := range row {
			if cell > rowMaxs[i] {
				rowMaxs[i] = cell
			}
			if cell < colMins[j] {
				colMins[j] = cell
			}
		}
	}

	// Return all cells that are both a row max and columm min
	for i, row := range *m {
		for j, cell := range row {
			if cell == rowMaxs[i] && cell == colMins[j] {
				pairs = append(pairs, Pair{i + 1, j + 1})
			}
		}
	}
	return pairs

}
