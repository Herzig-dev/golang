package romanconverter

import (
	"strings"
)

func RomanToInt(s string) int {
	s = strings.ToUpper(s)

	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		char := rune(s[i])
		value := romanMap[char]

		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}

	return total
}