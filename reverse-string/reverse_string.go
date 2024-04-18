package reverse

func Reverse(input string) string {
	in := []rune(input)
	le := len(in)
	// for i, j := 0, le-1; i < le; i, j = i+1, j-1 {
	// 	in[i], in[j] = in[j], in[i]
	// }
	i := 0
	j := le - 1
	for i < j {
		in[i], in[j] = in[j], in[i]
		i++
		j--
	}
	return string(in)
}
