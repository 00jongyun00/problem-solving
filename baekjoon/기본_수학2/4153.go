package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var a, b, c int
	for {
		fmt.Fscanf(rd, "%d %d %d\n", &a, &b, &c)
		if a == 0 && b == 0 && c == 0 {
			break
		}
		n1 := a * a
		n2 := b * b
		n3 := c * c
		if n1+n2 == n3 || n2+n3 == n1 || n1+n3 == n2 {
			fmt.Fprint(wr, "right\n")
		} else {
			fmt.Fprint(wr, "wrong\n")
		}
	}
	wr.Flush()
}
