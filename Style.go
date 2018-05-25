package imgui

// #include "StyleWrapper.h"
import "C"

// StyleColorID identifies a color in the UI style.
type StyleColorID int

// StyleColor identifier
const (
	StyleColorText StyleColorID = iota
	StyleColorTextDisabled
	StyleColorWindowBg
	StyleColorChildBg
	StyleColorPopupBg
	StyleColorBorder
	StyleColorBorderShadow
	StyleColorFrameBg
	StyleColorFrameBgHovered
	StyleColorFrameBgActive
	StyleColorTitleBg
	StyleColorTitleBgActive
	StyleColorTitleBgCollapsed
	StyleColorMenuBarBg
	StyleColorScrollbarBg
	StyleColorScrollbarGrab
	StyleColorScrollbarGrabHovered
	StyleColorScrollbarGrabActive
	StyleColorCheckMark
	StyleColorSliderGrab
	StyleColorSliderGrabActive
	StyleColorButton
	StyleColorButtonHovered
	StyleColorButtonActive
	StyleColorHeader
	StyleColorHeaderHovered
	StyleColorHeaderActive
	StyleColorSeparator
	StyleColorSeparatorHovered
	StyleColorSeparatorActive
	StyleColorResizeGrip
	StyleColorResizeGripHovered
	StyleColorResizeGripActive
	StyleColorPlotLines
	StyleColorPlotLinesHovered
	StyleColorPlotHistogram
	StyleColorPlotHistogramHovered
	StyleColorTextSelectedBg
	StyleColorModalWindowDarkening
	StyleColorDragDropTarget
	StyleColorNavHighlight
	StyleColorNavWindowingHighlight
)

// Style describes the overall graphical representation of the user interface.
type Style uintptr

func (style Style) handle() C.IggGuiStyle {
	return C.IggGuiStyle(style)
}

// SetColor sets a color value of the UI style.
func (style Style) SetColor(id StyleColorID, value Vec4) {
	valueArg, _ := value.wrapped()
	C.iggStyleSetColor(style.handle(), C.int(id), valueArg)
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
