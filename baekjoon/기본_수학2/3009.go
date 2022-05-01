package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var x, y int
	left := make(map[int]int)
	right := make(map[int]int)
	square := make([][]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscanf(rd, "%d %d\n", &x, &y)
		square[i] = make([]int, 2)
		square[i][0] = x
		square[i][1] = y
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if j == 0 {
				if _, ok := left[square[i][j]]; ok {
					left[square[i][j]]++
				} else {
					left[square[i][j]] = 1
				}
			} else {
				if _, ok := right[square[i][j]]; ok {
					right[square[i][j]]++
				} else {
					right[square[i][j]] = 1
				}
			}
		}
	}
	for key, val := range left {
		if val == 1 {
			fmt.Fprintf(wr, "%d ", key)
		}
	}
	for key, val := range right {
		if val == 1 {
			fmt.Fprintf(wr, "%d\n", key)
		}
	}

	wr.Flush()
}
