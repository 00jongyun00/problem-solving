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

	for i := 0; i < n; i++ {
		var cases []int
		c := 0
		sum := 0

		fmt.Fscan(reader, &c)
		for j := 0; j < c; j++ {
			input := 0
			fmt.Fscan(reader, &input)
			cases = append(cases, input)
			sum += input
		}
		average := sum / c
		count := 0
		for k := 0; k < c; k++ {
			if cases[k] > average {
				count++
			}
		}
		fmt.Printf("%0.3f%%\n", float64(count) / float64(c) * 100.0)
	}
}
