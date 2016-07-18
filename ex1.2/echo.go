// The Go Programming Language
// Exercise 1.2
//
// Echo: prints its command-line arguments by index and value, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		fmt.Printf("%d\t%s\n", index, arg)
	}
}

