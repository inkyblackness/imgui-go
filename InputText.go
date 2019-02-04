package imgui

// #include "InputTextCallbackDataWrapper.h"
import "C"
import "sync"

const (
	// InputTextFlagsNone sets everything default.
	InputTextFlagsNone = 0
	// InputTextFlagsCharsDecimal allows 0123456789.+-
	InputTextFlagsCharsDecimal = 1 << 0
	// InputTextFlagsCharsHexadecimal allow 0123456789ABCDEFabcdef
	InputTextFlagsCharsHexadecimal = 1 << 1
	// InputTextFlagsCharsUppercase turns a..z into A..Z.
	InputTextFlagsCharsUppercase = 1 << 2
	// InputTextFlagsCharsNoBlank filters out spaces, tabs.
	InputTextFlagsCharsNoBlank = 1 << 3
	// InputTextFlagsAutoSelectAll selects entire text when first taking mouse focus.
	InputTextFlagsAutoSelectAll = 1 << 4
	// InputTextFlagsEnterReturnsTrue returns 'true' when Enter is pressed (as opposed to when the value was modified).
	InputTextFlagsEnterReturnsTrue = 1 << 5
	// InputTextFlagsCallbackCompletion for callback on pressing TAB (for completion handling).
	InputTextFlagsCallbackCompletion = 1 << 6
	// InputTextFlagsCallbackHistory for callback on pressing Up/Down arrows (for history handling).
	InputTextFlagsCallbackHistory = 1 << 7
	// InputTextFlagsCallbackAlways for callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackAlways = 1 << 8
	// InputTextFlagsCallbackCharFilter for callback on character inputs to replace or discard them.
	// Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsCallbackCharFilter = 1 << 9
	// InputTextFlagsAllowTabInput when pressing TAB to input a '\t' character into the text field.
	InputTextFlagsAllowTabInput = 1 << 10
	// InputTextFlagsCtrlEnterForNewLine in multi-line mode, unfocus with Enter, add new line with Ctrl+Enter
	// (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsCtrlEnterForNewLine = 1 << 11
	// InputTextFlagsNoHorizontalScroll disables following the cursor horizontally.
	InputTextFlagsNoHorizontalScroll = 1 << 12
	// InputTextFlagsAlwaysInsertMode sets insert mode.
	InputTextFlagsAlwaysInsertMode = 1 << 13
	// InputTextFlagsReadOnly sets read-only mode.
	InputTextFlagsReadOnly = 1 << 14
	// InputTextFlagsPassword sets password mode, display all characters as '*'.
	InputTextFlagsPassword = 1 << 15
	// InputTextFlagsNoUndoRedo disables undo/redo. Note that input text owns the text data while active,
	// if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsNoUndoRedo = 1 << 16
	// InputTextFlagsCharsScientific allows 0123456789.+-*/eE (Scientific notation input).
	InputTextFlagsCharsScientific = 1 << 17
	// InputTextFlagsCallbackResize for callback on buffer capacity changes request (beyond 'buf_size' parameter value),
	// allowing the string to grow. Notify when the string wants to be resized (for string types which hold a cache of their Size).
	// You will be provided a new BufSize in the callback and NEED to honor it. (see misc/cpp/imgui_stdlib.h for an example of using this)
	InputTextFlagsCallbackResize = 1 << 18
)

// InputTextCallback is called for sharing state of an input field.
// By default, the callback should return 0.
type InputTextCallback func(InputTextCallbackData) int32

var inputTextCallbacks = make(map[C.int]InputTextCallback)
var inputTextCallbacksMutex sync.Mutex

//export iggInputTextCallback
func iggInputTextCallback(handle C.IggInputTextCallbackData, key C.int) C.int {
	cb := iggInputTextCallbackFor(key)
	return C.int(cb(InputTextCallbackData{handle: handle}))
}

func iggInputTextCallbackFor(key C.int) InputTextCallback {
	inputTextCallbacksMutex.Lock()
	defer inputTextCallbacksMutex.Unlock()
	cb, existing := inputTextCallbacks[key]
	if !existing || (cb == nil) {
		return func(InputTextCallbackData) int32 { return 0 }
	}
	return cb
}

func iggInputTextCallbackKeyFor(cb InputTextCallback) C.int {
	if cb == nil {
		return 0
	}
	inputTextCallbacksMutex.Lock()
	defer inputTextCallbacksMutex.Unlock()
	key := C.int(len(inputTextCallbacks) + 1)
	for _, existing := inputTextCallbacks[key]; existing; _, existing = inputTextCallbacks[key] {
		key++
	}
	inputTextCallbacks[key] = cb
	return key
}

func iggInputTextCallbackKeyRelease(key C.int) {
	inputTextCallbacksMutex.Lock()
	defer inputTextCallbacksMutex.Unlock()
	delete(inputTextCallbacks, key)
}

// InputTextCallbackData represents the shared state of InputText(), passed as an argument to your callback.
type InputTextCallbackData struct {
	handle C.IggInputTextCallbackData
}

// EventFlags returns a combination on InputTextFlagsCallback* constants.
func (data InputTextCallbackData) EventFlags() int {
	return int(C.iggInputTextCallbackDataGetEventFlags(data.handle))
}
