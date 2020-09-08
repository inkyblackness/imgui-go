package imgui

// #include "wrapper/Layout.h"
import "C"

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

// Separator is generally horizontal. Inside a menu bar or in horizontal layout mode, this becomes a vertical separator.
func Separator() {
	C.iggSeparator()
}

// SameLineV is between widgets or groups to layout them horizontally.
func SameLineV(posX float32, spacingW float32) {
	C.iggSameLine(C.float(posX), C.float(spacingW))
}

// SameLine calls SameLineV(0, -1).
func SameLine() {
	SameLineV(0, -1)
}

// Spacing adds vertical spacing.
func Spacing() {
	C.iggSpacing()
}

// Dummy adds a dummy item of given size.
func Dummy(size Vec2) {
	sizeArg, _ := size.wrapped()
	C.iggDummy(sizeArg)
}

// Indent moves content position toward the right by style.IndentSpacing.
func Indent() {
	C.iggIndent(0)
}

// Unindent moves content position back to the left by style.IndentSpacing.
func Unindent() {
	C.iggUnindent(0)
}

// IndentV moves content position toward the right, by style.IndentSpacing or indentW if not zero.
func IndentV(indentW float32) {
	C.iggIndent(C.float(indentW))
}

// UnindentV moves content position back to the left, by style.IndentSpacing or indentW if not zero.
func UnindentV(indentW float32) {
	C.iggUnindent(C.float(indentW))
}

// BeginGroup locks horizontal starting position + capture group bounding box into one "item";
// So you can use IsItemHovered() or layout primitives such as SameLine() on whole group, etc.
func BeginGroup() {
	C.iggBeginGroup()
}

// EndGroup must be called for each call to BeginGroup().
func EndGroup() {
	C.iggEndGroup()
}

// CursorPos returns the cursor position in window coordinates (relative to window position).
func CursorPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggCursorPos(valueArg)
	valueFin()
	return value
}

// CursorPosX returns the x-coordinate of the cursor position in window coordinates.
func CursorPosX() float32 {
	return float32(C.iggCursorPosX())
}

// CursorPosY returns the y-coordinate of the cursor position in window coordinates.
func CursorPosY() float32 {
	return float32(C.iggCursorPosY())
}

// CursorStartPos returns the initial cursor position in window coordinates.
func CursorStartPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggCursorStartPos(valueArg)
	valueFin()
	return value
}

// CursorScreenPos returns the cursor position in absolute screen coordinates.
func CursorScreenPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggCursorScreenPos(valueArg)
	valueFin()
	return value
}

// SetCursorPos sets the cursor relative to the current window.
func SetCursorPos(localPos Vec2) {
	localPosArg, _ := localPos.wrapped()
	C.iggSetCursorPos(localPosArg)
}

// SetCursorScreenPos sets the cursor position in absolute screen coordinates.
func SetCursorScreenPos(absPos Vec2) {
	absPosArg, _ := absPos.wrapped()
	C.iggSetCursorScreenPos(absPosArg)
}

// AlignTextToFramePadding vertically aligns upcoming text baseline to
// FramePadding.y so that it will align properly to regularly framed
// items. Call if you have text on a line before a framed item.
func AlignTextToFramePadding() {
	C.iggAlignTextToFramePadding()
}

// TextLineHeight returns ~ FontSize.
func TextLineHeight() float32 {
	return float32(C.iggGetTextLineHeight())
}

// TextLineHeightWithSpacing returns ~ FontSize + style.ItemSpacing.y (distance in pixels between 2 consecutive lines of text).
func TextLineHeightWithSpacing() float32 {
	return float32(C.iggGetTextLineHeightWithSpacing())
}

// FrameHeight returns the height of the current frame. This is equal to the
// font size plus the padding at the top and bottom.
func FrameHeight() float32 {
	return float32(C.iggGetFrameHeight())
}

// FrameHeightWithSpacing returns the height of the current frame with the item
// spacing added. This is equal to the font size plus the padding at the top
// and bottom, plus the value of style.ItemSpacing.y.
func FrameHeightWithSpacing() float32 {
	return float32(C.iggGetFrameHeightWithSpacing())
}
