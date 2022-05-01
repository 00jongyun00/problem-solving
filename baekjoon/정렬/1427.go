package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var rd *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	var n int
	var arr []int
	fmt.Fscanf(rd, "%d\n", &n)
	for ; n > 0; n /= 10 {
		tmp := n % 10
		arr = append(arr, tmp)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	for _, ele := range arr {
		fmt.Fprint(wr, ele)
	}
	wr.Flush()
}
