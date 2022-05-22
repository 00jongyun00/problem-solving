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
	// 모든 행에 Queen 이 놓였을때 result 를 +1 하고 return
	if depth == N {
		result++
		return
	}

	for i := 0; i < N; i++ {
		board[depth] = i

		// depth (행) i (열) 의 현재 행렬 좌표에 Queen 을 놓을 수 있는지 확인 한 다음 가능 하면 재귀호출
		if possible(depth) {
			DFS(depth + 1)
		}
	}
}

func possible(col int) bool {
	for i := 0; i < col; i++ {
		// 현재 재귀 호출된 depth (행) 의 열 과 depth 전 까지의 i (행) 의 열 이 같은 선상에 있으면 false
		if board[col] == board[i] {
			return false
		}

		// 행 끼리 의 차 와 열 끼리 의 차가 동일 하면 false
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
