package railfence

import (
	"bytes"
)

// Oscillator oscillates an integer between min and max values
type Oscillator struct {
	Min, Max, Direction, Value int
}

// NewOscillator creates a new Oscillator value
func NewOscillator(min, max, value int) *Oscillator {
	return &Oscillator{min, max, 1, value}
}

// Oscillate oscillates the oscillator!
func (o *Oscillator) Oscillate() {
	if (o.Value + o.Direction) < o.Min {
		o.Direction = 1
	} else if (o.Value + o.Direction) >= o.Max {
		o.Direction = -1
	}
	o.Value += o.Direction
}

func getOffsets(message string, numRails int) []byte {

	osc := NewOscillator(0, numRails, 0)
	rails := make([][]byte, numRails)

	for i := range message {
		rails[osc.Value] = append(rails[osc.Value], byte(i))
		osc.Oscillate()
	}

	return bytes.Join(rails, []byte{})
}

// Encode a message using rail fence encryption
func Encode(message string, numRails int) string {

	var buf bytes.Buffer

	for _, offset := range getOffsets(message, numRails) {
		buf.WriteByte(message[offset])
	}

	return buf.String()
}

// Decode a message using rail fence decryption
func Decode(message string, numRails int) string {

	out := make([]byte, len(message))

	for i, offset := range getOffsets(message, numRails) {
		out[offset] = message[i]
	}

	return string(out)
}
