package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var n float64
	fmt.Fscan(rd, &n)
	fmt.Fprintf(wr, "%.6f\n", n*n*math.Pi)
	fmt.Fprintf(wr, "%.6f\n", n*n*2)
	wr.Flush()
}
