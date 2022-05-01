package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var x, y, w, h int
	fmt.Fscanf(rd, "%d %d %d %d\n", &x, &y, &w, &h)
	fmt.Fprint(wr, min(min(w-x, x), min(h-y, y)))
	wr.Flush()
}
