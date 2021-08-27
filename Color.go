package imgui

// #include "wrapper/Color.h"
import "C"

// ColorEditFlags for ColorEdit3V(), etc.
type ColorEditFlags int

const (
	// ColorEditFlagsNone default = 0.
	ColorEditFlagsNone ColorEditFlags = 0
	// ColorEditFlagsNoAlpha ignores Alpha component (read 3 components from the input pointer).
	ColorEditFlagsNoAlpha ColorEditFlags = 1 << 1
	// ColorEditFlagsNoPicker disables picker when clicking on colored square.
	ColorEditFlagsNoPicker ColorEditFlags = 1 << 2
	// ColorEditFlagsNoOptions disables toggling options menu when right-clicking on inputs/small preview.
	ColorEditFlagsNoOptions ColorEditFlags = 1 << 3
	// ColorEditFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs).
	ColorEditFlagsNoSmallPreview ColorEditFlags = 1 << 4
	// ColorEditFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorEditFlagsNoInputs ColorEditFlags = 1 << 5
	// ColorEditFlagsNoTooltip disables tooltip when hovering the preview.
	ColorEditFlagsNoTooltip ColorEditFlags = 1 << 6
	// ColorEditFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorEditFlagsNoLabel ColorEditFlags = 1 << 7
	// ColorEditFlagsNoSidePreview disables bigger color preview on right side of the picker, use small colored square preview instead.
	ColorEditFlagsNoSidePreview ColorEditFlags = 1 << 8
	// ColorEditFlagsNoDragDrop disables drag and drop target. ColorButton: disable drag and drop source.
	ColorEditFlagsNoDragDrop ColorEditFlags = 1 << 9
	// ColorEditFlagsNoBorder disables border (which is enforced by default).
	ColorEditFlagsNoBorder ColorEditFlags = 1 << 10

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions().
	// The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call
	// SetColorEditOptions() during startup.

	// ColorEditFlagsAlphaBar shows vertical alpha bar/gradient in picker.
	ColorEditFlagsAlphaBar ColorEditFlags = 1 << 16
	// ColorEditFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreview ColorEditFlags = 1 << 17
	// ColorEditFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreviewHalf ColorEditFlags = 1 << 18
	// ColorEditFlagsHDR = (WIP) surrently only disable 0.0f..1.0f limits in RGBA edition.
	// Note: you probably want to use ImGuiColorEditFlags_Float flag as well.
	ColorEditFlagsHDR ColorEditFlags = 1 << 19
	// ColorEditFlagsRGB sets the format as RGB.
	ColorEditFlagsRGB ColorEditFlags = 1 << 20
	// ColorEditFlagsHSV sets the format as HSV.
	ColorEditFlagsHSV ColorEditFlags = 1 << 21
	// ColorEditFlagsHEX sets the format as HEX.
	ColorEditFlagsHEX ColorEditFlags = 1 << 22
	// ColorEditFlagsUint8 _display_ values formatted as 0..255.
	ColorEditFlagsUint8 ColorEditFlags = 1 << 23
	// ColorEditFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorEditFlagsFloat ColorEditFlags = 1 << 24

	// ColorEditFlagsPickerHueBar shows bar for Hue, rectangle for Sat/Value.
	ColorEditFlagsPickerHueBar ColorEditFlags = 1 << 25
	// ColorEditFlagsPickerHueWheel shows wheel for Hue, triangle for Sat/Value.
	ColorEditFlagsPickerHueWheel ColorEditFlags = 1 << 26
	// ColorEditFlagsInputRGB enables input and output data in RGB format.
	ColorEditFlagsInputRGB ColorEditFlags = 1 << 27
	// ColorEditFlagsInputHSV enables input and output data in HSV format.
	ColorEditFlagsInputHSV ColorEditFlags = 1 << 28
)

// ColorEdit3 calls ColorEdit3V(label, col, 0).
func ColorEdit3(label string, col *[3]float32) bool {
	return ColorEdit3V(label, col, 0)
}

// ColorEdit3V will show a clickable little square which will open a color picker window for 3D vector (rgb format).
func ColorEdit3V(label string, col *[3]float32, flags ColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorEdit3(labelArg, ccol, C.int(flags)) != 0
}

// ColorEdit4 calls ColorEdit4V(label, col, 0).
func ColorEdit4(label string, col *[4]float32) bool {
	return ColorEdit4V(label, col, 0)
}

// ColorEdit4V will show a clickable little square which will open a color picker window for 4D vector (rgba format).
func ColorEdit4V(label string, col *[4]float32, flags ColorEditFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorEdit4(labelArg, ccol, C.int(flags)) != 0
}

