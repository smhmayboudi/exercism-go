package clock

import (
	"fmt"
	"math"
)

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	mTota := h*60 + m
	for mTota < 0 {
		mTota += 24 * 60
	}
	h2 := int(math.Floor(float64(mTota/60))) % 24
	m2 := mTota % 60
	return Clock{
		h: h2,
		m: m2,
	}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
