package forth

import (
	"barrettotte/going-forth/stack"
	"fmt"
)

// try to pop n integers from stack
func popNInt(s *stack.Stack, n int) ([]int, error) {
	items := make([]int, n)

	for i := 0; i < n; i++ {
		n, err := popInt(s)
		if err != nil {
			return nil, err
		}
		items[i] = n
	}
	return items, nil
}

// try to pop an integer from the stack
func popInt(s *stack.Stack) (int, error) {
	n, err := s.Pop()
	if err != nil {
		return 0, err // underflow
	}
	i, ok := n.(int)
	if !ok {
		return 0, &unexpectedTypeError{}
	}
	return i, nil
}

// try to pop n items from stack
func popN(s *stack.Stack, n int) ([]stack.Item, error) {
	items := make([]stack.Item, n)

	for i := 0; i < n; i++ {
		x, err := s.Pop()
		if err != nil {
			return nil, err
		}
		items[i] = x
	}
	return items, nil
}

// convert true->1 and false->0
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// relational operators ( n1 n2 -- flag )
func relop(s *stack.Stack, operator string) error {
	ops, err := popNInt(s, 2)
	if err != nil {
		return err
	}
	switch operator {
	case "=":
		s.Push(boolToInt(ops[1] == ops[0]))
	case "<>":
		s.Push(boolToInt(ops[1] != ops[0]))
	case "<":
		s.Push(boolToInt(ops[1] < ops[0]))
	case ">":
		s.Push(boolToInt(ops[1] > ops[0]))
	case "<=":
		s.Push(boolToInt(ops[1] <= ops[0]))
	case ">=":
		s.Push(boolToInt(ops[1] >= ops[0]))
	default:
		return &unknownOperatorError{operator}
	}
	return nil
}

// bAdd  :+ ( n1 n2 -- n3 )
func bAdd(f *Forth) error {
	ops, err := popNInt(&f.ds, 2)
	if err != nil {
		return err
	}
	f.ds.Push(ops[1] + ops[0])
	return nil
}

// bAdd1  :1+ ( n1 -- n2 )
func bAdd1(f *Forth) error {
	n, err := popInt(&f.ds)
	if err != nil {
		return err
	}
	f.ds.Push(n + 1)
	return nil
}

// bSub  :- ( n1 n2 -- n3 )
func bSub(f *Forth) error {
	ops, err := popNInt(&f.ds, 2)
	if err != nil {
		return err
	}
	f.ds.Push(ops[1] - ops[0])
	return nil
}

// bSub1  :1- ( n1 -- n2 )
func bSub1(f *Forth) error {
	n, err := popInt(&f.ds)
	if err != nil {
		return err
	}
	f.ds.Push(n - 1)
	return nil
}

// bMul  :* ( n1 n2 -- n3 )
func bMul(f *Forth) error {
	ops, err := popNInt(&f.ds, 2)
	if err != nil {
		return err
	}
	f.ds.Push(ops[1] * ops[0])
	return nil
}

// bMul2  :2* ( n1 -- n2 )
func bMul2(f *Forth) error {
	n, err := popInt(&f.ds)
	if err != nil {
		return err
	}
	f.ds.Push(n * 2)
	return nil
}

// bDiv  :/ ( n1 n2 -- n3 )
func bDiv(f *Forth) error {
	ops, err := popNInt(&f.ds, 2)
	if err != nil {
		return err
	}
	if ops[0] == 0 {
		return &divByZeroError{}
	}
	f.ds.Push(ops[1] / ops[0])
	return nil
}

// bDiv2  :2/ ( n1 -- n2 )
func bDiv2(f *Forth) error {
	n, err := popInt(&f.ds)
	if err != nil {
		return err
	}
	f.ds.Push(n / 2)
	return nil
}

// bDivMod  :/MOD ( n1 n2 -- n3 n4 )
func bDivMod(f *Forth) error {
	ops, err := popNInt(&f.ds, 2)
	if err != nil {
		return err
	}
	if ops[1] == 0 {
		return &divByZeroError{}
	}
	f.ds.Push(ops[1] % ops[0])
	f.ds.Push(ops[1] / ops[0])
	return nil
}

