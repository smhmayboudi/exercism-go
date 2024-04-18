package romannumerals

import (
	"errors"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || 3999 < input {
		return "", errors.New("error")
	}
	romans := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	res := ""
	for index, value := range romans {
		for input < index {
			res += value
			input -= index
		}
	}
	return res, nil
}
