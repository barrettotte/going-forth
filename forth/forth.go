package forth

import (
	"barrettotte/going-forth/stack"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// compiled word
type word struct {
	fn         func(*Forth) error
	uncompiled string
}

// Forth environment
type Forth struct {
	ds   stack.Stack     // data stack
	rs   stack.Stack     // return stack
	dict map[string]word // dictionary of words available
}

// Interpreter states
const (
	stateNormal       = 0
	stateCommentCheck = 1
	stateComment      = 2
	stateCompileCheck = 3
	stateCompile      = 4
)

// NewForth creates and initializes a new Forth environment
func NewForth() *Forth {
	f := new(Forth)
	f.ds = *stack.NewStack()
	f.rs = *stack.NewStack()

	f.dict = make(map[string]word)
	f.addBuiltins()
	return f
}

// InterpretFile reads a file and interprets its contents
func (f *Forth) InterpretFile(path string) error {
	raw, err := readLines(path)
	if err != nil {
		return err
	}

	// first pass compiles user-defined words
	compiled, err := f.compileWords(strings.Join(raw[:], "\n"))
	if err != nil {
		return err
	}

	// use compiled source lines
	lines := strings.Split(compiled, "\n")

	for i, line := range lines {
		err = f.InterpretStmt(fmt.Sprintf(" %s ", line))
		if err != nil {
			if rte, ok := err.(*RuntimeError); ok {
				rte.line = i + 1
				return rte
			}
			if ute, ok := err.(*UnknownTokenError); ok {
				ute.line = i + 1
				return ute
			}
		}
	}
	return nil
}

// InterpretStmt interprets a Forth statement
func (f *Forth) InterpretStmt(s string) error {
	token, i := "", 0
	state := stateNormal

	for i < len(s) {
		c := s[i]

		if state == stateComment {
			// either do nothing, or end comment
			if c == ')' {
				state = stateNormal
			}
		} else if c == ' ' {
			if state == stateCommentCheck {
				state = stateComment
			} else if def, found := f.dict[strings.ToLower(token)]; found {
				// process word

				if def.uncompiled != "" {
					f.InterpretStmt(def.uncompiled)
				} else if err := def.fn(f); err != nil {
					return &RuntimeError{-1, i}
				}
				token = ""
			} else if n, err := strconv.Atoi(token); err == nil {
				// process literal

				f.ds.Push(n)
				token = ""
			} else if token != "" {
				return &UnknownTokenError{token, -1, i}
			}
		} else if c == '(' {
			state = stateCommentCheck
		} else if c == '\\' {
			return nil // line comment
		} else {
			token += string(c)

			if token == "--" {
				return nil // line comment
			}
		}
		i++
	}
	return nil
}

// compile user-define words
func (f *Forth) compileWords(src string) (string, error) {
	state := stateNormal
	i, newSrc := 0, ""
	token, compiled := "", ""

	for i < len(src) {
		c := src[i]

		if c == ':' {
			state = stateCompileCheck
		} else if state == stateCompileCheck && c == ' ' {
			state = stateCompile
		} else if state == stateCompile {
			// get word token
			if token == "" {
				for c != ' ' {
					c = src[i]
					token += string(c)
					i++
				}
				token = token[:len(token)-1] // trailing space
			}
			// get source of word
			for c != ';' {
				c = src[i]
				compiled += string(c)
				i++
			}
			compiled = compiled[:len(compiled)-1] // remove word delimiter
			f.dict[token] = word{func(f *Forth) error { return nil }, compiled}
			state = stateNormal
			// fmt.Printf(": %s %s ;\n", token, compiled)
			token, compiled = "", ""
		} else {
			newSrc += string(c)
		}
		i++
	}
	state = stateNormal
	return newSrc, nil
}

// compile a new forth word
func (f *Forth) compileWord(token string) error {
	w := word{func(f *Forth) error { return nil }, ""}
	f.dict[strings.ToLower(token)] = w
	return nil
}

// Add built-in words to Forth environment
func (f *Forth) addBuiltins() {
	f.dict["+"] = word{bAdd, ""}
	f.dict["-"] = word{bSub, ""}
	f.dict["*"] = word{bMul, ""}
	f.dict["/"] = word{bDiv, ""}
	f.dict["1+"] = word{bAdd1, ""}
	f.dict["1-"] = word{bSub1, ""}
	f.dict["2*"] = word{bMul2, ""}
	f.dict["2/"] = word{bDiv2, ""}
	f.dict["/mod"] = word{bDivMod, ""}
	f.dict["mod"] = word{bMod, ""}
	f.dict["="] = word{bEq, ""}
	f.dict["<>"] = word{bNe, ""}
	f.dict["<"] = word{bLt, ""}
	f.dict[">"] = word{bGt, ""}
	f.dict["<="] = word{bLe, ""}
	f.dict[">="] = word{bGe, ""}
	f.dict["true"] = word{bTrue, ""}
	f.dict["false"] = word{bFalse, ""}
	f.dict["dup"] = word{bDup, ""}
	f.dict["2dup"] = word{b2Dup, ""}
	f.dict["drop"] = word{bDrop, ""}
	f.dict["2drop"] = word{b2Drop, ""}
	f.dict["swap"] = word{bSwap, ""}
	f.dict["over"] = word{bOver, ""}
	f.dict["rot"] = word{bRot, ""}
	f.dict["."] = word{bDot, ""}
	f.dict[".s"] = word{bShow, ""}
	f.dict[".r"] = word{bShowR, ""}
	f.dict["cr"] = word{bCr, ""}
	f.dict["emit"] = word{bEmit, ""}
}

// read file into lines
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
