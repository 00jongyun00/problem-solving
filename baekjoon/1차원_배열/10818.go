package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	n := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &nums[i])
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Printf("%d %d\n", nums[0], nums[len(nums)-1])
}
