package main

import (
	"bufio"
	"fmt"
	"os"
)

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
		}
	}
	return arr
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanf(rd, "%d\n", &n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d\n", &arr[i])
	}

	sorted := sort(arr)
	for _, ele := range sorted {
		fmt.Fprintf(wr, "%d\n", ele)
	}

	wr.Flush()
}
