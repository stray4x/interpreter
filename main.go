package main

import (
	"fmt"
	"os"

	"github.com/stray4x/interpreter/repl"
)

func main() {
	fmt.Println("Welcome to the Shreks programming language!")
	repl.Start(os.Stdin, os.Stdout)
}
