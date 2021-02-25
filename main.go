package main

import (
	"barrettotte/going-forth/forth"
)

func main() {
	f := forth.NewForth()
	f.InterpretStmt("1 2 - .")
}
