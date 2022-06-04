package imgui

// #include "wrapper/Widgets.h"
import "C"

import (
	"fmt"
	"math"
	"strings"
)

// Text adds formatted text. See PushTextWrapPosV() or PushStyleColorV() for modifying the output.
// Without any modified style stack, the text is unformatted.
func Text(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()
	// Internally we use ImGui::TextUnformatted, for the most direct call.
	C.iggTextUnformatted(textArg)
}

// Textf calls Text(fmt.Sprintf(format, v...) .
func Textf(format string, v ...interface{}) {
	Text(fmt.Sprintf(format, v...))
}

// LabelText adds text+label aligned the same way as value+label widgets.
func LabelText(label, text string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	textArg, textFin := wrapString(text)
	defer textFin()
	C.iggLabelText(labelArg, textArg)
}

// LabelTextf calls LabelText(label, fmt.Sprintf(format, v...)) .
func LabelTextf(label, format string, v ...interface{}) {
	LabelText(label, fmt.Sprintf(format, v...))
}

// ButtonV returns true if it is clicked.
func ButtonV(id string, size Vec2) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	sizeArg, _ := size.wrapped()
	return C.iggButton(idArg, sizeArg) != 0
}

// Button calls ButtonV(id, Vec2{0,0}).
func Button(id string) bool {
	return ButtonV(id, Vec2{})
}

// ButtonFlags Flags for InvisibleButton().
type ButtonFlags int

const (
	// ButtonFlagsNone is no flag applied.
	ButtonFlagsNone ButtonFlags = 0
	// ButtonFlagsMouseButtonLeft reacts on left mouse button (default).
	ButtonFlagsMouseButtonLeft ButtonFlags = 1 << 0
	// ButtonFlagsMouseButtonRight reacts on right mouse button.
	ButtonFlagsMouseButtonRight ButtonFlags = 1 << 1
	// ButtonFlagsMouseButtonMiddle reacts on center mouse button.
	ButtonFlagsMouseButtonMiddle ButtonFlags = 1 << 2
)

// InvisibleButtonV returns true if it is clicked.
func InvisibleButtonV(id string, size Vec2, flags ButtonFlags) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	sizeArg, _ := size.wrapped()
	return C.iggInvisibleButton(idArg, sizeArg, C.int(flags)) != 0
}

// InvisibleButton calls InvisibleButtonV(id, size, ButtonFlagsNone).
func InvisibleButton(id string, size Vec2) bool {
	return InvisibleButtonV(id, size, ButtonFlagsNone)
}

// ImageV adds an image based on given texture ID.
// Refer to TextureID what this represents and how it is drawn.
func ImageV(id TextureID, size Vec2, uv0, uv1 Vec2, tintCol, borderCol Vec4) {
	sizeArg, _ := size.wrapped()
	uv0Arg, _ := uv0.wrapped()
	uv1Arg, _ := uv1.wrapped()
	tintColArg, _ := tintCol.wrapped()
	borderColArg, _ := borderCol.wrapped()
	C.iggImage(id.handle(), sizeArg, uv0Arg, uv1Arg, tintColArg, borderColArg)
}

// Image calls ImageV(id, size, Vec2{0,0}, Vec2{1,1}, Vec4{1,1,1,1}, Vec4{0,0,0,0}).
func Image(id TextureID, size Vec2) {
	ImageV(id, size, Vec2{X: 0, Y: 0}, Vec2{X: 1, Y: 1}, Vec4{X: 1, Y: 1, Z: 1, W: 1}, Vec4{X: 0, Y: 0, Z: 0, W: 0})
}

// ImageButtonV adds a button with an image, based on given texture ID.
// Refer to TextureID what this represents and how it is drawn.
// <0 framePadding uses default frame padding settings. 0 for no padding.
func ImageButtonV(id TextureID, size Vec2, uv0, uv1 Vec2, framePadding int, bgCol Vec4, tintCol Vec4) bool {
	sizeArg, _ := size.wrapped()
	uv0Arg, _ := uv0.wrapped()
	uv1Arg, _ := uv1.wrapped()
	bgColArg, _ := bgCol.wrapped()
	tintColArg, _ := tintCol.wrapped()
	return C.iggImageButton(id.handle(), sizeArg, uv0Arg, uv1Arg, C.int(framePadding), bgColArg, tintColArg) != 0
}

// ImageButton calls ImageButtonV(id, size, Vec2{0,0}, Vec2{1,1}, -1, Vec4{0,0,0,0}, Vec4{1,1,1,1}).
func ImageButton(id TextureID, size Vec2) bool {
	return ImageButtonV(id, size, Vec2{X: 0, Y: 0}, Vec2{X: 1, Y: 1}, -1, Vec4{X: 0, Y: 0, Z: 0, W: 0}, Vec4{X: 1, Y: 1, Z: 1, W: 1})
}

// Checkbox creates a checkbox in the selected state.
// The return value indicates if the selected state has changed.
func Checkbox(id string, selected *bool) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	selectedArg, selectedFin := wrapBool(selected)
	defer selectedFin()
	return C.iggCheckbox(idArg, selectedArg) != 0
}

// RadioButton returns true if it is clicked and active indicates if it is selected.
func RadioButton(id string, active bool) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	return C.iggRadioButton(idArg, castBool(active)) != 0
}

// RadioButtonInt modifies integer v. Returns true if it is selected.
//
// The radio button will be set if v == button. Useful for groups of radio
// buttons. In the example below, "radio b" will be selected.
//
//		v := 1
//		imgui.RadioButtonInt("radio a", &v, 0)
//		imgui.RadioButtonInt("radio b", &v, 1)
//		imgui.RadioButtonInt("radio c", &v, 2)
//
func RadioButtonInt(id string, v *int, button int) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	ok := C.iggRadioButton(idArg, castBool(button == *v)) != 0
	if ok {
		*v = button
	}
	return ok
}

// Bullet draws a small circle and keeps the cursor on the same line.
// Advance cursor x position by TreeNodeToLabelSpacing(), same distance that TreeNode() uses.
func Bullet() {
	C.iggBullet()
}

// ProgressBarV creates a progress bar.
// size (for each axis) is < 0.0f: align to end, 0.0f: auto, > 0.0f: specified size.
func ProgressBarV(fraction float32, size Vec2, overlay string) {
	sizeArg, _ := size.wrapped()
	overlayArg, overlayFin := wrapString(overlay)
	defer overlayFin()
	C.iggProgressBar(C.float(fraction), sizeArg, overlayArg)
}

// ProgressBar calls ProgressBarV(fraction, Vec2{X: -math.SmallestNonzeroFloat32, Y: 0}, "").
func ProgressBar(fraction float32) {
	ProgressBarV(fraction, Vec2{X: -math.SmallestNonzeroFloat32, Y: 0}, "")
}

