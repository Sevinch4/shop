package check

import "unicode"

func PhoneNumber(phone string) bool {
	for _, p := range phone {
		if p == '+' {
			continue
		} else if !unicode.IsNumber(p) {
			return false
		}
	}
	return true
}

func Price(price int) bool {
	if price < 0 {
		return false
	}
	return true
}
