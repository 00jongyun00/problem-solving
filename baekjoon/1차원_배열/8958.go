package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := 0
	fmt.Fscanln(reader, &n)
	quiz := make([]string, n)

	for i := 0; i < n; i++ {
		text, _ := reader.ReadString('\n')
		quiz[i] = text
	}
	for _, q := range quiz {
		score := 0
		total := 0
		for _, s := range q {
			if s == 79 {
				score = score + 1
				total += score
			} else {
				score = 0
			}
		}
		fmt.Println(total)
	}

}
