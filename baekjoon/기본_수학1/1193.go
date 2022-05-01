package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var N int
	i := 1
	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}
	for N > i {
		N -= i
		i++
	}
	if i%2 == 1 {
		fmt.Fprintf(wr, "%d/%d", i+1-N, N)
	} else {

		fmt.Fprintf(wr, "%d/%d", N, i+1-N)
	}
}
