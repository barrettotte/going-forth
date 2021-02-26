package forth

import (
	"fmt"
)

// unexpectedTypeError occurs when the wrong type is encountered
type unexpectedTypeError struct{}

func (e *unexpectedTypeError) Error() string {
	return fmt.Sprintf("Unexpected type encountered")
}

// divByZeroError occurs when an attempt to divide by zero is encountered
type divByZeroError struct{}

func (e *divByZeroError) Error() string {
	return "Division by zero encountered"
}

// unknownOperatorError occurs when an unknown operator is trying to be used
type unknownOperatorError struct {
	op string
}

func (e *unknownOperatorError) Error() string {
	return fmt.Sprintf("Unknown operator '%s' encountered", e.op)
}

// UnknownTokenError occurs when an invalid token is encountered
type UnknownTokenError struct {
	tok  string
	line int
	col  int
}

func (e *UnknownTokenError) Error() string {
	return fmt.Sprintf("Unknown token '%s' encountered at %d:%d", e.tok, e.line, e.col)
}

// RuntimeError is the client error when an error occurs at runtime
type RuntimeError struct {
	line int
	col  int
}

func (e *RuntimeError) Error() string {
	return fmt.Sprintf("Runtime error occurred at %d:%d", e.line, e.col)
}

// CompileError is the client error when an error occurs at compile
type CompileError struct {
	word string
	line int
	col  int
}

func (e *CompileError) Error() string {
	return fmt.Sprintf("Compile failed for word '%s' at %d:%d", e.word, e.line, e.col)
}
