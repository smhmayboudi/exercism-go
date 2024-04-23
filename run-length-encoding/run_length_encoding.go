package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	runes := []rune(input)
	var sb strings.Builder
	sb.Grow(len(runes))
	idC := 1
	for i := 0; i < len(runes); i++ {
		if i+1 != len(runes) && runes[i] == runes[i+1] {
			idC++
			continue
		}
		if 1 < idC {
			sb.WriteString(strconv.Itoa(idC))
		}
		sb.WriteRune(runes[i])
		idC = 1
	}
	return sb.String()
}

func RunLengthDecode(input string) string {
	runes := []rune(input)
	var sb strings.Builder
	sb.Grow(len(runes))
	var sbN strings.Builder
	sb.Grow(len(runes))
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' || ('a' <= runes[i] && runes[i] <= 'z') || ('A' <= runes[i] && runes[i] <= 'Z') {
			idC, _ := strconv.Atoi(sbN.String())
			sbN.Reset()
			if idC == 0 {
				idC = 1
			}
			for c := 0; c < idC; c++ {
				sb.WriteRune(runes[i])
			}
			continue
		}
		sbN.WriteRune(runes[i])
	}
	return sb.String()
}
