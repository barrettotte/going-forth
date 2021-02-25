package forth

// UnexpectedTypeError occurs when the wrong type is encountered
type UnexpectedTypeError struct{}

func (err *UnexpectedTypeError) Error() string {
	return "Unexpected type encountered"
}

// DivByZeroError occurs when an attempt to divide by zero is encountered
type DivByZeroError struct{}

func (err *DivByZeroError) Error() string {
	return "Division by zero encountered"
}

// UnknownOperatorError occurs when an unknown operator is trying to be used
type UnknownOperatorError struct{}

func (err *UnknownOperatorError) Error() string {
	return "Unknown operator encountered"
}
