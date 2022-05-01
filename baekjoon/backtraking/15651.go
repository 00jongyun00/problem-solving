package main

import (
	"bufio"
	"fmt"
	"os"
)

var rd *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, m int
var array [7]int

func solution(idx int) {
	if idx == m {
		for i := 0; i < m; i++ {
			fmt.Fprintf(wr, "%d ", array[i])
		}
		fmt.Fprintln(wr)
		return
	}
	for i := 1; i <= n; i++ {
		array[idx] = i
		solution(idx + 1)
	}
}

func main() {
	fmt.Fscanf(rd, "%d %d", &n, &m)
	solution(0)
	wr.Flush()
}
