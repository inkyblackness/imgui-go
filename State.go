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

// IsItemHovered calls IsItemHoveredV(HoveredFlagsNone)
func IsItemHovered() bool {
	return IsItemHoveredV(HoveredFlagsNone)
}

// IsItemActive returns if the last item is active. (e.g. button being held, text field being edited. This will continuously return true while holding mouse button on an item. Items that don't interact will always return false)
func IsItemActive() bool {
	return C.iggIsItemActive() != 0
}

// IsAnyItemActive returns true if the any item is active.
func IsAnyItemActive() bool {
	return C.iggIsAnyItemActive() != 0
}

// IsItemVisible returns true if the last item is visible
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

// IsWindowFocusedV returns if current window is focused or its root/child, depending on flags. See flags for options.
func IsWindowFocusedV(flags int) bool {
	return C.iggIsWindowFocused(C.int(flags)) != 0
}

// IsWindowFocused calls IsWindowFocusedV(FocusedFlagsNone)
func IsWindowFocused() bool {
	return IsWindowFocusedV(FocusedFlagsNone)
}

// IsWindowHoveredV returns if current window is hovered (and typically: not blocked by a popup/modal).
// See flags for options. NB: If you are trying to check whether your mouse should be dispatched to imgui or to your app,
// you should use the 'io.WantCaptureMouse' boolean for that!
func IsWindowHoveredV(flags int) bool {
	return C.iggIsWindowHovered(C.int(flags)) != 0
}

// IsWindowHovered calls IsWindowHoveredV(HoveredFlagsNone)
func IsWindowHovered() bool {
	return IsWindowHoveredV(HoveredFlagsNone)
}

// IsKeyDown returns true if the corresponding key is currently being held down.
func IsKeyDown(key int) bool {
	return C.iggIsKeyDown(C.int(key)) != 0
}

// IsKeyPressedV returns true if the corresponding key was pressed (went from !Down to Down).
// If repeat=true and the key is being held down then the press is repeated using io.KeyRepeatDelay and KeyRepeatRate
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
// If repeat=true and the mouse button is being held down then the click is repeated using io.KeyRepeatDelay and KeyRepeatRate
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
