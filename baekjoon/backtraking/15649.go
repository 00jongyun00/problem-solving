package main

import (
	"bufio"
	"fmt"
	"os"
)

var rd *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)
var n, m int
var visited []bool
var arr []int

func solution(cnt int) {
	if cnt == m {
		for i := 0; i < m; i++ {
			fmt.Fprintf(wr, "%d ", arr[i])
		}
		return
	}
	for i := 1; i <= n; i++ {
		if !visited[i] {
			visited[i] = true
			arr[cnt] = i
			solution(cnt + 1)
			visited[i] = false
		}
	}
}

func main() {
	fmt.Fscanf(rd, "%d %d\n", &n, &m)
	visited = make([]bool, 9)
	arr = make([]int, 9)
	solution(0)
	wr.Flush()
}
