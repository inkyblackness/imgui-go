package imgui

const (
	// ColumnsFlagsNone default = 0.
	ColumnsFlagsNone = 0
	// ColumnsFlagsNoBorder Disable column dividers.
	ColumnsFlagsNoBorder = 1 << 0
	// ColumnsFlagsNoResize Disable resizing columns when clicking on the dividers.
	ColumnsFlagsNoResize = 1 << 1
	// ColumnsFlagsNoPreserveWidths Disable column width preservation when adjusting columns.
	ColumnsFlagsNoPreserveWidths = 1 << 2
	// ColumnsFlagsNoForceWithinWindow Disable forcing columns to fit within window.
	ColumnsFlagsNoForceWithinWindow = 1 << 3
	// ColumnsFlagsGrowParentContentsSize (WIP) Restore pre-1.51 behavior of extending the parent window contents
	// size but _without affecting the columns width at all_. Will eventually remove.
	ColumnsFlagsGrowParentContentsSize = 1 << 4
)
