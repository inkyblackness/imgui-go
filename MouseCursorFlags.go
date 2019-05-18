package imgui

const (
	// Enumeration for MouseCursor()
	// User code may request binding to display given cursor by calling SetMouseCursor(), which is why we have some cursors that are marked unused here
	MouseCursorNone = iota - 1
	MouseCursorArrow
	// When hovering over InputText, etc.
	MouseCursorTextInput
	// (Unused by imgui functions)
	MouseCursorResizeAll
	// When hovering over an horizontal border
	MouseCursorResizeNS
	// When hovering over a vertical border or a column
	MouseCursorResizeEW
	// When hovering over the bottom-left corner of a window
	MouseCursorResizeNESW
	// When hovering over the bottom-right corner of a window
	MouseCursorResizeNWSE
	// (Unused by imgui functions. Use for e.g. hyperlinks)
	MouseCursorHand
	MouseCursorCount
)
