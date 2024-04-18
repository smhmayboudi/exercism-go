package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](cp []T, predicate func(T) bool) []T {
	var finalList []T
	for _, v := range cp {
		if predicate(v) {
			finalList = append(finalList, v)
		}
	}
	return finalList
}

func Discard[T any](cp []T, predicate func(T) bool) []T {
	var finalList []T
	for _, v := range cp {
		if !predicate(v) {
			finalList = append(finalList, v)
		}
	}
	return finalList
}
