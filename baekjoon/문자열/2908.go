package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var a, b string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanf(reader, "%s %s", &a, &b)
	a = reverse(a)
	b = reverse(b)

	inta, _ := strconv.ParseInt(a, 10, 64)
	intb, _ := strconv.ParseInt(b, 10, 64)
	if inta > intb {
		fmt.Fprintln(writer, inta)
	} else {
		fmt.Fprintln(writer, intb)
	}
	writer.Flush()
}