// ComboFlags for BeginComboV().
type ComboFlags int

const (
	// ComboFlagsNone default = 0.
	ComboFlagsNone ComboFlags = 0
	// ComboFlagsPopupAlignLeft aligns the popup toward the left by default.
	ComboFlagsPopupAlignLeft ComboFlags = 1 << 0
	// ComboFlagsHeightSmall has max ~4 items visible.
	// Tip: If you want your combo popup to be a specific size you can use SetNextWindowSizeConstraints() prior to calling BeginCombo().
	ComboFlagsHeightSmall ComboFlags = 1 << 1
	// ComboFlagsHeightRegular has max ~8 items visible (default).
	ComboFlagsHeightRegular ComboFlags = 1 << 2
	// ComboFlagsHeightLarge has max ~20 items visible.
	ComboFlagsHeightLarge ComboFlags = 1 << 3
	// ComboFlagsHeightLargest has as many fitting items as possible.
	ComboFlagsHeightLargest ComboFlags = 1 << 4
	// ComboFlagsNoArrowButton displays on the preview box without the square arrow button.
	ComboFlagsNoArrowButton ComboFlags = 1 << 5
	// ComboFlagsNoPreview displays only a square arrow button.
	ComboFlagsNoPreview ComboFlags = 1 << 6
)

// BeginComboV creates a combo box with complete control over the content to the user.
// Call EndCombo() if this function returns true.
// flags are the ComboFlags to apply.
func BeginComboV(label, previewValue string, flags ComboFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	previewValueArg, previewValueFin := wrapString(previewValue)
	defer previewValueFin()
	return C.iggBeginCombo(labelArg, previewValueArg, C.int(flags)) != 0
}

// BeginCombo calls BeginComboV(label, previewValue, 0).
func BeginCombo(label, previewValue string) bool {
	return BeginComboV(label, previewValue, 0)
}

// EndCombo must be called if BeginComboV() returned true.
func EndCombo() {
	C.iggEndCombo()
}

// Combo calls ComboV(id, value, list, -1).
func Combo(id string, value *int32, list []string) bool {
	return ComboV(id, value, list, -1)
}

// ComboV is a helper over BeginCombo()/EndCombo() which are kept available for convenience purpose.
// This is analogous to how ListBox are created.
func ComboV(id string, value *int32, list []string, heightInItems int) bool {
	valueArg, valueFin := wrapInt32(value)
	defer valueFin()
	return C.iggCombo(
		C.CString(id),
		valueArg,
		C.CString(strings.Join(list, string(byte(0)))+string(byte(0))),
		(C.int)(heightInItems),
	) != 0
}

// SliderFlags for DragFloat(), DragInt(), SliderFloat(), SliderInt() etc.
// We use the same sets of flags for DragXXX() and SliderXXX() functions as the features are the same and it makes it easier to swap them.
type SliderFlags int

const (
	// SliderFlagsNone is no flag applied.
	SliderFlagsNone SliderFlags = 0
	// SliderFlagsAlwaysClamp clamps value to min/max bounds when input manually with CTRL+Click. By default CTRL+Click allows going out of bounds.
	SliderFlagsAlwaysClamp SliderFlags = 1 << 4
	// SliderFlagsLogarithmic makes the widget logarithmic (linear otherwise). Consider using SliderFlagNoRoundToFormat with this if using a format-string with small amount of digits.
	SliderFlagsLogarithmic SliderFlags = 1 << 5
	// SliderFlagsNoRoundToFormat disables rounding underlying value to match precision of the display format string (e.g. %.3f values are rounded to those 3 digits).
	SliderFlagsNoRoundToFormat SliderFlags = 1 << 6
	// SliderFlagsNoInput disables CTRL+Click or Enter key allowing to input text directly into the widget.
	SliderFlagsNoInput SliderFlags = 1 << 7
)

// DragFloatV creates a draggable slider for floats.
func DragFloatV(label string, value *float32, speed, min, max float32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapFloat(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggDragFloat(labelArg, valueArg, C.float(speed), C.float(min), C.float(max), formatArg, C.int(flags)) != 0
}

// DragFloat calls DragFloatV(label, value, 1.0, 0.0, 0.0, "%.3f", SliderFlagsNone).
func DragFloat(label string, value *float32) bool {
	return DragFloatV(label, value, 1.0, 0.0, 0.0, "%.3f", SliderFlagsNone)
}

// DragFloat2V creates a draggable slider for a 2D vector.
func DragFloat2V(label string, values *[2]float32, speed, min, max float32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.float)(&values[0])
	return C.iggDragFloatN(labelArg, cvalues, 2, C.float(speed), C.float(min), C.float(max), formatArg, C.int(flags)) != 0
}

// DragFloat2 calls DragFloat2V(label, value, 1.0, 0.0, 0.0, "%.3f", SliderFlagsNone).
func DragFloat2(label string, value *[2]float32) bool {
	return DragFloat2V(label, value, 1.0, 0.0, 0.0, "%.3f", SliderFlagsNone)
}

// DragFloat3V creates a draggable slider for a 3D vector.
func DragFloat3V(label string, values *[3]float32, speed, min, max float32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.float)(&values[0])
	return C.iggDragFloatN(labelArg, cvalues, 3, C.float(speed), C.float(min), C.float(max), formatArg, C.int(flags)) != 0
}

// DragFloat3 calls DragFloat3V(label, value, 1.0, 0.0, 0.0, "%.3f", SliderFlagsNone).
func DragFloat3(label string, value *[3]float32) bool {
	return DragFloat3V(label, value, 1.0, 0.0, 0.0, "%.3f", SliderFlagsNone)
}

// DragFloat4V creates a draggable slider for a 4D vector.
func DragFloat4V(label string, values *[4]float32, speed, min, max float32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.float)(&values[0])
	return C.iggDragFloatN(labelArg, cvalues, 4, C.float(speed), C.float(min), C.float(max), formatArg, C.int(flags)) != 0
}

// DragFloat4 calls DragFloat4V(label, value, 1.0, 0.0, 0.0, "%.3f", SliderFlagsNone).
func DragFloat4(label string, value *[4]float32) bool {
	return DragFloat4V(label, value, 1.0, 0.0, 0.0, "%.3f", SliderFlagsNone)
}

// DragFloatRange2V creates a draggable slider in floats range.
func DragFloatRange2V(label string, currentMin *float32, currentMax *float32, speed float32, min float32, max float32, format string, formatMax string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	currentMinArg, currentMinFin := wrapFloat(currentMin)
	defer currentMinFin()
	currentMaxArg, currentMaxFin := wrapFloat(currentMax)
	defer currentMaxFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	formatMaxArg, formatMaxFin := wrapString(formatMax)
	defer formatMaxFin()
	return C.iggDragFloatRange2V(labelArg, currentMinArg, currentMaxArg, C.float(speed), C.float(min), C.float(max), formatArg, formatMaxArg, C.int(flags)) != 0
}

