package series

func All(n int, s string) []string {
	l := len(s)
	out := []string{}
	for i := 0; i < l; i++ {
		out = append(out, s[0+i:n+i])
	}
	return out
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}
