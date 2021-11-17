package imgui

// #include "wrapper/Window.h"
import "C"

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

// WindowFlags for BeginV(), etc.
type WindowFlags int

const (
	// WindowFlagsNone default = 0.
	WindowFlagsNone WindowFlags = 0
	// WindowFlagsNoTitleBar disables title-bar.
	WindowFlagsNoTitleBar WindowFlags = 1 << 0
	// WindowFlagsNoResize disables user resizing with the lower-right grip.
	WindowFlagsNoResize WindowFlags = 1 << 1
	// WindowFlagsNoMove disables user moving the window.
	WindowFlagsNoMove WindowFlags = 1 << 2
	// WindowFlagsNoScrollbar disables scrollbars. Window can still scroll with mouse or programmatically.
	WindowFlagsNoScrollbar WindowFlags = 1 << 3
	// WindowFlagsNoScrollWithMouse disables user vertically scrolling with mouse wheel. On child window, mouse wheel
	// will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoScrollWithMouse WindowFlags = 1 << 4
	// WindowFlagsNoCollapse disables user collapsing window by double-clicking on it.
	WindowFlagsNoCollapse WindowFlags = 1 << 5
	// WindowFlagsAlwaysAutoResize resizes every window to its content every frame.
	WindowFlagsAlwaysAutoResize WindowFlags = 1 << 6
	// WindowFlagsNoBackground disables drawing background color (WindowBg, etc.) and outside border. Similar as using
	// SetNextWindowBgAlpha(0.0f).
	WindowFlagsNoBackground WindowFlags = 1 << 7
	// WindowFlagsNoSavedSettings will never load/save settings in .ini file.
	WindowFlagsNoSavedSettings WindowFlags = 1 << 8
	// WindowFlagsNoMouseInputs disables catching mouse, hovering test with pass through.
	WindowFlagsNoMouseInputs WindowFlags = 1 << 9
	// WindowFlagsMenuBar has a menu-bar.
	WindowFlagsMenuBar WindowFlags = 1 << 10
	// WindowFlagsHorizontalScrollbar allows horizontal scrollbar to appear (off by default). You may use
	// SetNextWindowContentSize(ImVec2(width,0.0f)); prior to calling Begin() to specify width. Read code in imgui_demo
	// in the "Horizontal Scrolling" section.
	WindowFlagsHorizontalScrollbar WindowFlags = 1 << 11
	// WindowFlagsNoFocusOnAppearing disables taking focus when transitioning from hidden to visible state.
	WindowFlagsNoFocusOnAppearing WindowFlags = 1 << 12
	// WindowFlagsNoBringToFrontOnFocus disables bringing window to front when taking focus. e.g. clicking on it or
	// programmatically giving it focus.
	WindowFlagsNoBringToFrontOnFocus WindowFlags = 1 << 13
	// WindowFlagsAlwaysVerticalScrollbar always shows vertical scrollbar, even if ContentSize.y < Size.y .
	WindowFlagsAlwaysVerticalScrollbar WindowFlags = 1 << 14
	// WindowFlagsAlwaysHorizontalScrollbar always shows horizontal scrollbar, even if ContentSize.x < Size.x .
	WindowFlagsAlwaysHorizontalScrollbar WindowFlags = 1 << 15
	// WindowFlagsAlwaysUseWindowPadding ensures child windows without border uses style.WindowPadding (ignored by
	// default for non-bordered child windows, because more convenient).
	WindowFlagsAlwaysUseWindowPadding WindowFlags = 1 << 16
	// WindowFlagsNoNavInputs has no gamepad/keyboard navigation within the window.
	WindowFlagsNoNavInputs WindowFlags = 1 << 18
	// WindowFlagsNoNavFocus has no focusing toward this window with gamepad/keyboard navigation
	// (e.g. skipped by CTRL+TAB).
	WindowFlagsNoNavFocus WindowFlags = 1 << 19
	// WindowFlagsUnsavedDocument appends '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. When used in a tab/docking context, tab is selected on closure and closure is deferred by one
	// frame to allow code to cancel the closure (with a confirmation popup, etc.) without flicker.
	WindowFlagsUnsavedDocument WindowFlags = 1 << 20

	// WindowFlagsNoNav combines WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoNav = WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
	// WindowFlagsNoDecoration combines WindowFlagsNoTitleBar, WindowFlagsNoResize, WindowFlagsNoScrollbar and
	// WindowFlagsNoCollapse.
	WindowFlagsNoDecoration = WindowFlagsNoTitleBar | WindowFlagsNoResize | WindowFlagsNoScrollbar | WindowFlagsNoCollapse
	// WindowFlagsNoInputs combines WindowFlagsNoMouseInputs, WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoInputs = WindowFlagsNoMouseInputs | WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
)

