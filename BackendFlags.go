package imgui

const (
	// BackendFlagHasGamepad back-end supports gamepad and currently has one connected.
	BackendFlagHasGamepad = 1 << iota
	// BackendFlagHasMouseCursors back-end supports honoring MouseCursor() value to change the OS cursor shape.
	BackendFlagHasMouseCursors
	// BackendFlagHasSetMousePos back-end supports io.WantSetMousePos requests to reposition the OS mouse position
	// (only used if ImGuiConfigFlags_NavEnableSetMousePos is set).
	BackendFlagHasSetMousePos
)
