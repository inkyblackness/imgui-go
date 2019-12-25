package imgui

const (
	// SelectableFlagsNone default = 0
	SelectableFlagsNone = 0
	// SelectableFlagsDontClosePopups Clicking this don't close parent popup window
	SelectableFlagsDontClosePopups = 1 << 0
	// SelectableFlagsSpanAllColumns Selectable frame can span all columns (text will still fit in current column)
	SelectableFlagsSpanAllColumns = 1 << 1
	// SelectableFlagsAllowDoubleClick Generate press events on double clicks too
	SelectableFlagsAllowDoubleClick = 1 << 2
	// SelectableFlagsDisabled Cannot be selected, display grayed out text
	SelectableFlagsDisabled = 1 << 3
)
