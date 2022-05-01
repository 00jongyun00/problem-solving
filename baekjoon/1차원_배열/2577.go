package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	count := make([]int, 10)
	multiAllNum := 1

	for i := 0; i < 3; i++ {
		input := 0
		fmt.Fscanln(reader, &input)
		multiAllNum *= input
	}
	for {
		if multiAllNum == 0 {
			break
		}
		remainer := multiAllNum % 10
		multiAllNum /= 10
		count[remainer]++
	}
	for _, c := range count {
		fmt.Println(c)
	}
}