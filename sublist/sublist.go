package sublist

import "slices"

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	ll1 := len(l1)
	ll2 := len(l2)
	if len(l1) == 0 && len(l2) == 0 {
		return RelationEqual
	}
	slices.Sort(l1)
	slices.Sort(l2)
	if ll2 < ll1 {
		check := true
		for i := 0; i < ll2; i++ {
			check = check && (l1[i] == l2[i])
		}
		if !check {
			return RelationUnequal
		}
		return RelationSuperlist
	} else if ll1 < ll2 {
		check := true
		for i := 0; i < ll1; i++ {
			check = check && (l1[i] == l2[i])
		}
		if !check {
			return RelationUnequal
		}
		return RelationSublist
	} else {
		check := true
		for i := 0; i < ll1; i++ {
			check = check && (l1[i] == l2[i])
		}
		if !check {
			return RelationUnequal
		}
		return RelationEqual
	}
}
