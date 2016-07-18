// The Go Programming Language
// Exercise 1.1
//
// Echo: prints its command-line arguments, including the program name (os.Args[0])
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
