package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func solution(n int, m int, chessBoard [][]string) int {
	var min = n * m
	for i := 0; i < n-7; i++ {
		for j := 0; j < m-7; j++ {
			var cntChange1 = float64(0)
			var cntChange2 = float64(0)
			for k := i; k < i+8; k++ {
				for l := j; l < j+8; l++ {
					if (k+l)%2 == 0 {
						if string(chessBoard[k][l]) == "B" {
							cntChange1++
						} else {
							cntChange2++
						}
					} else {
						if string(chessBoard[k][l]) == "B" {
							cntChange2++
						} else {
							cntChange1++
						}
					}
				}
			}
			if min > int(math.Min(cntChange1, cntChange2)) {
				min = int(math.Min(cntChange1, cntChange2))
			}
		}
	}
	return min
}

func main() {
	var n, m int
	var line string
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	fmt.Fscanf(rd, "%d %d\n", &n, &m)
	chessBoard := make([][]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%s\n", &line)
		chessBoard[i] = strings.Split(line, "")
	}
	fmt.Fprint(wr, solution(n, m, chessBoard))

	wr.Flush()
}
