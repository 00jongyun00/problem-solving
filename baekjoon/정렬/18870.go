package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var rd *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)

func makeSliceUnique(xs []int) []int {
	keys := make(map[int]struct{})
	res := make([]int, 0)
	for _, ele := range xs {
		if _, ok := keys[ele]; ok {
			continue
		} else {
			keys[ele] = struct{}{}
			res = append(res, ele)
		}
	}
	return res
}

func main() {
	var n, x int
	xs := make([]int, 0)
	d := make(map[int]int)
	fmt.Fscanf(rd, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d ", &x)
		xs = append(xs, x)
	}
	set := makeSliceUnique(xs)
	sort.Ints(set)

	for i := range set {
		d[set[i]] = i
	}
	for _, x := range xs {
		fmt.Fprintf(wr, "%d ", d[x])
	}

	wr.Flush()
}
