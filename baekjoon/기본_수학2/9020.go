package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPrimes() []bool {
	primes := make([]bool, 10001)
	for i := 2; i < len(primes); i++ {
		primes[i] = true
	}

	for i := 2; i < len(primes); i++ {
		if primes[i] {
			for j := i * 2; j < len(primes); j += i {
				primes[j] = false
			}
		}
	}
	return primes
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	var t, n int
	primes := getPrimes()
	fmt.Fscanf(rd, "%d\n", &t)
	for i := 0; i < t; i++ {
		fmt.Fscanf(rd, "%d\n", &n)
		var a, b, minSub int
		minSub = 10001
		for j := n; j >= 2; j-- {
			if primes[j] && primes[n-j] {
				sub := n - j*2
				if sub < 0 {
					sub = -sub
				}
				if minSub > sub {
					a = j
					b = n - j
					minSub = sub
				}
			}
		}
		if a < b {
			fmt.Fprintf(wr, "%d %d\n", a, b)
		} else {
			fmt.Fprintf(wr, "%d %d\n", b, a)
		}
	}
	wr.Flush()
}
