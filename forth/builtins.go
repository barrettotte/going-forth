package forth

import (
	"barrettotte/going-forth/stack"
)

// UnexpectedTypeError occurs when the wrong type is encountered
type UnexpectedTypeError struct{}

func (err *UnexpectedTypeError) Error() string {
	return "Unexpected type encountered"
}

// try to pop stack item and assert type to integer
func popInt(s *stack.Stack) (int, error) {
	item, err := s.Pop()
	if err != nil {
		return 0, err // Underflow
	}

	i, ok := item.(int)
	if !ok {
		return 0, &UnexpectedTypeError{}
	}
	return i, nil
}

// try to pop two integers from a stack
func popTwoInts(s *stack.Stack) (int, int, error) {
	a, err := popInt(s)
	if err != nil {
		return 0, 0, err
	}

	b, err := popInt(s)
	if err != nil {
		return 0, 0, err
	}

	return a, b, nil
}

// bAdd  :+ ( a b -- c )
// Pop two operands from data stack,
// add them, and push sum back onto data stack.
func bAdd(f *Forth) error {
	a, b, err := popTwoInts(&f.ds)
	if err != nil {
		return err
	}

	f.ds.Push(a + b)
	return err
}

// TODO: bSub  :- ( a b -- c )

// TODO: bMul  :* ( a b -- c )

// TODO: bDiv  :/ ( a b -- c )

// TODO: bDup  :dup ( n -- n n )

// TODO: bDrop  :drop ( n -- )

// TODO: bSwap  :swap ( a b -- b a )

// TODO: bOver  :over ( a b -- a b a )

// TODO: bRot  :rot ( a b c -- b c a )

// TODO: bDot  :. ( n -- )

// TODO: bEmit :emit  ( n -- )

// TODO: bCr  :cr  ( -- )

// TODO: bDotQuote :." ( -- )

// relational ops ???
// and, or, invert ???
// if then else  ???
// do loop ??
