package imgui

// #include "wrapper/Style.h"
import "C"

// StyleVarID identifies a style variable in the UI style.
type StyleVarID int

const (
	// StyleVarAlpha is a float.
	StyleVarAlpha StyleVarID = iota
	// StyleVarDisabledAlpha is a float.
	StyleVarDisabledAlpha
	// StyleVarWindowPadding is a Vec2.
	StyleVarWindowPadding
	// StyleVarWindowRounding is a float.
	StyleVarWindowRounding
	// StyleVarWindowBorderSize is a float.
	StyleVarWindowBorderSize
	// StyleVarWindowMinSize is a Vec2.
	StyleVarWindowMinSize
	// StyleVarWindowTitleAlign is a Vec2.
	StyleVarWindowTitleAlign
	// StyleVarChildRounding is a float.
	StyleVarChildRounding
	// StyleVarChildBorderSize is a float.
	StyleVarChildBorderSize
	// StyleVarPopupRounding is a float.
	StyleVarPopupRounding
	// StyleVarPopupBorderSize is a float.
	StyleVarPopupBorderSize
	// StyleVarFramePadding is a Vec2.
	StyleVarFramePadding
	// StyleVarFrameRounding is a float.
	StyleVarFrameRounding
	// StyleVarFrameBorderSize is a float.
	StyleVarFrameBorderSize
	// StyleVarItemSpacing is a Vec2.
	StyleVarItemSpacing
	// StyleVarItemInnerSpacing is a Vec2.
	StyleVarItemInnerSpacing
	// StyleVarIndentSpacing is a float.
	StyleVarIndentSpacing
	// StyleVarCellPadding is a Vec2.
	StyleVarCellPadding
	// StyleVarScrollbarSize is a float.
	StyleVarScrollbarSize
	// StyleVarScrollbarRounding is a float.
	StyleVarScrollbarRounding
	// StyleVarGrabMinSize is a float.
	StyleVarGrabMinSize
	// StyleVarGrabRounding is a float.
	StyleVarGrabRounding
	// StyleVarTabRounding is a float.
	StyleVarTabRounding
	// StyleVarButtonTextAlign is a Vec2.
	StyleVarButtonTextAlign
	// StyleVarSelectableTextAlign is a Vec2.
	StyleVarSelectableTextAlign
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

// Dir is a cardinal direction.
type Dir int

// This is the list of Dir identifier.
const (
	DirNone  Dir = -1
	DirLeft  Dir = 0
	DirRight Dir = 1
	DirUp    Dir = 2
	DirDown  Dir = 3
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

// SetItemInnerSpacing sets the horizontal and vertical spacing between elements of
// a composed widget (e.g. a slider and its label).
func (style Style) SetItemInnerSpacing(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggStyleSetItemInnerSpacing(style.handle(), valueArg)
}

// SetItemSpacing sets horizontal and vertical spacing between widgets or lines.
func (style Style) SetItemSpacing(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggStyleSetItemSpacing(style.handle(), valueArg)
}

// SetFramePadding sets the padding within a framed rectangle (used by most widgets).
func (style Style) SetFramePadding(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggStyleSetFramePadding(style.handle(), valueArg)
}

// SetWindowPadding sets the padding within a window.
func (style Style) SetWindowPadding(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggStyleSetWindowPadding(style.handle(), valueArg)
}

// SetCellPadding sets the padding within a table cell.
func (style Style) SetCellPadding(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggStyleSetCellPadding(style.handle(), valueArg)
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

// TouchExtraPadding is the expansion for reactive bounding box for touch-based system where touch position is not accurate enough.
func (style Style) TouchExtraPadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggGetTouchExtraPadding(style.handle(), valueArg)
	valueFin()
	return value
}

// SetTouchExtraPadding expand reactive bounding box for touch-based system where touch position is not accurate enough.
func (style Style) SetTouchExtraPadding(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggSetTouchExtraPadding(style.handle(), valueArg)
}

// Alpha returns the global alpha that applies to everything in Dear ImGui.
func (style Style) Alpha() float32 {
	return float32(C.iggGetAlpha(style.handle()))
}

// SetAlpha sets the global alpha that applies to everything in Dear ImGui.
func (style Style) SetAlpha(v float32) {
	C.iggSetAlpha(style.handle(), C.float(v))
}

// DisabledAlpha returns the additional alpha multiplier applied by
// BeginDisabled(). Multiply over current value of Alpha.
func (style Style) DisabledAlpha() float32 {
	return float32(C.iggGetDisabledAlpha(style.handle()))
}

// SetDisabledAlpha sets the additional alpha multiplier applied by
// BeginDisabled(). Multiply over current value of Alpha.
func (style Style) SetDisabledAlpha(v float32) {
	C.iggSetDisabledAlpha(style.handle(), C.float(v))
}

// WindowRounding returns the radius of window corners rounding. Set to 0.0f to
// have rectangular windows. Large values tend to lead to variety of artifacts
// and are not recommended.
func (style Style) WindowRounding() float32 {
	return float32(C.iggGetWindowRounding(style.handle()))
}

// SetWindowRounding sets the radius of window corners rounding. Set to 0.0f to
// have rectangular windows. Large values tend to lead to variety of artifacts
// and are not recommended.
func (style Style) SetWindowRounding(v float32) {
	C.iggSetWindowRounding(style.handle(), C.float(v))
}

// WindowBorderSize returns the thickness of border around windows. Generally
// set to 0.0f or 1.0f. (Other values are not well tested and more CPU/GPU costly).
func (style Style) WindowBorderSize() float32 {
	return float32(C.iggGetWindowBorderSize(style.handle()))
}

// SetWindowBorderSize sets the thickness of border around windows. Generally
// set to 0.0f or 1.0f. (Other values are not well tested and more CPU/GPU costly).
func (style Style) SetWindowBorderSize(v float32) {
	C.iggSetWindowBorderSize(style.handle(), C.float(v))
}

// ChildRounding returns the radius of child window corners rounding. Set to
// 0.0f to have rectangular windows.
func (style Style) ChildRounding() float32 {
	return float32(C.iggGetChildRounding(style.handle()))
}

// SetChildRounding sets the radius of child window corners rounding. Set to
// 0.0f to have rectangular windows.
func (style Style) SetChildRounding(v float32) {
	C.iggSetChildRounding(style.handle(), C.float(v))
}

// ChildBorderSize returns the thickness of border around child windows.
// Generally set to 0.0f or 1.0f. (Other values are not well tested and more
// CPU/GPU costly).
func (style Style) ChildBorderSize() float32 {
	return float32(C.iggGetChildBorderSize(style.handle()))
}

// SetChildBorderSize sets the thickness of border around child windows.
// Generally set to 0.0f or 1.0f. (Other values are not well tested and more
// CPU/GPU costly).
func (style Style) SetChildBorderSize(v float32) {
	C.iggSetChildBorderSize(style.handle(), C.float(v))
}

// PopupRounding returns the radius of popup window corners rounding. (Note
// that tooltip windows use WindowRounding).
func (style Style) PopupRounding() float32 {
	return float32(C.iggGetPopupRounding(style.handle()))
}

// SetPopupRounding sets the radius of popup window corners rounding. (Note
// that tooltip windows use WindowRounding).
func (style Style) SetPopupRounding(v float32) {
	C.iggSetPopupRounding(style.handle(), C.float(v))
}

// PopupBorderSize returns the thickness of border around popup/tooltip windows.
// Generally set to 0.0f or 1.0f. (Other values are not well tested and more
// CPU/GPU costly).
func (style Style) PopupBorderSize() float32 {
	return float32(C.iggGetPopupBorderSize(style.handle()))
}

// SetPopupBorderSize sets the thickness of border around popup/tooltip windows.
// Generally set to 0.0f or 1.0f. (Other values are not well tested and more
// CPU/GPU costly).
func (style Style) SetPopupBorderSize(v float32) {
	C.iggSetPopupBorderSize(style.handle(), C.float(v))
}

// FrameRounding returns the radius of frame corners rounding. Set to 0.0f to
// have rectangular frame (used by most widgets).
func (style Style) FrameRounding() float32 {
	return float32(C.iggGetFrameRounding(style.handle()))
}

// SetFrameRounding sets the radius of frame corners rounding. Set to 0.0f to
// have rectangular frame (used by most widgets).
func (style Style) SetFrameRounding(v float32) {
	C.iggSetFrameRounding(style.handle(), C.float(v))
}

// FrameBorderSize returns the thickness of border around frames. Generally set
// to 0.0f or 1.0f. (Other values are not well tested and more CPU/GPU costly).
func (style Style) FrameBorderSize() float32 {
	return float32(C.iggGetFrameBorderSize(style.handle()))
}

// SetFrameBorderSize sets the thickness of border around frames. Generally set
// to 0.0f or 1.0f. (Other values are not well tested and more CPU/GPU costly).
func (style Style) SetFrameBorderSize(v float32) {
	C.iggSetFrameBorderSize(style.handle(), C.float(v))
}

// IndentSpacing returns the horizontal indentation when e.g. entering a tree
// node. Generally == (FontSize + FramePadding.x*2).
func (style Style) IndentSpacing() float32 {
	return float32(C.iggGetIndentSpacing(style.handle()))
}

// SetIndentSpacing sets the horizontal indentation when e.g. entering a tree
// node. Generally == (FontSize + FramePadding.x*2).
func (style Style) SetIndentSpacing(v float32) {
	C.iggSetIndentSpacing(style.handle(), C.float(v))
}

// ColumnsMinSpacing returns the minimum horizontal spacing between two columns.
// Preferably > (FramePadding.x + 1).
func (style Style) ColumnsMinSpacing() float32 {
	return float32(C.iggGetColumnsMinSpacing(style.handle()))
}

// SetColumnsMinSpacing sets the minimum horizontal spacing between two columns.
// Preferably > (FramePadding.x + 1).
func (style Style) SetColumnsMinSpacing(v float32) {
	C.iggSetColumnsMinSpacing(style.handle(), C.float(v))
}

// ScrollbarSize returns the width of the vertical scrollbar, Height of the
// horizontal scrollbar.
func (style Style) ScrollbarSize() float32 {
	return float32(C.iggGetColumnsMinSpacing(style.handle()))
}

// SetScrollbarSize sets the width of the vertical scrollbar, Height of the
// horizontal scrollbar.
func (style Style) SetScrollbarSize(v float32) {
	C.iggSetScrollbarSize(style.handle(), C.float(v))
}

// ScrollbarRounding returns the radius of grab corners for scrollbar.
func (style Style) ScrollbarRounding() float32 {
	return float32(C.iggGetColumnsMinSpacing(style.handle()))
}

// SetScrollbarRounding sets the radius of grab corners for scrollbar.
func (style Style) SetScrollbarRounding(v float32) {
	C.iggSetScrollbarRounding(style.handle(), C.float(v))
}

// GrabMinSize returns the minimum width/height of a grab box for
// slider/scrollbar.
func (style Style) GrabMinSize() float32 {
	return float32(C.iggGetGrabMinSize(style.handle()))
}

// SetGrabMinSize sets the minimum width/height of a grab box for
// slider/scrollbar.
func (style Style) SetGrabMinSize(v float32) {
	C.iggSetGrabMinSize(style.handle(), C.float(v))
}

// GrabRounding returns the radius of grabs corners rounding. Set to 0.0f to
// have rectangular slider grabs.
func (style Style) GrabRounding() float32 {
	return float32(C.iggGetGrabRounding(style.handle()))
}

// SetGrabRounding sets the radius of grabs corners rounding. Set to 0.0f to
// have rectangular slider grabs.
func (style Style) SetGrabRounding(v float32) {
	C.iggSetGrabRounding(style.handle(), C.float(v))
}

// LogSliderDeadzone returns the size in pixels of the dead-zone around zero on
// logarithmic sliders that cross zero.
func (style Style) LogSliderDeadzone() float32 {
	return float32(C.iggGetLogSliderDeadzone(style.handle()))
}

// SetLogSliderDeadzone sets the size in pixels of the dead-zone around zero on
// logarithmic sliders that cross zero.
func (style Style) SetLogSliderDeadzone(v float32) {
	C.iggSetLogSliderDeadzone(style.handle(), C.float(v))
}

// TabRounding returns the radius of upper corners of a tab. Set to 0.0f to have
// rectangular tabs.
func (style Style) TabRounding() float32 {
	return float32(C.iggGetTabRounding(style.handle()))
}

// SetTabRounding sets the radius of upper corners of a tab. Set to 0.0f to have
// rectangular tabs.
func (style Style) SetTabRounding(v float32) {
	C.iggSetTabRounding(style.handle(), C.float(v))
}

// TabBorderSize returns the thickness of border around tabs.
func (style Style) TabBorderSize() float32 {
	return float32(C.iggGetTabBorderSize(style.handle()))
}

// SetTabBorderSize sets the thickness of border around tabs.
func (style Style) SetTabBorderSize(v float32) {
	C.iggSetTabBorderSize(style.handle(), C.float(v))
}

// TabMinWidthForCloseButton returns the minimum width for close button to
// appears on an unselected tab when hovered. Set to 0.0f to always show when
// hovering, set to FLT_MAX to never show close button unless selected.
func (style Style) TabMinWidthForCloseButton() float32 {
	return float32(C.iggGetTabMinWidthForCloseButton(style.handle()))
}

// SetTabMinWidthForCloseButton sets the minimum width for close button to
// appears on an unselected tab when hovered. Set to 0.0f to always show when
// hovering, set to FLT_MAX to never show close button unless selected.
func (style Style) SetTabMinWidthForCloseButton(v float32) {
	C.iggSetTabMinWidthForCloseButton(style.handle(), C.float(v))
}

// MouseCursorScale returns the scale software rendered mouse cursor (when
// io.MouseDrawCursor is enabled). May be removed later.
func (style Style) MouseCursorScale() float32 {
	return float32(C.iggGetMouseCursorScale(style.handle()))
}

// SetMouseCursorScale sets the scale software rendered mouse cursor (when
// io.MouseDrawCursor is enabled). May be removed later.
func (style Style) SetMouseCursorScale(v float32) {
	C.iggSetMouseCursorScale(style.handle(), C.float(v))
}

// CurveTessellationTol returns the tessellation tolerance when using
// PathBezierCurveTo() without a specific number of segments. Decrease for
// highly tessellated curves (higher quality, more polygons), increase to reduce
// quality.
func (style Style) CurveTessellationTol() float32 {
	return float32(C.iggGetCurveTessellationTol(style.handle()))
}

// SetCurveTessellationTol sets the tessellation tolerance when using
// PathBezierCurveTo() without a specific number of segments. Decrease for
// highly tessellated curves (higher quality, more polygons), increase to reduce
// quality.
func (style Style) SetCurveTessellationTol(v float32) {
	C.iggSetCurveTessellationTol(style.handle(), C.float(v))
}

// CircleTessellationMaxError returns the maximum error (in pixels) allowed when
// using AddCircle()/AddCircleFilled() or drawing rounded corner rectangles with
// no explicit segment count specified. Decrease for higher quality but more
// geometry.
func (style Style) CircleTessellationMaxError() float32 {
	return float32(C.iggGetCircleTessellationMaxError(style.handle()))
}

// SetCircleTessellationMaxError sets the maximum error (in pixels) allowed when
// using AddCircle()/AddCircleFilled() or drawing rounded corner rectangles with
// no explicit segment count specified. Decrease for higher quality but more
// geometry.
func (style Style) SetCircleTessellationMaxError(v float32) {
	C.iggSetCircleTessellationMaxError(style.handle(), C.float(v))
}

// WindowMinSize returns the minimum window size. This is a global setting. If
// you want to constraint individual windows, use
// SetNextWindowSizeConstraints().
func (style Style) WindowMinSize() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetWindowMinSize(style.handle(), valueArg)
	valueFin()
	return value
}

// SetWindowMinSize sets the minimum window size. This is a global setting. If
// you want to constraint individual windows, use
// SetNextWindowSizeConstraints().
func (style Style) SetWindowMinSize(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggSetWindowMinSize(style.handle(), valueArg)
}

// WindowTitleAlign returns the alignment for title bar text. Defaults to
// (0.0f,0.5f) for left-aligned,vertically centered.
func (style Style) WindowTitleAlign() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetWindowTitleAlign(style.handle(), valueArg)
	valueFin()
	return value
}

