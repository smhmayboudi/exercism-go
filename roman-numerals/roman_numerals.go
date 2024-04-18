package romannumerals

import "math"

func ToRomanNumeral(input int) (string, error) {
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
		max := int(math.Floor(float64(input / index)))
		for i := 0; i < max; i++ {
			res += value
		}
		input -= index * max
	}
	return res, nil
}
