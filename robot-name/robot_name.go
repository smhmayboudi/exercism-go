package robotname

import (
	"math/rand"
	"strings"
	"time"
)

// Define the Robot type here.
type Robot string

func (r *Robot) Name() (string, error) {
	if *r == "" {
		var sb strings.Builder
		sb.Grow(5)
		var run rune
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < 2; i++ {
			run = 'A' + rune(rng.Intn(26))
			sb.WriteRune(run)
		}
		for i := 0; i < 3; i++ {
			run = '0' + rune(rng.Intn(9))
			sb.WriteRune(run)
		}
		*r = Robot(sb.String())
	}
	return string(*r), nil
}

func (r *Robot) Reset() {
	*r = ""
	r.Name()
}
