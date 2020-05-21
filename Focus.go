package imgui

// #include "wrapper/Focus.h"
import "C"

// SetItemDefaultFocus makes the last item the default focused item of a window.
func SetItemDefaultFocus() {
	C.iggSetItemDefaultFocus()
}

// IsItemFocused returns true if the last item is focused.
func IsItemFocused() bool {
	return C.iggIsItemFocused() != 0
}

// IsAnyItemFocused returns true if any item is focused.
func IsAnyItemFocused() bool {
	return C.iggIsAnyItemFocused() != 0
}

// SetKeyboardFocusHere calls SetKeyboardFocusHereV(0).
func SetKeyboardFocusHere() {
	C.iggSetKeyboardFocusHere(0)
}

// SetKeyboardFocusHereV gives keyboard focus to next item.
func SetKeyboardFocusHereV(offset int) {
	C.iggSetKeyboardFocusHere(C.int(offset))
}
