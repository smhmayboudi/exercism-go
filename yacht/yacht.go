package yacht

func Sum(dice []int) int {
	total := 0
	for _, n := range dice {
		total += n
	}
	return total
}

func Score(dice []int, category string) int {
	diceCounts := make(map[int]int)
	for _, d := range dice {
		diceCounts[d] += 1
	}

	switch category {
	case "ones":
		return diceCounts[1]
	case "twos":
		return diceCounts[2] * 2
	case "threes":
		return diceCounts[3] * 3
	case "fours":
		return diceCounts[4] * 4
	case "fives":
		return diceCounts[5] * 5
	case "sixes":
		return diceCounts[6] * 6
	case "full house":
		if len(diceCounts) == 2 {
			score := 0
			for d, c := range diceCounts {
				if c != 2 && c != 3 {
					return 0
				}
				score += d * c
			}
			return score
		}
	case "four of a kind":
		for d, c := range diceCounts {
			if c >= 4 {
				return d * 4
			}
		}
	case "little straight":
		for i := 1; i <= 5; i++ {
			if diceCounts[i] != 1 {
				return 0
			}
		}
		return 30
	case "big straight":
		for i := 2; i <= 6; i++ {
			if diceCounts[i] != 1 {
				return 0
			}
		}
		return 30
	case "choice":
		return Sum(dice)
	case "yacht":
		for _, c := range diceCounts {
			if c != 5 {
				return 0
			} else {
				return 50
			}
		}
	}
	return 0
}
