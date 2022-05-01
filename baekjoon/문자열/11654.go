package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var ch string
	fmt.Fscan(reader, &ch)
	fmt.Println(ch[0])
}