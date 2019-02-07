package imgui

// #include "InputTextCallbackDataWrapper.h"
import "C"
import (
	"sync"
	"unsafe"
)

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
	// inputTextFlagsCallbackResize for callback on buffer capacity change requests.
	inputTextFlagsCallbackResize = 1 << 18
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

// EventFlag returns one of the InputTextFlagsCallback* constants to indicate the nature of the callback.
func (data InputTextCallbackData) EventFlag() int {
	return int(C.iggInputTextCallbackDataGetEventFlag(data.handle))
}

// Flags returns the set of flags that the user originally passed to InputText.
func (data InputTextCallbackData) Flags() int {
	return int(C.iggInputTextCallbackDataGetFlags(data.handle)) & ^inputTextFlagsCallbackResize
}

// EventChar returns the current character input.
func (data InputTextCallbackData) EventChar() rune {
	return rune(C.iggInputTextCallbackDataGetEventChar(data.handle))
}

// SetEventChar overrides what the user entered. Set to zero do drop the current input.
// Returning 1 from the callback also drops the current input.
//
// Note: The internal representation of characters is based on uint16, so less than rune would provide.
func (data InputTextCallbackData) SetEventChar(value rune) {
	C.iggInputTextCallbackDataSetEventChar(data.handle, C.ushort(value))
}

// EventKey returns the currently pressed key. Valid for completion and history callbacks.
func (data InputTextCallbackData) EventKey() int {
	return int(C.iggInputTextCallbackDataGetEventKey(data.handle))
}

func (data InputTextCallbackData) setBuf(buf unsafe.Pointer, size, textLen int) {
	C.iggInputTextCallbackDataSetBuf(data.handle, (*C.char)(buf), C.int(size), C.int(textLen))
}

func (data InputTextCallbackData) bufSize() int {
	return int(C.iggInputTextCallbackDataGetBufSize(data.handle))
}

func (data InputTextCallbackData) bufTextLen() int {
	return int(C.iggInputTextCallbackDataGetBufTextLen(data.handle))
}

// CursorPos returns the byte-offset of the cursor within the buffer.
func (data InputTextCallbackData) CursorPos() int {
	return int(C.iggInputTextCallbackDataGetCursorPos(data.handle))
}

// SetCursorPos changes the current byte-offset of the cursor within the buffer.
func (data InputTextCallbackData) SetCursorPos(value int) {
	C.iggInputTextCallbackDataSetCursorPos(data.handle, C.int(value))
}

// SelectionStart returns the byte-offset of the selection start within the buffer.
func (data InputTextCallbackData) SelectionStart() int {
	return int(C.iggInputTextCallbackDataGetSelectionStart(data.handle))
}

// SetSelectionStart changes the current byte-offset of the selection start within the buffer.
func (data InputTextCallbackData) SetSelectionStart(value int) {
	C.iggInputTextCallbackDataSetSelectionStart(data.handle, C.int(value))
}

// SelectionEnd returns the byte-offset of the selection end within the buffer.
func (data InputTextCallbackData) SelectionEnd() int {
	return int(C.iggInputTextCallbackDataGetSelectionEnd(data.handle))
}

// SetSelectionEnd changes the current byte-offset of the selection end within the buffer.
func (data InputTextCallbackData) SetSelectionEnd(value int) {
	C.iggInputTextCallbackDataSetSelectionEnd(data.handle, C.int(value))
}
