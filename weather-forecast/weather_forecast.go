/*
Package weather provides tools to convert
temperatures to and from Kelvin.
*/
package weather

// CurrentCondition represents a certain temperature in degrees Celsius.
var CurrentCondition string

// CurrentLocation represents a certain temperature in degrees Celsius.
var CurrentLocation string

// Forecast returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
