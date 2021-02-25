package forth

import (
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
		t.Errorf("Unexpected error occured. %s", err)
	}
	if tos != 3 {
		t.Errorf("Expected sum of 3, but got %d", tos)
	}

	f.InterpretStmt("8 1+")
	tos, err = f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occured. %s", err)
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
		t.Errorf("Unexpected error occured. %s", err)
	}
	if tos != -1 {
		t.Errorf("Expected difference of -1, but got %d", tos)
	}

	f.InterpretStmt("6 1-")
	tos, err = f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occured. %s", err)
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
		t.Errorf("Unexpected error occured. %s", err)
	}
	if tos != 64 {
		t.Errorf("Expected product of 64, but got %d", tos)
	}

	f.InterpretStmt("50 2*")
	tos, err = f.ds.Peek()

	if err != nil {
		t.Errorf("Unexpected error occured. %s", err)
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
		t.Errorf("Unexpected error occured. %s", err)
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
		t.Errorf("Unexpected error occured. %s", err)
	}
	if tos != 4 {
		t.Errorf("Expected quotient of 4, but got %d", tos)
	}
}
