package imgui

const (
	// Enumeration for MouseCursor()
	// User code may request binding to display given cursor by calling SetMouseCursor(), which is why we have some cursors that are marked unused here
	MouseCursorNone = iota - 1
	MouseCursorArrow
	MouseCursorTextInput  // When hovering over InputText, etc.
	MouseCursorResizeAll  // (Unused by imgui functions)
	MouseCursorResizeNS   // When hovering over an horizontal border
	MouseCursorResizeEW   // When hovering over a vertical border or a column
	MouseCursorResizeNESW // When hovering over the bottom-left corner of a window
	MouseCursorResizeNWSE // When hovering over the bottom-right corner of a window
	MouseCursorHand       // (Unused by imgui functions. Use for e.g. hyperlinks)
	MouseCursorCount
)
