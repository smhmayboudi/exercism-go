package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](cp []T, predicate func(T) bool) []T {
	if cp == nil {
		return cp
	}
	output := make([]T, 0, len(cp))
	for _, value := range cp {
		if predicate(value) {
			output = append(output, value)
		}
	}
	return output
}

func Discard[T any](cp []T, predicate func(T) bool) []T {
	return Keep(cp, func(val T) bool { return !predicate(val) })
}
