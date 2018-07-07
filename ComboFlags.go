package imgui

const (
	// ComboFlagPopupAlignLeft aligns the popup toward the left by default.
	ComboFlagPopupAlignLeft = 1 << iota
	// ComboFlagHeightSmall has max ~4 items visible.
	// Tip: If you want your combo popup to be a specific size you can use SetNextWindowSizeConstraints() prior to calling BeginCombo().
	ComboFlagHeightSmall
	// ComboFlagHeightRegular has max ~8 items visible (default).
	ComboFlagHeightRegular
	// ComboFlagHeightLarge has max ~20 items visible.
	ComboFlagHeightLarge
	// ComboFlagHeightLargest has as many fitting items as possible.
	ComboFlagHeightLargest
	// ComboFlagNoArrowButton displays on the preview box without the square arrow button.
	ComboFlagNoArrowButton
	// ComboFlagNoPreview displays only a square arrow button.
	ComboFlagNoPreview
)