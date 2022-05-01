package main

import (
	"bufio"
	"fmt"
	"os"
)

func fact(n int) int {
	r := 1
	for i := 2; i <= n; i++ {
		r *= i
	}
	return r
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanf(rd, "%d\n", &n)
	fmt.Fprintf(wr, "%d", fact(n))
	wr.Flush()
}
