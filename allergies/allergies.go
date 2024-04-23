package allergies

const testVersion = 1

var allergens = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func AllergicTo(score uint, a string) bool {
	allergen := allergens[a]
	return score&allergen != 0
}

func Allergies(score uint) []string {
	allergies := []string{}
	for allergen, s := range allergens {
		if score&s != score {
			allergies = append(allergies, allergen)
		}
	}
	return allergies
}
