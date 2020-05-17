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

// MouseCursor returns desired cursor type, reset in imgui.NewFrame(), this is updated during the frame.
// Valid before Render(). If you use software rendering by setting io.MouseDrawCursor ImGui will render those for you.
func MouseCursor() int {
	return int(C.iggGetMouseCursor())
}

// SetMouseCursor sets desired cursor type.
func SetMouseCursor(cursor int) {
	C.iggSetMouseCursor(C.int(cursor))
}