// DragFloatRange2 calls DragFloatRange2V(label, currentMin, currentMax, 1, 0, 0, "%.3f", "%.3f", SliderFlagsNone).
func DragFloatRange2(label string, currentMin *float32, currentMax *float32) bool {
	return DragFloatRange2V(label, currentMin, currentMax, 1, 0, 0, "%.3f", "%.3f", SliderFlagsNone)
}

// DragIntV creates a draggable slider for integers.
func DragIntV(label string, value *int32, speed float32, min, max int32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapInt32(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggDragInt(labelArg, valueArg, C.float(speed), C.int(min), C.int(max), formatArg, C.int(flags)) != 0
}

// DragInt calls DragIntV(label, value, 1.0, 0, 0, "%d", SliderFlagsNone).
func DragInt(label string, value *int32) bool {
	return DragIntV(label, value, 1.0, 0, 0, "%d", SliderFlagsNone)
}

// DragInt2V creates a draggable slider for a 2D vector.
func DragInt2V(label string, values *[2]int32, speed float32, min, max int32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.int)(&values[0])
	return C.iggDragIntN(labelArg, cvalues, 2, C.float(speed), C.int(min), C.int(max), formatArg, C.int(flags)) != 0
}

// DragInt2 calls DragInt2V(label, value, 1.0, 0.0, 0.0, "%d", SliderFlagsNone).
func DragInt2(label string, value *[2]int32) bool {
	return DragInt2V(label, value, 1.0, 0.0, 0.0, "%d", SliderFlagsNone)
}

// DragInt3V creates a draggable slider for a 3D vector.
func DragInt3V(label string, values *[3]int32, speed float32, min, max int32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.int)(&values[0])
	return C.iggDragIntN(labelArg, cvalues, 3, C.float(speed), C.int(min), C.int(max), formatArg, C.int(flags)) != 0
}

// DragInt3 calls DragInt3V(label, value, 1.0, 0.0, 0.0, "%d", SliderFlagsNone).
func DragInt3(label string, value *[3]int32) bool {
	return DragInt3V(label, value, 1.0, 0.0, 0.0, "%d", SliderFlagsNone)
}

// DragInt4V creates a draggable slider for a 4D vector.
func DragInt4V(label string, values *[4]int32, speed float32, min, max int32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.int)(&values[0])
	return C.iggDragIntN(labelArg, cvalues, 4, C.float(speed), C.int(min), C.int(max), formatArg, C.int(flags)) != 0
}

// DragInt4 calls DragInt4V(label, value, 1.0, 0.0, 0.0, "%d", SliderFlagsNone).
func DragInt4(label string, value *[4]int32) bool {
	return DragInt4V(label, value, 1.0, 0.0, 0.0, "%d", SliderFlagsNone)
}

// DragIntRange2V creates a draggable slider in ints range.
func DragIntRange2V(label string, currentMin *int32, currentMax *int32, speed float32, min int, max int, format string, formatMax string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	currentMinArg, currentMinFin := wrapInt32(currentMin)
	defer currentMinFin()
	currentMaxArg, currentMaxFin := wrapInt32(currentMax)
	defer currentMaxFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	formatMaxArg, formatMaxFin := wrapString(formatMax)
	defer formatMaxFin()
	return C.iggDragIntRange2V(labelArg, currentMinArg, currentMaxArg, C.float(speed), C.int(min), C.int(max), formatArg, formatMaxArg, C.int(flags)) != 0
}

// DragIntRange2 calls DragIntRange2V(label, currentMin, currentMax, 1, 0, 0, "%d", "%d", SliderFlagsNone).
func DragIntRange2(label string, currentMin *int32, currentMax *int32) bool {
	return DragIntRange2V(label, currentMin, currentMax, 1, 0, 0, "%d", "%d", SliderFlagsNone)
}

// SliderFloatV creates a slider for floats.
func SliderFloatV(label string, value *float32, min, max float32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapFloat(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggSliderFloat(labelArg, valueArg, C.float(min), C.float(max), formatArg, C.int(flags)) != 0
}

// SliderFloat calls SliderIntV(label, value, min, max, "%.3f", SliderFlagsNone).
func SliderFloat(label string, value *float32, min, max float32) bool {
	return SliderFloatV(label, value, min, max, "%.3f", SliderFlagsNone)
}

// SliderFloat2V creates slider for a 2D vector.
func SliderFloat2V(label string, values *[2]float32, min, max float32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.float)(&values[0])
	return C.iggSliderFloatN(labelArg, cvalues, 2, C.float(min), C.float(max), formatArg, C.int(flags)) != 0
}

// SliderFloat2 calls SliderFloat2V(label, values, min, max, "%.3f", SliderFlagsNone).
func SliderFloat2(label string, values *[2]float32, min, max float32) bool {
	return SliderFloat2V(label, values, min, max, "%.3f", SliderFlagsNone)
}

// SliderFloat3V creates slider for a 3D vector.
func SliderFloat3V(label string, values *[3]float32, min, max float32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.float)(&values[0])
	return C.iggSliderFloatN(labelArg, cvalues, 3, C.float(min), C.float(max), formatArg, C.int(flags)) != 0
}

// SliderFloat3 calls SliderFloat3V(label, values, min, max, "%.3f", SliderFlagsNone).
func SliderFloat3(label string, values *[3]float32, min, max float32) bool {
	return SliderFloat3V(label, values, min, max, "%.3f", SliderFlagsNone)
}

// SliderFloat4V creates slider for a 4D vector.
func SliderFloat4V(label string, values *[4]float32, min, max float32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.float)(&values[0])
	return C.iggSliderFloatN(labelArg, cvalues, 4, C.float(min), C.float(max), formatArg, C.int(flags)) != 0
}

// SliderFloat4 calls SliderFloat3V(label, values, min, max, "%.3f", SliderFlagsNone).
func SliderFloat4(label string, values *[4]float32, min, max float32) bool {
	return SliderFloat4V(label, values, min, max, "%.3f", SliderFlagsNone)
}

// SliderIntV creates a slider for integers.
func SliderIntV(label string, value *int32, min, max int32, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapInt32(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggSliderInt(labelArg, valueArg, C.int(min), C.int(max), formatArg, C.int(flags)) != 0
}

// SliderInt calls SliderIntV(label, value, min, max, "%d", SliderFlagsNone).
func SliderInt(label string, value *int32, min, max int32) bool {
	return SliderIntV(label, value, min, max, "%d", SliderFlagsNone)
}

// SliderInt2V creates slider for a 2D vector.
func SliderInt2V(label string, values *[2]int32, min, max int, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.int)(&values[0])
	return C.iggSliderIntN(labelArg, cvalues, 2, C.int(min), C.int(max), formatArg, C.int(flags)) != 0
}