// ColorButton displays a color square/button, hover for details, returns true when pressed.
func ColorButton(id string, col Vec4, flags ColorEditFlags, size Vec2) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	sizeArg, _ := size.wrapped()
	colArg, _ := col.wrapped()
	return C.iggColorButton(idArg, colArg, C.int(flags), sizeArg) != 0
}

// ColorPickerFlags for ColorPicker3V(), etc.
type ColorPickerFlags int

const (
	// ColorPickerFlagsNone default = 0.
	ColorPickerFlagsNone ColorPickerFlags = 0
	// ColorPickerFlagsNoPicker disables picker when clicking on colored square.
	ColorPickerFlagsNoPicker ColorPickerFlags = 1 << 2
	// ColorPickerFlagsNoOptions disables toggling options menu when right-clicking on inputs/small preview.
	ColorPickerFlagsNoOptions ColorPickerFlags = 1 << 3
	// ColorPickerFlagsNoAlpha ignoreс Alpha component (read 3 components from the input pointer).
	ColorPickerFlagsNoAlpha ColorPickerFlags = 1 << 1
	// ColorPickerFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs).
	ColorPickerFlagsNoSmallPreview ColorPickerFlags = 1 << 4
	// ColorPickerFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorPickerFlagsNoInputs ColorPickerFlags = 1 << 5
	// ColorPickerFlagsNoTooltip disables tooltip when hovering the preview.
	ColorPickerFlagsNoTooltip ColorPickerFlags = 1 << 6
	// ColorPickerFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorPickerFlagsNoLabel ColorPickerFlags = 1 << 7
	// ColorPickerFlagsNoSidePreview disables bigger color preview on right side of the picker, use small colored square preview instead.
	ColorPickerFlagsNoSidePreview ColorPickerFlags = 1 << 8

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions().
	// The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call
	// SetColorPickerOptions() during startup.

	// ColorPickerFlagsAlphaBar shows vertical alpha bar/gradient in picker.
	ColorPickerFlagsAlphaBar ColorPickerFlags = 1 << 16
	// ColorPickerFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
	ColorPickerFlagsAlphaPreview ColorPickerFlags = 1 << 17
	// ColorPickerFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
	ColorPickerFlagsAlphaPreviewHalf ColorPickerFlags = 1 << 18
	// ColorPickerFlagsRGB sets the format as RGB.
	ColorPickerFlagsRGB ColorPickerFlags = 1 << 20
	// ColorPickerFlagsHSV sets the format as HSV.
	ColorPickerFlagsHSV ColorPickerFlags = 1 << 21
	// ColorPickerFlagsHEX sets the format as HEX.
	ColorPickerFlagsHEX ColorPickerFlags = 1 << 22
	// ColorPickerFlagsUint8 _display_ values formatted as 0..255.
	ColorPickerFlagsUint8 ColorPickerFlags = 1 << 23
	// ColorPickerFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorPickerFlagsFloat ColorPickerFlags = 1 << 24
	// ColorPickerFlagsPickerHueBar bar for Hue, rectangle for Sat/Value.
	ColorPickerFlagsPickerHueBar ColorPickerFlags = 1 << 25
	// ColorPickerFlagsPickerHueWheel wheel for Hue, triangle for Sat/Value.
	ColorPickerFlagsPickerHueWheel ColorPickerFlags = 1 << 26
	// ColorPickerFlagsInputRGB enables input and output data in RGB format.
	ColorPickerFlagsInputRGB ColorPickerFlags = 1 << 27
	// ColorPickerFlagsInputHSV enables input and output data in HSV format.
	ColorPickerFlagsInputHSV ColorPickerFlags = 1 << 28
)

// ColorPicker3 calls ColorPicker3V(label, col, 0).
func ColorPicker3(label string, col *[3]float32) bool {
	return ColorPicker3V(label, col, 0)
}

// ColorPicker3V will show directly a color picker control for editing a color in 3D vector (rgb format).
func ColorPicker3V(label string, col *[3]float32, flags ColorPickerFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorPicker3(labelArg, ccol, C.int(flags)) != 0
}

// ColorPicker4 calls ColorPicker4V(label, col, 0).
func ColorPicker4(label string, col *[4]float32) bool {
	return ColorPicker4V(label, col, 0)
}

// ColorPicker4V will show directly a color picker control for editing a color in 4D vector (rgba format).
func ColorPicker4V(label string, col *[4]float32, flags ColorPickerFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorPicker4(labelArg, ccol, C.int(flags)) != 0
}
