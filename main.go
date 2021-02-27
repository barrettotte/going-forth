package main

import (
	"barrettotte/going-forth/forth"
	"fmt"
)

func main() {
	f := forth.NewForth()
	var err error

	// err := f.InterpretStmt("1 2 - .")

	// err = f.InterpretFile("examples/first.fth")
	// err = f.InterpretFile("examples/arithmetic.fth")
	// err = f.InterpretFile("examples/moreops.fth")
	err = f.InterpretFile("examples/define.fth")
	// err = f.InterpretFile("examples/bugs.fth")

	if err != nil {
		panic(err)
	}
	fmt.Println()
}