// SetWindowTitleAlign sets the alignment for title bar text. Defaults to
// (0.0f,0.5f) for left-aligned,vertically centered.
func (style Style) SetWindowTitleAlign(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggSetWindowTitleAlign(style.handle(), valueArg)
}

// ButtonTextAlign returns the alignment of button text when button is larger
// than text. Defaults to (0.5f, 0.5f) (centered).
func (style Style) ButtonTextAlign() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetButtonTextAlign(style.handle(), valueArg)
	valueFin()
	return value
}

// SetButtonTextAlign sets the alignment of button text when button is larger
// than text. Defaults to (0.5f, 0.5f) (centered).
func (style Style) SetButtonTextAlign(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggSetButtonTextAlign(style.handle(), valueArg)
}

// SelectableTextAlign returns the alignment of selectable text. Defaults to
// (0.0f, 0.0f) (top-left aligned). It's generally important to keep this
// left-aligned if you want to lay multiple items on a same line.
func (style Style) SelectableTextAlign() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetSelectableTextAlign(style.handle(), valueArg)
	valueFin()
	return value
}

// SetSelectableTextAlign sets the alignment of selectable text. Defaults to
// (0.0f, 0.0f) (top-left aligned). It's generally important to keep this
// left-aligned if you want to lay multiple items on a same line.
func (style Style) SetSelectableTextAlign(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggSetSelectableTextAlign(style.handle(), valueArg)
}

