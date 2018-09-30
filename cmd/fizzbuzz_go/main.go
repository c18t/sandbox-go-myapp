package main

import (
	"flag"
	"fmt"
)

var n *int = flag.Int("n", 100, "request count")

func main() {
	flag.Parse()
	for i := 1; i <= *n; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
