package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	var a, b big.Int
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	fmt.Fscanf(rd, "%d %d", &a, &b)
	sum := new(big.Int)
	sum = sum.Add(&a, &b)
	fmt.Fprintf(wr, "%d\n", sum)
}
