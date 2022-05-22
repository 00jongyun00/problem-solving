package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var N, result int
var board []int

func DFS(depth int) {
	if depth == N {
		result++
		return
	}

	for i := 0; i < N; i++ {
		board[depth] = i

		if possible(depth) {
			DFS(depth + 1)
		}
	}
}

func possible(col int) bool {
	for i := 0; i < col; i++ {
		if board[col] == board[i] {
			return false
		}

		if math.Abs(float64(col-i)) == math.Abs(float64(board[col]-board[i])) {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fscan(reader, &N)
	board = make([]int, N)
	DFS(0)
	fmt.Fprintln(writer, result)
}
