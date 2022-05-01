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

	var n, cnt int
	if sc.Scan() {
		n, _ = strconv.Atoi(sc.Text())
	}

	for n > 0 {
		if n%5 == 0 {
			cnt += n / 5
			n = 0
			break
		}
		n -= 3
		cnt++
	}

	if n != 0 {
		fmt.Fprint(wr, "-1\n")
	} else {
		fmt.Fprintf(wr, "%d\n", cnt)
	}
}