// SliderInt2 calls SliderInt2V(label, values, min, max, "%d", SliderFlagsNone).
func SliderInt2(label string, values *[2]int32, min, max int) bool {
	return SliderInt2V(label, values, min, max, "%d", SliderFlagsNone)
}

// SliderInt3V creates slider for a 3D vector.
func SliderInt3V(label string, values *[3]int32, min, max int, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.int)(&values[0])
	return C.iggSliderIntN(labelArg, cvalues, 3, C.int(min), C.int(max), formatArg, C.int(flags)) != 0
}

// SliderInt3 calls SliderInt3V(label, values, min, max, "%d", SliderFlagsNone).
func SliderInt3(label string, values *[3]int32, min, max int) bool {
	return SliderInt3V(label, values, min, max, "%d", SliderFlagsNone)
}

// SliderInt4V creates slider for a 4D vector.
func SliderInt4V(label string, values *[4]int32, min, max int, format string, flags SliderFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.int)(&values[0])
	return C.iggSliderIntN(labelArg, cvalues, 4, C.int(min), C.int(max), formatArg, C.int(flags)) != 0
}

// SliderInt4 calls SliderInt4V(label, values, min, max, "%d", SliderFlagsNone).
func SliderInt4(label string, values *[4]int32, min, max int) bool {
	return SliderInt4V(label, values, min, max, "%d", SliderFlagsNone)
}

