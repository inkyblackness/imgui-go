package imgui

import "C"
import (
	"fmt"
)

// AssertHandler is a handler for an assertion that happened in the native part of ImGui.
type AssertHandler func(expression string, file string, line int)

// AssertionError is the standard error being thrown by the default handler.
type AssertionError struct {
	Expression string
	File       string
	Line       int
}

// Error returns the string representation.
func (err AssertionError) Error() string {
	return fmt.Sprintf(`Assertion failed!
File: %s, Line %d

Expression: %s
`, err.File, err.Line, err.Expression)
}

var assertHandler AssertHandler = func(expression string, file string, line int) {
	panic(AssertionError{
		Expression: expression,
		File:       file,
		Line:       line,
	})
}

// SetAssertHandler registers a handler function for all future assertions.
// Setting nil will disable special handling.
// The default handler panics.
func SetAssertHandler(handler AssertHandler) {
	assertHandler = handler
}

//export iggAssert
func iggAssert(expression *C.char, file *C.char, line C.int) {
	if assertHandler != nil {
		assertHandler(C.GoString(expression), C.GoString(file), int(line))
	}
}
