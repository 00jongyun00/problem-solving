package main

import (
	"bufio"
	"fmt"
	"os"
)

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}


func main() {
	croatia := []string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var word string
	fmt.Fscanln(reader, &word)

	cnt := 0
	for i := 0; i < len(word)-1; i++ {
		var ch string
		for j := 0; j < 3; j++ {
			if i + j >= len(word) {
				ch = ch + string(word[i:])
			} else {
				ch = ch + string(word[i + j])
			}
			if contains(croatia, ch) {
				cnt += j
				i += j
				break
			}
		}
	}
	fmt.Fprintln(writer, len(word) - cnt)
	writer.Flush()
}
