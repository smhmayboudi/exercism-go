package luhn

func Valid(id string) bool {
	var sum, digit int
	for i := len(id) - 1; 0 <= i; i-- {
		c := id[i]
		switch {
		case c == ' ':
			continue
		case c >= '0' && c <= '9':
			num := int(c - '0')
			if digit%2 == 1 {
				num <<= 1
			}
			if num > 9 {
				num -= 9
			}
			sum += num
			digit++
		default:
			return false
		}
	}
	return digit > 1 && sum%10 == 0
}
