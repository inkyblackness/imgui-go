package imgui

// #include "wrapper/DrawList.h"
import "C"
import (
	"unsafe"
)

// DrawList is a draw-command list.
// This is the low-level list of polygons that ImGui functions are filling.
// At the end of the frame, all command lists are passed to your render function for rendering.
//
// Each ImGui window contains its own DrawList. You can use GetWindowDrawList() to access
// the current window draw list and draw custom primitives.
//
// You can interleave normal ImGui calls and adding primitives to the current draw list.
//
// All positions are generally in pixel coordinates (top-left at (0,0), bottom-right at io.DisplaySize),
// however you are totally free to apply whatever transformation matrix to want to the data
// (if you apply such transformation you'll want to apply it to ClipRect as well)
//
// Important: Primitives are always added to the list and not culled (culling is done at
// higher-level by ImGui functions), if you use this API a lot consider coarse culling your drawn objects.
type DrawList uintptr

func (list DrawList) handle() C.IggDrawList {
	return C.IggDrawList(list)
}

// Commands returns the list of draw commands.
// Typically 1 command = 1 GPU draw call, unless the command is a callback.
func (list DrawList) Commands() []DrawCommand {
	count := int(C.iggDrawListGetCommandCount(list.handle()))
	commands := make([]DrawCommand, count)
	for i := 0; i < count; i++ {
		commands[i] = DrawCommand(C.iggDrawListGetCommand(list.handle(), C.int(i)))
	}
	return commands
}

// VertexBufferLayout returns the byte sizes necessary to select fields in a vertex buffer of a DrawList.
func VertexBufferLayout() (entrySize int, posOffset int, uvOffset int, colOffset int) {
	var entrySizeArg C.size_t
	var posOffsetArg C.size_t
	var uvOffsetArg C.size_t
	var colOffsetArg C.size_t
	C.iggGetVertexBufferLayout(&entrySizeArg, &posOffsetArg, &uvOffsetArg, &colOffsetArg)
	entrySize = int(entrySizeArg)
	posOffset = int(posOffsetArg)
	uvOffset = int(uvOffsetArg)
	colOffset = int(colOffsetArg)
	return
}

// VertexBuffer returns the handle information of the whole vertex buffer.
// Returned are the handle pointer and the total byte size.
// The buffer is a packed array of vertex entries, each consisting of a 2D position vector, a 2D UV vector,
// and a 4-byte color value. To determine the byte size and offset values, call VertexBufferLayout.
func (list DrawList) VertexBuffer() (unsafe.Pointer, int) {
	var data unsafe.Pointer
	var size C.int

	C.iggDrawListGetRawVertexBuffer(list.handle(), &data, &size)

	return data, int(size)
}

// IndexBufferLayout returns the byte size necessary to select fields in an index buffer of DrawList.
func IndexBufferLayout() (entrySize int) {
	var entrySizeArg C.size_t
	C.iggGetIndexBufferLayout(&entrySizeArg)
	entrySize = int(entrySizeArg)
	return
}

// IndexBuffer returns the handle information of the whole index buffer.
// Returned are the handle pointer and the total byte size.
// The buffer is a packed array of index entries, each consisting of an integer offset.
// To determine the byte size, call IndexBufferLayout.
func (list DrawList) IndexBuffer() (unsafe.Pointer, int) {
	var data unsafe.Pointer
	var size C.int

	C.iggDrawListGetRawIndexBuffer(list.handle(), &data, &size)

	return data, int(size)
}

// WindowDrawList returns the DrawList for the current window.
func WindowDrawList() DrawList {
	return DrawList(C.iggGetWindowDrawList())
}

// This is a list of DrawCornerFlags.
const (
	DrawCornerFlagsNone     = 0x0
	DrawCornerFlagsTopLeft  = 0x1
	DrawCornerFlagsTopRight = 0x2
	DrawCornerFlagsBotLeft  = 0x4
	DrawCornerFlagsBotRight = 0x8
	DrawCornerFlagsTop      = DrawCornerFlagsTopLeft | DrawCornerFlagsTopRight
	DrawCornerFlagsBot      = DrawCornerFlagsBotLeft | DrawCornerFlagsBotRight
	DrawCornerFlagsLeft     = DrawCornerFlagsTopLeft | DrawCornerFlagsBotLeft
	DrawCornerFlagsRight    = DrawCornerFlagsTopRight | DrawCornerFlagsBotRight
	DrawCornerFlagsAll      = 0xF
)

// AddLine call AddLineV with a thickness value of 1.0.
func (list DrawList) AddLine(p1 Vec2, p2 Vec2, col PackedColor) {
	list.AddLineV(p1, p2, col, 1.0)
}

