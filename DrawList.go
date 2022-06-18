package imgui

// #include "wrapper/DrawList.h"
import "C"
import (
	"image/color"
	"unsafe"
)

// DrawList is a draw-command list.
// This is the low-level list of polygons that ImGui functions are filling.
// At the end of the frame, all command lists are passed to your render function for rendering.
//
// Each ImGui window contains its own DrawList. You can use WindowDrawList() to access
// the current window draw list and draw custom primitives.
//
// You can interleave normal ImGui calls and adding primitives to the current draw list.
//
// In single viewport mode, top-left is == MainViewport().Pos() (generally 0,0), bottom-right is == MainViewport().Pos()+Size (generally io.DisplaySize).
// You are totally free to apply whatever transformation matrix to want to the data
// (depending on the use of the transformation you may want to apply it to ClipRect as well!)
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

// ForegroundDrawList returns the DrawList for over all windows.
func ForegroundDrawList() DrawList {
	return DrawList(C.iggGetForegroundDrawList())
}

// BackgroundDrawList returns the DrawList for the background behind all windows.
func BackgroundDrawList() DrawList {
	return DrawList(C.iggGetBackgroundDrawList())
}

// DrawFlags for DrawList.AddRectV, etc.
type DrawFlags int

const (
	// DrawFlagsNone specified the default behaviour.
	DrawFlagsNone DrawFlags = 0
	// DrawFlagsClosed for PathStroke(), AddPolyline(): specify that shape should be closed.
	DrawFlagsClosed DrawFlags = 1 << 0
	// DrawFlagsRoundCornersTopLeft for AddRect(), AddRectFilled(), PathRect(): enable rounding top-left corner only (when rounding > 0.0f, we default to all corners).
	DrawFlagsRoundCornersTopLeft DrawFlags = 1 << 4
	// DrawFlagsRoundCornersTopRight for AddRect(), AddRectFilled(), PathRect(): enable rounding top-right corner only (when rounding > 0.0f, we default to all corners).
	DrawFlagsRoundCornersTopRight DrawFlags = 1 << 5
	// DrawFlagsRoundCornersBottomLeft for AddRect(), AddRectFilled(), PathRect(): enable rounding bottom-left corner only (when rounding > 0.0f, we default to all corners).
	DrawFlagsRoundCornersBottomLeft DrawFlags = 1 << 6
	// DrawFlagsRoundCornersBottomRight for AddRect(), AddRectFilled(), PathRect(): enable rounding bottom-right corner only (when rounding > 0.0f, we default to all corners).
	DrawFlagsRoundCornersBottomRight DrawFlags = 1 << 7
	// DrawFlagsRoundCornersNone for AddRect(), AddRectFilled(), PathRect(): disable rounding on all corners (when rounding > 0.0f).
	DrawFlagsRoundCornersNone DrawFlags = 1 << 8
	// DrawFlagsRoundCornersTop for AddRect(), AddRectFilled(), PathRect(): enable rounding top corners only (when rounding > 0.0f, we default to all corners).
	DrawFlagsRoundCornersTop DrawFlags = DrawFlagsRoundCornersTopLeft | DrawFlagsRoundCornersTopRight
	// DrawFlagsRoundCornersBottom for AddRect(), AddRectFilled(), PathRect(): enable rounding bottom corners only (when rounding > 0.0f, we default to all corners).
	DrawFlagsRoundCornersBottom DrawFlags = DrawFlagsRoundCornersBottomLeft | DrawFlagsRoundCornersBottomRight
	// DrawFlagsRoundCornersLeft for AddRect(), AddRectFilled(), PathRect(): enable rounding left corners only (when rounding > 0.0f, we default to all corners).
	DrawFlagsRoundCornersLeft DrawFlags = DrawFlagsRoundCornersBottomLeft | DrawFlagsRoundCornersTopLeft
	// DrawFlagsRoundCornersRight for AddRect(), AddRectFilled(), PathRect(): enable rounding right corners only (when rounding > 0.0f, we default to all corners).
	DrawFlagsRoundCornersRight DrawFlags = DrawFlagsRoundCornersBottomRight | DrawFlagsRoundCornersTopRight
	// DrawFlagsRoundCornersAll  for AddRect(), AddRectFilled(), PathRect(): enable rounding for all corners.
	DrawFlagsRoundCornersAll DrawFlags = DrawFlagsRoundCornersTopLeft | DrawFlagsRoundCornersTopRight | DrawFlagsRoundCornersBottomLeft | DrawFlagsRoundCornersBottomRight
	// DrawFlagsRoundCornersDefault default to ALL corners if none of the RoundCornersXX flags are specified.
	DrawFlagsRoundCornersDefault DrawFlags = DrawFlagsRoundCornersAll
	// DrawFlagsRoundCornersMask is the bitmask containing the corner flags.
	DrawFlagsRoundCornersMask DrawFlags = DrawFlagsRoundCornersAll | DrawFlagsRoundCornersNone
)

