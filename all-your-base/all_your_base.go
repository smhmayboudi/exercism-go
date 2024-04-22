// Package allyourbase provides number conversions.
package allyourbase

import (
	"errors"
)

// ConvertToBase converts a nb expressed using a base into another using another base.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {

	// Error cases
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	// First, convert to decimal number
	nb := 0
	mul := 1
	idlen := len(inputDigits)
	for i := 0; i < idlen; i++ {
		nb += mul * inputDigits[idlen-1-i]
		mul *= inputBase
	}

	// Shortcut
	if nb < outputBase {
		return []int{nb}, nil
	}

	// Else computes dest number right-to-left
	tmp := []int{}
	for nb >= outputBase {
		tmp = append(tmp, nb%outputBase)
		nb /= outputBase
	}
	if nb != 0 {
		tmp = append(tmp, nb)
	}

	// And reverse result
	result := make([]int, len(tmp))
	for i := 0; i < len(tmp); i++ {
		result[i] = tmp[len(tmp)-1-i]
	}

	// Done!
	return result, nil
}
