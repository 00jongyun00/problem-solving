package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	mansion := make([][]int, 15)
	for i := 0; i < 15; i++ {
		mansion[i] = make([]int, 15)
		mansion[i][0] = 1
	}
	defer wr.Flush()

	for i := 0; i < 15; i++ {
		for j := 1; j < 15; j++ {
			if i == 0 {
				mansion[i][j] = j + 1
			} else {
				mansion[i][j] = mansion[i-1][j] + mansion[i][j-1]
			}
		}
	}
	var t, k, n int
	fmt.Fscanln(rd, &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(rd, "%d\n%d\n", &k, &n)
		fmt.Fprintln(wr, mansion[k][n-1])
	}
}
