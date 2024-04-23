package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(inputString string) (Matrix, error) {
	rows := strings.Split(inputString, "\n")
	result := make(Matrix, len(rows))

	for index, line := range rows {
		columns := strings.Split(strings.TrimSpace(line), " ")
		result[index] = make([]int, len(columns))

		for position, element := range columns {
			num, err := strconv.Atoi(element)
			if err != nil {
				return nil, errors.New("Problem converting number")
			}
			result[index][position] = num
		}
	}
	return result, result.validate()
}
func (result Matrix) validate() error {
	for r := 0; r < len(numMR)-1; r++ {
		if len(numMR[r]) != len(numMR[r+1]) {
			return errors.New("Invalid matrix")
		}
	}
	return nil
}

func (matrix Matrix) Rows() [][]int {
	rows := make([][]int, len(matrix))

	for index, row := range matrix {
		rows[index] = make([]int, len(row))

		for position, element := range row {
			rows[index][position] = element
		}
	}
	return rows
}

func (matrix Matrix) Cols() [][]int {
	cols := make([][]int, len(matrix[0]))

	for position, _ := range cols {
		cols[position] = make([]int, len(matrix))

		for index, _ := range matrix {
			cols[position][index] = matrix[index][position]
		}
	}
	return cols
}
func (matrix Matrix) Set(row, col, val int) (ok bool) {
	if ok = row >= 0 && col >= 0 && row < len(matrix) && col < len(matrix[0]); ok {
		matrix[row][col] = val
	}
	return
}
