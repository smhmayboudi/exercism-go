package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	out := []Triplet{}
	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			for k := j; k <= max; k++ {
				if i*i+j*j == k*k {
					out = append(out, Triplet{i, j, k})
				}
			}
		}
	}
	return out
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	res := []Triplet{}
	for _, t := range Range(1, p/2) {
		if t[0]+t[1]+t[2] == p {
			res = append(res, t)
		}
	}
	return res
}
