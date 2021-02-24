package forth

import (
	"barrettotte/going-forth/stack"
	"fmt"
)

// compiled word
type definition func(*Forth) error

// Forth environment
type Forth struct {
	ds    stack.Stack           // data stack
	rs    stack.Stack           // return stack
	words map[string]definition // dictionary of words available
}

// NewForth creates and initializes a new Forth environment
func NewForth() *Forth {
	f := new(Forth)
	f.ds = *stack.NewStack()
	f.rs = *stack.NewStack()

	f.words = make(map[string]definition)
	f.addBuiltins()
	return f
}

// Add built-in words to Forth environment
func (f *Forth) addBuiltins() {
	f.words["+"] = bAdd

	// TODO: test
	f.ds.Push(1)
	f.ds.Push(5)

	f.words["+"](f)

	x, _ := f.ds.Pop()
	fmt.Println(x)
}
