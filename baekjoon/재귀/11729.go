package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func hanoi(n int, from int, mid int, to int, wr *bufio.Writer) {
	if n == 1 {
		fmt.Fprintf(wr, "%d %d\n", from, to)
		return
	}
	hanoi(n-1, from, to, mid, wr)
	fmt.Fprintf(wr, "%d %d\n", from, to)
	hanoi(n-1, mid, from, to, wr)
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanf(rd, "%d\n", &n)
	fmt.Fprintf(wr, "%d\n", int(math.Pow(2, float64(n)))-1)
	hanoi(n, 1, 2, 3, wr)

	wr.Flush()
}
