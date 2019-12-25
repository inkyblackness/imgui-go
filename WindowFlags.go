package imgui

const (
	// WindowFlagsNone default = 0
	WindowFlagsNone = 0
	// WindowFlagsNoTitleBar Disable title-bar
	WindowFlagsNoTitleBar = 1 << 0
	// WindowFlagsNoResize Disable user resizing with the lower-right grip
	WindowFlagsNoResize = 1 << 1
	// WindowFlagsNoMove Disable user moving the window
	WindowFlagsNoMove = 1 << 2
	// WindowFlagsNoScrollbar Disable scrollbars (window can still scroll with mouse or programmatically)
	WindowFlagsNoScrollbar = 1 << 3
	// WindowFlagsNoScrollWithMouse Disable user vertically scrolling with mouse wheel. On child window, mouse wheel
	// will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoScrollWithMouse = 1 << 4
	// WindowFlagsNoCollapse Disable user collapsing window by double-clicking on it
	WindowFlagsNoCollapse = 1 << 5
	// WindowFlagsAlwaysAutoResize Resize every window to its content every frame
	WindowFlagsAlwaysAutoResize = 1 << 6
	// WindowFlagsNoBackground Disable drawing background color (WindowBg, etc.) and outside border. Similar as using
	// SetNextWindowBgAlpha(0.0f).
	WindowFlagsNoBackground = 1 << 7
	// WindowFlagsNoSavedSettings Never load/save settings in .ini file
	WindowFlagsNoSavedSettings = 1 << 8
	// WindowFlagsNoMouseInputs Disable catching mouse, hovering test with pass through.
	WindowFlagsNoMouseInputs = 1 << 9
	// WindowFlagsMenuBar Has a menu-bar
	WindowFlagsMenuBar = 1 << 10
	// WindowFlagsHorizontalScrollbar Allow horizontal scrollbar to appear (off by default). You may use
	// SetNextWindowContentSize(ImVec2(width,0.0f)); prior to calling Begin() to specify width. Read code in imgui_demo
	// in the "Horizontal Scrolling" section.
	WindowFlagsHorizontalScrollbar = 1 << 11
	// WindowFlagsNoFocusOnAppearing Disable taking focus when transitioning from hidden to visible state
	WindowFlagsNoFocusOnAppearing = 1 << 12
	// WindowFlagsNoBringToFrontOnFocus Disable bringing window to front when taking focus (e.g. clicking on it or
	// programmatically giving it focus)
	WindowFlagsNoBringToFrontOnFocus = 1 << 13
	// WindowFlagsAlwaysVerticalScrollbar Always show vertical scrollbar (even if ContentSize.y < Size.y)
	WindowFlagsAlwaysVerticalScrollbar = 1 << 14
	// WindowFlagsAlwaysHorizontalScrollbar Always show horizontal scrollbar (even if ContentSize.x < Size.x)
	WindowFlagsAlwaysHorizontalScrollbar = 1 << 15
	// WindowFlagsAlwaysUseWindowPadding Ensure child windows without border uses style.WindowPadding (ignored by
	// default for non-bordered child windows, because more convenient)
	WindowFlagsAlwaysUseWindowPadding = 1 << 16
	// WindowFlagsNoNavInputs No gamepad/keyboard navigation within the window
	WindowFlagsNoNavInputs = 1 << 18
	// WindowFlagsNoNavFocus No focusing toward this window with gamepad/keyboard navigation (e.g. skipped by CTRL+TAB)
	WindowFlagsNoNavFocus = 1 << 19
	// WindowFlagsUnsavedDocument Append '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. When used in a tab/docking context, tab is selected on closure and closure is deferred by one
	// frame to allow code to cancel the closure (with a confirmation popup, etc.) without flicker.
	WindowFlagsUnsavedDocument = 1 << 20

	// WindowFlagsNoNav combines WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoNav = WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
	// WindowFlagsNoDecoration combines WindowFlagsNoTitleBar, WindowFlagsNoResize, WindowFlagsNoScrollbar and
	// WindowFlagsNoCollapse.
	WindowFlagsNoDecoration = WindowFlagsNoTitleBar | WindowFlagsNoResize | WindowFlagsNoScrollbar | WindowFlagsNoCollapse
	// WindowFlagsNoInputs combines WindowFlagsNoMouseInputs, WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoInputs = WindowFlagsNoMouseInputs | WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
)
