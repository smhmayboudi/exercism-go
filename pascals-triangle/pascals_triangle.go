package pascal

func Triangle(rows int) [][]int {
	t := 1
	out := [][]int{}
	for i := 0; i < rows; i++ {
		r := []int{}
		for j := 0; j < i+1; j++ {
			if i == 0 || j == 0 {
				t = 1
			} else {
				t = t * (i - j + 1) / j
			}
			r = append(r, t)
		}
		out = append(out, r)
	}
	return out
}
