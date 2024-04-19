package protein

import (
	"errors"
)

func FromRNA(rna string) ([]string, error) {
	s := []rune(rna)
	le := len(s)
	res := []string{}
	for i := 0; i < le; i = i + 3 {
		str, err := FromCodon(rna[i : i+3])
		if err != nil {
			if err == ErrStop {
				break
			}
			return nil, err
		}
		res = append(res, str)
	}
	return res, nil
}

var ErrInvalidBase = errors.New("invalid base")
var ErrStop = errors.New("stop")

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
