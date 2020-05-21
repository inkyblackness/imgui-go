package imgui

// #include "wrapper/State.h"
import "C"

// IsItemClicked returns true if the current item is clicked with the left
// mouse button.
func IsItemClicked() bool {
	return C.iggIsItemClicked() != 0
}

// IsItemHoveredV returns true if the last item is hovered.
// (and usable, aka not blocked by a popup, etc.). See HoveredFlags for more options.
func IsItemHoveredV(flags int) bool {
	return C.iggIsItemHovered(C.int(flags)) != 0
}

// IsItemHovered calls IsItemHoveredV(HoveredFlagsNone).
func IsItemHovered() bool {
	return IsItemHoveredV(HoveredFlagsNone)
}

// IsItemActive returns if the last item is active.
// e.g. button being held, text field being edited.
//
// This will continuously return true while holding mouse button on an item.
// Items that don't interact will always return false.
func IsItemActive() bool {
	return C.iggIsItemActive() != 0
}

// IsAnyItemActive returns true if the any item is active.
func IsAnyItemActive() bool {
	return C.iggIsAnyItemActive() != 0
}

// IsItemVisible returns true if the last item is visible.
func IsItemVisible() bool {
	return C.iggIsItemVisible() != 0
}

// IsWindowAppearing returns whether the current window is appearing.
func IsWindowAppearing() bool {
	return C.iggIsWindowAppearing() != 0
}

// IsWindowCollapsed returns whether the current window is collapsed.
func IsWindowCollapsed() bool {
	return C.iggIsWindowCollapsed() != 0
}

// This is a list of FocusedFlags combinations.
const (
	// FocusedFlagsNone Return true if directly over the item/window, not obstructed by another window,
	// not obstructed by an active popup or modal blocking inputs under them.
	FocusedFlagsNone = 0
	// FocusedFlagsChildWindows returns true if any children of the window is focused
	FocusedFlagsChildWindows = 1 << 0
	// FocusedFlagsRootWindow tests from root window (top most parent of the current hierarchy)
	FocusedFlagsRootWindow = 1 << 1
	// FocusedFlagsAnyWindow returns true if any window is focused.
	// Important: If you are trying to tell how to dispatch your low-level inputs, do NOT use this.
	// Use WantCaptureMouse instead.
	FocusedFlagsAnyWindow = 1 << 2

	FocusedFlagsRootAndChildWindows = FocusedFlagsRootWindow | FocusedFlagsChildWindows
)

// IsWindowFocusedV returns if current window is focused or its root/child, depending on flags. See flags for options.
func IsWindowFocusedV(flags int) bool {
	return C.iggIsWindowFocused(C.int(flags)) != 0
}

// IsWindowFocused calls IsWindowFocusedV(FocusedFlagsNone).
func IsWindowFocused() bool {
	return IsWindowFocusedV(FocusedFlagsNone)
}

// This is a list of HoveredFlags combinations.
const (
	// HoveredFlagsNone Return true if directly over the item/window, not obstructed by another window,
	// not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsNone = 0
	// HoveredFlagsChildWindows IsWindowHovered() only: Return true if any children of the window is hovered.
	HoveredFlagsChildWindows = 1 << 0
	// HoveredFlagsRootWindow IsWindowHovered() only: Test from root window (top most parent of the current hierarchy).
	HoveredFlagsRootWindow = 1 << 1
	// HoveredFlagsAnyWindow IsWindowHovered() only: Return true if any window is hovered.
	HoveredFlagsAnyWindow = 1 << 2
	// HoveredFlagsAllowWhenBlockedByPopup Return true even if a popup window is normally blocking access to this item/window.
	HoveredFlagsAllowWhenBlockedByPopup = 1 << 3
	// HoveredFlagsAllowWhenBlockedByActiveItem Return true even if an active item is blocking access to this item/window.
	// Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenBlockedByActiveItem = 1 << 5
	// HoveredFlagsAllowWhenOverlapped Return true even if the position is overlapped by another window
	HoveredFlagsAllowWhenOverlapped = 1 << 6
	// HoveredFlagsAllowWhenDisabled Return true even if the item is disabled
	HoveredFlagsAllowWhenDisabled = 1 << 7

	HoveredFlagsRectOnly            = HoveredFlagsAllowWhenBlockedByPopup | HoveredFlagsAllowWhenBlockedByActiveItem | HoveredFlagsAllowWhenOverlapped
	HoveredFlagsRootAndChildWindows = HoveredFlagsRootWindow | HoveredFlagsChildWindows
)

