package imgui

const (
	// WindowFlagsNoTitleBar disables the title-bar.
	WindowFlagsNoTitleBar = 1 << 0
	// WindowFlagsNoResize disables user resizing with the lower-right grip.
	WindowFlagsNoResize = 1 << 1
	// WindowFlagsNoMove disables user moving the window.
	WindowFlagsNoMove = 1 << 2
	// WindowFlagsNoScrollbar disables scrollbars (window can still scroll with mouse or programatically).
	WindowFlagsNoScrollbar = 1 << 3
	// WindowFlagsNoScrollWithMouse disables user vertically scrolling with mouse wheel.
	// On child window, mouse wheel will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoScrollWithMouse = 1 << 4
	// WindowFlagsNoCollapse disables user collapsing window by double-clicking on it.
	WindowFlagsNoCollapse = 1 << 5
	// WindowFlagsAlwaysAutoResize resizes every window to its content every frame.
	WindowFlagsAlwaysAutoResize = 1 << 6
	// WindowFlagsNoSavedSettings prohibits load/save settings in .ini file.
	WindowFlagsNoSavedSettings = 1 << 8
	// WindowFlagsNoInputs disables catching mouse or keyboard inputs, hovering test with pass through.
	WindowFlagsNoInputs = 1 << 9
	// WindowFlagsMenuBar has a menu-bar.
	WindowFlagsMenuBar = 1 << 10
	// WindowFlagsHorizontalScrollbar allows horizontal scrollbar to appear (off by default).
	// You may use SetNextWindowContentSize(Vec2(width,0.0f)); prior to calling Begin() to specify width.
	// Read code in imgui_demo in the "Horizontal Scrolling" section.
	WindowFlagsHorizontalScrollbar = 1 << 11
	// WindowFlagsNoFocusOnAppearing disables taking focus when transitioning from hidden to visible state.
	WindowFlagsNoFocusOnAppearing = 1 << 12
	// WindowFlagsNoBringToFrontOnFocus disables bringing window to front when taking focus.
	// (e.g. clicking on it or programatically giving it focus)
	WindowFlagsNoBringToFrontOnFocus = 1 << 13
	// WindowFlagsAlwaysVerticalScrollbar shows always a vertical scrollbar (even if ContentSize.y < Size.y)
	WindowFlagsAlwaysVerticalScrollbar = 1 << 14
	// WindowFlagsAlwaysHorizontalScrollbar shows always a horizontal scrollbar (even if ContentSize.x < Size.x).
	WindowFlagsAlwaysHorizontalScrollbar = 1 << 15
	// WindowFlagsAlwaysUseWindowPadding ensures child windows without border uses style.WindowPadding.
	// (ignored by default for non-bordered child windows, because more convenient)
	WindowFlagsAlwaysUseWindowPadding = 1 << 16
	// WindowFlagsResizeFromAnySide [BETA] enables resize from any corners and borders.
	// Your back-end needs to honor the different values of io.MouseCursor set by imgui.
	WindowFlagsResizeFromAnySide = 1 << 17
	// WindowFlagsNoNavInputs disables gamepad/keyboard navigation within the window.
	WindowFlagsNoNavInputs = 1 << 18
	// WindowFlagsNoNavFocus disables focusing toward this window with gamepad/keyboard navigation (e.g. skipped by CTRL+TAB)
	WindowFlagsNoNavFocus = 1 << 19
	// WindowFlagsNoNav combines WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoNav = WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
)
