package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	len1, len2 := len(l1), len(l2)
	switch {
	case len1 < len2:
		return testSubList(l1, l2)
	case len1 > len2:
		return testSuperlist(l1, l2)
	default:
		return testEqual(l1, l2)
	}
}
func testEqual(l1, l2 []int) Relation {
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return RelationUnequal
		}
	}
	return RelationEqual
}
func testSubList(l1, l2 []int) Relation {
	len1, len2 := len(l1), len(l2)
	for i := 0; i+len1 <= len2; i++ {
		if testEqual(l1, l2[i:i+len1]) == RelationEqual {
			return RelationSublist
		}
	}
	return RelationUnequal
}
func testSuperlist(l1, l2 []int) Relation {
	if testSubList(l2, l1) == RelationSublist {
		return RelationSuperlist
	}
	return RelationUnequal
}
