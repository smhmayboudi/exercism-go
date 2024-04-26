package secret

func Handshake(code uint) []string {
	str := map[uint]string{
		0b00001: "wink",
		0b00010: "double blink",
		0b00100: "close your eyes",
		0b01000: "jump",
		0b10000: "Reverse the order of the operations in the secret handshake.",
	}
	out := []string{}
	for index, value := range str {
		if code|index == code {
			out = append(out, value)
		}
	}
	return out
}
