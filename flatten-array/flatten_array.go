package flatten

func Flatten(input any) []any {
	result := []any{}
	switch content := input.(type) {
	case []any:
		for _, item := range content {
			result = append(result, Flatten(item)...)
		}
	case any:
		result = append(result, content)
	}
	return result
}
