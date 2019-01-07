package imgui

// #cgo CXXFLAGS: -std=c++11
// #include "imguiWrapper.h"
import "C"

// User fill ImGuiIO.KeyMap[] array with indices into the ImGuiIO.KeysDown[512] array
const (
	KeyTab        = 0
	KeyLeftArrow  = 1
	KeyRightArrow = 2
	KeyUpArrow    = 3
	KeyDownArrow  = 4
	KeyPageUp     = 5
	KeyPageDown   = 6
	KeyHome       = 7
	KeyEnd        = 8
	KeyInsert     = 9
	KeyDelete     = 10
	KeyBackspace  = 11
	KeySpace      = 12
	KeyEnter      = 13
	KeyEscape     = 14
	KeyA          = 15 // for text edit CTRL+A: select all
	KeyC          = 16 // for text edit CTRL+C: copy
	KeyV          = 17 // for text edit CTRL+V: paste
	KeyX          = 18 // for text edit CTRL+X: cut
	KeyY          = 19 // for text edit CTRL+Y: redo
	KeyZ          = 20 // for text edit CTRL+Z: undo
	KeyCOUNT      = 21
)

// Version returns a version string e.g. "1.23".
func Version() string {
	return C.GoString(C.iggGetVersion())
}

// CurrentIO returns access to the ImGui communication struct for the currently active context.
func CurrentIO() IO {
	return IO{handle: C.iggGetCurrentIO()}
}

// CurrentStyle returns the UI Style for the currently active context.
func CurrentStyle() Style {
	return Style(C.iggGetCurrentStyle())
}

// NewFrame starts a new ImGui frame, you can submit any command from this point until Render()/EndFrame().
func NewFrame() {
	C.iggNewFrame()
}

// Render ends the ImGui frame, finalize the draw data.
// After this method, call RenderedDrawData to retrieve the draw commands and execute them.
func Render() {
	C.iggRender()
}

// EndFrame ends the ImGui frame. Automatically called by Render(), so most likely don't need to ever
// call that yourself directly. If you don't need to render you may call EndFrame() but you'll have
// wasted CPU already. If you don't need to render, better to not create any imgui windows instead!
func EndFrame() {
	C.iggEndFrame()
}

// RenderedDrawData returns the created draw commands, which are valid after Render() and
// until the next call to NewFrame(). This is what you have to render.
func RenderedDrawData() DrawData {
	return DrawData(C.iggGetDrawData())
}

// ShowDemoWindow creates a demo/test window. Demonstrates most ImGui features.
// Call this to learn about the library! Try to make it always available in your application!
func ShowDemoWindow(open *bool) {
	openArg, openFin := wrapBool(open)
	defer openFin()
	C.iggShowDemoWindow(openArg)
}

// ShowUserGuide adds basic help/info block (not a window): how to manipulate ImGui as a end-user (mouse/keyboard controls).
func ShowUserGuide() {
	C.iggShowUserGuide()
}

// BeginV pushes a new window to the stack and start appending to it.
// You may append multiple times to the same window during the same frame.
// If the open argument is provided, the window can be closed, in which case the value will be false after the call.
//
// Returns false if the window is currently not visible.
// Regardless of the return value, End() must be called for each call to Begin().
func BeginV(id string, open *bool, flags int) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	openArg, openFin := wrapBool(open)
	defer openFin()
	return C.iggBegin(idArg, openArg, C.int(flags)) != 0
}

// Begin calls BeginV(id, nil, 0).
func Begin(id string) bool {
	return BeginV(id, nil, 0)
}

// End closes the scope for the previously opened window.
// Every call to Begin() must be matched with a call to End().
func End() {
	C.iggEnd()
}

// BeginChildV pushes a new child to the stack and starts appending to it.
func BeginChildV(id string, size Vec2, border bool, flags int) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	sizeArg, _ := size.wrapped()
	return C.iggBeginChild(idArg, sizeArg, castBool(border), C.int(flags)) != 0
}

// BeginChild calls BeginChildV(id, Vec2{0,0}, false, 0).
func BeginChild(id string) bool {
	return BeginChildV(id, Vec2{}, false, 0)
}

// EndChild closes the scope for the previously opened child.
// Every call to BeginChild() must be matched with a call to EndChild().
func EndChild() {
	C.iggEndChild()
}

// SetNextWindowPosV sets next window position.
// Call before Begin(). Use pivot=(0.5,0.5) to center on given point, etc.
func SetNextWindowPosV(pos Vec2, cond Condition, pivot Vec2) {
	posArg, _ := pos.wrapped()
	pivotArg, _ := pivot.wrapped()
	C.iggSetNextWindowPos(posArg, C.int(cond), pivotArg)
}

// SetNextWindowPos calls SetNextWindowPosV(pos, 0, Vec{0,0})
func SetNextWindowPos(pos Vec2) {
	SetNextWindowPosV(pos, 0, Vec2{})
}

// SetNextWindowSizeV sets next window size.
// Set axis to 0.0 to force an auto-fit on this axis. Call before Begin().
func SetNextWindowSizeV(size Vec2, cond Condition) {
	sizeArg, _ := size.wrapped()
	C.iggSetNextWindowSize(sizeArg, C.int(cond))
}

// SetNextWindowSize calls SetNextWindowSizeV(size, 0)
func SetNextWindowSize(size Vec2) {
	SetNextWindowSizeV(size, 0)
}

// SetNextWindowFocus sets next window to be focused / front-most. Call before Begin().
func SetNextWindowFocus() {
	C.iggSetNextWindowFocus()
}

