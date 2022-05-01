package main

import (
	"bufio"
	"fmt"
	"os"
)

func blackjack(arr []int, m int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				sum := arr[i] + arr[j] + arr[k]
				if sum > result && sum <= m {
					result = sum
				}
			}
		}
	}
	return result
}

func main() {
	var n, m int
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	fmt.Fscanf(rd, "%d %d\n", &n, &m)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d ", &arr[i])
	}
	fmt.Fprintf(wr, "%d\n", blackjack(arr, m))
	wr.Flush()
}
