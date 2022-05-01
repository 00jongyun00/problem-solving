package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wr = bufio.NewWriter(os.Stdout)
)

func solve2(n int) {
	// fast
	i := 666
	for true {
		cur := i
		for cur >= 666 {
			if cur%1000 == 666 {
				n--
				break
			}
			cur /= 10
		}
		if n <= 0 {
			fmt.Fprintln(wr, i)
			break
		}
		i++
	}
}

func main() {
	defer wr.Flush()

	var n int
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}
	solve2(n)
}