// DisplayWindowPadding returns the window position are clamped to be visible
// within the display area or monitors by at least this amount. Only applies to
// regular windows.
func (style Style) DisplayWindowPadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetDisplayWindowPadding(style.handle(), valueArg)
	valueFin()
	return value
}

// SetDisplayWindowPadding sets the window position are clamped to be visible
// within the display area or monitors by at least this amount. Only applies to
// regular windows.
func (style Style) SetDisplayWindowPadding(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggSetDisplayWindowPadding(style.handle(), valueArg)
}

// DisplaySafeAreaPadding returns the if you cannot see the edges of your screen
// (e.g. on a TV) increase the safe area padding. Apply to popups/tooltips as
// well regular windows. NB: Prefer configuring your TV sets correctly.
func (style Style) DisplaySafeAreaPadding() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggStyleGetDisplaySafeAreaPadding(style.handle(), valueArg)
	valueFin()
	return value
}

// SetDisplaySafeAreaPadding sets the if you cannot see the edges of your screen
// (e.g. on a TV) increase the safe area padding. Apply to popups/tooltips as
// well regular windows. NB: Prefer configuring your TV sets correctly.
func (style Style) SetDisplaySafeAreaPadding(value Vec2) {
	valueArg, _ := value.wrapped()
	C.iggSetDisplaySafeAreaPadding(style.handle(), valueArg)
}

