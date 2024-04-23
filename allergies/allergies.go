package allergies

var allergens = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

func Allergies(score uint) (allergies []string) {
	for i := uint(1); i <= 128; i *= 2 {
		if score&i == i {
			allergies = append(allergies, allergens[i])
		}
	}
	return
}

func AllergicTo(score uint, substance string) bool {
	allergens := Allergies(score)
	for _, al := range allergens {
		if al == substance {
			return true
		}
	}
	return false
}
