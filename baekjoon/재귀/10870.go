package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibonacci(n int) int {
	nums := make([]int, 21)
	for i := 0; i < len(nums); i++ {
		nums[i] = 0
	}

	for i := 1; i <= n; i++ {
		if i <= 2 {
			nums[i] = 1
		} else {
			nums[i] = nums[i-1] + nums[i-2]
		}
	}
	return nums[n]
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanf(rd, "%d\n", &n)
	fmt.Fprintf(wr, "%d", fibonacci(n))

	wr.Flush()
}
