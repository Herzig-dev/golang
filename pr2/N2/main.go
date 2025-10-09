package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func romanToInt(s string) int {
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите римское число: ")
	scanner.Scan()
	input := scanner.Text()

	input = strings.TrimSpace(input)

	result := romanToInt(input)
	fmt.Printf("Арабское число: %d\n", result)
}