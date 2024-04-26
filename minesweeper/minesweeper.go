package minesweeper

import (
	"fmt"
	"strings"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {
	lb := len(board)
	out := make([]string, lb)
	for r := 0; r < lb; r++ {
		var sb strings.Builder
		lbr := len(board[r])
		for c := 0; c < lbr; c++ {
			s := board[r][c]
			if s == ' ' {
				idx := 0
				for r2 := r - 1; r2 <= r+1; r2++ {
					for c2 := c - 1; c2 <= c2+1; c2++ {
						top := 0 <= r2-1
						left := 0 <= c2-1
						right := c2 <= lbr
						botton := r2 <= lb
						if top && left && right && botton {
							s2 := board[r2][c2]
							if s2 == '*' {
								idx++
							}
						}
					}
				}
				if idx != 0 {
					sb.WriteString(fmt.Sprintf("%d", idx))
				} else {
					sb.WriteString(" ")
				}
			}
		}
		out[r] = sb.String()
	}
	return out
}
