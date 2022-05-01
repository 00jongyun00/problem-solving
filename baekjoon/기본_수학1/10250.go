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

	var r int
	fmt.Fscanln(rd, &r)
	for i := 0; i < r; i++ {
		var h, w, n int
		fmt.Fscanf(rd, "%d %d %d\n", &h, &w, &n)
		if n%h == 0 {
			fmt.Fprintf(wr, "%d%02d\n", h, n/h)
		} else {
			fmt.Fprintf(wr, "%d%02d\n", n%h, n/h+1)
		}
	}
}
