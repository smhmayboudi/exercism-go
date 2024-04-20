package cipher

import (
	"strings"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

func shiftRune(r rune, distance int) rune {
	if distance < 0 {
		distance = 26 + distance
	}
	idx := (int(r-'a') + distance) % 26
	return rune('a' + idx)
}

func NewCaesar() Cipher {
	return &shift{
		distance: 3,
	}
}

type shift struct {
	distance int
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return &shift{
		distance: distance,
	}
}

func (c shift) Encode(input string) string {
	var sb strings.Builder
	sb.Grow(len(input))
	in := []rune(strings.ToLower(input))
	for i := 0; i < len(in); i++ {
		r := in[i]
		if 'a' <= r && r <= 'z' {
			sb.WriteRune(shiftRune(r, c.distance))
		}
	}
	return sb.String()
}

func (c shift) Decode(input string) string {
	var sb strings.Builder
	sb.Grow(len(input))
	in := []rune(strings.ToLower(input))
	for i := 0; i < len(in); i++ {
		r := in[i]
		if 'a' <= r && r <= 'z' {
			sb.WriteRune(shiftRune(r, 26-c.distance))
		}
	}
	return sb.String()
}

type vigenere struct {
	distances []int
}

func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}
	var valid bool = false
	distances := make([]int, 0, len(key))

	in := []rune(key)
	for i := 0; i < len(in); i++ {
		r := in[i]
		if r < 'a' || 'z' < r {
			valid = false
			break
		}
		if r != 'a' {
			valid = true
		}
		distances = append(distances, int(r-'a'))
	}
	if !valid {
		return nil
	}
	return &vigenere{
		distances: distances,
	}
}

func (v vigenere) Encode(input string) string {
	var sb strings.Builder
	sb.Grow(len(input))
	idx := 0
	in := []rune(strings.ToLower(input))
	for i := 0; i < len(in); i++ {
		r := in[i]
		if 'a' <= r && r <= 'z' {
			sb.WriteRune(shiftRune(r, v.distances[idx%len(v.distances)]))
			idx++
		}
	}
	return sb.String()
}

func (v vigenere) Decode(input string) string {
	var sb strings.Builder
	sb.Grow(len(input))
	idx := 0
	in := []rune(strings.ToLower(input))
	for i := 0; i < len(in); i++ {
		r := in[i]
		if 'a' <= r && r <= 'z' {
			sb.WriteRune(shiftRune(r, -v.distances[idx%len(v.distances)]))
			idx++
		}
	}
	return sb.String()
}
