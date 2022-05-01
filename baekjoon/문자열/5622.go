package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dial := []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 10, 10, 10, 10}
	var result int
	var input string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &input)
	for i := 0; i < len(input); i++ {
		result += dial[int(input[i])-65]
	}

	fmt.Fprintln(writer, result)
	writer.Flush()
}
