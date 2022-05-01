package main

import (
	"bufio"
	"fmt"
	"os"
)

func recursion(i, j, n int, s *string) {
	if (i/n)%3 == 1 && (j/n)%3 == 1 {
		*s = " "
	} else if n/3 == 0 {
		*s = "*"
	} else {
		recursion(i, j, n/3, s)
	}
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var s string
	defer wr.Flush()
	var n int
	fmt.Fscanf(rd, "%d\n", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			recursion(i, j, n, &s)
			wr.WriteString(s)
		}
		wr.WriteString("\n")
	}
}
