package perfect

import "errors"

// Define the Classification type here.
type Classification byte

const (
	ClassificationUndefined Classification = iota // undefined
	ClassificationDeficient                       // aliquot sum < number
	ClassificationPerfect                         // aliquot sum = number
	ClassificationAbundant                        // aliquot sum > number
)

var ErrOnlyPositive = errors.New("only positive numbers")

func Classify(number int64) (Classification, error) {
	if number <= 0 {
		return ClassificationUndefined, ErrOnlyPositive
	}
	sum := int64(0)
	for i := int64(1); i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}

	switch sum := sum; {
	case sum < number:
		return ClassificationDeficient, nil
	case sum == number:
		return ClassificationPerfect, nil
	case sum > number:
		return ClassificationAbundant, nil
	}

	return ClassificationDeficient, nil
}