// VSliderFloatV creates a vertically oriented slider for floats.
func VSliderFloatV(label string, size Vec2, value *float32, min, max float32, format string, flags SliderFlags) bool {
	sizeArg, _ := size.wrapped()
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapFloat(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggVSliderFloat(labelArg, sizeArg, valueArg, C.float(min), C.float(max), formatArg, C.int(flags)) != 0
}

// VSliderFloat calls VSliderIntV(label, size, value, min, max, "%.3f", SliderFlagsNone).
func VSliderFloat(label string, size Vec2, value *float32, min, max float32) bool {
	return VSliderFloatV(label, size, value, min, max, "%.3f", SliderFlagsNone)
}

// VSliderIntV creates a vertically oriented slider for integers.
func VSliderIntV(label string, size Vec2, value *int32, min, max int32, format string, flags SliderFlags) bool {
	sizeArg, _ := size.wrapped()
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapInt32(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggVSliderInt(labelArg, sizeArg, valueArg, C.int(min), C.int(max), formatArg, C.int(flags)) != 0
}

// VSliderInt calls VSliderIntV(label, size, value, min, max, "%d", SliderFlagsNone).
func VSliderInt(label string, size Vec2, value *int32, min, max int32) bool {
	return VSliderIntV(label, size, value, min, max, "%d", SliderFlagsNone)
}

// InputTextFlags for InputTextV(), etc.
type InputTextFlags int

const (
	// InputTextFlagsNone sets everything default.
	InputTextFlagsNone InputTextFlags = 0
	// InputTextFlagsCharsDecimal allows 0123456789.+-.
	InputTextFlagsCharsDecimal InputTextFlags = 1 << 0
	// InputTextFlagsCharsHexadecimal allow 0123456789ABCDEFabcdef.
	InputTextFlagsCharsHexadecimal InputTextFlags = 1 << 1
	// InputTextFlagsCharsUppercase turns a..z into A..Z.
	InputTextFlagsCharsUppercase InputTextFlags = 1 << 2
	// InputTextFlagsCharsNoBlank filters out spaces, tabs.
	InputTextFlagsCharsNoBlank InputTextFlags = 1 << 3
	// InputTextFlagsAutoSelectAll selects entire text when first taking mouse focus.
	InputTextFlagsAutoSelectAll InputTextFlags = 1 << 4
	// InputTextFlagsEnterReturnsTrue returns 'true' when Enter is pressed (as opposed to when the value was modified).
	InputTextFlagsEnterReturnsTrue InputTextFlags = 1 << 5
	// InputTextFlagsCallbackCompletion for callback on pressing TAB (for completion handling).
	InputTextFlagsCallbackCompletion InputTextFlags = 1 << 6
	// InputTextFlagsCallbackHistory for callback on pressing Up/Down arrows (for history handling).
	InputTextFlagsCallbackHistory InputTextFlags = 1 << 7
	// InputTextFlagsCallbackAlways for callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackAlways InputTextFlags = 1 << 8
	// InputTextFlagsCallbackCharFilter for callback on character inputs to replace or discard them.
	// Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsCallbackCharFilter InputTextFlags = 1 << 9
	// InputTextFlagsAllowTabInput when pressing TAB to input a '\t' character into the text field.
	InputTextFlagsAllowTabInput InputTextFlags = 1 << 10
	// InputTextFlagsCtrlEnterForNewLine in multi-line mode, unfocus with Enter, add new line with Ctrl+Enter
	// (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsCtrlEnterForNewLine InputTextFlags = 1 << 11
	// InputTextFlagsNoHorizontalScroll disables following the cursor horizontally.
	InputTextFlagsNoHorizontalScroll InputTextFlags = 1 << 12
	// InputTextFlagsAlwaysInsertMode was renamed to InputTextFlagsAlwaysOverwriteMode to reflect its actual behavior and will be removed in v5.
	// Deprecated: Use InputTextFlagsAlwaysOverwriteMode.
	InputTextFlagsAlwaysInsertMode InputTextFlags = 1 << 13
	// InputTextFlagsAlwaysOverwriteMode sets overwrite mode.
	InputTextFlagsAlwaysOverwriteMode InputTextFlags = 1 << 13
	// InputTextFlagsReadOnly sets read-only mode.
	InputTextFlagsReadOnly InputTextFlags = 1 << 14
	// InputTextFlagsPassword sets password mode, display all characters as '*'.
	InputTextFlagsPassword InputTextFlags = 1 << 15
	// InputTextFlagsNoUndoRedo disables undo/redo. Note that input text owns the text data while active,
	// if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsNoUndoRedo InputTextFlags = 1 << 16
	// InputTextFlagsCharsScientific allows 0123456789.+-*/eE (Scientific notation input).
	InputTextFlagsCharsScientific InputTextFlags = 1 << 17
	// inputTextFlagsCallbackResize for callback on buffer capacity change requests.
	inputTextFlagsCallbackResize InputTextFlags = 1 << 18
	// ImGuiInputTextFlagsCallbackEdit for callback on any edit (note that InputText() already returns true on edit, the callback is useful mainly to manipulate the underlying buffer while focus is active).
	ImGuiInputTextFlagsCallbackEdit InputTextFlags = 1 << 19
)

// InputTextV creates a text field for dynamic text input.
//
// Contrary to the original library, this wrapper does not limit the maximum number of possible characters.
// Dynamic resizing of the internal buffer is handled within the wrapper and the user will never be called for such requests.
//
// The provided callback is called for any of the requested InputTextFlagsCallback* flags.
//
// To implement a character limit, provide a callback that drops input characters when the requested length has been reached.
func InputTextV(label string, text *string, flags InputTextFlags, cb InputTextCallback) bool {
	return inputTextSingleline(label, nil, text, flags, cb)
}

// InputText calls InputTextV(label, text, 0, nil).
func InputText(label string, text *string) bool {
	return InputTextV(label, text, 0, nil)
}

// InputTextWithHintV creates a text field for dynamic text input with a hint.
//
// Contrary to the original library, this wrapper does not limit the maximum number of possible characters.
// Dynamic resizing of the internal buffer is handled within the wrapper and the user will never be called for such requests.
//
// The provided callback is called for any of the requested InputTextFlagsCallback* flags.
//
// To implement a character limit, provide a callback that drops input characters when the requested length has been reached.
func InputTextWithHintV(label string, hint string, text *string, flags InputTextFlags, cb InputTextCallback) bool {
	return inputTextSingleline(label, &hint, text, flags, cb)
}

// InputTextWithHint calls InputTextWithHintV(label, hint, text, 0, nil).
func InputTextWithHint(label string, hint string, text *string) bool {
	return InputTextWithHintV(label, hint, text, 0, nil)
}

func inputTextSingleline(label string, hint *string, text *string, flags InputTextFlags, cb InputTextCallback) bool {
	if text == nil {
		panic("text can't be nil")
	}
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	var hintArg *C.char
	var hintFin func()
	if hint != nil {
		hintArg, hintFin = wrapString(*hint)
		defer hintFin()
	}
	state := newInputTextState(*text, cb)
	defer func() {
		*text = state.buf.toGo()
		state.release()
	}()

	return C.iggInputTextSingleline(labelArg, hintArg, (*C.char)(state.buf.ptr), C.uint(state.buf.size),
		C.int(flags|inputTextFlagsCallbackResize), state.key) != 0
}

// InputTextMultilineV provides a field for dynamic text input of multiple lines.
//
// Contrary to the original library, this wrapper does not limit the maximum number of possible characters.
// Dynamic resizing of the internal buffer is handled within the wrapper and the user will never be called for such requests.
//
// The provided callback is called for any of the requested InputTextFlagsCallback* flags.
//
// To implement a character limit, provide a callback that drops input characters when the requested length has been reached.
func InputTextMultilineV(label string, text *string, size Vec2, flags InputTextFlags, cb InputTextCallback) bool {
	if text == nil {
		panic("text can't be nil")
	}
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	sizeArg, _ := size.wrapped()
	state := newInputTextState(*text, cb)
	defer func() {
		*text = state.buf.toGo()
		state.release()
	}()

	return C.iggInputTextMultiline(labelArg, (*C.char)(state.buf.ptr), C.uint(state.buf.size), sizeArg,
		C.int(flags|inputTextFlagsCallbackResize), state.key) != 0
}

// InputTextMultiline calls InputTextMultilineV(label, text, Vec2{0,0}, 0, nil).
func InputTextMultiline(label string, text *string) bool {
	return InputTextMultilineV(label, text, Vec2{}, 0, nil)
}

// InputIntV creates a input field for integer type.
func InputIntV(label string, value *int32, step int, stepFast int, flags InputTextFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapInt32(value)
	defer valueFin()

	return C.iggInputInt(labelArg, valueArg, C.int(step), C.int(stepFast), C.int(flags)) != 0
}

// InputInt calls InputIntV(label, value, 1, 100, 0).
func InputInt(label string, value *int32) bool {
	return InputIntV(label, value, 1, 100, 0)
}

// CollapsingHeader calls CollapsingHeaderV(label, 0).
func CollapsingHeader(label string) bool {
	return CollapsingHeaderV(label, 0)
}

// CollapsingHeaderV adds a collapsing header with TreeNode flags.
func CollapsingHeaderV(label string, flags TreeNodeFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggCollapsingHeader(labelArg, C.int(flags)) != 0
}

// TreeNodeFlags for TreeNodeV(), CollapsingHeaderV(), etc.
type TreeNodeFlags int

const (
	// TreeNodeFlagsNone default = 0.
	TreeNodeFlagsNone TreeNodeFlags = 0
	// TreeNodeFlagsSelected draws as selected.
	TreeNodeFlagsSelected TreeNodeFlags = 1 << 0
	// TreeNodeFlagsFramed draws frame with background (e.g. for CollapsingHeader).
	TreeNodeFlagsFramed TreeNodeFlags = 1 << 1
	// TreeNodeFlagsAllowItemOverlap hit testing to allow subsequent widgets to overlap this one.
	TreeNodeFlagsAllowItemOverlap TreeNodeFlags = 1 << 2
	// TreeNodeFlagsNoTreePushOnOpen doesn't do a TreePush() when open
	// (e.g. for CollapsingHeader) = no extra indent nor pushing on ID stack.
	TreeNodeFlagsNoTreePushOnOpen TreeNodeFlags = 1 << 3
	// TreeNodeFlagsNoAutoOpenOnLog doesn't automatically and temporarily open node when Logging is active
	// (by default logging will automatically open tree nodes).
	TreeNodeFlagsNoAutoOpenOnLog TreeNodeFlags = 1 << 4
	// TreeNodeFlagsDefaultOpen defaults node to be open.
	TreeNodeFlagsDefaultOpen TreeNodeFlags = 1 << 5
	// TreeNodeFlagsOpenOnDoubleClick needs double-click to open node.
	TreeNodeFlagsOpenOnDoubleClick TreeNodeFlags = 1 << 6
	// TreeNodeFlagsOpenOnArrow opens only when clicking on the arrow part.
	// If TreeNodeFlagsOpenOnDoubleClick is also set, single-click arrow or double-click all box to open.
	TreeNodeFlagsOpenOnArrow TreeNodeFlags = 1 << 7
	// TreeNodeFlagsLeaf allows no collapsing, no arrow (use as a convenience for leaf nodes).
	TreeNodeFlagsLeaf TreeNodeFlags = 1 << 8
	// TreeNodeFlagsBullet displays a bullet instead of an arrow.
	TreeNodeFlagsBullet TreeNodeFlags = 1 << 9
	// TreeNodeFlagsFramePadding uses FramePadding (even for an unframed text node) to
	// vertically align text baseline to regular widget height. Equivalent to calling AlignTextToFramePadding().
	TreeNodeFlagsFramePadding TreeNodeFlags = 1 << 10
	// TreeNodeFlagsSpanAvailWidth extends hit box to the right-most edge, even if not framed.
	// This is not the default in order to allow adding other items on the same line.
	// In the future we may refactor the hit system to be front-to-back, allowing natural overlaps
	// and then this can become the default.
	TreeNodeFlagsSpanAvailWidth TreeNodeFlags = 1 << 11
	// TreeNodeFlagsSpanFullWidth extends hit box to the left-most and right-most edges (bypass the indented area).
	TreeNodeFlagsSpanFullWidth TreeNodeFlags = 1 << 12
	// TreeNodeFlagsNavLeftJumpsBackHere (WIP) Nav: left direction may move to this TreeNode() from any of its child
	// (items submitted between TreeNode and TreePop).
	TreeNodeFlagsNavLeftJumpsBackHere TreeNodeFlags = 1 << 13
	// TreeNodeFlagsCollapsingHeader combines TreeNodeFlagsFramed and TreeNodeFlagsNoAutoOpenOnLog.
	TreeNodeFlagsCollapsingHeader = TreeNodeFlagsFramed | TreeNodeFlagsNoTreePushOnOpen | TreeNodeFlagsNoAutoOpenOnLog
)

// TreeNodeV returns true if the tree branch is to be rendered. Call TreePop() in this case.
func TreeNodeV(label string, flags TreeNodeFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggTreeNode(labelArg, C.int(flags)) != 0
}

// TreeNode calls TreeNodeV(label, 0).
func TreeNode(label string) bool {
	return TreeNodeV(label, 0)
}

// TreeNodef calls TreeNode(fmt.Sprintf(format, v...)) .
func TreeNodef(format string, v ...interface{}) bool {
	return TreeNode(fmt.Sprintf(format, v...))
}

// TreePop finishes a tree branch. This has to be called for a matching TreeNodeV call returning true.
func TreePop() {
	C.iggTreePop()
}

// SetNextItemOpen sets the open/collapsed state of the following tree node.
func SetNextItemOpen(open bool, cond Condition) {
	C.iggSetNextItemOpen(castBool(open), C.int(cond))
}

// TreeNodeToLabelSpacing returns the horizontal distance preceding label for a regular unframed TreeNode.
func TreeNodeToLabelSpacing() float32 {
	return float32(C.iggGetTreeNodeToLabelSpacing())
}

// SelectableFlags for SelectableV().
type SelectableFlags int

const (
	// SelectableFlagsNone default = 0.
	SelectableFlagsNone SelectableFlags = 0
	// SelectableFlagsDontClosePopups makes clicking the selectable not close any parent popup windows.
	SelectableFlagsDontClosePopups SelectableFlags = 1 << 0
	// SelectableFlagsSpanAllColumns allows the selectable frame to span all columns (text will still fit in current column).
	SelectableFlagsSpanAllColumns SelectableFlags = 1 << 1
	// SelectableFlagsAllowDoubleClick generates press events on double clicks too.
	SelectableFlagsAllowDoubleClick SelectableFlags = 1 << 2
	// SelectableFlagsDisabled disallows selection and displays text in a greyed out color.
	SelectableFlagsDisabled SelectableFlags = 1 << 3
	// SelectableFlagsAllowItemOverlap hit testing to allow subsequent widgets to overlap this one (WIP).
	SelectableFlagsAllowItemOverlap SelectableFlags = 1 << 4
)

// SelectableV returns true if the user clicked it, so you can modify your selection state.
// flags are the SelectableFlags to apply.
// size.x==0.0: use remaining width, size.x>0.0: specify width.
// size.y==0.0: use label height, size.y>0.0: specify height.
func SelectableV(label string, selected bool, flags SelectableFlags, size Vec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	sizeArg, _ := size.wrapped()
	return C.iggSelectable(labelArg, castBool(selected), C.int(flags), sizeArg) != 0
}

// Selectable calls SelectableV(label, false, 0, Vec2{0, 0}).
func Selectable(label string) bool {
	return SelectableV(label, false, 0, Vec2{})
}

// BeginListBoxV opens a framed scrolling region.
// - This is essentially a thin wrapper to using BeginChild/EndChild with some stylistic changes.
// - The BeginListBox()/EndListBox() api allows you to manage your contents and selection state however you want it, by creating e.g. Selectable() or any items.
// - The simplified/old ListBox() api are helpers over BeginListBox()/EndListBox() which are kept available for convenience purpose. This is analoguous to how Combos are created.
// - Choose frame width:   size.x > 0.0f: custom  /  size.x < 0.0f or -FLT_MIN: right-align   /  size.x = 0.0f (default): use current ItemWidth
// - Choose frame height:  size.y > 0.0f: custom  /  size.y < 0.0f or -FLT_MIN: bottom-align  /  size.y = 0.0f (default): arbitrary default height which can fit ~7 items.
func BeginListBoxV(label string, size Vec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	sizeArg, _ := size.wrapped()
	return C.iggBeginListBox(labelArg, sizeArg) != 0
}

// BeginListBox calls BeginListBoxV(label, Vec2{}).
func BeginListBox(label string) bool {
	return BeginListBoxV(label, Vec2{})
}

// EndListBox closes the scope for the previously opened ListBox.
// only call EndListBox() if BeginListBox() returned true!
func EndListBox() {
	C.iggEndListBox()
}

// ListBoxV creates a list of selectables of given items with equal height, enclosed with header and footer.
// This version accepts a custom item height.
// The function returns true if the selection was changed. The value of currentItem will indicate the new selected item.
func ListBoxV(label string, currentItem *int32, items []string, heightItems int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	valueArg, valueFin := wrapInt32(currentItem)
	defer valueFin()

	itemsCount := len(items)

	argv := make([]*C.char, itemsCount)
	itemFins := make([]func(), 0, itemsCount)
	defer func() {
		for _, itemFin := range itemFins {
			itemFin()
		}
	}()
	for i, item := range items {
		itemArg, itemFin := wrapString(item)
		itemFins = append(itemFins, itemFin)
		argv[i] = itemArg
	}

	return C.iggListBox(labelArg, valueArg, &argv[0], C.int(itemsCount), C.int(heightItems)) != 0
}

// ListBox calls ListBoxV(label, currentItem, items, -1)
// The function returns true if the selection was changed. The value of currentItem will indicate the new selected item.
func ListBox(label string, currentItem *int32, items []string) bool {
	return ListBoxV(label, currentItem, items, -1)
}

// PlotLines draws an array of floats as a line graph.
// It calls PlotLinesV using no overlay text and automatically calculated scale and graph size.
func PlotLines(label string, values []float32) {
	PlotLinesV(label, values, 0, "", math.MaxFloat32, math.MaxFloat32, Vec2{})
}

// PlotLinesV draws an array of floats as a line graph with additional options.
// valuesOffset specifies an offset into the values array at which to start drawing, wrapping around when the end of the values array is reached.
// overlayText specifies a string to print on top of the graph.
// scaleMin and scaleMax define the scale of the y axis, if either is math.MaxFloat32 that value is calculated from the input data.
// graphSize defines the size of the graph, if either coordinate is zero the default size for that direction is used.
func PlotLinesV(label string, values []float32, valuesOffset int, overlayText string, scaleMin float32, scaleMax float32, graphSize Vec2) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	valuesCount := len(values)
	valuesArray := make([]C.float, valuesCount)
	for i, value := range values {
		valuesArray[i] = C.float(value)
	}

	var overlayTextArg *C.char
	if overlayText != "" {
		var overlayTextFinisher func()
		overlayTextArg, overlayTextFinisher = wrapString(overlayText)
		defer overlayTextFinisher()
	}

	graphSizeArg, _ := graphSize.wrapped()

	C.iggPlotLines(labelArg, &valuesArray[0], C.int(valuesCount), C.int(valuesOffset), overlayTextArg, C.float(scaleMin), C.float(scaleMax), graphSizeArg)
}

// PlotHistogram draws an array of floats as a bar graph.
// It calls PlotHistogramV using no overlay text and automatically calculated scale and graph size.
func PlotHistogram(label string, values []float32) {
	PlotHistogramV(label, values, 0, "", math.MaxFloat32, math.MaxFloat32, Vec2{})
}

// PlotHistogramV draws an array of floats as a bar graph with additional options.
// valuesOffset specifies an offset into the values array at which to start drawing, wrapping around when the end of the values array is reached.
// overlayText specifies a string to print on top of the graph.
// scaleMin and scaleMax define the scale of the y axis, if either is math.MaxFloat32 that value is calculated from the input data.
// graphSize defines the size of the graph, if either coordinate is zero the default size for that direction is used.
func PlotHistogramV(label string, values []float32, valuesOffset int, overlayText string, scaleMin float32, scaleMax float32, graphSize Vec2) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	valuesCount := len(values)
	valuesArray := make([]C.float, valuesCount)
	for i, value := range values {
		valuesArray[i] = C.float(value)
	}

	var overlayTextArg *C.char
	if overlayText != "" {
		var overlayTextFinisher func()
		overlayTextArg, overlayTextFinisher = wrapString(overlayText)
		defer overlayTextFinisher()
	}

	graphSizeArg, _ := graphSize.wrapped()

	C.iggPlotHistogram(labelArg, &valuesArray[0], C.int(valuesCount), C.int(valuesOffset), overlayTextArg, C.float(scaleMin), C.float(scaleMax), graphSizeArg)
}

