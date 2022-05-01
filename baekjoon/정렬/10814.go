package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var rd *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)

type client struct {
	age  int
	name string
	idx  int
}

func main() {
	var n, inputAge int
	var inputName string
	fmt.Fscanf(rd, "%d\n", &n)
	clients := make([]client, 0)
	for i := 0; i < n; i++ {
		fmt.Fscanf(rd, "%d %s\n", &inputAge, &inputName)
		clients = append(clients, client{inputAge, inputName, i})
	}

	sort.Slice(clients, func(i, j int) bool {
		if clients[i].age == clients[j].age {
			return clients[i].idx < clients[j].idx
		}
		return clients[i].age < clients[j].age
	})

	for _, c := range clients {
		fmt.Fprintf(wr, "%d %s\n", c.age, c.name)
	}

	wr.Flush()
}
