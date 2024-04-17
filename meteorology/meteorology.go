package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (temperatureUnit TemperatureUnit) String() string {
	if temperatureUnit == Celsius {
		return "°C"
	}
	return "°F"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temperature Temperature) String() string {
	return fmt.Sprintf("%d %s", temperature.degree, temperature.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speedUnit SpeedUnit) String() string {
	if speedUnit == KmPerHour {
		return "km/h"
	}
	return "mph"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed Speed) String() string {
	return fmt.Sprintf("%d %s", speed.magnitude, speed.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (meteorologyData MeteorologyData) String() string {
	// San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
	return fmt.Sprintf(
		"%s: %s, Wind %s at %s, %d%% Humidity",
		meteorologyData.location,
		meteorologyData.temperature.String(),
		meteorologyData.windDirection,
		meteorologyData.windSpeed.String(),
		meteorologyData.humidity,
	)
}
