package perfect

// Define the Classification type here.
type Classification byte

const (
	ClassificationDeficient Classification = iota // aliquot sum < number
	ClassificationPerfect                         // aliquot sum = number
	ClassificationAbundant                        // aliquot sum > number
)

func Classify(number int64) (Classification, error) {
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
