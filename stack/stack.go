package stack

// Item struct for generic stack item
type Item interface{}

// Stack data structure
type Stack struct {
	items []Item
}

// UnderflowError occurs when peeking or popping an empty stack
type UnderflowError struct{}

func (err *UnderflowError) Error() string {
	return "No items on stack"
}

// NewStack creates a new stack
func NewStack() *Stack {
	s := new(Stack)
	s.items = []Item{}
	return s
}

// Push adds item to top of stack
func (s *Stack) Push(item Item) {
	s.items = append(s.items, item)
}

// Pop removes item from top of stack
func (s *Stack) Pop() (Item, *UnderflowError) {
	if len(s.items) == 0 {
		return nil, &UnderflowError{}
	}
	tos := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return tos, nil
}

// Peek item on top of stack
func (s *Stack) Peek() (Item, *UnderflowError) {
	if len(s.items) == 0 {
		return nil, &UnderflowError{}
	}
	return s.items[len(s.items)-1], nil
}

// Reset clears stack contents
func (s *Stack) Reset() {
	s.items = nil
}

// Size gets amount of items on stack
func (s *Stack) Size() int {
	return len(s.items)
}

// IsEmpty checks if stack has any items
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Dump creates a copy of stack contents
func (s *Stack) Dump() []Item {
	var copied = make([]Item, len(s.items))
	copy(copied, s.items)
	return copied
}
