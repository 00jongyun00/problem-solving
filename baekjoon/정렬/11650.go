package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var rd *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)

type coordinate struct {
	x int
	y int
}

func main() {
	var n, x, y int
	fmt.Fscanf(rd, "%d\n", &n)
	coors := make([]coordinate, 0)
	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d %d\n", &x, &y)
		coors = append(coors, coordinate{x, y})
	}
	sort.Slice(coors, func(i, j int) bool {
		if coors[i].x == coors[j].x {
			return coors[i].y < coors[j].y
		}
		return coors[i].x < coors[j].x
	})
	for _, coor := range coors {
		fmt.Fprintf(wr, "%d %d\n", coor.x, coor.y)
	}
	wr.Flush()
}
