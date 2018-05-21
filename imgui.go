package imgui

// #include "imguiWrapper.h"
import "C"

// Version returns a version string e.g. "1.23".
func Version() string {
	return C.GoString(C.iggGetVersion())
}

// CurrentIO returns access to the ImGui communication struct for the currently active context.
func CurrentIO() IO {
	return IO{C.iggGetCurrentIO()}
}

// CurrentStyle returns the UI Style for the currently active context.
func CurrentStyle() Style {
	return Style(C.iggGetCurrentStyle())
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

// RenderedDrawData returns the created draw commands, which are valid after Render() and
// until the next call to NewFrame(). This is what you have to render.
func RenderedDrawData() DrawData {
	return DrawData(C.iggGetDrawData())
}

// ShowDemoWindow creates a demo/test window. Demonstrates most ImGui features.
// Call this to learn about the library! Try to make it always available in your application!
func ShowDemoWindow(open *bool) {
	openArg, openFin := wrapBool(open)
	defer openFin()
	C.iggShowDemoWindow(openArg)
}
