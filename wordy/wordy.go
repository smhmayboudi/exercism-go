package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	qLength := len(question)
	if qLength < 10 {
		return 0, false
	}
	if question[qLength-1] == '?' {
		question = question[:qLength-1]
	}
	words := strings.Split(question, " ")
	length := len(words)
	if length < 3 {
		return 0, false
	}
	if words[0] != "What" || words[1] != "is" {
		return 0, false
	}
	words = words[2:]
	length -= 2
	total, err := strconv.Atoi(words[0])
	if err != nil {
		return 0, false
	}
	length -= 1
	if length == 0 {
		return total, true
	}
	words = words[1:]
	wordsNoBy := make([]string, 0, len(words))
	for _, w := range words {
		if w != "by" {
			wordsNoBy = append(wordsNoBy, w)
		}
	}
	nums := make([]int, 0, len(words))
	ops := make([]string, 0, len(words))
	for i, x := range wordsNoBy {
		if i%2 == 0 {
			ops = append(ops, x)
		} else {
			n, err := strconv.Atoi(x)
			if err == nil {
				nums = append(nums, n)
			} else {
				return 0, false
			}
		}
	}
	if len(ops) != len(nums) {
		return 0, false
	}
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		switch op {
		case "plus":
			total = add(total, nums[i])
		case "minus":
			total = subtract(total, nums[i])
		case "multiplied":
			total = multiply(total, nums[i])
		case "divided":
			total = divide(total, nums[i])
		default:
			return 0, false
		}
	}

	return total, true
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}
