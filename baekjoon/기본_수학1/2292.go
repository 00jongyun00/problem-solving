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

	var N, i int
	s := 1
	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}
	for i = 1; i <= N; i++ {
		s += (i - 1) * 6
		if N <= s {
			break
		}
	}
	fmt.Fprintln(wr, i)
}
