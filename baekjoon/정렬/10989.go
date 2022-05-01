package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rd *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	var n, num int
	fmt.Fscanf(rd, "%d\n", &n)
	count := make([]int, 10001)

	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d\n", &num)
		count[num]++
	}

	for idx, num := range count {
		s := strconv.Itoa(idx) + "\n"
		if num > 0 {
			fmt.Fprint(wr, strings.Repeat(s, num))
		}
	}

	wr.Flush()
}
