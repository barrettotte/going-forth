package forth

import (
	"barrettotte/going-forth/stack"
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

// InterpretFile reads a file and interprets its contents
func (f *Forth) InterpretFile() {

}

// InterpretLine interprets a line of text
func (f *Forth) InterpretLine() {

}

// Add built-in words to Forth environment
func (f *Forth) addBuiltins() {
	f.words["+"] = bAdd
	f.words["-"] = bSub
	f.words["*"] = bMul
	f.words["/"] = bDiv
	f.words["1+"] = bAdd1
	f.words["1-"] = bSub1
	f.words["2*"] = bMul2
	f.words["2/"] = bDiv2
	f.words["/mod"] = bMod
	f.words["="] = bEq
	f.words["<>"] = bNe
	f.words["<"] = bLt
	f.words[">"] = bGt
	f.words["<="] = bLe
	f.words[">="] = bGe
	f.words["true"] = bTrue
	f.words["false"] = bFalse
	f.words["dup"] = bDup
	f.words["2dup"] = b2Dup
	f.words["drop"] = bDrop
	f.words["2drop"] = b2Drop
	f.words["swap"] = bSwap
	f.words["over"] = bOver
	f.words["rot"] = bRot
	f.words["."] = bDot
	f.words[".s"] = bShow
	f.words[".r"] = bShowR
	f.words["cr"] = bCr
	f.words["emit"] = bEmit
}
