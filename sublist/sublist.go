package sublist

import "reflect"

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if reflect.DeepEqual(l1, l2) {
		return RelationEqual
	}
	if contains(l2, l1) {
		return RelationSublist
	}
	if contains(l1, l2) {
		return RelationSuperlist
	}
	return RelationUnequal
}

func contains(s1, s2 []int) bool {
	i, j := 0, 0
	for ; i < len(s1) && j < len(s2); i++ {
		if s1[i] == s2[j] {
			j++
		} else {
			i, j = i-j, 0
		}
	}
	return j == len(s2)
}
