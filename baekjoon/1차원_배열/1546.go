package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := 0
	fmt.Fscanln(reader, &n)

	scores := make([]float64, n)
	max := 0.0
	accu := 0.0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &scores[i])
		if scores[i] > max {
			max = scores[i]
		}
	}
	for i := 0; i < n; i++ {
		accu += scores[i] / max
	}
	fmt.Printf("%f", accu / float64(n) * 100)
}