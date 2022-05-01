package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var rd *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)

type word struct {
	str string
}

func contains(words []word, word string) bool {
	for _, w := range words {
		if w.str == word {
			return true
		}
	}
	return false
}

func main() {
	var input string
	var n int
	words := make([]word, 0)
	fmt.Fscanf(rd, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%s\n", &input)
		if contains(words, input) {
			continue
		}
		words = append(words, word{input})
	}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i].str) == len(words[j].str) {
			cnt := 0
			for {
				if words[i].str[cnt] == words[j].str[cnt] {
					cnt++
				} else {
					return words[i].str[cnt] < words[j].str[cnt]
				}
			}
		}
		return len(words[i].str) < len(words[j].str)
	})

	for _, word := range words {
		fmt.Fprintln(wr, word.str)
	}

	wr.Flush()
}
