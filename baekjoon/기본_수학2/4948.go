package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMax(arr []int) int {
	max := -1
	for i, _ := range arr {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

func initArr(arr []bool) []bool {
	for i := 0; i < len(arr); i++ {
		if i < 2 {
			arr[i] = false
		} else {
			arr[i] = true
		}
	}
	return arr
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var n int
	inputs := make([]int, 0)
	for {
		fmt.Fscanf(rd, "%d\n", &n)
		if n == 0 {
			break
		}
		inputs = append(inputs, n)
	}
	max := findMax(inputs)
	arr := make([]bool, max*2+1)
	arr = initArr(arr)

	for i := 2; i < len(arr); i++ {
		if arr[i] {
			for j := i * 2; j < len(arr); j += i {
				arr[j] = false
			}
		}
	}
	for _, x := range inputs {
		cnt := 0
		for i := x + 1; i <= x*2; i++ {
			if arr[i] {
				cnt++
			}
		}
		fmt.Fprintf(wr, "%d\n", cnt)
	}
	wr.Flush()
}
