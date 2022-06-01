package dfs

import (
	"fmt"
	"math"
	"testing"
)

var C, N int
var array []int
var result int

func badook(level, sum int) {
	if sum > C {
		return
	}
	if level == N {
		result = int(math.Max(float64(result), float64(sum)))
		return
	}
	badook(level+1, sum+array[level])
	badook(level+1, sum)
}

func TestBadook(t *testing.T) {
	C = 259
	N = 5
	result = 0
	array = []int{81, 58, 42, 33, 61}
	badook(0, 0)
	fmt.Println(result)
}
