package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Define the Robot type here.
type Robot string

func generateRobotNames() []string {
	pos := 0
	const maxRobotNames = 26 * 26 * 10 * 10 * 10
	names := make([]string, maxRobotNames)

	for i := 'A'; i <= 'Z'; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			for k := 0; k < 1000; k++ {
				names[pos] = fmt.Sprintf("%s%s%03d", string(i), string(j), k)
				pos++
			}
		}
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
	return names
}

var (
	idx      = 0
	namePool = generateRobotNames()
)

func (r *Robot) Name() (string, error) {
	if *r != "" {
		return string(*r), nil
	}

	*r = Robot(namePool[idx])
	idx++

	return string(*r), nil
}

func (r *Robot) Reset() {
	*r = ""
}
