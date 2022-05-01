package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, result int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		ch := make(map[string]int)
		var word string
		flag := true
		fmt.Fscanln(reader, &word)
		for j := 0; j < len(word); j++ {
			str := string(word[j])
			if c, ok := ch[str]; ok {
				if j - c == 1 {
					ch[string(word[j])] = j
				} else {
					flag = false
					break
				}
			} else {
				ch[string(word[j])] = j
			}
		}
		if flag {
			result++
		}
	}
	fmt.Fprintln(writer, result)
	writer.Flush()
}