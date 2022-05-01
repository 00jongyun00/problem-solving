package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m, n, sum int
	min := 10001
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	fmt.Fscanf(rd, "%d\n%d\n", &m, &n)
	if m == 1 {
		m++
	}
	for ; m <= n; m++ {
		flag := true
		for j := 2; j <= m/2; j++ {
			if m%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			sum += m
			if min > m {
				min = m
			}
		}
	}
	if sum != 0 {
		fmt.Fprintf(wr, "%d\n%d", sum, min)
	} else {
		fmt.Fprint(wr, -1)
	}
}
