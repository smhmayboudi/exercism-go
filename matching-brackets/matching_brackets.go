package brackets

func Bracket(input string) bool {
	out := map[rune]int{}
	for _, value := range input {
		if value == '{' || value == '}' || value == '[' || value == ']' || value == '(' || value == ')' {
			out[value]++
		}
	}
	return out['{'] == out['}'] && out['['] == out[']'] && out['('] == out[')']
}
