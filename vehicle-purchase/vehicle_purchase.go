package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	kind2 := strings.ToLower(kind)
	return kind2 == "car" || kind2 == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return fmt.Sprintf("%s is clearly the better choice.", option1)
	}
	return fmt.Sprintf("%s is clearly the better choice.", option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch age := age; {
	case age < 3:
		return originalPrice * 0.8
	case 10 <= age:
		return originalPrice * 0.5
	default:
		return originalPrice * 0.7
	}
}
