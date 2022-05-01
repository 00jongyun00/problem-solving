package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str, result string
	alpha := make([]int, 26)
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(rd, &str)
	str = strings.ToUpper(str)

	for _, s := range str {
		alpha[s - 65]++
	}
	max := -1
	for i, x := range alpha {
		if x > max {
			max = x
			result = string(rune(i)+ 65)
		} else if x == max {
			result = "?"
		}
	}
	fmt.Fprintf(wr, "%s", result)
	wr.Flush()
}