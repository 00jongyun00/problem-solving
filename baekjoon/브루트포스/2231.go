package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int) int {
	result := 0
	flag := false
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		if sum+i == n {
			result = i
			flag = true
		}

		if i == n || flag {
			break
		}
	}
	return result
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanf(rd, "%d\n", &n)
	fmt.Fprintf(wr, "%d\n", solution(n))
	wr.Flush()
}
