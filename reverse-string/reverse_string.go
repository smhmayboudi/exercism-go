package reverse

func Reverse(input string) string {
	in := []rune(input)
	le := len(in)
	for i := 0; i < le/2; i++ {
		in[i], in[le-1-i] = in[le-1-i], in[i]
	}
	return string(in)
}