// IsWindowHoveredV returns if current window is hovered (and typically: not blocked by a popup/modal).
// See flags for options. NB: If you are trying to check whether your mouse should be dispatched to imgui or to your app,
// you should use the 'io.WantCaptureMouse' boolean for that!
func IsWindowHoveredV(flags int) bool {
	return C.iggIsWindowHovered(C.int(flags)) != 0
}

// IsWindowHovered calls IsWindowHoveredV(HoveredFlagsNone).
func IsWindowHovered() bool {
	return IsWindowHoveredV(HoveredFlagsNone)
}

// IsKeyDown returns true if the corresponding key is currently being held down.
func IsKeyDown(key int) bool {
	return C.iggIsKeyDown(C.int(key)) != 0
}

// IsKeyPressedV returns true if the corresponding key was pressed (went from !Down to Down).
// If repeat=true and the key is being held down then the press is repeated using io.KeyRepeatDelay and KeyRepeatRate.
func IsKeyPressedV(key int, repeat bool) bool {
	return C.iggIsKeyPressed(C.int(key), castBool(repeat)) != 0
}

// IsKeyPressed calls IsKeyPressedV(key, true).
func IsKeyPressed(key int) bool {
	return IsKeyPressedV(key, true)
}

// IsKeyReleased returns true if the corresponding key was released (went from Down to !Down).
func IsKeyReleased(key int) bool {
	return C.iggIsKeyReleased(C.int(key)) != 0
}

// IsMouseDown returns true if the corresponding mouse button is currently being held down.
func IsMouseDown(button int) bool {
	return C.iggIsMouseDown(C.int(button)) != 0
}

// IsAnyMouseDown returns true if any mouse button is currently being held down.
func IsAnyMouseDown() bool {
	return C.iggIsAnyMouseDown() != 0
}

// IsMouseClickedV returns true if the mouse button was clicked (0=left, 1=right, 2=middle)
// If repeat=true and the mouse button is being held down then the click is repeated using io.KeyRepeatDelay and KeyRepeatRate.
func IsMouseClickedV(button int, repeat bool) bool {
	return C.iggIsMouseClicked(C.int(button), castBool(repeat)) != 0
}

// IsMouseClicked calls IsMouseClickedV(key, false).
func IsMouseClicked(button int) bool {
	return IsMouseClickedV(button, false)
}

// IsMouseReleased returns true if the mouse button was released (went from Down to !Down).
func IsMouseReleased(button int) bool {
	return C.iggIsMouseReleased(C.int(button)) != 0
}

// IsMouseDoubleClicked returns true if the mouse button was double-clicked (0=left, 1=right, 2=middle).
func IsMouseDoubleClicked(button int) bool {
	return C.iggIsMouseDoubleClicked(C.int(button)) != 0
}

// MousePos returns the current window position in screen space.
func MousePos() Vec2 {
	var value Vec2
	valueArg, valueFin := value.wrapped()
	C.iggMousePos(valueArg)
	valueFin()
	return value
}

// Enumeration for MouseCursor()
// User code may request binding to display given cursor by calling SetMouseCursor(),
// which is why we have some cursors that are marked unused here.
const (
	// MouseCursorNone no mouse cursor
	MouseCursorNone = -1
	// MouseCursorArrow standard arrow mouse cursor
	MouseCursorArrow = 0
	// MouseCursorTextInput when hovering over InputText, etc.
	MouseCursorTextInput = 1
	// MouseCursorResizeAll (Unused by imgui functions)
	MouseCursorResizeAll = 2
	// MouseCursorResizeNS when hovering over an horizontal border
	MouseCursorResizeNS = 3
	// MouseCursorResizeEW when hovering over a vertical border or a column
	MouseCursorResizeEW = 4
	// MouseCursorResizeNESW when hovering over the bottom-left corner of a window
	MouseCursorResizeNESW = 5
	// MouseCursorResizeNWSE when hovering over the bottom-right corner of a window
	MouseCursorResizeNWSE = 6
	// MouseCursorHand (Unused by imgui functions. Use for e.g. hyperlinks)
	MouseCursorHand  = 7
	MouseCursorCount = 8
)

// MouseCursor returns desired cursor type, reset in imgui.NewFrame(), this is updated during the frame.
// Valid before Render(). If you use software rendering by setting io.MouseDrawCursor ImGui will render those for you.
func MouseCursor() int {
	return int(C.iggGetMouseCursor())
}

// SetMouseCursor sets desired cursor type.
func SetMouseCursor(cursor int) {
	C.iggSetMouseCursor(C.int(cursor))
}
