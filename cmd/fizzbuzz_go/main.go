package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var n *int = flag.Int("n", 100, "request count")

var writer io.Writer

func init() {
	writer = os.Stdout
}

func fprintln(a ...interface{}) {
	fmt.Fprintln(writer, a...)
}

func main() {
	flag.Parse()
	for i := 1; i <= *n; i++ {
		fprintln(FizzBuzz(i))
	}
}