// DrawCornerFlags is replaced by DrawFlags and will be removed in v5.
// Deprecated: Use DrawFlags.
type DrawCornerFlags = DrawFlags

const (
	// DrawCornerFlagsNone specifies the default behaviour.
	// Deprecated: Use DrawFlagsRoundCornersNone
	DrawCornerFlagsNone DrawCornerFlags = DrawFlagsRoundCornersNone
	// DrawCornerFlagsTopLeft draw corner in the top left.
	// Deprecated: Use DrawFlagsRoundCornersTopLeft
	DrawCornerFlagsTopLeft DrawCornerFlags = DrawFlagsRoundCornersTopLeft
	// DrawCornerFlagsTopRight draw corner in the top right.
	// Deprecated: Use DrawFlagsRoundCornersTopRight
	DrawCornerFlagsTopRight DrawCornerFlags = DrawFlagsRoundCornersTopRight
	// DrawCornerFlagsBotLeft draw corner in the bottom left.
	// Deprecated: Use DrawFlagsRoundCornersBottomLeft
	DrawCornerFlagsBotLeft DrawCornerFlags = DrawFlagsRoundCornersBottomLeft
	// DrawCornerFlagsBotRight draw corner in the bottom right.
	// Deprecated: Use DrawFlagsRoundCornersBottomRight
	DrawCornerFlagsBotRight DrawCornerFlags = DrawFlagsRoundCornersBottomRight
	// DrawCornerFlagsAll draws all corners.
	// Deprecated: Use DrawFlagsRoundCornersAll
	DrawCornerFlagsAll DrawCornerFlags = DrawFlagsRoundCornersAll
	// DrawCornerFlagsTop draw corners at the top of the area.
	// Deprecated: Use DrawFlagsRoundCornersTop
	DrawCornerFlagsTop DrawCornerFlags = DrawCornerFlagsTopLeft | DrawCornerFlagsTopRight
	// DrawCornerFlagsBot draw corners at the bottom of the area.
	// Deprecated: Use DrawFlagsRoundCornersBottom
	DrawCornerFlagsBot DrawCornerFlags = DrawCornerFlagsBotLeft | DrawCornerFlagsBotRight
	// DrawCornerFlagsLeft draw corners on the left of the area.
	// Deprecated: Use DrawFlagsRoundCornersLeft
	DrawCornerFlagsLeft DrawCornerFlags = DrawCornerFlagsTopLeft | DrawCornerFlagsBotLeft
	// DrawCornerFlagsRight draw corners on the rigth of the area.
	// Deprecated: Use DrawFlagsRoundCornersRight
	DrawCornerFlagsRight DrawCornerFlags = DrawCornerFlagsTopRight | DrawCornerFlagsBotRight
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
func (list DrawList) AddRectV(min Vec2, max Vec2, col PackedColor, rounding float32, flags DrawFlags, thickness float32) {
	minArg, _ := min.wrapped()
	maxArg, _ := max.wrapped()
	C.iggAddRect(list.handle(), minArg, maxArg, C.IggPackedColor(col), C.float(rounding), C.int(flags), C.float(thickness))
}

// AddRectFilled calls AddRectFilledV(min, max, col, 1.0, DrawCornerFlagsAll).
func (list DrawList) AddRectFilled(min Vec2, max Vec2, col PackedColor) {
	list.AddRectFilledV(min, max, col, 1.0, DrawCornerFlagsAll)
}

// AddRectFilledMultiColor adds a multicolor filled rectangle to the draw list.
// min is the upper-left corner of the rectangle, and max is the lower right
// corner. rectangles with dimensions of 1 pixel are not rendered properly.
func (list DrawList) AddRectFilledMultiColor(min Vec2, max Vec2, colUpperLeft, colUpperRight, colBottomRight, colBottomLeft PackedColor) {
	minArg, _ := min.wrapped()
	maxArg, _ := max.wrapped()
	C.iggAddRectFilledMultiColor(list.handle(), minArg, maxArg, C.IggPackedColor(colUpperLeft), C.IggPackedColor(colUpperRight), C.IggPackedColor(colBottomRight), C.IggPackedColor(colBottomLeft))
}

// AddRectFilledV adds a filled rectangle to the draw list. min is the
// upper-left corner of the rectangle, and max is the lower right corner.
// rectangles with dimensions of 1 pixel are not rendered properly.
func (list DrawList) AddRectFilledV(min Vec2, max Vec2, col PackedColor, rounding float32, flags DrawFlags) {
	minArg, _ := min.wrapped()
	maxArg, _ := max.wrapped()
	C.iggAddRectFilled(list.handle(), minArg, maxArg, C.IggPackedColor(col), C.float(rounding), C.int(flags))
}

// AddCircleFilled calls AddCircleFilledV(center, radius, col, 0).
func (list DrawList) AddCircleFilled(center Vec2, radius float32, col PackedColor) {
	list.AddCircleFilledV(center, radius, col, 0)
}

// AddCircleFilledV adds a filled circle to the draw list. min is the
// upper-left corner of the rectangle, and max is the lower right corner.
func (list DrawList) AddCircleFilledV(center Vec2, radius float32, col PackedColor, numSegments int) {
	centerArg, _ := center.wrapped()
	C.iggAddCircleFilled(list.handle(), centerArg, C.float(radius), C.IggPackedColor(col), C.int(numSegments))
}

// AddCircle calls AddCircleV(center, radius, col, 0, 1.0).
func (list DrawList) AddCircle(center Vec2, radius float32, col PackedColor) {
	list.AddCircleV(center, radius, col, 0, 1.0)
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

// AddText adds a text in specified color at given position pos.
func (list DrawList) AddText(pos Vec2, col PackedColor, text string) {
	CString := newStringBuffer(text)
	defer CString.free()
	posArg, _ := pos.wrapped()
	C.iggAddText(list.handle(), posArg, C.IggPackedColor(col), (*C.char)(CString.ptr), C.int(CString.size)-1)
}

// AddImage calls AddImageV(textureId, posMin, posMax, Vec2{0,0}, Vec2{1,1}, Packed(color.White)).
func (list DrawList) AddImage(textureID TextureID, posMin Vec2, posMax Vec2) {
	list.AddImageV(textureID, posMin, posMax, Vec2{X: 0, Y: 0}, Vec2{X: 1, Y: 1}, Packed(color.White))
}

// AddImageV adds an image based on given texture ID.
func (list DrawList) AddImageV(textureID TextureID, posMin Vec2, posMax Vec2, uvMin Vec2, uvMax Vec2, tintCol PackedColor) {
	posMinArg, _ := posMin.wrapped()
	posMaxArg, _ := posMax.wrapped()
	uvMinArg, _ := uvMin.wrapped()
	uvMaxArg, _ := uvMax.wrapped()
	C.iggAddImage(list.handle(), C.IggTextureID(textureID), posMinArg, posMaxArg, uvMinArg, uvMaxArg, C.IggPackedColor(tintCol))
}

// PushClipRect performs render-level scissoring.
// It calls PushClipRectV(min, max, false).
func (list DrawList) PushClipRect(min, max Vec2) {
	list.PushClipRectV(min, max, false)
}

// PushClipRectV performs render-level scissoring.
//
// This is passed down to your render function but not used for CPU-side coarse
// clipping. Prefer using higher-level imgui.PushClipRect() to affect logic
// (hit-testing and widget culling).
func (list DrawList) PushClipRectV(min, max Vec2, intersectWithCurrentClipRect bool) {
	minArg, _ := min.wrapped()
	maxArg, _ := max.wrapped()
	C.iggPushClipRect(list.handle(), minArg, maxArg, castBool(intersectWithCurrentClipRect))
}

// PopClipRect removes the current clip rect and returns to the previous one.
func (list DrawList) PopClipRect() {
	C.iggPopClipRect(list.handle())
}
