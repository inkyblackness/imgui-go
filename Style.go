package imgui

// #include "wrapper/Style.h"
import "C"

// StyleVarID identifies a style variable in the UI style.
type StyleVarID int

const (
	// StyleVarAlpha is a float.
	StyleVarAlpha StyleVarID = 0
	// StyleVarWindowPadding is a Vec2.
	StyleVarWindowPadding StyleVarID = 1
	// StyleVarWindowRounding is a float.
	StyleVarWindowRounding StyleVarID = 2
	// StyleVarWindowBorderSize is a float.
	StyleVarWindowBorderSize StyleVarID = 3
	// StyleVarWindowMinSize is a Vec2.
	StyleVarWindowMinSize StyleVarID = 4
	// StyleVarWindowTitleAlign is a Vec2.
	StyleVarWindowTitleAlign StyleVarID = 5
	// StyleVarChildRounding is a float.
	StyleVarChildRounding StyleVarID = 6
	// StyleVarChildBorderSize is a float.
	StyleVarChildBorderSize StyleVarID = 7
	// StyleVarPopupRounding is a float.
	StyleVarPopupRounding StyleVarID = 8
	// StyleVarPopupBorderSize is a float.
	StyleVarPopupBorderSize StyleVarID = 9
	// StyleVarFramePadding is a Vec2.
	StyleVarFramePadding StyleVarID = 10
	// StyleVarFrameRounding is a float.
	StyleVarFrameRounding StyleVarID = 11
	// StyleVarFrameBorderSize is a float.
	StyleVarFrameBorderSize StyleVarID = 12
	// StyleVarItemSpacing is a Vec2.
	StyleVarItemSpacing StyleVarID = 13
	// StyleVarItemInnerSpacing is a Vec2.
	StyleVarItemInnerSpacing StyleVarID = 14
	// StyleVarIndentSpacing is a float.
	StyleVarIndentSpacing StyleVarID = 15
	// StyleVarCellPadding is a Vec2.
	StyleVarCellPadding StyleVarID = 16
	// StyleVarScrollbarSize is a float.
	StyleVarScrollbarSize StyleVarID = 17
	// StyleVarScrollbarRounding is a float.
	StyleVarScrollbarRounding StyleVarID = 18
	// StyleVarGrabMinSize is a float.
	StyleVarGrabMinSize StyleVarID = 19
	// StyleVarGrabRounding is a float.
	StyleVarGrabRounding StyleVarID = 20
	// StyleVarTabRounding is a float.
	StyleVarTabRounding StyleVarID = 21
	// StyleVarButtonTextAlign is a Vec2.
	StyleVarButtonTextAlign StyleVarID = 22
	// StyleVarSelectableTextAlign is a Vec2.
	StyleVarSelectableTextAlign StyleVarID = 23
)

// StyleColorID identifies a color in the UI style.
type StyleColorID int

// This is the list of StyleColor identifier.
const (
	StyleColorText                  StyleColorID = 0
	StyleColorTextDisabled          StyleColorID = 1
	StyleColorWindowBg              StyleColorID = 2
	StyleColorChildBg               StyleColorID = 3
	StyleColorPopupBg               StyleColorID = 4
	StyleColorBorder                StyleColorID = 5
	StyleColorBorderShadow          StyleColorID = 6
	StyleColorFrameBg               StyleColorID = 7
	StyleColorFrameBgHovered        StyleColorID = 8
	StyleColorFrameBgActive         StyleColorID = 9
	StyleColorTitleBg               StyleColorID = 10
	StyleColorTitleBgActive         StyleColorID = 11
	StyleColorTitleBgCollapsed      StyleColorID = 12
	StyleColorMenuBarBg             StyleColorID = 13
	StyleColorScrollbarBg           StyleColorID = 14
	StyleColorScrollbarGrab         StyleColorID = 15
	StyleColorScrollbarGrabHovered  StyleColorID = 16
	StyleColorScrollbarGrabActive   StyleColorID = 17
	StyleColorCheckMark             StyleColorID = 18
	StyleColorSliderGrab            StyleColorID = 19
	StyleColorSliderGrabActive      StyleColorID = 20
	StyleColorButton                StyleColorID = 21
	StyleColorButtonHovered         StyleColorID = 22
	StyleColorButtonActive          StyleColorID = 23
	StyleColorHeader                StyleColorID = 24
	StyleColorHeaderHovered         StyleColorID = 25
	StyleColorHeaderActive          StyleColorID = 26
	StyleColorSeparator             StyleColorID = 27
	StyleColorSeparatorHovered      StyleColorID = 28
	StyleColorSeparatorActive       StyleColorID = 29
	StyleColorResizeGrip            StyleColorID = 30
	StyleColorResizeGripHovered     StyleColorID = 31
	StyleColorResizeGripActive      StyleColorID = 32
	StyleColorTab                   StyleColorID = 33
	StyleColorTabHovered            StyleColorID = 34
	StyleColorTabActive             StyleColorID = 35
	StyleColorTabUnfocused          StyleColorID = 36
	StyleColorTabUnfocusedActive    StyleColorID = 37
	StyleColorPlotLines             StyleColorID = 38
	StyleColorPlotLinesHovered      StyleColorID = 39
	StyleColorPlotHistogram         StyleColorID = 40
	StyleColorPlotHistogramHovered  StyleColorID = 41
	StyleColorTableHeaderBg         StyleColorID = 42 // Table header background
	StyleColorTableBorderStrong     StyleColorID = 43 // Table outer and header borders (prefer using Alpha=1.0 here)
	StyleColorTableBorderLight      StyleColorID = 44 // Table inner borders (prefer using Alpha=1.0 here)
	StyleColorTableRowBg            StyleColorID = 45 // Table row background (even rows)
	StyleColorTableRowBgAlt         StyleColorID = 46 // Table row background (odd rows)
	StyleColorTextSelectedBg        StyleColorID = 47
	StyleColorDragDropTarget        StyleColorID = 48
	StyleColorNavHighlight          StyleColorID = 49 // Gamepad/keyboard: current highlighted item
	StyleColorNavWindowingHighlight StyleColorID = 50 // Highlight window when using CTRL+TAB
	StyleColorNavWindowingDarkening StyleColorID = 51 // Darken/colorize entire screen behind the CTRL+TAB window list, when active
	StyleColorModalWindowDarkening  StyleColorID = 52 // Darken/colorize entire screen behind a modal window, when one is active
)

