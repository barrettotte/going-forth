package main

import (
	"barrettotte/going-forth/forth"
	"fmt"
)

func main() {
	f := forth.NewForth()
	var err error

	// err := f.InterpretStmt("1 2 - .")

	// err = f.InterpretFile("forth-pgms/first.fth")
	// err = f.InterpretFile("forth-pgms/arithmetic.fth")
	// err = f.InterpretFile("forth-pgms/moreops.fth")
	err = f.InterpretFile("forth-pgms/define.fth")
	// err = f.InterpretFile("forth-pgms/bugs.fth")

	if err != nil {
		panic(err)
	}
	fmt.Println()
}
