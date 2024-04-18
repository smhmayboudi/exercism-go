package romannumerals

import (
	"errors"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || 3999 < input {
		return "", errors.New("error")
	}
	var romanNum = map[int]string{
		1000: "M", 2000: "MM", 3000: "MMM",
		100: "C", 200: "CC", 300: "CCC", 400: "CD", 500: "D", 600: "DC", 700: "DCC", 800: "DCCC", 900: "CM",
		10: "X", 20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC",
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
	}
	roman := ""
	// thousands
	k := input / 1000
	roman += romanNum[k*1000]
	// hundreds
	c := (input / 100) % 10
	roman += romanNum[c*100]
	// tens
	d := (input / 10) % 10
	roman += romanNum[d*10]
	// units
	u := (input / 1) % 10
	roman += romanNum[u*1]
	return roman, nil
}
