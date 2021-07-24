package imgui

// #include "wrapper/DrawData.h"
import "C"
import "unsafe"

// RenderedDrawData returns the created draw commands, which are valid after Render() and
// until the next call to NewFrame(). This is what you have to render.
func RenderedDrawData() DrawData {
	return DrawData(C.iggGetDrawData())
}

// DrawData contains all draw data to render an ImGui frame.
type DrawData uintptr

func (data DrawData) handle() C.IggDrawData {
	return C.IggDrawData(data)
}

// Valid indicates whether the structure is usable.
// It is valid only after Render() is called and before the next NewFrame() is called.
func (data DrawData) Valid() bool {
	return (data.handle() != nil) && (C.iggDrawDataValid(data.handle()) != 0)
}

// CommandLists is an array of DrawList to render.
// The DrawList are owned by the context and only pointed to from here.
func (data DrawData) CommandLists() []DrawList {
	var handles unsafe.Pointer
	var count C.int

	C.iggDrawDataGetCommandLists(data.handle(), &handles, &count)
	list := make([]DrawList, int(count))
	for i := 0; i < int(count); i++ {
		list[i] = DrawList(*((*uintptr)(handles)))
		handles = unsafe.Pointer(uintptr(handles) + unsafe.Sizeof(handles)) // nolint: gas
	}

	return list
}

// DisplayPos returns the top-left position of the viewport to render.
// Use this as the top-left of the orthogonal projection matrix.
// For the main viewport this is equal to MainViewport().Pos().
// Usually {0, 0} in a single-viewport application.
func (data DrawData) DisplayPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggDrawDataDisplayPos(data.handle(), valueArg)
	valueFin()
	return value
}

// DisplaySize returns the size of the viewport to render.
// For the main viewport this is equal to MainViewport().Size().
// Usually set by IO.SetDisplaySize().
func (data DrawData) DisplaySize() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggDrawDataDisplaySize(data.handle(), valueArg)
	valueFin()
	return value
}

// FrameBufferScale returns the amount of pixels for each unit of DisplaySize().
// Generally (1,1) on normal display, (2,2) on OSX with Retina display.
// See also IO.DisplayFrameBufferScale().
func (data DrawData) FrameBufferScale() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggDrawDataFrameBufferScale(data.handle(), valueArg)
	valueFin()
	return value
}

// ScaleClipRects is a helper to scale the ClipRect field of each DrawCmd.
// Use if your final output buffer is at a different scale than ImGui expects,
// or if there is a difference between your window resolution and framebuffer resolution.
func (data DrawData) ScaleClipRects(scale Vec2) {
	scaleArg, _ := scale.wrapped()
	C.iggDrawDataScaleClipRects(data.handle(), scaleArg)
}
