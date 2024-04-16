package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || len(slice) <= index {
		return -1
	}
	item := slice[index]
	return item
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || len(slice) <= index {
		return append(slice, value)
	}
	slice2 := slice[:index]
	slice2 = append(slice2, value)
	slice2 = append(slice2, slice[index+1:]...)
	return slice2
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	slice2 := append([]int{}, values...)
	slice2 = append(slice2, slice...)
	return slice2
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || len(slice) <= index {
		return slice
	}
	slice2 := slice[:index]
	slice2 = append(slice2, slice[index+1:]...)
	return slice2
}
