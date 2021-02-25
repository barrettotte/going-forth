package forth

import (
	"barrettotte/going-forth/stack"
	"testing"
)

func newForth() *Forth {
	f := NewForth()
	return f
}

func TestAdd(t *testing.T) {
	f := newForth()
	f.InterpretStmt("1 2 +")
	tos, err := f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 3 {
		t.Errorf("Expected sum of 3, but got %d", tos)
	}

	f.InterpretStmt("8 1+")
	tos, err = f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 9 {
		t.Errorf("Expected sum of 9, but got %d", tos)
	}
}

func TestSub(t *testing.T) {
	f := newForth()
	f.InterpretStmt("1 2 -")
	tos, err := f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != -1 {
		t.Errorf("Expected difference of -1, but got %d", tos)
	}

	f.InterpretStmt("6 1-")
	tos, err = f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 5 {
		t.Errorf("Expected difference of 5, but got %d", tos)
	}
}

func TestMul(t *testing.T) {
	f := newForth()
	f.InterpretStmt("4 4 4 * *")
	tos, err := f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 64 {
		t.Errorf("Expected product of 64, but got %d", tos)
	}

	f.InterpretStmt("50 2*")
	tos, err = f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 100 {
		t.Errorf("Expected product of 5, but got %d", tos)
	}
}

func TestDiv(t *testing.T) {
	f := newForth()
	f.InterpretStmt("50 2 /")
	tos, err := f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 25 {
		t.Errorf("Expected quotient of 25, but got %d", tos)
	}

	dbz := f.InterpretStmt("drop 4 0 / drop")
	if dbz == nil {
		t.Errorf("Expected division by zero error, but did not get one.")
	}

	f.InterpretStmt("8 2/")
	tos, err = f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 4 {
		t.Errorf("Expected quotient of 4, but got %d", tos)
	}
}

func TestMod(t *testing.T) {
	f := newForth()
	f.InterpretStmt("3 2 mod")
	tos, err := f.ds.Pop()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 1 {
		t.Errorf("Expected remainder of 1, but got %d", tos)
	}

	f.InterpretStmt("19 6 /mod")
	tos, err = f.ds.Pop()
	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 3 {
		t.Errorf("Expected quotient of 3, but got %d", tos)
	}

	tos, err = f.ds.Pop()
	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 1 {
		t.Errorf("Expected remainder of 1, but got %d", tos)
	}
}

func TestRelop(t *testing.T) {
	f := newForth()

	f.InterpretStmt("10 10 =")
	tos, err := f.ds.Pop()
	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 1 {
		t.Errorf("Expected flag to be 1, but got %d", tos)
	}

	f.InterpretStmt("10 10 <>")
	tos, err = f.ds.Pop()
	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 0 {
		t.Errorf("Expected flag to be 0, but got %d", tos)
	}

	f.InterpretStmt("5 10 <")
	tos, err = f.ds.Pop()
	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 1 {
		t.Errorf("Expected flag to be 1, but got %d", tos)
	}

	f.InterpretStmt("5 10 >")
	tos, err = f.ds.Pop()
	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 0 {
		t.Errorf("Expected flag to be 0, but got %d", tos)
	}
}

func TestDup(t *testing.T) {
	f := newForth()
	f.InterpretStmt("8 dup")
	tos, err := f.ds.Pop()
	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err)
	}
	if tos != 8 {
		t.Errorf("Expected 8, but got %d", tos)
	}
}

func TestDrop(t *testing.T) {
	f := newForth()
	f.InterpretStmt("8 drop")
	_, err := f.ds.Pop()
	if err == nil {
		t.Errorf("Expected error to occur, but did not happen.")
	}
}

func TestSwap(t *testing.T) {
	f := newForth()
	f.InterpretStmt("1 2 3 4 swap")
	const n = 4
	items := make([]stack.Item, n)

	for i := 0; i < n; i++ {
		item, err := f.ds.Pop()
		if err != nil {
			t.Errorf("Unexpected error occurred. %s", err)
		}
		items[i] = item
	}
	expected := [n]int{3, 4, 2, 1} // stack {1,2,4,3} <-- top
	for i := 0; i < n; i++ {
		if expected[i] != items[i] {
			t.Errorf("Expected %d, but got %d", expected[i], items[i])
		}
	}
}

func TestOver(t *testing.T) {
	f := newForth()
	f.InterpretStmt("1 2 3 over")
	const n = 4
	items := make([]stack.Item, n)

	for i := 0; i < n; i++ {
		item, err := f.ds.Pop()
		if err != nil {
			t.Errorf("Unexpected error occurred. %s", err)
		}
		items[i] = item
	}
	expected := [n]int{2, 3, 2, 1} // stack {1,2,3,2} <-- top
	for i := 0; i < n; i++ {
		if expected[i] != items[i] {
			t.Errorf("Expected %d, but got %d", expected[i], items[i])
		}
	}
}

func TestRot(t *testing.T) {
	f := newForth()
	f.InterpretStmt("1 2 3 rot")
	const n = 3
	items := make([]stack.Item, n)

	for i := 0; i < n; i++ {
		item, err := f.ds.Pop()
		if err != nil {
			t.Errorf("Unexpected error occurred. %s", err)
		}
		items[i] = item
	}
	expected := [n]int{1, 3, 2} // stack {2,3,1} <-- top
	for i := 0; i < n; i++ {
		if expected[i] != items[i] {
			t.Errorf("Expected %d, but got %d", expected[i], items[i])
		}
	}
}