// SetTooltip sets a text tooltip under the mouse-cursor, typically use with IsItemHovered().
// Overrides any previous call to SetTooltip().
func SetTooltip(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()
	C.iggSetTooltip(textArg)
}

// SetTooltipf calls SetTooltip(fmt.Sprintf(format, v...)) .
func SetTooltipf(format string, v ...interface{}) {
	SetTooltip(fmt.Sprintf(format, v...))
}

// BeginTooltip begins/appends to a tooltip window. Used to create full-featured tooltip (with any kind of contents).
// Requires a call to EndTooltip().
func BeginTooltip() {
	C.iggBeginTooltip()
}

// EndTooltip closes the previously started tooltip window.
func EndTooltip() {
	C.iggEndTooltip()
}

// BeginMainMenuBar creates and appends to a full screen menu-bar.
// If the return value is true, then EndMainMenuBar() must be called!
func BeginMainMenuBar() bool {
	return C.iggBeginMainMenuBar() != 0
}

// EndMainMenuBar finishes a main menu bar.
// Only call EndMainMenuBar() if BeginMainMenuBar() returns true!
func EndMainMenuBar() {
	C.iggEndMainMenuBar()
}

// BeginMenuBar appends to menu-bar of current window.
// This requires WindowFlagsMenuBar flag set on parent window.
// If the return value is true, then EndMenuBar() must be called!
func BeginMenuBar() bool {
	return C.iggBeginMenuBar() != 0
}

