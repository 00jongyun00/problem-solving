package easy

import (
	"fmt"
)

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	num, lastInt, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastInt {
			total -= num
		} else {
			total += num
		}
		lastInt = num
	}
	return total
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