// AntiAliasedLines returns the state of anti-aliased lines/borders. Disable if
// you are really tight on CPU/GPU. Latched at the beginning of the frame
// (copied to ImDrawList).
func (style Style) AntiAliasedLines() bool {
	return C.iggStyleGetAntiAliasedLines(style.handle()) != 0
}

// SetAntiAliasedLines enables of disables anti-aliased lines/borders. Disable
// if you are really tight on CPU/GPU. Latched at the beginning of the frame
// (copied to ImDrawList).
func (style Style) SetAntiAliasedLines(value bool) {
	i := 0
	if value {
		i = 1
	}
	C.iggStyleSetAntiAliasedLines(style.handle(), C.int(i))
}

// AntiAliasedLinesUseTex returns the state of anti-aliased lines/borders using
// textures where possible. Require backend to render with bilinear filtering.
// Latched at the beginning of the frame (copied to ImDrawList).
func (style Style) AntiAliasedLinesUseTex() bool {
	return C.iggStyleGetAntiAliasedLinesUseTex(style.handle()) != 0
}

// SetAntiAliasedLinesUseTex enables of disables anti-aliased lines/borders
// using textures where possible. Require backend to render with bilinear
// filtering. Latched at the beginning of the frame (copied to ImDrawList).
func (style Style) SetAntiAliasedLinesUseTex(value bool) {
	i := 0
	if value {
		i = 1
	}
	C.iggStyleSetAntiAliasedLinesUseTex(style.handle(), C.int(i))
}

