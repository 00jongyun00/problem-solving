package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := 9
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &nums[i])
	}

	max := -1
	idx := -1
	for i, num := range nums {
		if num > max {
			max = num
			idx = i + 1
		}
	}
	fmt.Printf("%d\n%d", max, idx)
}