// AddLineV adds a line to draw list, extending from point p1 to p2.
func (list DrawList) AddLineV(p1 Vec2, p2 Vec2, col PackedColor, thickness float32) {
	p1Arg, _ := p1.wrapped()
	p2Arg, _ := p2.wrapped()
	C.iggAddLine(list.handle(), p1Arg, p2Arg, C.IggPackedColor(col), C.float(thickness))
}

// AddRect calls AddRectV with rounding and thickness values of 1.0 and DrawCornerFlagsAll.
func (list DrawList) AddRect(min Vec2, max Vec2, col PackedColor) {
	list.AddRectV(min, max, col, 1.0, DrawCornerFlagsAll, 1.0)
}

// AddRectV adds a rectangle to draw list. min is the upper-left corner of the
// rectangle, and max is the lower right corner. rectangles with dimensions of
// 1 pixel are not rendered properly.
//
// drawCornerFlags indicate which corners of the rectangle are to be rounded.
func (list DrawList) AddRectV(min Vec2, max Vec2, col PackedColor, rounding float32, drawCornerFlags int, thickness float32) {
	minArg, _ := min.wrapped()
	maxArg, _ := max.wrapped()
	C.iggAddRect(list.handle(), minArg, maxArg, C.IggPackedColor(col), C.float(rounding), C.int(drawCornerFlags), C.float(thickness))
}

// AddRectFilled calls AddRectFilledV(min, max, col, 1.0, DrawCornerFlagsAll).
func (list DrawList) AddRectFilled(min Vec2, max Vec2, col PackedColor) {
	list.AddRectFilledV(min, max, col, 1.0, DrawCornerFlagsAll)
}

// AddRectFilledV adds a filled rectangle to the draw list. min is the
// upper-left corner of the rectangle, and max is the lower right corner.
// rectangles with dimensions of 1 pixel are not rendered properly.
func (list DrawList) AddRectFilledV(min Vec2, max Vec2, col PackedColor, rounding float32, drawCornerFlags int) {
	minArg, _ := min.wrapped()
	maxArg, _ := max.wrapped()
	C.iggAddRectFilled(list.handle(), minArg, maxArg, C.IggPackedColor(col), C.float(rounding), C.int(drawCornerFlags))
}

// AddCircleFilled calls AddCircleFilledV(center, radius, col, 12).
func (list DrawList) AddCircleFilled(center Vec2, radius float32, col PackedColor) {
	list.AddCircleFilledV(center, radius, col, 12)
}

// AddCircleFilledV adds a filled circle to the draw list. min is the
// upper-left corner of the rectangle, and max is the lower right corner.
func (list DrawList) AddCircleFilledV(center Vec2, radius float32, col PackedColor, numSegments int) {
	centerArg, _ := center.wrapped()
	C.iggAddCircleFilled(list.handle(), centerArg, C.float(radius), C.IggPackedColor(col), C.int(numSegments))
}

// AddCircle calls AddCircleV(center, radius, col, 12, 1.0).
func (list DrawList) AddCircle(center Vec2, radius float32, col PackedColor) {
	list.AddCircleV(center, radius, col, 12, 1.0)
}

// AddCircleV adds a unfilled circle to the draw list. min is the upper-left
// corner of the rectangle, and max is the lower right corner.
func (list DrawList) AddCircleV(center Vec2, radius float32, col PackedColor, numSegments int, thickness float32) {
	centerArg, _ := center.wrapped()
	C.iggAddCircle(list.handle(), centerArg, C.float(radius), C.IggPackedColor(col), C.int(numSegments), C.float(thickness))
}

// AddTriangle calls AddTriangleV(p1, p2, p3, col, 1.0).
func (list DrawList) AddTriangle(p1 Vec2, p2 Vec2, p3 Vec2, col PackedColor) {
	list.AddTriangleV(p1, p2, p3, col, 1.0)
}

// AddTriangleV adds an unfilled triangle of points p1, p2, p3 to the draw list.
func (list DrawList) AddTriangleV(p1 Vec2, p2 Vec2, p3 Vec2, col PackedColor, thickness float32) {
	p1Arg, _ := p1.wrapped()
	p2Arg, _ := p2.wrapped()
	p3Arg, _ := p3.wrapped()
	C.iggAddTriangle(list.handle(), p1Arg, p2Arg, p3Arg, C.IggPackedColor(col), C.float(thickness))
}

// AddTriangleFilled adds an filled triangle of points p1, p2, p3 to the draw list.
func (list DrawList) AddTriangleFilled(p1 Vec2, p2 Vec2, p3 Vec2, col PackedColor) {
	p1Arg, _ := p1.wrapped()
	p2Arg, _ := p2.wrapped()
	p3Arg, _ := p3.wrapped()
	C.iggAddTriangleFilled(list.handle(), p1Arg, p2Arg, p3Arg, C.IggPackedColor(col))
}