// AntiAliasedFill returns the state of anti-aliased edges around filled shapes
// (rounded rectangles, circles, etc.). Disable if you are really tight on
// CPU/GPU. Latched at the beginning of the frame (copied to ImDrawList).
func (style Style) AntiAliasedFill() bool {
	return C.iggStyleGetAntiAliasedFill(style.handle()) != 0
}

// SetAntiAliasedFill enables of disables anti-aliased edges around filled
// shapes (rounded rectangles, circles, etc.). Disable if you are really tight
// on CPU/GPU. Latched at the beginning of the frame (copied to ImDrawList).
func (style Style) SetAntiAliasedFill(value bool) {
	i := 0
	if value {
		i = 1
	}
	C.iggStyleSetAntiAliasedFill(style.handle(), C.int(i))
}

// WindowMenuButtonPosition sets the side of the collapsing/docking button in
// the title bar (None/Left/Right). Defaults to DirLeft.
func (style Style) WindowMenuButtonPosition() Dir {
	return Dir(C.iggStyleGetWindowMenuButtonPosition(style.handle()))
}

// SetWindowMenuButtonPosition returns the side of the collapsing/docking button
// in the title bar (None/Left/Right). Defaults to DirLeft.
func (style Style) SetWindowMenuButtonPosition(value Dir) {
	C.iggStyleSetWindowMenuButtonPosition(style.handle(), C.int(value))
}

// ColorButtonPosition sets the side of the color button in the ColorEdit4
// widget (left/right). Defaults to DirRight.
func (style Style) ColorButtonPosition() Dir {
	return Dir(C.iggStyleGetColorButtonPosition(style.handle()))
}

// SetColorButtonPosition returns the side of the color button in the ColorEdit4
// widget (left/right). Defaults to DirRight.
func (style Style) SetColorButtonPosition(value Dir) {
	C.iggStyleSetColorButtonPosition(style.handle(), C.int(value))
}
