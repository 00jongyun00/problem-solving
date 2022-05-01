package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var rd *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)

func solution(arr []int) []int {
	sort.Ints(arr)
	var result []int
	var sum, max int
	count := make(map[int]int)
	modes := make([]int, 0)

	for _, ele := range arr {
		sum += ele
		if _, exists := count[ele]; !exists {
			count[ele] = 1
		} else {
			count[ele]++
		}
		if max < count[ele] {
			max = count[ele]
		}
	}

	for key, value := range count {
		if max == value {
			modes = append(modes, key)
		}
	}
	sort.Ints(modes)
	average := float64(sum) / float64(len(arr))
	result = append(result, int(math.Round(average)))
	result = append(result, arr[len(arr)/2])
	if len(modes) > 1 {
		result = append(result, modes[1])
	} else {
		result = append(result, modes[0])
	}
	result = append(result, arr[len(arr)-1]-arr[0])

	return result
}

func main() {
	var n int
	fmt.Fscanf(rd, "%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d\n", &arr[i])
	}
	result := solution(arr)

	for _, ele := range result {
		fmt.Fprintln(wr, ele)
	}
	wr.Flush()
}
