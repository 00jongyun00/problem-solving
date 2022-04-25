package easy

import "fmt"

func isPalindrome(x int) bool {
	sum := 0
	compare := x

	for x > 0 {
		r := x % 10
		sum = (sum * 10) + r
		x /= 10
	}
	return sum == compare
}

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}
