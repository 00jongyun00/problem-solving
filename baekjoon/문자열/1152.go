package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	line, _ := reader.ReadString('\n')
	words := strings.Split(line, " ")
	count := 0
	for i := range words {
		if words[i] != "\n" && words[i] != "" {
			count++
		}
	}
	fmt.Fprintln(writer, count)
	writer.Flush()
}