// BeginV pushes a new window to the stack and start appending to it.
// You may append multiple times to the same window during the same frame.
// If the open argument is provided, the window can be closed, in which case the value will be false after the call.
//
// Returns false if the window is currently not visible.
// Regardless of the return value, End() must be called for each call to Begin().
func BeginV(id string, open *bool, flags WindowFlags) bool {
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
// flags are the WindowFlags to apply.
func BeginChildV(id string, size Vec2, border bool, flags WindowFlags) bool {
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

// WindowPos returns the current window position in screen space.
// This is useful if you want to do your own drawing via the DrawList API.
func WindowPos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggWindowPos(valueArg)
	valueFin()
	return value
}

// WindowSize returns the size of the current window.
func WindowSize() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggWindowSize(valueArg)
	valueFin()
	return value
}

// WindowWidth returns the width of the current window.
func WindowWidth() float32 {
	return float32(C.iggWindowWidth())
}

// WindowHeight returns the height of the current window.
func WindowHeight() float32 {
	return float32(C.iggWindowHeight())
}

// ContentRegionAvail returns the size of the content region that is available (based on the current cursor position).
func ContentRegionAvail() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggContentRegionAvail(valueArg)
	valueFin()
	return value
}

// ContentRegionMax returns current content boundaries (typically window boundaries including scrolling,
// or current column boundaries), in windows coordinates.
func ContentRegionMax() Vec2 {
	out := Vec2{}
	outArg, outFin := out.wrapped()
	C.iggGetContentRegionMax(outArg)
	outFin()
	return out
}

// WindowContentRegionMin returns the content boundaries min (roughly (0,0)-Scroll), in window coordinates.
func WindowContentRegionMin() Vec2 {
	out := Vec2{}
	outArg, outFin := out.wrapped()
	C.iggGetWindowContentRegionMin(outArg)
	outFin()
	return out
}

// WindowContentRegionMax returns the content boundaries max (roughly (0,0)+Size-Scroll) where Size can be overridden
// with SetNextWindowContentSize(), in window coordinates.
func WindowContentRegionMax() Vec2 {
	out := Vec2{}
	outArg, outFin := out.wrapped()
	C.iggGetWindowContentRegionMax(outArg)
	outFin()
	return out
}

// WindowContentRegionWidth returns the width of the content boundary, in window coordinates.
//
// DEPRECATED: in v1.85, GetWindowContentRegionWidth() function has been removed.
// This function now returns GetWindowContentRegionMax().x - GetWindowContentRegionMin().x,
// as per v1.85 release notes. Please note it's not very useful in practice, and
// using GetContentRegionAvail().x is generally a better choice.
func WindowContentRegionWidth() float32 {
	return WindowContentRegionMax().X - WindowContentRegionMin().X
}

// SetNextWindowPosV sets next window position.
// Call before Begin(). Use pivot=(0.5,0.5) to center on given point, etc.
func SetNextWindowPosV(pos Vec2, cond Condition, pivot Vec2) {
	posArg, _ := pos.wrapped()
	pivotArg, _ := pivot.wrapped()
	C.iggSetNextWindowPos(posArg, C.int(cond), pivotArg)
}

// SetNextWindowPos calls SetNextWindowPosV(pos, 0, Vec{0,0}).
func SetNextWindowPos(pos Vec2) {
	SetNextWindowPosV(pos, 0, Vec2{})
}

// SetNextWindowCollapsed sets the next window collapsed state.
func SetNextWindowCollapsed(collapsed bool, cond Condition) {
	C.iggSetNextWindowCollapsed(castBool(collapsed), C.int(cond))
}

