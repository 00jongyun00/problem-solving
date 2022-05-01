package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, r int
	var s string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &t)
	defer writer.Flush()

	for i := 0; i < t; i++ {
		fmt.Fscanf(reader, "%d %s\n", &r, &s)
		for _, char := range s {
			for j := 0; j < r; j++ {
				fmt.Fprintf(writer, "%s", string(char))
			}
		}
		fmt.Fprintln(writer)
	}
}
