package binarysearch

// SearchInts performs binary search returning index of the match.
// -1 if the key is not found or array is 0 length.
func SearchInts(list []int, key int) int {
	l := 0
	r := len(list) - 1
	for l <= r {
		c := (l + r) / 2
		if list[c] == key {
			return c
		} else if list[c] < key {
			l = c + 1
		} else {
			r = c - 1
		}
	}
	return -1
}
