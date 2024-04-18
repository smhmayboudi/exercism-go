package reverse

func Reverse(input string) string {
	in := []rune(input)
	le := len(in)
	i := 0
	j := le - 1
	for i < j {
		in[i], in[j] = in[j], in[i]
		i++
		j--
	}
	return string(in)
}
