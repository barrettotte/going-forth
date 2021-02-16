package stack

import (
	"testing"
)

func newStack() *Stack {
	s := Stack{}
	s.New()
	return &s
}

func TestPushPop(t *testing.T) {
	s := newStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Errorf("Expected size 3, but got %v", s.Size())
	}

	x, err := s.Pop()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err.msg)
	}
	if x != 3 {
		t.Errorf("Expected 3, but got %v", x)
	}
	if s.Size() != 2 {
		t.Errorf("Expected size 2, but got %v", s.Size())
	}

	// force underflow error
	_, err = s.Pop()
	_, err = s.Pop()
	_, err = s.Pop()

	if err == nil {
		t.Errorf("Expected UnderflowError, but no error occurred.")
	}
}

func TestPeek(t *testing.T) {
	s := newStack()
	s.Push("AAA")
	s.Push("BBB")
	s.Push("CCC")

	x, err := s.Peek()

	if err != nil {
		t.Errorf("Unexpected error occurred. %s", err.msg)
	}
	if x != "CCC" {
		t.Errorf("Expected 'AAA', but got %v", x)
	}

	// force underflow error
	s.Pop()
	s.Pop()
	s.Pop()

	_, err = s.Peek()

	if err == nil {
		t.Errorf("Expected UnderflowError, but no error occurred.")
	}
}

func TestIsEmpty(t *testing.T) {
	s := newStack()

	if !s.IsEmpty() {
		t.Errorf("Expected empty stack, but it was not empty.")
	}
	s.Push(1)

	if s.IsEmpty() {
		t.Errorf("Expected stack with contents, but the stack was empty")
	}
}
