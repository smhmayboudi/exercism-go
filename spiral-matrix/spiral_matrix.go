package spiralmatrix

func SpiralMatrix(size int) [][]int {
	out := make([][]int, size)
	for r := 0; r < size; r++ {
		out[r] = make([]int, size)
	}
	lastRow := 0
	lastC := 0
	for i := 0; i < size*size; i++ {
		out[lastRow][lastC] = i + 1
		if lastRow == 0 && lastC+1 < size {
			lastC++
		} else if lastC == size-1 && lastRow+1 < size {
			lastRow++
		} else if lastRow == size && 0 < lastC+1 {
			lastC--
		} else if lastC == 0 && 0 < lastRow+1 {
			lastRow--
		}
	}
	return out
}
