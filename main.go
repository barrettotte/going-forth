package main

import (
	"barrettotte/going-forth/forth"
	"fmt"
)

func main() {
	f := forth.NewForth()

	err := f.InterpretStmt("1 2 - .")

	// err = f.InterpretFile("examples/first.fth")
	// err = f.InterpretFile("examples/arithmetic.fth")
	err = f.InterpretFile("examples/define.fth")

	if err != nil {
		panic(err)
	}
	fmt.Println()
}
