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
	defer wr.Flush()

	var A, B, V int64
	fmt.Fscanf(rd, "%d %d %d", &A, &B, &V)
	day := int64(math.Ceil(float64(V-B) / float64(A-B)))
	fmt.Fprintln(wr, day)
}
