package allergies

func Allergies(allergies uint) []string {
	out := []string{}

	eggs := uint(1)
	peanuts := uint(2)
	shellfish := uint(4)
	strawberries := uint(8)
	tomatoes := uint(16)
	chocolate := uint(32)
	pollen := uint(64)
	cats := uint(128)

	if allergies|eggs == allergies {
		out = append(out, "eggs")
	}
	if allergies|peanuts == allergies {
		out = append(out, "peanuts")
	}
	if allergies|shellfish == allergies {
		out = append(out, "shellfish")
	}
	if allergies|strawberries == allergies {
		out = append(out, "strawberries")
	}
	if allergies|tomatoes == allergies {
		out = append(out, "tomatoes")
	}
	if allergies|chocolate == allergies {
		out = append(out, "chocolate")
	}
	if allergies|pollen == allergies {
		out = append(out, "pollen")
	}
	if allergies|cats == allergies {
		out = append(out, "cats")
	}
	return out
}

func AllergicTo(allergies uint, allergen string) bool {
	a := uint(0)
	switch allergen {
	case "eggs":
		a = uint(1)
		break
	case "peanuts":
		a = uint(2)
		break
	case "shellfish":
		a = uint(4)
		break
	case "strawberries":
		a = uint(8)
		break
	case "tomatoes":
		a = uint(16)
		break
	case "chocolate":
		a = uint(32)
		break
	case "pollen":
		a = uint(64)
		break
	case "cats":
		a = uint(128)
		break

	}
	return allergies|a == allergies
}
