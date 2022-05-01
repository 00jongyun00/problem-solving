package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	asciiFisrt := 97
	alpha := make([]int, 26)
	for i := 0; i < len(alpha); i++ {
		alpha[i] = -1
	}

	var input string
	fmt.Fscanln(reader, &input)

	for i := 0; i < len(input); i++ {
		idx := int(input[i]) - asciiFisrt
		if alpha[idx] == -1 {
			alpha[idx] = i
		}
	}

	for _, al := range alpha {
		fmt.Printf("%d ", al)
	}
}
