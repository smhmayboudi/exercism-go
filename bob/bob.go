package bob

import (
	"strings"
)

// "Sure." This is his response if you ask him a question, such as "How are you?" The convention used for questions is that it ends with a question mark.
// "Whoa, chill out!" This is his answer if you YELL AT HIM. The convention used for yelling is ALL CAPITAL LETTERS.
// "Calm down, I know what I'm doing!" This is what he says if you yell a question at him.
// "Fine. Be that way!" This is how he responds to silence. The convention used for silence is nothing, or various combinations of whitespace characterstrings.
// "Whatever." This is what he answers to anything else.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark {
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(remark, "?") {
		return "Sure."
	}
	if remark == "" {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
