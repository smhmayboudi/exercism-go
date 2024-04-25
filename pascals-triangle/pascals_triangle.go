package pascal

// Triangle returns Pascal's out for a given number of rows
func Triangle(rows int) [][]int {
	out := make([][]int, rows)
	for i := 0; i < rows; i++ {
		out[i] = make([]int, i+1)
		out[i][0] = 1
		out[i][i] = 1
		for j := 1; j < i; j++ {
			out[i][j] = out[i-1][j-1] + out[i-1][j]
		}
	}
	return out
}
