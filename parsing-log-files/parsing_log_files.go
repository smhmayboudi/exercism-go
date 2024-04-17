package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<\W*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, value := range lines {
		re := regexp.MustCompile(`".*(?i)password.*"`)
		if re.MatchString(value) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, ``)
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +(\w+)`)
	for i := 0; i < len(lines); i++ {
		if re.MatchString(lines[i]) {
			matches := re.FindStringSubmatch(lines[i])
			lines[i] = fmt.Sprintf("[USR] %s %s", matches[1], lines[i])
		}
	}
	return lines
}
