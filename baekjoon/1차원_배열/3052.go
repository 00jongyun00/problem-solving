package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	x := 42
	reader := bufio.NewReader(os.Stdin)
	result := make(map[int]bool)

	for i := 0; i < 10; i++ {
		input := 0
		fmt.Fscan(reader, &input)
		result[input % x] = true
	}
	fmt.Println(len(result))
}