// SetNextWindowSizeV sets next window size.
// Set axis to 0.0 to force an auto-fit on this axis. Call before Begin().
func SetNextWindowSizeV(size Vec2, cond Condition) {
	sizeArg, _ := size.wrapped()
	C.iggSetNextWindowSize(sizeArg, C.int(cond))
}

// SetNextWindowSize calls SetNextWindowSizeV(size, 0).
func SetNextWindowSize(size Vec2) {
	SetNextWindowSizeV(size, 0)
}

// SetNextWindowSizeConstraints set next window size limits.
// Use -1,-1 on either X/Y axis to preserve the current size.
// Use callback to apply non-trivial programmatic constraints.
func SetNextWindowSizeConstraints(sizeMin Vec2, sizeMax Vec2) {
	sizeMinArg, _ := sizeMin.wrapped()
	sizeMaxArg, _ := sizeMax.wrapped()
	C.iggSetNextWindowSizeConstraints(sizeMinArg, sizeMaxArg)
}

// SetNextWindowContentSize sets next window content size (~ enforce the range of scrollbars).
// Does not include window decorations (title bar, menu bar, etc.).
// Set one axis to 0.0 to leave it automatic. This function must be called before Begin() to take effect.
func SetNextWindowContentSize(size Vec2) {
	sizeArg, _ := size.wrapped()
	C.iggSetNextWindowContentSize(sizeArg)
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

// PushItemWidth pushes width of items for common large "item+label" widgets.
// >0.0f: width in pixels, <0.0f align xx pixels to the right of window (so -math.SmallestNonzeroFloat32 always align width to the right side).
func PushItemWidth(width float32) {
	C.iggPushItemWidth(C.float(width))
}

// PopItemWidth must be called for each call to PushItemWidth().
func PopItemWidth() {
	C.iggPopItemWidth()
}

// SetNextItemWidth sets width of the _next_ common large "item+label" widget.
// >0.0f: width in pixels, <0.0f align xx pixels to the right of window (so -math.SmallestNonzeroFloat32 always align width to the right side).
func SetNextItemWidth(width float32) {
	C.iggSetNextItemWidth(C.float(width))
}

// ItemFlags for PushItemFlag().
type ItemFlags int

const (
	// ItemFlagsNone default = 0.
	ItemFlagsNone ItemFlags = 0
	// ItemFlagsNoTabStop has no tab stop.
	ItemFlagsNoTabStop ItemFlags = 1 << 0
	// ItemFlagsButtonRepeat will return true multiple times based on io.KeyRepeatDelay and io.KeyRepeatRate settings.
	ItemFlagsButtonRepeat ItemFlags = 1 << 1
	// ItemFlagsDisabled [BETA] disable interactions but doesn't affect visuals yet. See github.com/ocornut/imgui/issues/211.
	ItemFlagsDisabled ItemFlags = 1 << 2
	// ItemFlagsNoNav has no nav.
	ItemFlagsNoNav ItemFlags = 1 << 3
	// ItemFlagsNoNavDefaultFocus has no nav default focus.
	ItemFlagsNoNavDefaultFocus ItemFlags = 1 << 4
	// ItemFlagsSelectableDontClosePopup automatically closes current Popup window.
	ItemFlagsSelectableDontClosePopup ItemFlags = 1 << 5
	// ItemFlagsMixedValue [BETA] represent a mixed/indeterminate value, generally multi-selection where values differ.
	// Currently only supported by Checkbox() (later should support all sorts of widgets).
	ItemFlagsMixedValue ItemFlags = 1 << 6
	// ItemFlagsDefault default = 0.
	ItemFlagsDefault ItemFlags = 0
)

// PushItemFlag changes flags in the existing options for the next items until PopItemFlag() is called.
func PushItemFlag(options ItemFlags, enabled bool) {
	C.iggPushItemFlag(C.int(options), castBool(enabled))
}

// PopItemFlag restores flags that were changed by the previous call to PushItemFlag().
func PopItemFlag() {
	C.iggPopItemFlag()
}

// CalcItemWidth returns the width of items given pushed settings and current cursor position.
func CalcItemWidth() float32 {
	return float32(C.iggCalcItemWidth())
}

// PushTextWrapPosV defines word-wrapping for Text() commands.
// < 0.0f: no wrapping; 0.0f: wrap to end of window (or column); > 0.0f: wrap at 'wrapPosX' position in window local space.
// Requires a matching call to PopTextWrapPos().
func PushTextWrapPosV(wrapPosX float32) {
	C.iggPushTextWrapPos(C.float(wrapPosX))
}

// PushTextWrapPos calls PushTextWrapPosV(0).
func PushTextWrapPos() {
	PushTextWrapPosV(0)
}

// PopTextWrapPos resets the last pushed position.
func PopTextWrapPos() {
	C.iggPopTextWrapPos()
}

// PushAllowKeyboardFocus allow the next window to take focus of the keyboard.
func PushAllowKeyboardFocus(allow bool) {
	C.iggPushAllowKeyboardFocus(castBool(allow))
}

//PopAllowKeyboardFocus pops most recent allow keyboard focus setting.
func PopAllowKeyboardFocus() {
	C.iggPopAllowKeyboardFocus()
}

// PushButtonRepeat enables button to repeat press if held.
func PushButtonRepeat(repeat bool) {
	C.iggPushButtonRepeat(castBool(repeat))
}

// PopButtonRepeat pops most recent button repeat setting.
func PopButtonRepeat() {
	C.iggPopButtonRepeat()
}

// Viewport A Platform Window (always only one in 'master' branch), in the future may represent Platform Monitor.
type Viewport uintptr

// ViewportFlags flags for viewport.
type ViewportFlags int

const (
	// ViewportFlagsNone default = 0.
	ViewportFlagsNone ViewportFlags = 0
	// ViewportFlagsIsPlatformWindow represents a Platform Window.
	ViewportFlagsIsPlatformWindow ViewportFlags = 1 << 0
	// ViewportFlagsIsPlatformMonitor represents a Platform Monitor (unused yet).
	ViewportFlagsIsPlatformMonitor ViewportFlags = 1 << 1
	// ViewportFlagsOwnedByApp Platform Window: is created/managed by the application (rather than a dear imgui backend).
	ViewportFlagsOwnedByApp ViewportFlags = 1 << 2
)

// MainViewport returns primary/default viewport.
// - Currently represents the Platform Window created by the application which is hosting our Dear ImGui windows.
// - In 'docking' branch with multi-viewport enabled, we extend this concept to have multiple active viewports.
// - In the future we will extend this concept further to also represent Platform Monitor and support a "no main platform window" operation mode.
func MainViewport() Viewport {
	return Viewport(C.iggGetMainViewport())
}

func (viewport Viewport) handle() C.IggViewport {
	return C.IggViewport(viewport)
}

// Flags returns viewports flags value.
func (viewport Viewport) Flags() ViewportFlags {
	if viewport == 0 {
		return ViewportFlagsNone
	}
	return ViewportFlags(C.iggViewportGetFlags(viewport.handle()))
}

// Pos returns viewports Main Area: Position of the viewport (Dear Imgui coordinates are the same as OS desktop/native coordinates).
func (viewport Viewport) Pos() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetPos(viewport.handle(), valueArg)
	valueFin()
	return value
}

// Size returns viewports Main Area: Size of the viewport.
func (viewport Viewport) Size() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetSize(viewport.handle(), valueArg)
	valueFin()
	return value
}

// WorkPos returns viewports Work Area: Position of the viewport minus task bars, menus bars, status bars (>= Pos).
func (viewport Viewport) WorkPos() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetWorkPos(viewport.handle(), valueArg)
	valueFin()
	return value
}

// WorkSize returns viewports Work Area: Size of the viewport minus task bars, menu bars, status bars (<= Size).
func (viewport Viewport) WorkSize() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetWorkSize(viewport.handle(), valueArg)
	valueFin()
	return value
}

// Center returns center of the viewport.
func (viewport Viewport) Center() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetCenter(viewport.handle(), valueArg)
	valueFin()
	return value
}

// WorkCenter returns center of the viewport minus task bars, menu bars, status bars.
func (viewport Viewport) WorkCenter() Vec2 {
	if viewport == 0 {
		return Vec2{}
	}

	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggViewportGetWorkCenter(viewport.handle(), valueArg)
	valueFin()
	return value
}
