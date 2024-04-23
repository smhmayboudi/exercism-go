package piglatin

import (
	"regexp"
	"strings"
)

func Sentence(sentence string) string {
	// A, E, I, O, and U
	// Rule 1: xray -> xray+ay | yttria -> yttria+ay | xrita -> xrita+ay
	re := regexp.MustCompile(`^([aeiou]|xr|yt)`)
	if matched := re.MatchString(sentence); matched {
		return sentence + "ay"
	}
	// B, C, D, F, G, H, J, K, L, M, N, P, Q, R, S, T, V, W, X, Y, and Z
	// Rule 2: chair -> air+ch+ay
	re = regexp.MustCompile(`^[bcdfghjklmnpqrstvwxyz]+`)
	if matched := re.MatchString(sentence); matched {
		submatch := re.FindStringSubmatch(sentence)
		lastSubmatch := submatch[len(submatch)-1]
		sentence = strings.Replace(sentence, lastSubmatch, "", 1)
		return sentence + lastSubmatch + "ay"
	}
	// Rule 3: square -> are+squ+ay
	re = regexp.MustCompile(`^[bcdfghjklmnpqrstvwxyz](qu)?`)
	if matched := re.MatchString(sentence); matched {
		submatch := re.FindStringSubmatch(sentence)
		lastSubmatch := submatch[len(submatch)-1]
		sentence = strings.Replace(sentence, lastSubmatch, "", 1)
		return sentence + lastSubmatch + "ay"
	}
	// Rule 4: rhythm -> ythm+rh+ay | my -> y+m+ay
	re = regexp.MustCompile(`^([bcdfghjklmnpqrstvwxyz]+|[aeiou])y`)
	if matched := re.MatchString(sentence); matched {
		submatch := re.FindStringSubmatch(sentence)
		lastSubmatch := submatch[len(submatch)-1]
		sentence = strings.Replace(sentence, lastSubmatch, "", 1)
		return sentence + lastSubmatch + "ay"
	}
	return sentence
}
