package stringset

import (
	"reflect"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set []string

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	return Set(l)
}

func (s Set) String() string {
	return strings.Join(s, "")
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	return strings.Contains(s.String(), elem)
}

func (s Set) Add(elem string) {
	s = append(s, elem)
}

func Subset(s1, s2 Set) bool {
	panic("Please implement the Subset function")
}

func Disjoint(s1, s2 Set) bool {
	panic("Please implement the Disjoint function")
}

func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	panic("Please implement the Intersection function")
}

func Difference(s1, s2 Set) Set {
	panic("Please implement the Difference function")
}

func Union(s1, s2 Set) Set {
	panic("Please implement the Union function")
}