// bMod  :MOD ( n1 n2 -- n3 )
func bMod(f *Forth) error {
	ops, err := popNInt(&f.ds, 2)
	if err != nil {
		return err
	}
	if ops[1] == 0 {
		return &divByZeroError{}
	}
	f.ds.Push(ops[1] % ops[0])
	return nil
}

// TODO: zero-relational operators
// 0=  0<>  0<  0>  0<=  0>=

// bEq  := ( n1 n2 -- flag )
func bEq(f *Forth) error {
	return relop(&f.ds, "=")
}

// bNe  :<> ( n1 n2 -- flag )
func bNe(f *Forth) error {
	return relop(&f.ds, "<>")
}

// bLt  :< ( n1 n2 -- flag )
func bLt(f *Forth) error {
	return relop(&f.ds, "<")
}

// bGt  :> ( n1 n2 -- flag )
func bGt(f *Forth) error {
	return relop(&f.ds, ">")
}

// bLe  :<= ( n1 n2 -- flag )
func bLe(f *Forth) error {
	return relop(&f.ds, "<=")
}

// bGe  :>= ( n1 n2 -- flag )
func bGe(f *Forth) error {
	return relop(&f.ds, ">=")
}

// bTrue  :true ( -- flag )
func bTrue(f *Forth) error {
	f.ds.Push(1)
	return nil
}

// bFalse  :false ( -- flag )
func bFalse(f *Forth) error {
	f.ds.Push(0)
	return nil
}

// bDup  :dup ( n -- n n )
func bDup(f *Forth) error {
	n, err := f.ds.Pop()
	if err != nil {
		return err
	}
	f.ds.Push(n)
	f.ds.Push(n)
	return nil
}

// b2Dup  :2dup ( n1 n2 -- n1 n2 n1 n2 )
func b2Dup(f *Forth) error {
	pair, err := popN(&f.ds, 2)
	if err != nil {
		return err
	}
	f.ds.Push(pair[1])
	f.ds.Push(pair[0])
	f.ds.Push(pair[1])
	f.ds.Push(pair[0])
	return nil
}

// bDrop  :drop ( n -- )
func bDrop(f *Forth) error {
	_, err := f.ds.Pop()
	if err != nil {
		return err
	}
	return nil
}

// b2Drop  :2drop (  n1 n2 -- )
func b2Drop(f *Forth) error {
	err := bDrop(f)
	if err == nil {
		err = bDrop(f)
	}
	return err
}

// bSwap  :swap ( n1 n2 -- n2 n1 )
func bSwap(f *Forth) error {
	pair, err := popN(&f.ds, 2)
	if err != nil {
		return err
	}
	f.ds.Push(pair[0])
	f.ds.Push(pair[1])
	return nil
}

// bOver  :over ( n1 n2 -- n1 n2 n1 )
func bOver(f *Forth) error {
	pair, err := popN(&f.ds, 2)
	if err != nil {
		return err
	}
	f.ds.Push(pair[1])
	f.ds.Push(pair[0])
	f.ds.Push(pair[1])
	return nil
}

// bRot  :rot ( n1 n2 n3 -- n2 n3 n1 )
func bRot(f *Forth) error {
	trio, err := popN(&f.ds, 3)
	if err != nil {
		return err
	}
	f.ds.Push(trio[1])
	f.ds.Push(trio[0])
	f.ds.Push(trio[2])
	return nil
}

// bDot  :. ( n -- )
func bDot(f *Forth) error {
	n, err := f.ds.Pop()
	if err != nil {
		return err
	}
	fmt.Printf("%v", n)
	return nil
}

// bShow  :.s ( -- )
func bShow(f *Forth) error {
	for _, item := range f.ds.Dump() {
		fmt.Println(item)
	}
	return nil
}

// bShowR  :.r ( -- )
func bShowR(f *Forth) error {
	for _, item := range f.rs.Dump() {
		fmt.Println(item)
	}
	return nil
}

// bCr  :cr  ( -- )
func bCr(f *Forth) error {
	fmt.Println()
	return nil
}

// bEmit :emit  ( n -- )
func bEmit(f *Forth) error {
	n, err := popInt(&f.ds)
	if err != nil {
		return err
	}
	fmt.Print(string(rune(n)))
	return nil
}

// TODO: bDotQuote :." ( -- )
