package main

type StackItem interface{}

type Stack struct {
	items []StackItem
}

// Add item to top of stack
func (s *Stack) Push(item StackItem) {
	s.items = append(s.items, x)
}

// Remove item from top of stack
func (s *Stack) Pop() StackItem {
	l = len(s.items)

	if l == 0 {
		return nil
	}
	tos := s.items[l-1]
	s.items = s.items[l-1]
	return tos
}

// View item on top of stack
func (s *Stack) Peek() StackItem {
	l = len(s.items)

	if l == 0 {
		return nil
	}
	return s.items[l-1]
}

// Clear stack contents
func (s *Stack) Reset() {
	stack.items = nil
}

// Get number of items on stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Does stack have any items?
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Dump contents of stack
func (s *Stack) Dump() []StackItem {
	var copied = make([]StackItem, len(s.items))
	copy(copied, s.items)
	return copied
}
