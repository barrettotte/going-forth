package forth

import (
	"barrettotte/going-forth/stack"
	"fmt"
	"strconv"
	"strings"
)

// compiled word
type word func(*Forth) error

// Forth environment
type Forth struct {
	ds   stack.Stack     // data stack
	rs   stack.Stack     // return stack
	dict map[string]word // dictionary of words available
}

// NewForth creates and initializes a new Forth environment
func NewForth() *Forth {
	f := new(Forth)
	f.ds = *stack.NewStack()
	f.rs = *stack.NewStack()

	f.dict = make(map[string]word)
	f.addBuiltins()
	return f
}

// InterpretFile reads a file and interprets its contents
func (f *Forth) InterpretFile() {
	// TODO:
	// read file
	// for each line call InterpretStmt
}

// InterpretStmt interprets a Forth statement
func (f *Forth) InterpretStmt(s string) error {
	var err error

	for _, token := range strings.Split(s, " ") {
		if def, found := f.dict[strings.ToLower(token)]; found {
			err = def(f)

			if err != nil {
				break
			}
		} else if n, err := strconv.Atoi(token); err == nil {
			f.ds.Push(n)
		} else {
			fmt.Printf("Invalid token encountered '%s'\n", token)
		}
	}
	return err
}

// Add built-in words to Forth environment
func (f *Forth) addBuiltins() {
	f.dict["+"] = bAdd
	f.dict["-"] = bSub
	f.dict["*"] = bMul
	f.dict["/"] = bDiv
	f.dict["1+"] = bAdd1
	f.dict["1-"] = bSub1
	f.dict["2*"] = bMul2
	f.dict["2/"] = bDiv2
	f.dict["/mod"] = bDivMod
	f.dict["mod"] = bMod
	f.dict["="] = bEq
	f.dict["<>"] = bNe
	f.dict["<"] = bLt
	f.dict[">"] = bGt
	f.dict["<="] = bLe
	f.dict[">="] = bGe
	f.dict["true"] = bTrue
	f.dict["false"] = bFalse
	f.dict["dup"] = bDup
	f.dict["2dup"] = b2Dup
	f.dict["drop"] = bDrop
	f.dict["2drop"] = b2Drop
	f.dict["swap"] = bSwap
	f.dict["over"] = bOver
	f.dict["rot"] = bRot
	f.dict["."] = bDot
	f.dict[".s"] = bShow
	f.dict[".r"] = bShowR
	f.dict["cr"] = bCr
	f.dict["emit"] = bEmit
}
