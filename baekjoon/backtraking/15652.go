package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, M int
var reader *bufio.Reader
var writer *bufio.Writer
var result, check []int

func DFS(at, depth int) {
	if depth == M {
		for _, r := range result {
			fmt.Fprintf(writer, "%d ", r)
		}
		fmt.Fprintln(writer)
		return
	}
	for i := at; i <= N; i++ {
		result[depth] = i
		DFS(i, depth+1)
	}
}

func main() {

	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	fmt.Fscanf(reader, "%d %d\n", &N, &M)
	result = make([]int, M)
	check = make([]int, N+1)
	DFS(1, 0)
}
