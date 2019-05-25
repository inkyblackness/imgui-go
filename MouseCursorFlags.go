package imgui

// Enumeration for MouseCursor()
// User code may request binding to display given cursor by calling SetMouseCursor(), which is why we have some cursors that are marked unused here
const (
	// MouseCursorNone no mouse cursor
	MouseCursorNone = iota - 1
	// MouseCursorArrow standard arrow mouse cursor
	MouseCursorArrow
	// MouseCursorTextInput when hovering over InputText, etc.
	MouseCursorTextInput
	// MouseCursorResizeAll (Unused by imgui functions)
	MouseCursorResizeAll
	// MouseCursorResizeNS when hovering over an horizontal border
	MouseCursorResizeNS
	// MouseCursorResizeEW when hovering over a vertical border or a column
	MouseCursorResizeEW
	// MouseCursorResizeNESW when hovering over the bottom-left corner of a window
	MouseCursorResizeNESW
	// MouseCursorResizeNWSE when hovering over the bottom-right corner of a window
	MouseCursorResizeNWSE
	// MouseCursorHand (Unused by imgui functions. Use for e.g. hyperlinks)
	MouseCursorHand
	MouseCursorCount
)
