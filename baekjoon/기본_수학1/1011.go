package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var t, x, y int
	fmt.Fscanf(rd, "%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(rd, "%d %d\n", &x, &y)
		distance := y - x
		move := 1
		cnt := 0
		for ;distance / 2 >= move;{
			distance -= move * 2
			cnt += 2
			move++
		}
		if 1 <= distance && distance <= move {
			cnt++
		} else if move < distance {
			cnt += 2
		}
		fmt.Fprintf(wr, "%d\n", cnt)
	}
}