// EndMenuBar finishes a menu bar.
// Only call EndMenuBar() if BeginMenuBar() returns true!
func EndMenuBar() {
	C.iggEndMenuBar()
}

// BeginMenuV creates a sub-menu entry.
// If the return value is true, then EndMenu() must be called!
func BeginMenuV(label string, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggBeginMenu(labelArg, castBool(enabled)) != 0
}

// BeginMenu calls BeginMenuV(label, true).
func BeginMenu(label string) bool {
	return BeginMenuV(label, true)
}

// EndMenu finishes a sub-menu entry.
// Only call EndMenu() if BeginMenu() returns true!
func EndMenu() {
	C.iggEndMenu()
}

// MenuItemV adds a menu item with given label.
// Returns true if the item is selected.
// If selected is not nil, it will be toggled when true is returned.
// Shortcuts are displayed for convenience but not processed by ImGui at the moment.
func MenuItemV(label string, shortcut string, selected bool, enabled bool) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	shortcutArg, shortcutFin := wrapString(shortcut)
	defer shortcutFin()
	return C.iggMenuItem(labelArg, shortcutArg, castBool(selected), castBool(enabled)) != 0
}

// MenuItem calls MenuItemV(label, "", false, true).
func MenuItem(label string) bool {
	return MenuItemV(label, "", false, true)
}

// Columns API is Legacy (2021: prefer using Tables!).
// - You can also use SameLineV(pos_x, 0) to mimic simplified columns.
//
// Columns calls ColumnsV(1, "", false).
func Columns() {
	ColumnsV(1, "", false)
}

// ColumnsV creates a column layout of the specified number of columns.
func ColumnsV(count int, label string, border bool) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	C.iggColumns(C.int(count), labelArg, castBool(border))
}

// NextColumn next column, defaults to current row or next row if the current row is finished.
func NextColumn() {
	C.iggNextColumn()
}

// ColumnIndex get current column index.
func ColumnIndex() int {
	return int(C.iggGetColumnIndex())
}

// ColumnWidth calls ColumnWidthV(-1).
func ColumnWidth() int {
	return ColumnWidthV(-1)
}

// ColumnWidthV get column width (in pixels). pass -1 to use current column.
func ColumnWidthV(index int) int {
	return int(C.iggGetColumnWidth(C.int(index)))
}

// SetColumnWidth sets column width (in pixels). pass -1 to use current column.
func SetColumnWidth(index int, width float32) {
	C.iggSetColumnWidth(C.int(index), C.float(width))
}