// Style describes the overall graphical representation of the user interface.
type Style uintptr

// CurrentStyle returns the UI Style for the currently active context.
func CurrentStyle() Style {
	return Style(C.iggGetCurrentStyle())
}

// StyleColorsDark sets the new, recommended style (default).
func StyleColorsDark() {
	C.iggStyleColorsDark()
}

// StyleColorsClassic sets the classic style.
func StyleColorsClassic() {
	C.iggStyleColorsClassic()
}

// StyleColorsLight sets the light style, best used with borders and a custom, thicker font.
func StyleColorsLight() {
	C.iggStyleColorsLight()
}

// PushStyleColor pushes the current style color for given ID on a stack and sets the given one.
// To revert to the previous color, call PopStyleColor().
func PushStyleColor(id StyleColorID, color Vec4) {
	colorArg, _ := color.wrapped()
	C.iggPushStyleColor(C.int(id), colorArg)
}

// PopStyleColorV reverts the given amount of style color changes.
func PopStyleColorV(count int) {
	C.iggPopStyleColor(C.int(count))
}

// PopStyleColor calls PopStyleColorV(1).
func PopStyleColor() {
	PopStyleColorV(1)
}

// PushStyleVarFloat pushes a float value on the stack to temporarily modify a style variable.
func PushStyleVarFloat(id StyleVarID, value float32) {
	C.iggPushStyleVarFloat(C.int(id), C.float(value))
}

// PushStyleVarVec2 pushes a Vec2 value on the stack to temporarily modify a style variable.
func PushStyleVarVec2(id StyleVarID, value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggPushStyleVarVec2(C.int(id), valueArg)
}

// PopStyleVarV reverts the given amount of style variable changes.
func PopStyleVarV(count int) {
	C.iggPopStyleVar(C.int(count))
}

// PopStyleVar calls PopStyleVarV(1).
func PopStyleVar() {
	PopStyleVarV(1)
}

func (style Style) handle() C.IggGuiStyle {
	return C.IggGuiStyle(style)
}

// ItemInnerSpacing is the horizontal and vertical spacing between elements of
// a composed widget (e.g. a slider and its label).
func (style Style) ItemInnerSpacing() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetItemInnerSpacing(style.handle(), valueArg)
	valueFin()
	return value
}

// ItemSpacing returns horizontal and vertical spacing between widgets or lines.
func (style Style) ItemSpacing() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetItemSpacing(style.handle(), valueArg)
	valueFin()
	return value
}

// FramePadding is the padding within a framed rectangle (used by most widgets).
func (style Style) FramePadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetFramePadding(style.handle(), valueArg)
	valueFin()
	return value
}

// WindowPadding is the padding within a window.
func (style Style) WindowPadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetWindowPadding(style.handle(), valueArg)
	valueFin()
	return value
}

// CellPadding is the padding within a table cell.
func (style Style) CellPadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetCellPadding(style.handle(), valueArg)
	valueFin()
	return value
}

// SetColor sets a color value of the UI style.
func (style Style) SetColor(id StyleColorID, value Vec4) {
	valueArg, _ := value.wrapped()
	C.iggStyleSetColor(style.handle(), C.int(id), valueArg)
}

// Color gets a color value from the UI style.
func (style Style) Color(id StyleColorID) Vec4 {
	var value Vec4
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetColor(style.handle(), C.int(id), valueArg)
	valueFin()
	return value
}

// ScaleAllSizes applies a scaling factor to all sizes.
// To scale your entire UI (e.g. if you want your app to use High DPI or generally be DPI aware) you may use this helper function.
// Scaling the fonts is done separately and is up to you.
//
// Important: This operation is lossy because all sizes are rounded to integer.
// If you need to change your scale multiples, call this over a freshly initialized style rather than scaling multiple times.
func (style Style) ScaleAllSizes(scale float32) {
	C.iggStyleScaleAllSizes(style.handle(), C.float(scale))
}

// SetTouchExtraPadding expand reactive bounding box for touch-based system where touch position is not accurate enough.
func (style Style) SetTouchExtraPadding(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggSetTouchExtraPadding(style.handle(), valueArg)
}
