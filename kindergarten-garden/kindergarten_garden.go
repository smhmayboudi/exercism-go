package kindergarten

import (
	"errors"
	"slices"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

type Garden struct {
	diagram  string
	children []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	child := ""
	invalid := false
	for i := 0; i < len(children); i++ {
		if child == children[i] {
			invalid = true
		} else {
			child = children[i]
		}
	}
	lenchi := len(children)
	x := []rune(diagram)
	for i := 0; i < len(x); i++ {
		if !(x[i] == '\n' || ('A' <= x[i] && x[i] <= 'Z')) {
			invalid = invalid || true
		}
	}
	invalid = invalid || len(x)%2 != 0
	invalid = invalid || (lenchi*4)+2 != len(x)
	invalid = invalid || x[0] != '\n'
	invalid = invalid || x[len(x)/2] != '\n'
	if invalid {
		return nil, errors.New("error")
	}
	children2 := make([]string, len(children))
	copy(children2, children)
	slices.Sort(children2)
	return &Garden{
		diagram:  diagram,
		children: children2,
	}, nil
}

// Alice
// Bob
// Charlie
// David
// Eve
// Fred
// Ginny
// Harriet
// Ileana
// Joseph
// Kincaid
// Larry

// Grass
// Clover
// Radishes
// Violets

func (g *Garden) Plants(child string) (out []string, ok bool) {
	cID := slices.Index(g.children, child)
	if cID < 0 {
		return
	}
	dig := []rune(g.diagram)

	lenDia := len(g.diagram)/2 - 1
	lenChi := len(g.children)
	plaPerCh := lenDia / lenChi
	start := 1 + cID*plaPerCh
	for i := start; i < plaPerCh+start; i++ {
		pln := Name(dig[i])
		out = append(out, pln)
	}
	next := start + lenDia + 1
	for i := next; i < plaPerCh+next; i++ {
		pln := Name(dig[i])
		out = append(out, pln)
	}
	ok = true
	return
}

func Name(plant rune) string {
	switch plant {
	case 'G':
		return "grass"
	case 'C':
		return "clover"
	case 'R':
		return "radishes"
	case 'V':
		return "violets"
	default:
		return ""
	}
}