// ColumnOffset calls ColumnOffsetV(-1).
func ColumnOffset() float32 {
	return ColumnOffsetV(-1)
}

// ColumnOffsetV get position of column line (in pixels, from the left side of the contents region). pass -1 to use
// current column, otherwise 0..GetColumnsCount() inclusive. column 0 is typically 0.0.
func ColumnOffsetV(index int) float32 {
	return float32(C.iggGetColumnOffset(C.int(index)))
}

// SetColumnOffset set position of column line (in pixels, from the left side of the contents region). pass -1 to use
// current column.
func SetColumnOffset(index int, offsetX float32) {
	C.iggSetColumnOffset(C.int(index), C.float(offsetX))
}

// ColumnsCount returns number of current columns.
func ColumnsCount() int {
	return int(C.iggGetColumnsCount())
}

// TabBarFlags for BeginTabBarV().
type TabBarFlags int

const (
	// TabBarFlagsNone default = 0.
	TabBarFlagsNone TabBarFlags = 0
	// TabBarFlagsReorderable Allow manually dragging tabs to re-order them + New tabs are appended at the end of list.
	TabBarFlagsReorderable TabBarFlags = 1 << 0
	// TabBarFlagsAutoSelectNewTabs Automatically select new tabs when they appear.
	TabBarFlagsAutoSelectNewTabs TabBarFlags = 1 << 1
	// TabBarFlagsTabListPopupButton Disable buttons to open the tab list popup.
	TabBarFlagsTabListPopupButton TabBarFlags = 1 << 2
	// TabBarFlagsNoCloseWithMiddleMouseButton Disable behavior of closing tabs (that are submitted with p_open != NULL)
	// with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabBarFlagsNoCloseWithMiddleMouseButton TabBarFlags = 1 << 3
	// TabBarFlagsNoTabListScrollingButtons Disable scrolling buttons (apply when fitting policy is
	// TabBarFlagsFittingPolicyScroll).
	TabBarFlagsNoTabListScrollingButtons TabBarFlags = 1 << 4
	// TabBarFlagsNoTooltip Disable tooltips when hovering a tab.
	TabBarFlagsNoTooltip TabBarFlags = 1 << 5
	// TabBarFlagsFittingPolicyResizeDown Resize tabs when they don't fit.
	TabBarFlagsFittingPolicyResizeDown TabBarFlags = 1 << 6
	// TabBarFlagsFittingPolicyScroll Add scroll buttons when tabs don't fit.
	TabBarFlagsFittingPolicyScroll TabBarFlags = 1 << 7
	// TabBarFlagsFittingPolicyMask combines
	// TabBarFlagsFittingPolicyResizeDown and TabBarFlagsFittingPolicyScroll.
	TabBarFlagsFittingPolicyMask = TabBarFlagsFittingPolicyResizeDown | TabBarFlagsFittingPolicyScroll
	// TabBarFlagsFittingPolicyDefault alias for TabBarFlagsFittingPolicyResizeDown.
	TabBarFlagsFittingPolicyDefault = TabBarFlagsFittingPolicyResizeDown
)

// BeginTabBarV create and append into a TabBar.
func BeginTabBarV(strID string, flags TabBarFlags) bool {
	idArg, idFin := wrapString(strID)
	defer idFin()

	return C.iggBeginTabBar(idArg, C.int(flags)) != 0
}

// BeginTabBar calls BeginTabBarV(strId, 0).
func BeginTabBar(strID string) bool {
	return BeginTabBarV(strID, 0)
}

// EndTabBar only call EndTabBar() if BeginTabBar() returns true!
func EndTabBar() {
	C.iggEndTabBar()
}

// TabItemFlags for BeginTabItemV(), BeginTabButtonV().
type TabItemFlags int

const (
	// TabItemFlagsNone default = 0.
	TabItemFlagsNone TabItemFlags = 0
	// TabItemFlagsUnsavedDocument Append '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. Also: tab is selected on closure and closure is deferred by one frame to allow code to undo it
	// without flicker.
	TabItemFlagsUnsavedDocument TabItemFlags = 1 << 0
	// TabItemFlagsSetSelected Trigger flag to programmatically make the tab selected when calling BeginTabItem().
	TabItemFlagsSetSelected TabItemFlags = 1 << 1
	// TabItemFlagsNoCloseWithMiddleMouseButton  Disable behavior of closing tabs (that are submitted with
	// p_open != NULL) with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabItemFlagsNoCloseWithMiddleMouseButton TabItemFlags = 1 << 2
	// TabItemFlagsNoPushID Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem().
	TabItemFlagsNoPushID TabItemFlags = 1 << 3
	// TabItemFlagsNoTooltip Disable tooltip for the given tab.
	TabItemFlagsNoTooltip TabItemFlags = 1 << 4
	// TabItemFlagsNoReorder Disable reordering this tab or having another tab cross over this tab.
	TabItemFlagsNoReorder TabItemFlags = 1 << 5
	// TabItemFlagsLeading Enforce the tab position to the left of the tab bar (after the tab list popup button).
	TabItemFlagsLeading TabItemFlags = 1 << 6
	// TabItemFlagsTrailing Enforce the tab position to the right of the tab bar (before the scrolling buttons).
	TabItemFlagsTrailing TabItemFlags = 1 << 7
)

// BeginTabItemV create a Tab. Returns true if the Tab is selected.
func BeginTabItemV(label string, open *bool, flags TabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	openArg, openFin := wrapBool(open)
	defer openFin()

	return C.iggBeginTabItem(labelArg, openArg, C.int(flags)) != 0
}

// BeginTabItem calls BeginTabItemV(label, nil, 0).
func BeginTabItem(label string) bool {
	return BeginTabItemV(label, nil, 0)
}

// EndTabItem finishes a tab item.
// Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem().
func EndTabItem() {
	C.iggEndTabItem()
}

// TabItemButtonV create a Tab behaving like a button. return true when clicked. cannot be selected in the tab bar.
func TabItemButtonV(label string, flags TabItemFlags) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggTabItemButton(labelArg, C.int(flags)) != 0
}

// TabItemButton calls TabItemButtonV(label, 0).
func TabItemButton(label string) bool {
	return TabItemButtonV(label, 0)
}

// SetTabItemClosed notify TabBar or Docking system of a closed tab/window ahead
// (useful to reduce visual flicker on reorderable tab bars). For tab-bar: call
// after BeginTabBar() and before Tab submissions. Otherwise call with a window name.
func SetTabItemClosed(tabOrDockedWindowLabel string) {
	labelArg, labelFin := wrapString(tabOrDockedWindowLabel)
	defer labelFin()
	C.iggSetTabItemClosed(labelArg)
}
