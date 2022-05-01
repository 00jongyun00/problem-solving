package bfs

import (
	"testing"
)

var n int
var ch []int
var test *testing.T

func findASubSet(input int) {
	if input == n+1 {
		for i, c := range ch {
			if c == 1 {
				test.Logf("%d ", i)
			}
		}
		test.Log()
		return
	}
	ch[input] = 1
	findASubSet(input + 1)
	ch[input] = 0
	findASubSet(input + 1)
}

func TestFindASubSet(t *testing.T) {
	test = t
	n = 3
	ch = make([]int, n+1)
	findASubSet(1)
}