// SetNextWindowBgAlpha sets next window background color alpha.
// Helper to easily modify ImGuiCol_WindowBg/ChildBg/PopupBg.
func SetNextWindowBgAlpha(value float32) {
	C.iggSetNextWindowBgAlpha(C.float(value))
}

// PushFont adds the given font on the stack. Use DefaultFont to refer to the default font.
func PushFont(font Font) {
	C.iggPushFont(font.handle())
}

// PopFont removes the previously pushed font from the stack.
func PopFont() {
	C.iggPopFont()
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

// PushItemWidth sets width of items for the common item+label case, in pixels.
// 0.0f = default to ~2/3 of windows width, >0.0f: width in pixels,
// <0.0f align xx pixels to the right of window (so -1.0f always align width to the right side).
func PushItemWidth(width float32) {
	C.iggPushItemWidth(C.float(width))
}

// PopItemWidth must be called for each call to PushItemWidth().
func PopItemWidth() {
	C.iggPopItemWidth()
}

// PushTextWrapPosV defines word-wrapping for Text() commands.
// < 0.0f: no wrapping; 0.0f: wrap to end of window (or column); > 0.0f: wrap at 'wrapPosX' position in window local space.
// Requires a matching call to PopTextWrapPos().
func PushTextWrapPosV(wrapPosX float32) {
	C.iggPushTextWrapPos(C.float(wrapPosX))
}

// PushTextWrapPos calls PushTextWrapPosV(0.0).
func PushTextWrapPos() {
	PushTextWrapPosV(0.0)
}

// PopTextWrapPos resets the last pushed position.
func PopTextWrapPos() {
	C.iggPopTextWrapPos()
}

// PushID pushes the given identifier into the ID stack. IDs are hash of the entire stack!
func PushID(id string) {
	idArg, idFin := wrapString(id)
	defer idFin()
	C.iggPushID(idArg)
}

// PopID removes the last pushed identifier from the ID stack.
func PopID() {
	C.iggPopID()
}

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

// BeginComboV creates a combo box with complete control over the content to the user.
// Call EndCombo() if this function returns true.
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

// BeginGroup locks horizontal starting position + capture group bounding box into one "item"
// (so you can use IsItemHovered() or layout primitives such as SameLine() on whole group, etc.)
func BeginGroup() {
	C.iggBeginGroup()
}

// EndGroup must be called for each call to BeginGroup().
func EndGroup() {
	C.iggEndGroup()
}

// SetCursorPos sets the cursor relative to the current window.
func SetCursorPos(localPos Vec2) {
	localPosArg, _ := localPos.wrapped()
	C.iggSetCursorPos(localPosArg)
}

// TextLineHeight returns ~ FontSize.
func TextLineHeight() float32 {
	return float32(C.iggGetTextLineHeight())
}

// TextLineHeightWithSpacing returns ~ FontSize + style.ItemSpacing.y (distance in pixels between 2 consecutive lines of text).
func TextLineHeightWithSpacing() float32 {
	return float32(C.iggGetTextLineHeightWithSpacing())
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

// SetNextTreeNodeOpen sets the open/collapsed state of the following tree node.
func SetNextTreeNodeOpen(open bool, cond Condition) {
	C.iggSetNextTreeNodeOpen(castBool(open), C.int(cond))
}

// SelectableV returns true if the user clicked it, so you can modify your selection state.
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

// OpenPopup marks popup as open (don't call every frame!).
// Popups are closed when user click outside, or if CloseCurrentPopup() is called within a BeginPopup()/EndPopup() block.
// By default, Selectable()/MenuItem() are calling CloseCurrentPopup().
// Popup identifiers are relative to the current ID-stack (so OpenPopup and BeginPopup needs to be at the same level).
func OpenPopup(id string) {
	idArg, idFin := wrapString(id)
	defer idFin()
	C.iggOpenPopup(idArg)
}

// BeginPopupModalV creates modal dialog (regular window with title bar, block interactions behind the modal window,
// can't close the modal window by clicking outside).
func BeginPopupModalV(name string, open *bool, flags int) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()
	openArg, openFin := wrapBool(open)
	defer openFin()
	return C.iggBeginPopupModal(nameArg, openArg, C.int(flags)) != 0
}

// BeginPopupModal calls BeginPopupModalV(name, nil, 0)
func BeginPopupModal(name string) bool {
	return BeginPopupModalV(name, nil, 0)
}

// BeginPopupContextItemV returns true if the identified mouse button was pressed
// while hovering over the last item.
func BeginPopupContextItemV(label string, mouseButton int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggBeginPopupContextItem(labelArg, C.int(mouseButton)) != 0
}

// BeginPopupContextItem calls BeginPopupContextItemV("", 1)
func BeginPopupContextItem() bool {
	return BeginPopupContextItemV("", 1)
}

// EndPopup finshes a popup. Only call EndPopup() if BeginPopupXXX() returns true!
func EndPopup() {
	C.iggEndPopup()
}

// CloseCurrentPopup closes the popup we have begin-ed into.
// Clicking on a MenuItem or Selectable automatically close the current popup.
func CloseCurrentPopup() {
	C.iggCloseCurrentPopup()
}

// IsItemHoveredV returns true if the last item is hovered.
// (and usable, aka not blocked by a popup, etc.). See HoveredFlags for more options.
func IsItemHoveredV(flags int) bool {
	return C.iggIsItemHovered(C.int(flags)) != 0
}

// IsItemHovered calls IsItemHoveredV(HoveredFlagsDefault)
func IsItemHovered() bool {
	return IsItemHoveredV(HoveredFlagsDefault)
}

// IsKeyPressed returns true if the corresponding key is currently pressed.
func IsKeyPressed(key int) bool {
	return C.iggIsKeyPressed(C.int(key)) != 0
}
