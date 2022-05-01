package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var n, x, r, i, j int
	fmt.Fscanf(rd, "%d\n", &n)
	for i = 0; i < n; i++ {
		fmt.Fscanf(rd, "%d ", &x)
		for j = 2; j < x; j++ {
			if x%j == 0 {
				break
			}
		}
		if j == x {
			r++
		}
	}
	fmt.Fprintf(wr, "%d", r)
}
