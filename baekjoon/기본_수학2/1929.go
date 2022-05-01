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

	var m, n, i int
	fmt.Fscanf(rd, "%d %d\n", &m, &n)

	for ; m <= n; m++ {
		if m == 1 {
			continue
		} else if m == 2 {
			fmt.Fprintf(wr, "%d\n", m)
			continue
		}
		flag := true
		for i = 2; i*i <= m; i++ {
			if m%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Fprintf(wr, "%d\n", m)
		}
	}
}
