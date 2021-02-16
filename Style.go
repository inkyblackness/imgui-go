package imgui

// #include "wrapper/Style.h"
import "C"

// StyleVar identifies a style variable in the UI style.
type StyleVar int

const (
	// StyleVarAlpha is a float
	StyleVarAlpha StyleVar = 0
	// StyleVarWindowPadding is a Vec2
	StyleVarWindowPadding StyleVar = 1
	// StyleVarWindowRounding is a float
	StyleVarWindowRounding StyleVar = 2
	// StyleVarWindowBorderSize is a float
	StyleVarWindowBorderSize StyleVar = 3
	// StyleVarWindowMinSize is a Vec2
	StyleVarWindowMinSize StyleVar = 4
	// StyleVarWindowTitleAlign is a Vec2
	StyleVarWindowTitleAlign StyleVar = 5
	// StyleVarChildRounding is a float
	StyleVarChildRounding StyleVar = 6
	// StyleVarChildBorderSize is a float
	StyleVarChildBorderSize StyleVar = 7
	// StyleVarPopupRounding is a float
	StyleVarPopupRounding StyleVar = 8
	// StyleVarPopupBorderSize is a float
	StyleVarPopupBorderSize StyleVar = 9
	// StyleVarFramePadding is a Vec2
	StyleVarFramePadding StyleVar = 10
	// StyleVarFrameRounding is a float
	StyleVarFrameRounding StyleVar = 11
	// StyleVarFrameBorderSize is a float
	StyleVarFrameBorderSize StyleVar = 12
	// StyleVarItemSpacing is a Vec2
	StyleVarItemSpacing StyleVar = 13
	// StyleVarItemInnerSpacing is a Vec2
	StyleVarItemInnerSpacing StyleVar = 14
	// StyleVarIndentSpacing is a float
	StyleVarIndentSpacing StyleVar = 15
	// StyleVarCellPadding is a Vec2
	StyleVarCellPadding StyleVar = 16
	// StyleVarScrollbarSize is a float
	StyleVarScrollbarSize StyleVar = 17
	// StyleVarScrollbarRounding is a float
	StyleVarScrollbarRounding StyleVar = 18
	// StyleVarGrabMinSize is a float
	StyleVarGrabMinSize StyleVar = 19
	// StyleVarGrabRounding is a float
	StyleVarGrabRounding StyleVar = 20
	// StyleVarTabRounding is a float
	StyleVarTabRounding StyleVar = 21
	// StyleVarButtonTextAlign is a Vec2
	StyleVarButtonTextAlign StyleVar = 22
	// StyleVarSelectableTextAlign is a Vec2
	StyleVarSelectableTextAlign StyleVar = 23
)

// StyleColor identifies a color in the UI style.
type StyleColor int

// This is the list of StyleColor identifier.
const (
	StyleColorText                  StyleColor = 0
	StyleColorTextDisabled          StyleColor = 1
	StyleColorWindowBg              StyleColor = 2
	StyleColorChildBg               StyleColor = 3
	StyleColorPopupBg               StyleColor = 4
	StyleColorBorder                StyleColor = 5
	StyleColorBorderShadow          StyleColor = 6
	StyleColorFrameBg               StyleColor = 7
	StyleColorFrameBgHovered        StyleColor = 8
	StyleColorFrameBgActive         StyleColor = 9
	StyleColorTitleBg               StyleColor = 10
	StyleColorTitleBgActive         StyleColor = 11
	StyleColorTitleBgCollapsed      StyleColor = 12
	StyleColorMenuBarBg             StyleColor = 13
	StyleColorScrollbarBg           StyleColor = 14
	StyleColorScrollbarGrab         StyleColor = 15
	StyleColorScrollbarGrabHovered  StyleColor = 16
	StyleColorScrollbarGrabActive   StyleColor = 17
	StyleColorCheckMark             StyleColor = 18
	StyleColorSliderGrab            StyleColor = 19
	StyleColorSliderGrabActive      StyleColor = 20
	StyleColorButton                StyleColor = 21
	StyleColorButtonHovered         StyleColor = 22
	StyleColorButtonActive          StyleColor = 23
	StyleColorHeader                StyleColor = 24
	StyleColorHeaderHovered         StyleColor = 25
	StyleColorHeaderActive          StyleColor = 26
	StyleColorSeparator             StyleColor = 27
	StyleColorSeparatorHovered      StyleColor = 28
	StyleColorSeparatorActive       StyleColor = 29
	StyleColorResizeGrip            StyleColor = 30
	StyleColorResizeGripHovered     StyleColor = 31
	StyleColorResizeGripActive      StyleColor = 32
	StyleColorTab                   StyleColor = 33
	StyleColorTabHovered            StyleColor = 34
	StyleColorTabActive             StyleColor = 35
	StyleColorTabUnfocused          StyleColor = 36
	StyleColorTabUnfocusedActive    StyleColor = 37
	StyleColorPlotLines             StyleColor = 38
	StyleColorPlotLinesHovered      StyleColor = 39
	StyleColorPlotHistogram         StyleColor = 40
	StyleColorPlotHistogramHovered  StyleColor = 41
	StyleColorTableHeaderBg         StyleColor = 42 // Table header background
	StyleColorTableBorderStrong     StyleColor = 43 // Table outer and header borders (prefer using Alpha=1.0 here)
	StyleColorTableBorderLight      StyleColor = 44 // Table inner borders (prefer using Alpha=1.0 here)
	StyleColorTableRowBg            StyleColor = 45 // Table row background (even rows)
	StyleColorTableRowBgAlt         StyleColor = 46 // Table row background (odd rows)
	StyleColorTextSelectedBg        StyleColor = 47
	StyleColorDragDropTarget        StyleColor = 48
	StyleColorNavHighlight          StyleColor = 49 // Gamepad/keyboard: current highlighted item
	StyleColorNavWindowingHighlight StyleColor = 50 // Highlight window when using CTRL+TAB
	StyleColorNavWindowingDarkening StyleColor = 51 // Darken/colorize entire screen behind the CTRL+TAB window list, when active
	StyleColorModalWindowDarkening  StyleColor = 52 // Darken/colorize entire screen behind a modal window, when one is active
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
func PushStyleColor(id StyleColor, color Vec4) {
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
func PushStyleVarFloat(id StyleVar, value float32) {
	C.iggPushStyleVarFloat(C.int(id), C.float(value))
}

// PushStyleVarVec2 pushes a Vec2 value on the stack to temporarily modify a style variable.
func PushStyleVarVec2(id StyleVar, value Vec2) {
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
func (style Style) SetColor(id StyleColor, value Vec4) {
	valueArg, _ := value.wrapped()
	C.iggStyleSetColor(style.handle(), C.int(id), valueArg)
}

// Color gets a color value from the UI style.
func (style Style) Color(id StyleColor) Vec4 {
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
