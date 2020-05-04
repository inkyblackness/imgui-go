package imgui

const (
	// ColorEditFlagsNone default = 0
	ColorEditFlagsNone = 0
	// ColorEditFlagsNoAlpha ignores Alpha component (read 3 components from the input pointer).
	ColorEditFlagsNoAlpha = 1 << 1
	// ColorEditFlagsNoPicker disables picker when clicking on colored square.
	ColorEditFlagsNoPicker = 1 << 2
	// ColorEditFlagsNoOptions disables toggling options menu when right-clicking on inputs/small preview.
	ColorEditFlagsNoOptions = 1 << 3
	// ColorEditFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs)
	ColorEditFlagsNoSmallPreview = 1 << 4
	// ColorEditFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorEditFlagsNoInputs = 1 << 5
	// ColorEditFlagsNoTooltip disables tooltip when hovering the preview.
	ColorEditFlagsNoTooltip = 1 << 6
	// ColorEditFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorEditFlagsNoLabel = 1 << 7
	// ColorEditFlagsNoSidePreview disables bigger color preview on right side of the picker, use small colored square preview instead.
	ColorEditFlagsNoSidePreview = 1 << 8
	// ColorEditFlagsNoDragDrop disables drag and drop target. ColorButton: disable drag and drop source.
	ColorEditFlagsNoDragDrop = 1 << 9
	// ColorEditFlagsNoBorder disables border (which is enforced by default)
	ColorEditFlagsNoBorder = 1 << 10

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions(). The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call SetColorEditOptions() during startup.

	// ColorEditFlagsAlphaBar shows vertical alpha bar/gradient in picker.
	ColorEditFlagsAlphaBar = 1 << 16
	// ColorEditFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreview = 1 << 17
	// ColorEditFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreviewHalf = 1 << 18
	// ColorEditFlagsHDR = (WIP) surrently only disable 0.0f..1.0f limits in RGBA edition (note: you probably want to use ImGuiColorEditFlags_Float flag as well).
	ColorEditFlagsHDR = 1 << 19
	// ColorEditFlagsRGB sets the format as RGB
	ColorEditFlagsRGB = 1 << 20
	// ColorEditFlagsHSV sets the format as HSV
	ColorEditFlagsHSV = 1 << 21
	// ColorEditFlagsHEX sets the format as HEX
	ColorEditFlagsHEX = 1 << 22
	// ColorEditFlagsUint8 _display_ values formatted as 0..255.
	ColorEditFlagsUint8 = 1 << 23
	// ColorEditFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorEditFlagsFloat = 1 << 24

	// ColorEditFlagsPickerHueBar shows bar for Hue, rectangle for Sat/Value.
	ColorEditFlagsPickerHueBar = 1 << 25
	// ColorEditFlagsPickerHueWheel shows wheel for Hue, triangle for Sat/Value.
	ColorEditFlagsPickerHueWheel = 1 << 26
	// ColorEditFlagsInputRGB enables input and output data in RGB format.
	ColorEditFlagsInputRGB = 1 << 27
	// ColorEditFlagsInputHSV enables input and output data in HSV format.
	ColorEditFlagsInputHSV = 1 << 28
)
