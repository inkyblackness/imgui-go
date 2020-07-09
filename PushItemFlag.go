package imgui

// #include "wrapper/PushItemFlag.h"
import "C"

// Transient per-window flags, reset at the beginning of the frame. For child window, inherited from parent on first Begin().
// This is going to be exposed in imgui.h when stabilized enough.
const (
	ItemFlagsNone                     = 0
	ItemFlagsNoTabStop                = 1 << 0 // false
	ItemFlagsButtonRepeat             = 1 << 1 // false    // Button() will return true multiple times based on io.KeyRepeatDelay and io.KeyRepeatRate settings.
	ItemFlagsDisabled                 = 1 << 2 // false    // [BETA] Disable interactions but doesn't affect visuals yet. See github.com/ocornut/imgui/issues/211
	ItemFlagsNoNav                    = 1 << 3 // false
	ItemFlagsNoNavDefaultFocus        = 1 << 4 // false
	ItemFlagsSelectableDontClosePopup = 1 << 5 // false    // MenuItem/Selectable() automatically closes current Popup window
	ItemFlagsMixedValue               = 1 << 6 // false    // [BETA] Represent a mixed/indeterminate value, generally multi-selection where values differ. Currently only supported by Checkbox() (later should support all sorts of widgets)
	ItemFlagsDefault_                 = 0
)

// PushItemFlag pushes an option flag value on the stack to temporarily modify a state.
func PushItemFlag(option int) {
	C.iggPushItemFlag(C.int(option))
}

// PopItemFlag reverts the PushItemFlag change.
func PopItemFlag() {
	C.iggPopItemFlag()
}
