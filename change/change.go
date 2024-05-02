package change

import (
	"errors"
)

// Change returns minimum number of coins returned
func Change(coins []int, amount int) ([]int, error) {
	if amount < 0 {
		return nil, errors.New("amount must be >= 0")
	}

	amounts := make([][]int, amount+1)
	amounts[0] = []int{}

	for i := range amounts {
		for _, coin := range coins {
			if i-coin >= 0 && amounts[i-coin] != nil && (amounts[i] == nil || len(amounts[i-coin])+1 < len(amounts[i])) {
				amounts[i] = append([]int{coin}, amounts[i-coin]...)
			}
		}
	}

	if amounts[amount] == nil {
		return nil, errors.New("can not change")
	}

	return amounts[amount], nil
}
