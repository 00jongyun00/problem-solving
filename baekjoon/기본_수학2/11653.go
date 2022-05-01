package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	fmt.Fscanf(rd, "%d\n", &n)
	i := 2
	for n >= i {
		if n%i == 0 {
			n /= i
			fmt.Fprintf(wr, "%d\n", i)
		} else {
			i++
		}
	}
}
