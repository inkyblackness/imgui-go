package imgui

const (
	// FocusedFlagsNone Return true if directly over the item/window, not obstructed by another window,
	// not obstructed by an active popup or modal blocking inputs under them.
	FocusedFlagsNone = 0
	// FocusedFlagsChildWindows returns true if any children of the window is focused
	FocusedFlagsChildWindows = 1 << 0
	// FocusedFlagsRootWindow tests from root window (top most parent of the current hierarchy)
	FocusedFlagsRootWindow = 1 << 1
	// FocusedFlagsAnyWindow returns true if any window is focused. Important: If you are trying to tell how to dispatch your low-level inputs, do NOT use this. Use WantCaptureMouse instead.
	FocusedFlagsAnyWindow = 1 << 2
)

// FocusedFlags combinations
const (
	FocusedFlagsRootAndChildWindows = FocusedFlagsRootWindow | FocusedFlagsChildWindows
)
