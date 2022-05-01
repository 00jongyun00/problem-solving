package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, person [][]int) []int {
	check := make([]int, n)

	for i := 0; i < len(person); i++ {
		for j := 0; j < len(person); j++ {
			flag := 0
			for k := 0; k < 2; k++ {
				if person[i][k] < person[j][k] {
					flag++
				}
			}
			if flag == 2 {
				check[i]++
			}
		}
	}
	return check
}

func main() {
	var n, x, y int
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	fmt.Fscanf(rd, "%d\n", &n)
	person := make([][]int, n)
	for i := 0; i < n; i++ {
		person[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d %d\n", &x, &y)
		person[i] = append(person[i], x)
		person[i] = append(person[i], y)
	}

	check := solution(n, person)

	for i := 0; i < len(check); i++ {
		fmt.Fprintf(wr, "%d ", check[i]+1)
	}

	wr.Flush()
}
