package imgui

// #include "wrapper/Misc.h"
import "C"

// Version returns a version string e.g. "1.23".
func Version() string {
	return C.GoString(C.iggGetVersion())
}

// NewFrame starts a new ImGui frame, you can submit any command from this point until Render()/EndFrame().
func NewFrame() {
	C.iggNewFrame()
}

// Render ends the ImGui frame, finalize the draw data.
// After this method, call RenderedDrawData to retrieve the draw commands and execute them.
func Render() {
	C.iggRender()
}

// EndFrame ends the ImGui frame. Automatically called by Render(), so most likely don't need to ever
// call that yourself directly. If you don't need to render you may call EndFrame() but you'll have
// wasted CPU already. If you don't need to render, better to not create any imgui windows instead!
func EndFrame() {
	C.iggEndFrame()
}

// PushID pushes the given identifier into the ID stack. IDs are hash of the entire stack!
func PushID(id string) {
	idArg, idFin := wrapString(id)
	defer idFin()
	C.iggPushID(idArg)
}

// PushIDInt pushes the given identifier into the ID stack. IDs are hash of the entire stack!
func PushIDInt(id int) {
	C.iggPushIDInt(C.int(id))
}

// PopID removes the last pushed identifier from the ID stack.
func PopID() {
	C.iggPopID()
}

// Enumeration for MouseCursor()
// User code may request binding to display given cursor by calling SetMouseCursor(),
// which is why we have some cursors that are marked unused here.
const (
	// MouseCursorNone no mouse cursor
	MouseCursorNone = -1
	// MouseCursorArrow standard arrow mouse cursor
	MouseCursorArrow = 0
	// MouseCursorTextInput when hovering over InputText, etc.
	MouseCursorTextInput = 1
	// MouseCursorResizeAll (Unused by imgui functions)
	MouseCursorResizeAll = 2
	// MouseCursorResizeNS when hovering over an horizontal border
	MouseCursorResizeNS = 3
	// MouseCursorResizeEW when hovering over a vertical border or a column
	MouseCursorResizeEW = 4
	// MouseCursorResizeNESW when hovering over the bottom-left corner of a window
	MouseCursorResizeNESW = 5
	// MouseCursorResizeNWSE when hovering over the bottom-right corner of a window
	MouseCursorResizeNWSE = 6
	// MouseCursorHand (Unused by imgui functions. Use for e.g. hyperlinks)
	MouseCursorHand  = 7
	MouseCursorCount = 8
)

// MouseCursor returns desired cursor type, reset in imgui.NewFrame(), this is updated during the frame.
// Valid before Render(). If you use software rendering by setting io.MouseDrawCursor ImGui will render those for you.
func MouseCursor() int {
	return int(C.iggGetMouseCursor())
}

// SetMouseCursor sets desired cursor type.
func SetMouseCursor(cursor int) {
	C.iggSetMouseCursor(C.int(cursor))
}
