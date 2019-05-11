package imgui

const (
	// Back-end supports gamepad and currently has one connected.
	BackendFlagHasGamepad = 1 << iota
	// Back-end supports honoring GetMouseCursor() value to change the OS cursor shape.
	BackendFlagHasMouseCursors
	// Back-end supports io.WantSetMousePos requests to reposition the OS mouse position (only used if ImGuiConfigFlags_NavEnableSetMousePos is set).
	BackendFlagHasSetMousePos
)
