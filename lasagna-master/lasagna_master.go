package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePreparationTimePerLayer int) int {
	if averagePreparationTimePerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * averagePreparationTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodleLayerCount := 0
	sauceLayerCount := 0
	for _, value := range layers {
		if value == "noodles" {
			noodleLayerCount++
		} else if value == "sauce" {
			sauceLayerCount++
		}
	}
	return noodleLayerCount * 50, float64(sauceLayerCount) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	// friendsList = append(friendsList[:len(friendsList)-1], myList[len(myList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, cookbookPosition int) []float64 {
	quantities2 := make([]float64, len(quantities))
	copy(quantities2, quantities)
	for i := 0; i < len(quantities2); i++ {
		quantities2[i] *= float64(cookbookPosition) / 2
	}
	return quantities2
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
