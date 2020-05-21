package imgui

// #include "wrapper/Widgets.h"
import "C"

import "math"

// Text adds formatted text. See PushTextWrapPosV() or PushStyleColorV() for modifying the output.
// Without any modified style stack, the text is unformatted.
func Text(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()
	// Internally we use ImGui::TextUnformatted, for the most direct call.
	C.iggTextUnformatted(textArg)
}

// LabelText adds text+label aligned the same way as value+label widgets.
func LabelText(label, text string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	textArg, textFin := wrapString(text)
	defer textFin()
	C.iggLabelText(labelArg, textArg)
}

// ButtonV returning true if it is pressed.
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

// InvisibleButtonV returning true if it is pressed.
func InvisibleButtonV(id string, size Vec2) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	sizeArg, _ := size.wrapped()
	return C.iggInvisibleButton(idArg, sizeArg) != 0
}

// InvisibleButton calls InvisibleButtonV(id, Vec2{0,0}).
func InvisibleButton(id string) bool {
	return InvisibleButtonV(id, Vec2{})
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

// RadioButton returning true if it is pressed and active indicates if it is selected.
func RadioButton(id string, active bool) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	return C.iggRadioButton(idArg, castBool(active)) != 0
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

// ProgressBar calls ProgressBarV(fraction, Vec2{X: -1, Y: 0}, "").
func ProgressBar(fraction float32) {
	ProgressBarV(fraction, Vec2{X: -1, Y: 0}, "")
}

const (
	// ComboFlagNone default = 0
	ComboFlagNone = 0
	// ComboFlagPopupAlignLeft aligns the popup toward the left by default.
	ComboFlagPopupAlignLeft = 1 << 0
	// ComboFlagHeightSmall has max ~4 items visible.
	// Tip: If you want your combo popup to be a specific size you can use SetNextWindowSizeConstraints() prior to calling BeginCombo().
	ComboFlagHeightSmall = 1 << 1
	// ComboFlagHeightRegular has max ~8 items visible (default).
	ComboFlagHeightRegular = 1 << 2
	// ComboFlagHeightLarge has max ~20 items visible.
	ComboFlagHeightLarge = 1 << 3
	// ComboFlagHeightLargest has as many fitting items as possible.
	ComboFlagHeightLargest = 1 << 4
	// ComboFlagNoArrowButton displays on the preview box without the square arrow button.
	ComboFlagNoArrowButton = 1 << 5
	// ComboFlagNoPreview displays only a square arrow button.
	ComboFlagNoPreview = 1 << 6
)

// BeginComboV creates a combo box with complete control over the content to the user.
// Call EndCombo() if this function returns true.
// flags are the ComboFlags to apply.
func BeginComboV(label, previewValue string, flags int) bool {
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

// DragFloatV creates a draggable slider for floats.
func DragFloatV(label string, value *float32, speed, min, max float32, format string, power float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapFloat(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggDragFloat(labelArg, valueArg, C.float(speed), C.float(min), C.float(max), formatArg, C.float(power)) != 0
}

// DragFloat calls DragFloatV(label, value, 1.0, 0.0, 0.0, "%.3f", 1.0).
func DragFloat(label string, value *float32) bool {
	return DragFloatV(label, value, 1.0, 0.0, 0.0, "%.3f", 1.0)
}

// DragIntV creates a draggable slider for integers.
func DragIntV(label string, value *int32, speed float32, min, max int32, format string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapInt32(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggDragInt(labelArg, valueArg, C.float(speed), C.int(min), C.int(max), formatArg) != 0
}

// DragInt calls DragIntV(label, value, 1.0, 0, 0, "%d").
func DragInt(label string, value *int32) bool {
	return DragIntV(label, value, 1.0, 0, 0, "%d")
}

// SliderFloatV creates a slider for floats.
func SliderFloatV(label string, value *float32, min, max float32, format string, power float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapFloat(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggSliderFloat(labelArg, valueArg, C.float(min), C.float(max), formatArg, C.float(power)) != 0
}

// SliderFloat calls SliderIntV(label, value, min, max, "%.3f", 1.0).
func SliderFloat(label string, value *float32, min, max float32) bool {
	return SliderFloatV(label, value, min, max, "%.3f", 1.0)
}

// SliderFloat3V creates slider for a 3D vector.
func SliderFloat3V(label string, values *[3]float32, min, max float32, format string, power float32) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	cvalues := (*C.float)(&values[0])
	return C.iggSliderFloatN(labelArg, cvalues, 3, C.float(min), C.float(max), formatArg, C.float(power)) != 0
}

// SliderFloat3 calls SliderFloat3V(label, values, min, max, "%.3f", 1,0).
func SliderFloat3(label string, values *[3]float32, min, max float32) bool {
	return SliderFloat3V(label, values, min, max, "%.3f", 1.0)
}

// SliderIntV creates a slider for integers.
func SliderIntV(label string, value *int32, min, max int32, format string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapInt32(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggSliderInt(labelArg, valueArg, C.int(min), C.int(max), formatArg) != 0
}

// SliderInt calls SliderIntV(label, value, min, max, "%d").
func SliderInt(label string, value *int32, min, max int32) bool {
	return SliderIntV(label, value, min, max, "%d")
}

// VSliderFloatV creates a vertically oriented slider for floats.
func VSliderFloatV(label string, size Vec2, value *float32, min, max float32, format string, power float32) bool {
	sizeArg, _ := size.wrapped()
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapFloat(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggVSliderFloat(labelArg, sizeArg, valueArg, C.float(min), C.float(max), formatArg, C.float(power)) != 0
}

// VSliderFloat calls VSliderIntV(label, size, value, min, max, "%.3f", 1.0).
func VSliderFloat(label string, size Vec2, value *float32, min, max float32) bool {
	return VSliderFloatV(label, size, value, min, max, "%.3f", 1.0)
}

// VSliderIntV creates a vertically oriented slider for integers.
func VSliderIntV(label string, size Vec2, value *int32, min, max int32, format string) bool {
	sizeArg, _ := size.wrapped()
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	valueArg, valueFin := wrapInt32(value)
	defer valueFin()
	formatArg, formatFin := wrapString(format)
	defer formatFin()
	return C.iggVSliderInt(labelArg, sizeArg, valueArg, C.int(min), C.int(max), formatArg) != 0
}

// VSliderInt calls VSliderIntV(label, size, value, min, max, "%d").
func VSliderInt(label string, size Vec2, value *int32, min, max int32) bool {
	return VSliderIntV(label, size, value, min, max, "%d")
}

const (
	// InputTextFlagsNone sets everything default.
	InputTextFlagsNone = 0
	// InputTextFlagsCharsDecimal allows 0123456789.+-
	InputTextFlagsCharsDecimal = 1 << 0
	// InputTextFlagsCharsHexadecimal allow 0123456789ABCDEFabcdef
	InputTextFlagsCharsHexadecimal = 1 << 1
	// InputTextFlagsCharsUppercase turns a..z into A..Z.
	InputTextFlagsCharsUppercase = 1 << 2
	// InputTextFlagsCharsNoBlank filters out spaces, tabs.
	InputTextFlagsCharsNoBlank = 1 << 3
	// InputTextFlagsAutoSelectAll selects entire text when first taking mouse focus.
	InputTextFlagsAutoSelectAll = 1 << 4
	// InputTextFlagsEnterReturnsTrue returns 'true' when Enter is pressed (as opposed to when the value was modified).
	InputTextFlagsEnterReturnsTrue = 1 << 5
	// InputTextFlagsCallbackCompletion for callback on pressing TAB (for completion handling).
	InputTextFlagsCallbackCompletion = 1 << 6
	// InputTextFlagsCallbackHistory for callback on pressing Up/Down arrows (for history handling).
	InputTextFlagsCallbackHistory = 1 << 7
	// InputTextFlagsCallbackAlways for callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackAlways = 1 << 8
	// InputTextFlagsCallbackCharFilter for callback on character inputs to replace or discard them.
	// Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsCallbackCharFilter = 1 << 9
	// InputTextFlagsAllowTabInput when pressing TAB to input a '\t' character into the text field.
	InputTextFlagsAllowTabInput = 1 << 10
	// InputTextFlagsCtrlEnterForNewLine in multi-line mode, unfocus with Enter, add new line with Ctrl+Enter
	// (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsCtrlEnterForNewLine = 1 << 11
	// InputTextFlagsNoHorizontalScroll disables following the cursor horizontally.
	InputTextFlagsNoHorizontalScroll = 1 << 12
	// InputTextFlagsAlwaysInsertMode sets insert mode.
	InputTextFlagsAlwaysInsertMode = 1 << 13
	// InputTextFlagsReadOnly sets read-only mode.
	InputTextFlagsReadOnly = 1 << 14
	// InputTextFlagsPassword sets password mode, display all characters as '*'.
	InputTextFlagsPassword = 1 << 15
	// InputTextFlagsNoUndoRedo disables undo/redo. Note that input text owns the text data while active,
	// if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsNoUndoRedo = 1 << 16
	// InputTextFlagsCharsScientific allows 0123456789.+-*/eE (Scientific notation input).
	InputTextFlagsCharsScientific = 1 << 17
	// inputTextFlagsCallbackResize for callback on buffer capacity change requests.
	inputTextFlagsCallbackResize = 1 << 18
)

// InputTextV creates a text field for dynamic text input.
//
// Contrary to the original library, this wrapper does not limit the maximum number of possible characters.
// Dynamic resizing of the internal buffer is handled within the wrapper and the user will never be called for such requests.
//
// The provided callback is called for any of the requested InputTextFlagsCallback* flags.
//
// To implement a character limit, provide a callback that drops input characters when the requested length has been reached.
func InputTextV(label string, text *string, flags int, cb InputTextCallback) bool {
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
func InputTextWithHintV(label string, hint string, text *string, flags int, cb InputTextCallback) bool {
	return inputTextSingleline(label, &hint, text, flags, cb)
}

// InputTextWithHint calls InputTextWithHintV(label, hint, text, 0, nil).
func InputTextWithHint(label string, hint string, text *string) bool {
	return InputTextWithHintV(label, hint, text, 0, nil)
}

func inputTextSingleline(label string, hint *string, text *string, flags int, cb InputTextCallback) bool {
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
func InputTextMultilineV(label string, text *string, size Vec2, flags int, cb InputTextCallback) bool {
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
func InputIntV(label string, value *int32, step int, stepFast int, flags int) bool {
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

const (
	// ColorEditFlagsNone default = 0
	ColorEditFlagsNone = 0
	// ColorEditFlagsNoAlpha ignores Alpha component (read 3 components from the input pointer).
	ColorEditFlagsNoAlpha = 1 << 1
	// ColorEditFlagsNoPicker disables picker when clicking on colored square.
	ColorEditFlagsNoPicker = 1 << 2
	// ColorEditFlagsNoOptions disables toggling options menu when right-clicking on inputs/small preview.
	ColorEditFlagsNoOptions = 1 << 3
	// ColorEditFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs)
	ColorEditFlagsNoSmallPreview = 1 << 4
	// ColorEditFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorEditFlagsNoInputs = 1 << 5
	// ColorEditFlagsNoTooltip disables tooltip when hovering the preview.
	ColorEditFlagsNoTooltip = 1 << 6
	// ColorEditFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorEditFlagsNoLabel = 1 << 7
	// ColorEditFlagsNoSidePreview disables bigger color preview on right side of the picker, use small colored square preview instead.
	ColorEditFlagsNoSidePreview = 1 << 8
	// ColorEditFlagsNoDragDrop disables drag and drop target. ColorButton: disable drag and drop source.
	ColorEditFlagsNoDragDrop = 1 << 9
	// ColorEditFlagsNoBorder disables border (which is enforced by default)
	ColorEditFlagsNoBorder = 1 << 10

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions().
	// The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call
	// SetColorEditOptions() during startup.

	// ColorEditFlagsAlphaBar shows vertical alpha bar/gradient in picker.
	ColorEditFlagsAlphaBar = 1 << 16
	// ColorEditFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreview = 1 << 17
	// ColorEditFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreviewHalf = 1 << 18
	// ColorEditFlagsHDR = (WIP) surrently only disable 0.0f..1.0f limits in RGBA edition.
	// Note: you probably want to use ImGuiColorEditFlags_Float flag as well.
	ColorEditFlagsHDR = 1 << 19
	// ColorEditFlagsRGB sets the format as RGB
	ColorEditFlagsRGB = 1 << 20
	// ColorEditFlagsHSV sets the format as HSV
	ColorEditFlagsHSV = 1 << 21
	// ColorEditFlagsHEX sets the format as HEX
	ColorEditFlagsHEX = 1 << 22
	// ColorEditFlagsUint8 _display_ values formatted as 0..255.
	ColorEditFlagsUint8 = 1 << 23
	// ColorEditFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorEditFlagsFloat = 1 << 24

	// ColorEditFlagsPickerHueBar shows bar for Hue, rectangle for Sat/Value.
	ColorEditFlagsPickerHueBar = 1 << 25
	// ColorEditFlagsPickerHueWheel shows wheel for Hue, triangle for Sat/Value.
	ColorEditFlagsPickerHueWheel = 1 << 26
	// ColorEditFlagsInputRGB enables input and output data in RGB format.
	ColorEditFlagsInputRGB = 1 << 27
	// ColorEditFlagsInputHSV enables input and output data in HSV format.
	ColorEditFlagsInputHSV = 1 << 28
)

// ColorEdit3 calls ColorEdit3V(label, col, 0).
func ColorEdit3(label string, col *[3]float32) bool {
	return ColorEdit3V(label, col, 0)
}

// ColorEdit3V will show a clickable little square which will open a color picker window for 3D vector (rgb format).
func ColorEdit3V(label string, col *[3]float32, flags int) bool {
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
func ColorEdit4V(label string, col *[4]float32, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorEdit4(labelArg, ccol, C.int(flags)) != 0
}

const (
	// ColorPickerFlagsNone default = 0
	ColorPickerFlagsNone = 0
	// ColorPickerFlagsNoPicker disables picker when clicking on colored square.
	ColorPickerFlagsNoPicker = 1 << 2
	// ColorPickerFlagsNoOptions disables toggling options menu when right-clicking on inputs/small preview.
	ColorPickerFlagsNoOptions = 1 << 3
	// ColorPickerFlagsNoAlpha ignoreс Alpha component (read 3 components from the input pointer).
	ColorPickerFlagsNoAlpha = 1 << 1
	// ColorPickerFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs)
	ColorPickerFlagsNoSmallPreview = 1 << 4
	// ColorPickerFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorPickerFlagsNoInputs = 1 << 5
	// ColorPickerFlagsNoTooltip disables tooltip when hovering the preview.
	ColorPickerFlagsNoTooltip = 1 << 6
	// ColorPickerFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorPickerFlagsNoLabel = 1 << 7
	// ColorPickerFlagsNoSidePreview disables bigger color preview on right side of the picker, use small colored square preview instead.
	ColorPickerFlagsNoSidePreview = 1 << 8

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions().
	// The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call
	// SetColorPickerOptions() during startup.

	// ColorPickerFlagsAlphaBar shows vertical alpha bar/gradient in picker.
	ColorPickerFlagsAlphaBar = 1 << 16
	// ColorPickerFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
	ColorPickerFlagsAlphaPreview = 1 << 17
	// ColorPickerFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
	ColorPickerFlagsAlphaPreviewHalf = 1 << 18
	// ColorPickerFlagsRGB sets the format as RGB
	ColorPickerFlagsRGB = 1 << 20
	// ColorPickerFlagsHSV sets the format as HSV
	ColorPickerFlagsHSV = 1 << 21
	// ColorPickerFlagsHEX sets the format as HEX
	ColorPickerFlagsHEX = 1 << 22
	// ColorPickerFlagsUint8 _display_ values formatted as 0..255.
	ColorPickerFlagsUint8 = 1 << 23
	// ColorPickerFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorPickerFlagsFloat = 1 << 24
	// ColorPickerFlagsPickerHueBar bar for Hue, rectangle for Sat/Value.
	ColorPickerFlagsPickerHueBar = 1 << 25
	// ColorPickerFlagsPickerHueWheel wheel for Hue, triangle for Sat/Value.
	ColorPickerFlagsPickerHueWheel = 1 << 26
	// ColorPickerFlagsInputRGB enables input and output data in RGB format.
	ColorPickerFlagsInputRGB = 1 << 27
	// ColorPickerFlagsInputHSV enables input and output data in HSV format.
	ColorPickerFlagsInputHSV = 1 << 28
)

// ColorPicker3 calls ColorPicker3V(label, col, 0).
func ColorPicker3(label string, col *[3]float32) bool {
	return ColorPicker3V(label, col, 0)
}

// ColorPicker3V will show directly a color picker control for editing a color in 3D vector (rgb format).
func ColorPicker3V(label string, col *[3]float32, flags int) bool {
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
func ColorPicker4V(label string, col *[4]float32, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorPicker4(labelArg, ccol, C.int(flags)) != 0
}

// CollapsingHeader adds a collapsing header.
func CollapsingHeader(label string) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggCollapsingHeader(labelArg) != 0
}

const (
	// TreeNodeFlagsNone default = 0
	TreeNodeFlagsNone = 0
	// TreeNodeFlagsSelected draws as selected.
	TreeNodeFlagsSelected = 1 << 0
	// TreeNodeFlagsFramed draws full colored frame (e.g. for CollapsingHeader).
	TreeNodeFlagsFramed = 1 << 1
	// TreeNodeFlagsAllowItemOverlap hit testing to allow subsequent widgets to overlap this one.
	TreeNodeFlagsAllowItemOverlap = 1 << 2
	// TreeNodeFlagsNoTreePushOnOpen doesn't do a TreePush() when open
	// (e.g. for CollapsingHeader) = no extra indent nor pushing on ID stack.
	TreeNodeFlagsNoTreePushOnOpen = 1 << 3
	// TreeNodeFlagsNoAutoOpenOnLog doesn't automatically and temporarily open node when Logging is active
	// (by default logging will automatically open tree nodes).
	TreeNodeFlagsNoAutoOpenOnLog = 1 << 4
	// TreeNodeFlagsDefaultOpen defaults node to be open.
	TreeNodeFlagsDefaultOpen = 1 << 5
	// TreeNodeFlagsOpenOnDoubleClick needs double-click to open node.
	TreeNodeFlagsOpenOnDoubleClick = 1 << 6
	// TreeNodeFlagsOpenOnArrow opens only when clicking on the arrow part.
	// If TreeNodeFlagsOpenOnDoubleClick is also set, single-click arrow or double-click all box to open.
	TreeNodeFlagsOpenOnArrow = 1 << 7
	// TreeNodeFlagsLeaf allows no collapsing, no arrow (use as a convenience for leaf nodes).
	TreeNodeFlagsLeaf = 1 << 8
	// TreeNodeFlagsBullet displays a bullet instead of an arrow.
	TreeNodeFlagsBullet = 1 << 9
	// TreeNodeFlagsFramePadding uses FramePadding (even for an unframed text node) to
	// vertically align text baseline to regular widget height. Equivalent to calling AlignTextToFramePadding().
	TreeNodeFlagsFramePadding = 1 << 10
	// TreeNodeFlagsSpanAvailWidth extends hit box to the right-most edge, even if not framed.
	// This is not the default in order to allow adding other items on the same line.
	// In the future we may refactor the hit system to be front-to-back, allowing natural overlaps
	// and then this can become the default.
	TreeNodeFlagsSpanAvailWidth = 1 << 11
	// TreeNodeFlagsSpanFullWidth extends hit box to the left-most and right-most edges (bypass the indented area).
	TreeNodeFlagsSpanFullWidth = 1 << 12
	// TreeNodeFlagsNavLeftJumpsBackHere (WIP) Nav: left direction may move to this TreeNode() from any of its child
	// (items submitted between TreeNode and TreePop)
	TreeNodeFlagsNavLeftJumpsBackHere = 1 << 13
	// TreeNodeFlagsCollapsingHeader combines TreeNodeFlagsFramed and TreeNodeFlagsNoAutoOpenOnLog.
	TreeNodeFlagsCollapsingHeader = TreeNodeFlagsFramed | TreeNodeFlagsNoTreePushOnOpen | TreeNodeFlagsNoAutoOpenOnLog
)

// TreeNodeV returns true if the tree branch is to be rendered. Call TreePop() in this case.
func TreeNodeV(label string, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggTreeNode(labelArg, C.int(flags)) != 0
}

// TreeNode calls TreeNodeV(label, 0).
func TreeNode(label string) bool {
	return TreeNodeV(label, 0)
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

const (
	// SelectableFlagsNone default = 0
	SelectableFlagsNone = 0
	// SelectableFlagsDontClosePopups makes clicking the selectable not close any parent popup windows.
	SelectableFlagsDontClosePopups = 1 << 0
	// SelectableFlagsSpanAllColumns allows the selectable frame to span all columns (text will still fit in current column).
	SelectableFlagsSpanAllColumns = 1 << 1
	// SelectableFlagsAllowDoubleClick generates press events on double clicks too.
	SelectableFlagsAllowDoubleClick = 1 << 2
	// SelectableFlagsDisabled disallows selection and displays text in a greyed out color.
	SelectableFlagsDisabled = 1 << 3
)

// SelectableV returns true if the user clicked it, so you can modify your selection state.
// flags are the SelectableFlags to apply.
// size.x==0.0: use remaining width, size.x>0.0: specify width.
// size.y==0.0: use label height, size.y>0.0: specify height.
func SelectableV(label string, selected bool, flags int, size Vec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	sizeArg, _ := size.wrapped()
	return C.iggSelectable(labelArg, castBool(selected), C.int(flags), sizeArg) != 0
}

// Selectable calls SelectableV(label, false, 0, Vec2{0, 0}).
func Selectable(label string) bool {
	return SelectableV(label, false, 0, Vec2{})
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

	return C.iggListBoxV(labelArg, valueArg, &argv[0], C.int(itemsCount), C.int(heightItems)) != 0
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

// Columns calls ColumnsV(1, "", false).
func Columns() {
	ColumnsV(1, "", false)
}

// ColumnsV creates a column layout of the specified number of columns.
// The brittle columns API will be superseded by an upcoming 'table' API.
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

const (
	// TabBarFlagsNone default = 0.
	TabBarFlagsNone = 0
	// TabBarFlagsReorderable Allow manually dragging tabs to re-order them + New tabs are appended at the end of list
	TabBarFlagsReorderable = 1 << 0
	// TabBarFlagsAutoSelectNewTabs Automatically select new tabs when they appear
	TabBarFlagsAutoSelectNewTabs = 1 << 1
	// TabBarFlagsTabListPopupButton Disable buttons to open the tab list popup
	TabBarFlagsTabListPopupButton = 1 << 2
	// TabBarFlagsNoCloseWithMiddleMouseButton Disable behavior of closing tabs (that are submitted with p_open != NULL)
	// with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabBarFlagsNoCloseWithMiddleMouseButton = 1 << 3
	// TabBarFlagsNoTabListScrollingButtons Disable scrolling buttons (apply when fitting policy is
	// TabBarFlagsFittingPolicyScroll)
	TabBarFlagsNoTabListScrollingButtons = 1 << 4
	// TabBarFlagsNoTooltip Disable tooltips when hovering a tab
	TabBarFlagsNoTooltip = 1 << 5
	// TabBarFlagsFittingPolicyResizeDown Resize tabs when they don't fit
	TabBarFlagsFittingPolicyResizeDown = 1 << 6
	// TabBarFlagsFittingPolicyScroll Add scroll buttons when tabs don't fit
	TabBarFlagsFittingPolicyScroll = 1 << 7
	// TabBarFlagsFittingPolicyMask combines
	// TabBarFlagsFittingPolicyResizeDown and TabBarFlagsFittingPolicyScroll
	TabBarFlagsFittingPolicyMask = TabBarFlagsFittingPolicyResizeDown | TabBarFlagsFittingPolicyScroll
	// TabBarFlagsFittingPolicyDefault alias for TabBarFlagsFittingPolicyResizeDown
	TabBarFlagsFittingPolicyDefault = TabBarFlagsFittingPolicyResizeDown
)

// BeginTabBarV create and append into a TabBar.
func BeginTabBarV(strID string, flags int) bool {
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

const (
	// TabItemFlagsNone default = 0
	TabItemFlagsNone = 0
	// TabItemFlagsUnsavedDocument Append '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. Also: tab is selected on closure and closure is deferred by one frame to allow code to undo it
	// without flicker.
	TabItemFlagsUnsavedDocument = 1 << 0
	// TabItemFlagsSetSelected Trigger flag to programmatically make the tab selected when calling BeginTabItem()
	TabItemFlagsSetSelected = 1 << 1
	// TabItemFlagsNoCloseWithMiddleMouseButton  Disable behavior of closing tabs (that are submitted with
	// p_open != NULL) with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabItemFlagsNoCloseWithMiddleMouseButton = 1 << 2
	// TabItemFlagsNoPushID Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem()
	TabItemFlagsNoPushID = 1 << 3
)

// BeginTabItemV create a Tab. Returns true if the Tab is selected.
func BeginTabItemV(label string, open *bool, flags int) bool {
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

// SetTabItemClosed notify TabBar or Docking system of a closed tab/window ahead
// (useful to reduce visual flicker on reorderable tab bars). For tab-bar: call
// after BeginTabBar() and before Tab submissions. Otherwise call with a window name.
func SetTabItemClosed(tabOrDockedWindowLabel string) {
	labelArg, labelFin := wrapString(tabOrDockedWindowLabel)
	defer labelFin()
	C.iggSetTabItemClosed(labelArg)
}
