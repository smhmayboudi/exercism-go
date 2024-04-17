package space

import "strings"

type Planet string

func Age(seconds float64, planet Planet) float64 {
	years := seconds / 31557600.00
	switch strings.ToLower(string(planet)) {
	case "mercury":
		return years / 0.2408467
	case "venus":
		return years / 0.61519726
	case "earth":
		return years / 1.0
	case "mars":
		return years / 1.8808158
	case "jupiter":
		return years / 11.862615
	case "saturn":
		return years / 29.447498
	case "uranus":
		return years / 84.016846
	case "neptune":
		return years / 164.79132
	default:
		return -1
	}
}
