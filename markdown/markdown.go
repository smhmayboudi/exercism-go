package markdown

// implementation to refactor
import (
	"fmt"
	"regexp"
	"strings"
)

var headingRegexp = regexp.MustCompile(`^(#{1,6})\s(.+)$`)
var listRegexp = regexp.MustCompile(`^(\*)\s(.+)$`)

// Render translates markdown to HTML
func Render(markdown string) string {
	html := ""
	list := false
	for _, line := range strings.Split(markdown, "\n") {
		line = replaceInlineElements(line)
		if g := listRegexp.FindStringSubmatch(line); g != nil {
			if !list {
				html += "<ul>"
				list = true
			}
			html += "<li>" + g[2] + "</li>"
			continue
		}
		if list {
			html += "</ul>"
			list = false
		}
		if g := headingRegexp.FindStringSubmatch(line); g != nil {
			tag := fmt.Sprintf("h%d", len(g[1]))
			html += "<" + tag + ">" + g[2] + "</" + tag + ">"
			continue
		}
		html += "<p>" + line + "</p>"
	}
	if list {
		html += "</ul>"
	}
	return html
}

var strongRegexp = regexp.MustCompile("__(.*?)__")
var emRegexp = regexp.MustCompile("_(.*?)_")

func replaceInlineElements(markdown string) string {
	markdown = strongRegexp.ReplaceAllString(markdown, "<strong>$1</strong>")
	markdown = emRegexp.ReplaceAllString(markdown, "<em>$1</em>")
	return markdown
}
