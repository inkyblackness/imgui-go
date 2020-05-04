package imgui

// #include "wrapper/imguiWrapper.h"
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

// Text adds text in disabled style.
func TextDisabled(text string) {
	textArg, textFin := wrapString(text)
	defer textFin()
	C.iggTextDisabled(textArg)
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

// ProgressBarV creates a progress bar.
// size (for each axis) < 0.0f: align to end, 0.0f: auto, > 0.0f: specified size
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

// VSliderFloat calls VSliderIntV(label, size, value, min, max, "%.3f", 1.0
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

// VSliderInt calls VSliderIntV(label, size, value, min, max, "%d"
func VSliderInt(label string, size Vec2, value *int32, min, max int32) bool {
	return VSliderIntV(label, size, value, min, max, "%d")
}

// InputTextV creates a text field for dynamic text input.
//
// Contrary to the original library, this wrapper does not limit the maximum number of possible characters.
// Dynamic resizing of the internal buffer is handled within the wrapper and the user will never be called for such requests.
//
// The provided callback is called for any of the requested InputTextFlagsCallback* flags.
//
// To implement a character limit, provide a callback that drops input characters when the requested length has been reached.
func InputTextV(label string, text *string, flags int, cb InputTextCallback) bool {
	if text == nil {
		panic("text can't be nil")
	}
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	state := newInputTextState(*text, cb)
	defer func() {
		*text = state.buf.toGo()
		state.release()
	}()

	return C.iggInputText(labelArg, (*C.char)(state.buf.ptr), C.uint(state.buf.size),
		C.int(flags|inputTextFlagsCallbackResize), state.key) != 0
}

// InputText calls InputTextV(label, string, 0, nil)
func InputText(label string, text *string) bool {
	return InputTextV(label, text, 0, nil)
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

// InputTextMultiline calls InputTextMultilineV(label, text, Vec2{0,0}, 0, nil)
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

// ColorEdit3 calls ColorEdit3V(label, col, 0)
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

// ColorEdit4 calls ColorEdit4V(label, col, 0)
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

// ColorPicker3 calls ColorPicker3(label, col, 0)
func ColorPicker3(label string, col *[3]float32, flags int) bool {
	return ColorPicker3V(label, col, 0)
}

// ColorPicker3V will show directly a color picker control for editing a color in 3D vector (rgb format).
func ColorPicker3V(label string, col *[3]float32, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	ccol := (*C.float)(&col[0])
	return C.iggColorPicker3(labelArg, ccol, C.int(flags)) != 0
}

// ColorPicker4 calls ColorPicker4(label, col, 0)
func ColorPicker4(label string, col *[4]float32, flags int) bool {
	return ColorPicker4(label, col, 0)
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

// SelectableV returns true if the user clicked it, so you can modify your selection state.
// flags are the SelectableFlags to apply.
// size.x==0.0: use remaining width, size.x>0.0: specify width.
// size.y==0.0: use label height, size.y>0.0: specify height
func SelectableV(label string, selected bool, flags int, size Vec2) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	sizeArg, _ := size.wrapped()
	return C.iggSelectable(labelArg, castBool(selected), C.int(flags), sizeArg) != 0
}

// Selectable calls SelectableV(label, false, 0, Vec2{0, 0})
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
	for i, item := range items {
		itemArg, itemDeleter := wrapString(item)
		defer itemDeleter()
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

// ColumnOffset calls ColumnOffsetV(-1)
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

// BeginTabBarV create and append into a TabBar
func BeginTabBarV(strID string, flags int) bool {
	idArg, idFin := wrapString(strID)
	defer idFin()

	return C.iggBeginTabBar(idArg, C.int(flags)) != 0
}

// BeginTabBar calls BeginTabBarV(strId, 0)
func BeginTabBar(strID string) bool {
	return BeginTabBarV(strID, 0)
}

// EndTabBar only call EndTabBar() if BeginTabBar() returns true!
func EndTabBar() {
	C.iggEndTabBar()
}

// BeginTabItemV create a Tab. Returns true if the Tab is selected.
func BeginTabItemV(label string, open *bool, flags int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	openArg, openFin := wrapBool(open)
	defer openFin()

	return C.iggBeginTabItem(labelArg, openArg, C.int(flags)) != 0
}

// BeginTabItem calls BeginTabItemV(label, nil, 0)
func BeginTabItem(label string) bool {
	return BeginTabItemV(label, nil, 0)
}

// EndTabItem Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem()
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
