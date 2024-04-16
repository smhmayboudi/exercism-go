package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	timeParse, err := time.Parse("1/02/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return timeParse
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	timeParse, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return timeParse.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	timeParse, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return 12 <= timeParse.Hour() && timeParse.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	timeParse, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("You have an appointment on %s", timeParse.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	timeParse, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return timeParse
}
