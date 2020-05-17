package imgui

const (
	// ConditionAlways sets the variable.
	ConditionAlways Condition = 1 << 0
	// ConditionOnce sets the variable once per runtime session (only the first call with succeed).
	ConditionOnce = 1 << 1
	// ConditionFirstUseEver sets the variable if the object/window has no persistently saved data (no entry in .ini file).
	ConditionFirstUseEver = 1 << 2
	// ConditionAppearing sets the variable if the object/window is appearing after being hidden/inactive (or the first time).
	ConditionAppearing = 1 << 3
)

// Constants to fill IO.KeyMap() lookup with indices into the IO.KeysDown[512] array.
// The mapped indices are then the ones reported to IO.KeyPress() and IO.KeyRelease().
const (
	KeyTab         = 0
	KeyLeftArrow   = 1
	KeyRightArrow  = 2
	KeyUpArrow     = 3
	KeyDownArrow   = 4
	KeyPageUp      = 5
	KeyPageDown    = 6
	KeyHome        = 7
	KeyEnd         = 8
	KeyInsert      = 9
	KeyDelete      = 10
	KeyBackspace   = 11
	KeySpace       = 12
	KeyEnter       = 13
	KeyEscape      = 14
	KeyKeyPadEnter = 15
	KeyA           = 16 // for text edit CTRL+A: select all
	KeyC           = 17 // for text edit CTRL+C: copy
	KeyV           = 18 // for text edit CTRL+V: paste
	KeyX           = 19 // for text edit CTRL+X: cut
	KeyY           = 20 // for text edit CTRL+Y: redo
	KeyZ           = 21 // for text edit CTRL+Z: undo
)

// StyleVarID identifies a style variable in the UI style.
type StyleVarID int

const (
	// StyleVarAlpha is a float
	StyleVarAlpha StyleVarID = 0
	// StyleVarWindowPadding is a Vec2
	StyleVarWindowPadding StyleVarID = 1
	// StyleVarWindowRounding is a float
	StyleVarWindowRounding StyleVarID = 2
	// StyleVarWindowBorderSize is a float
	StyleVarWindowBorderSize StyleVarID = 3
	// StyleVarWindowMinSize is a Vec2
	StyleVarWindowMinSize StyleVarID = 4
	// StyleVarWindowTitleAlign is a Vec2
	StyleVarWindowTitleAlign StyleVarID = 5
	// StyleVarChildRounding is a float
	StyleVarChildRounding StyleVarID = 6
	// StyleVarChildBorderSize is a float
	StyleVarChildBorderSize StyleVarID = 7
	// StyleVarPopupRounding is a float
	StyleVarPopupRounding StyleVarID = 8
	// StyleVarPopupBorderSize is a float
	StyleVarPopupBorderSize StyleVarID = 9
	// StyleVarFramePadding is a Vec2
	StyleVarFramePadding StyleVarID = 10
	// StyleVarFrameRounding is a float
	StyleVarFrameRounding StyleVarID = 11
	// StyleVarFrameBorderSize is a float
	StyleVarFrameBorderSize StyleVarID = 12
	// StyleVarItemSpacing is a Vec2
	StyleVarItemSpacing StyleVarID = 13
	// StyleVarItemInnerSpacing is a Vec2
	StyleVarItemInnerSpacing StyleVarID = 14
	// StyleVarIndentSpacing is a float
	StyleVarIndentSpacing StyleVarID = 15
	// StyleVarScrollbarSize is a float
	StyleVarScrollbarSize StyleVarID = 16
	// StyleVarScrollbarRounding is a float
	StyleVarScrollbarRounding StyleVarID = 17
	// StyleVarGrabMinSize is a float
	StyleVarGrabMinSize StyleVarID = 18
	// StyleVarGrabRounding is a float
	StyleVarGrabRounding StyleVarID = 19
	// StyleVarTabRounding is a float
	StyleVarTabRounding StyleVarID = 20
	// StyleVarButtonTextAlign is a Vec2
	StyleVarButtonTextAlign StyleVarID = 21
	// StyleVarSelectableTextAlign is a Vec2
	StyleVarSelectableTextAlign StyleVarID = 22
)

// StyleColor identifier
const (
	StyleColorText                  StyleColorID = 0
	StyleColorTextDisabled          StyleColorID = 1
	StyleColorWindowBg              StyleColorID = 2
	StyleColorChildBg               StyleColorID = 3
	StyleColorPopupBg               StyleColorID = 4
	StyleColorBorder                StyleColorID = 5
	StyleColorBorderShadow          StyleColorID = 6
	StyleColorFrameBg               StyleColorID = 7
	StyleColorFrameBgHovered        StyleColorID = 8
	StyleColorFrameBgActive         StyleColorID = 9
	StyleColorTitleBg               StyleColorID = 10
	StyleColorTitleBgActive         StyleColorID = 11
	StyleColorTitleBgCollapsed      StyleColorID = 12
	StyleColorMenuBarBg             StyleColorID = 13
	StyleColorScrollbarBg           StyleColorID = 14
	StyleColorScrollbarGrab         StyleColorID = 15
	StyleColorScrollbarGrabHovered  StyleColorID = 16
	StyleColorScrollbarGrabActive   StyleColorID = 17
	StyleColorCheckMark             StyleColorID = 18
	StyleColorSliderGrab            StyleColorID = 19
	StyleColorSliderGrabActive      StyleColorID = 20
	StyleColorButton                StyleColorID = 21
	StyleColorButtonHovered         StyleColorID = 22
	StyleColorButtonActive          StyleColorID = 23
	StyleColorHeader                StyleColorID = 24
	StyleColorHeaderHovered         StyleColorID = 25
	StyleColorHeaderActive          StyleColorID = 26
	StyleColorSeparator             StyleColorID = 27
	StyleColorSeparatorHovered      StyleColorID = 28
	StyleColorSeparatorActive       StyleColorID = 29
	StyleColorResizeGrip            StyleColorID = 30
	StyleColorResizeGripHovered     StyleColorID = 31
	StyleColorResizeGripActive      StyleColorID = 32
	StyleColorTab                   StyleColorID = 33
	StyleColorTabHovered            StyleColorID = 34
	StyleColorTabActive             StyleColorID = 35
	StyleColorTabUnfocused          StyleColorID = 36
	StyleColorTabUnfocusedActive    StyleColorID = 37
	StyleColorPlotLines             StyleColorID = 38
	StyleColorPlotLinesHovered      StyleColorID = 39
	StyleColorPlotHistogram         StyleColorID = 40
	StyleColorPlotHistogramHovered  StyleColorID = 41
	StyleColorTextSelectedBg        StyleColorID = 42
	StyleColorDragDropTarget        StyleColorID = 43
	StyleColorNavHighlight          StyleColorID = 44 // Gamepad/keyboard: current highlighted item
	StyleColorNavWindowingHighlight StyleColorID = 45 // Highlight window when using CTRL+TAB
	StyleColorNavWindowingDarkening StyleColorID = 46 // Darken/colorize entire screen behind the CTRL+TAB window list, when active
	StyleColorModalWindowDarkening  StyleColorID = 47 // Darken/colorize entire screen behind a modal window, when one is active
)

const (
	// BackendFlagNone default = 0
	BackendFlagNone = 0
	// BackendFlagHasGamepad back-end Platform supports gamepad and currently has one connected.
	BackendFlagHasGamepad = 1 << 0
	// BackendFlagHasMouseCursors back-end Platform supports honoring GetMouseCursor() value to change the OS cursor
	// shape.
	BackendFlagHasMouseCursors = 1 << 1
	// BackendFlagHasSetMousePos back-end Platform supports io.WantSetMousePos requests to reposition the OS mouse
	// position (only used if ImGuiConfigFlags_NavEnableSetMousePos is set).
	BackendFlagHasSetMousePos = 1 << 2
	// BackendFlagsRendererHasVtxOffset back-end Renderer supports ImDrawCmd::VtxOffset. This enables output of large
	// meshes (64K+ vertices) while still using 16-bits indices.
	BackendFlagsRendererHasVtxOffset = 1 << 3
)

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

const (
	// ColorPickerFlagsNone default = 0
	ColorPickerFlagsNone = 0
	// ColorPickerFlagsNoPicker disables picker when clicking on colored square.
	ColorPickerFlagsNoPicker = 1 << 2
	// ColorPickerFlagsNoOptions disables toggling options menu when right-clicking on inputs/small preview.
	ColorPickerFlagsNoOptions = 1 << 3
	// ColorPickerFlagsNoAlpha ignoreс Alpha component (read 3 components from the input pointer).
	ColorPickerFlagsNoAlpha = 1 << 1
	// ColorPickerFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs)
	ColorPickerFlagsNoSmallPreview = 1 << 4
	// ColorPickerFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorPickerFlagsNoInputs = 1 << 5
	// ColorPickerFlagsNoTooltip disables tooltip when hovering the preview.
	ColorPickerFlagsNoTooltip = 1 << 6
	// ColorPickerFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorPickerFlagsNoLabel = 1 << 7
	// ColorPickerFlagsNoSidePreview disables bigger color preview on right side of the picker, use small colored square preview instead.
	ColorPickerFlagsNoSidePreview = 1 << 8

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions(). The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call SetColorPickerOptions() during startup.

	// ColorPickerFlagsAlphaBar shows vertical alpha bar/gradient in picker.
	ColorPickerFlagsAlphaBar = 1 << 16
	// ColorPickerFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
	ColorPickerFlagsAlphaPreview = 1 << 17
	// ColorPickerFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
	ColorPickerFlagsAlphaPreviewHalf = 1 << 18
	// ColorPickerFlagsRGB sets the format as RGB
	ColorPickerFlagsRGB = 1 << 20
	// ColorPickerFlagsHSV sets the format as HSV
	ColorPickerFlagsHSV = 1 << 21
	// ColorPickerFlagsHEX sets the format as HEX
	ColorPickerFlagsHEX = 1 << 22
	// ColorPickerFlagsUint8 _display_ values formatted as 0..255.
	ColorPickerFlagsUint8 = 1 << 23
	// ColorPickerFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorPickerFlagsFloat = 1 << 24
	// ColorPickerFlagsPickerHueBar bar for Hue, rectangle for Sat/Value.
	ColorPickerFlagsPickerHueBar = 1 << 25
	// ColorPickerFlagsPickerHueWheel wheel for Hue, triangle for Sat/Value.
	ColorPickerFlagsPickerHueWheel = 1 << 26
	// ColorPickerFlagsInputRGB enables input and output data in RGB format.
	ColorPickerFlagsInputRGB = 1 << 27
	// ColorPickerFlagsInputHSV enables input and output data in HSV format.
	ColorPickerFlagsInputHSV = 1 << 28
)

const (
	// ComboFlagNone default = 0
	ComboFlagNone = 0
	// ComboFlagPopupAlignLeft aligns the popup toward the left by default.
	ComboFlagPopupAlignLeft = 1 << 0
	// ComboFlagHeightSmall has max ~4 items visible.
	// Tip: If you want your combo popup to be a specific size you can use SetNextWindowSizeConstraints() prior to calling BeginCombo().
	ComboFlagHeightSmall = 1 << 1
	// ComboFlagHeightRegular has max ~8 items visible (default).
	ComboFlagHeightRegular = 1 << 2
	// ComboFlagHeightLarge has max ~20 items visible.
	ComboFlagHeightLarge = 1 << 3
	// ComboFlagHeightLargest has as many fitting items as possible.
	ComboFlagHeightLargest = 1 << 4
	// ComboFlagNoArrowButton displays on the preview box without the square arrow button.
	ComboFlagNoArrowButton = 1 << 5
	// ComboFlagNoPreview displays only a square arrow button.
	ComboFlagNoPreview = 1 << 6
)

const (
	// ConfigFlagNone default = 0
	ConfigFlagNone = 0
	// ConfigFlagNavEnableKeyboard master keyboard navigation enable flag. NewFrame() will automatically fill
	// io.NavInputs[] based on io.KeysDown[].
	ConfigFlagNavEnableKeyboard = 1 << 0
	// ConfigFlagNavEnableGamepad master gamepad navigation enable flag.
	// This is mostly to instruct your imgui back-end to fill io.NavInputs[]. Back-end also needs to set
	// BackendFlagHasGamepad.
	ConfigFlagNavEnableGamepad = 1 << 1
	// ConfigFlagNavEnableSetMousePos instruct navigation to move the mouse cursor. May be useful on TV/console systems
	// where moving a virtual mouse is awkward. Will update io.MousePos and set io.WantSetMousePos=true. If enabled you
	// MUST honor io.WantSetMousePos requests in your binding, otherwise ImGui will react as if the mouse is jumping
	// around back and forth.
	ConfigFlagNavEnableSetMousePos = 1 << 2
	// ConfigFlagNavNoCaptureKeyboard instruct navigation to not set the io.WantCaptureKeyboard flag when io.NavActive
	// is set.
	ConfigFlagNavNoCaptureKeyboard = 1 << 3
	// ConfigFlagNoMouse instruct imgui to clear mouse position/buttons in NewFrame(). This allows ignoring the mouse
	// information set by the back-end.
	ConfigFlagNoMouse = 1 << 4
	// ConfigFlagNoMouseCursorChange instruct back-end to not alter mouse cursor shape and visibility. Use if the
	// back-end cursor changes are interfering with yours and you don't want to use SetMouseCursor() to change mouse
	// cursor. You may want to honor requests from imgui by reading GetMouseCursor() yourself instead.
	ConfigFlagNoMouseCursorChange = 1 << 5

	// User storage (to allow your back-end/engine to communicate to code that may be shared between multiple projects.
	// Those flags are not used by core Dear ImGui)

	// ConfigFlagIsSRGB application is SRGB-aware.
	ConfigFlagIsSRGB = 1 << 20
	// ConfigFlagIsTouchScreen application is using a touch screen instead of a mouse.
	ConfigFlagIsTouchScreen = 1 << 21
)

// BeginDragDropSource flags
const (
	// DragDropFlagsNone specifies the default behaviour.
	DragDropFlagsNone = 0
	// DragDropFlagsSourceNoPreviewTooltip hides the tooltip that is open so you can display a preview or description of the source contents.
	DragDropFlagsSourceNoPreviewTooltip = 1 << 0
	// DragDropFlagsSourceNoDisableHover preserves the behaviour of IsItemHovered. By default, when dragging we clear data so that IsItemHovered() will return true, to avoid subsequent user code submitting tooltips.
	DragDropFlagsSourceNoDisableHover = 1 << 1
	// DragDropFlagsSourceNoHoldToOpenOthers disables the behavior that allows to open tree nodes and collapsing header by holding over them while dragging a source item.
	DragDropFlagsSourceNoHoldToOpenOthers = 1 << 2
	// DragDropFlagsSourceAllowNullID allows items such as Text(), Image() that have no unique identifier to be used as drag source, by manufacturing a temporary identifier based on their window-relative position. This is extremely unusual within the dear ecosystem and so we made it explicit.
	DragDropFlagsSourceAllowNullID = 1 << 3
	// DragDropFlagsSourceExtern specifies external source (from outside of), won't attempt to read current item/window info. Will always return true. Only one Extern source can be active simultaneously.
	DragDropFlagsSourceExtern = 1 << 4
)

// AcceptDragDropPayload flags
const (
	// DragDropFlagsAcceptBeforeDelivery makes AcceptDragDropPayload() return true even before the mouse button is released. You can then call IsDelivery() to test if the payload needs to be delivered.
	DragDropFlagsAcceptBeforeDelivery = 1 << 10
	// DragDropFlagsAcceptNoDrawDefaultRect does not draw the default highlight rectangle when hovering over target.
	DragDropFlagsAcceptNoDrawDefaultRect = 1 << 11
	// DragDropFlagsAcceptPeekOnly is for peeking ahead and inspecting the payload before delivery.
	DragDropFlagsAcceptPeekOnly = DragDropFlagsAcceptBeforeDelivery | DragDropFlagsAcceptNoDrawDefaultRect
)

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

	FocusedFlagsRootAndChildWindows = FocusedFlagsRootWindow | FocusedFlagsChildWindows
)

// Flags for FreeType rasterizer. By default, hinting is enabled and the font's native hinter is preferred over the auto-hinter.
const (
	// FreeTypeRasterizerFlagsNoHinting disables hinting.
	// This generally generates 'blurrier' bitmap glyphs when the glyph are rendered in any of the anti-aliased modes.
	FreeTypeRasterizerFlagsNoHinting = 1 << 0
	// FreeTypeRasterizerFlagsNoAutoHint disables auto-hinter.
	FreeTypeRasterizerFlagsNoAutoHint = 1 << 1
	// FreeTypeRasterizerFlagsForceAutoHint indicates that the auto-hinter is preferred over the font's native hinter.
	FreeTypeRasterizerFlagsForceAutoHint = 1 << 2
	// FreeTypeRasterizerFlagsLightHinting is a lighter hinting algorithm for gray-level modes.
	// Many generated glyphs are fuzzier but better resemble their original shape.
	// This is achieved by snapping glyphs to the pixel grid only vertically (Y-axis),
	// as is done by Microsoft's ClearType and Adobe's proprietary font renderer.
	// This preserves inter-glyph spacing in horizontal text.
	FreeTypeRasterizerFlagsLightHinting = 1 << 3
	// FreeTypeRasterizerFlagsMonoHinting is a strong hinting algorithm that should only be used for monochrome output.
	FreeTypeRasterizerFlagsMonoHinting = 1 << 4
	// FreeTypeRasterizerFlagsBold is for styling: Should we artificially embolden the font?
	FreeTypeRasterizerFlagsBold = 1 << 5
	// FreeTypeRasterizerFlagsOblique is for styling: Should we slant the font, emulating italic style?
	FreeTypeRasterizerFlagsOblique = 1 << 6
	// FreeTypeRasterizerFlagsMonochrome disables anti-aliasing. Combine this with MonoHinting for best results!
	FreeTypeRasterizerFlagsMonochrome = 1 << 7
)

const (
	// HoveredFlagsNone Return true if directly over the item/window, not obstructed by another window,
	// not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsNone = 0
	// HoveredFlagsChildWindows IsWindowHovered() only: Return true if any children of the window is hovered.
	HoveredFlagsChildWindows = 1 << 0
	// HoveredFlagsRootWindow IsWindowHovered() only: Test from root window (top most parent of the current hierarchy).
	HoveredFlagsRootWindow = 1 << 1
	// HoveredFlagsAnyWindow IsWindowHovered() only: Return true if any window is hovered.
	HoveredFlagsAnyWindow = 1 << 2
	// HoveredFlagsAllowWhenBlockedByPopup Return true even if a popup window is normally blocking access to this item/window.
	HoveredFlagsAllowWhenBlockedByPopup = 1 << 3
	// HoveredFlagsAllowWhenBlockedByActiveItem Return true even if an active item is blocking access to this item/window.
	// Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenBlockedByActiveItem = 1 << 5
	// HoveredFlagsAllowWhenOverlapped Return true even if the position is overlapped by another window
	HoveredFlagsAllowWhenOverlapped = 1 << 6
	// HoveredFlagsAllowWhenDisabled Return true even if the item is disabled
	HoveredFlagsAllowWhenDisabled = 1 << 7

	HoveredFlagsRectOnly            = HoveredFlagsAllowWhenBlockedByPopup | HoveredFlagsAllowWhenBlockedByActiveItem | HoveredFlagsAllowWhenOverlapped
	HoveredFlagsRootAndChildWindows = HoveredFlagsRootWindow | HoveredFlagsChildWindows
)

const (
	// InputTextFlagsNone sets everything default.
	InputTextFlagsNone = 0
	// InputTextFlagsCharsDecimal allows 0123456789.+-
	InputTextFlagsCharsDecimal = 1 << 0
	// InputTextFlagsCharsHexadecimal allow 0123456789ABCDEFabcdef
	InputTextFlagsCharsHexadecimal = 1 << 1
	// InputTextFlagsCharsUppercase turns a..z into A..Z.
	InputTextFlagsCharsUppercase = 1 << 2
	// InputTextFlagsCharsNoBlank filters out spaces, tabs.
	InputTextFlagsCharsNoBlank = 1 << 3
	// InputTextFlagsAutoSelectAll selects entire text when first taking mouse focus.
	InputTextFlagsAutoSelectAll = 1 << 4
	// InputTextFlagsEnterReturnsTrue returns 'true' when Enter is pressed (as opposed to when the value was modified).
	InputTextFlagsEnterReturnsTrue = 1 << 5
	// InputTextFlagsCallbackCompletion for callback on pressing TAB (for completion handling).
	InputTextFlagsCallbackCompletion = 1 << 6
	// InputTextFlagsCallbackHistory for callback on pressing Up/Down arrows (for history handling).
	InputTextFlagsCallbackHistory = 1 << 7
	// InputTextFlagsCallbackAlways for callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackAlways = 1 << 8
	// InputTextFlagsCallbackCharFilter for callback on character inputs to replace or discard them.
	// Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsCallbackCharFilter = 1 << 9
	// InputTextFlagsAllowTabInput when pressing TAB to input a '\t' character into the text field.
	InputTextFlagsAllowTabInput = 1 << 10
	// InputTextFlagsCtrlEnterForNewLine in multi-line mode, unfocus with Enter, add new line with Ctrl+Enter
	// (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsCtrlEnterForNewLine = 1 << 11
	// InputTextFlagsNoHorizontalScroll disables following the cursor horizontally.
	InputTextFlagsNoHorizontalScroll = 1 << 12
	// InputTextFlagsAlwaysInsertMode sets insert mode.
	InputTextFlagsAlwaysInsertMode = 1 << 13
	// InputTextFlagsReadOnly sets read-only mode.
	InputTextFlagsReadOnly = 1 << 14
	// InputTextFlagsPassword sets password mode, display all characters as '*'.
	InputTextFlagsPassword = 1 << 15
	// InputTextFlagsNoUndoRedo disables undo/redo. Note that input text owns the text data while active,
	// if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsNoUndoRedo = 1 << 16
	// InputTextFlagsCharsScientific allows 0123456789.+-*/eE (Scientific notation input).
	InputTextFlagsCharsScientific = 1 << 17
	// inputTextFlagsCallbackResize for callback on buffer capacity change requests.
	inputTextFlagsCallbackResize = 1 << 18
)

// Enumeration for MouseCursor()
// User code may request binding to display given cursor by calling SetMouseCursor(), which is why we have some cursors that are marked unused here
const (
	// MouseCursorNone no mouse cursor
	MouseCursorNone = -1
	// MouseCursorArrow standard arrow mouse cursor
	MouseCursorArrow = 0
	// MouseCursorTextInput when hovering over InputText, etc.
	MouseCursorTextInput = 1
	// MouseCursorResizeAll (Unused by imgui functions)
	MouseCursorResizeAll = 2
	// MouseCursorResizeNS when hovering over an horizontal border
	MouseCursorResizeNS = 3
	// MouseCursorResizeEW when hovering over a vertical border or a column
	MouseCursorResizeEW = 4
	// MouseCursorResizeNESW when hovering over the bottom-left corner of a window
	MouseCursorResizeNESW = 5
	// MouseCursorResizeNWSE when hovering over the bottom-right corner of a window
	MouseCursorResizeNWSE = 6
	// MouseCursorHand (Unused by imgui functions. Use for e.g. hyperlinks)
	MouseCursorHand  = 7
	MouseCursorCount = 8
)

const (
	// SelectableFlagsNone default = 0
	SelectableFlagsNone = 0
	// SelectableFlagsDontClosePopups makes clicking the selectable not close any parent popup windows.
	SelectableFlagsDontClosePopups = 1 << 0
	// SelectableFlagsSpanAllColumns allows the selectable frame to span all columns (text will still fit in current column).
	SelectableFlagsSpanAllColumns = 1 << 1
	// SelectableFlagsAllowDoubleClick generates press events on double clicks too.
	SelectableFlagsAllowDoubleClick = 1 << 2
	// SelectableFlagsDisabled disallows selection and displays text in a greyed out color.
	SelectableFlagsDisabled = 1 << 3
)

const (
	// TabBarFlagsNone default = 0.
	TabBarFlagsNone = 0
	// TabBarFlagsReorderable Allow manually dragging tabs to re-order them + New tabs are appended at the end of list
	TabBarFlagsReorderable = 1 << 0
	// TabBarFlagsAutoSelectNewTabs Automatically select new tabs when they appear
	TabBarFlagsAutoSelectNewTabs = 1 << 1
	// TabBarFlagsTabListPopupButton Disable buttons to open the tab list popup
	TabBarFlagsTabListPopupButton = 1 << 2
	// TabBarFlagsNoCloseWithMiddleMouseButton Disable behavior of closing tabs (that are submitted with p_open != NULL)
	// with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabBarFlagsNoCloseWithMiddleMouseButton = 1 << 3
	// TabBarFlagsNoTabListScrollingButtons Disable scrolling buttons (apply when fitting policy is
	// TabBarFlagsFittingPolicyScroll)
	TabBarFlagsNoTabListScrollingButtons = 1 << 4
	// TabBarFlagsNoTooltip Disable tooltips when hovering a tab
	TabBarFlagsNoTooltip = 1 << 5
	// TabBarFlagsFittingPolicyResizeDown Resize tabs when they don't fit
	TabBarFlagsFittingPolicyResizeDown = 1 << 6
	// TabBarFlagsFittingPolicyScroll Add scroll buttons when tabs don't fit
	TabBarFlagsFittingPolicyScroll = 1 << 7
	// TabBarFlagsFittingPolicyMask combines
	// TabBarFlagsFittingPolicyResizeDown and TabBarFlagsFittingPolicyScroll
	TabBarFlagsFittingPolicyMask = TabBarFlagsFittingPolicyResizeDown | TabBarFlagsFittingPolicyScroll
	// TabBarFlagsFittingPolicyDefault alias for TabBarFlagsFittingPolicyResizeDown
	TabBarFlagsFittingPolicyDefault = TabBarFlagsFittingPolicyResizeDown
)

const (
	// TabItemFlagsNone default = 0
	TabItemFlagsNone = 0
	// TabItemFlagsUnsavedDocument Append '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. Also: tab is selected on closure and closure is deferred by one frame to allow code to undo it
	// without flicker.
	TabItemFlagsUnsavedDocument = 1 << 0
	// TabItemFlagsSetSelected Trigger flag to programmatically make the tab selected when calling BeginTabItem()
	TabItemFlagsSetSelected = 1 << 1
	// TabItemFlagsNoCloseWithMiddleMouseButton  Disable behavior of closing tabs (that are submitted with
	// p_open != NULL) with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabItemFlagsNoCloseWithMiddleMouseButton = 1 << 2
	// TabItemFlagsNoPushID Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem()
	TabItemFlagsNoPushID = 1 << 3
)

const (
	// TreeNodeFlagsNone default = 0
	TreeNodeFlagsNone = 0
	// TreeNodeFlagsSelected draws as selected.
	TreeNodeFlagsSelected = 1 << 0
	// TreeNodeFlagsFramed draws full colored frame (e.g. for CollapsingHeader).
	TreeNodeFlagsFramed = 1 << 1
	// TreeNodeFlagsAllowItemOverlap hit testing to allow subsequent widgets to overlap this one.
	TreeNodeFlagsAllowItemOverlap = 1 << 2
	// TreeNodeFlagsNoTreePushOnOpen doesn't do a TreePush() when open
	// (e.g. for CollapsingHeader) = no extra indent nor pushing on ID stack.
	TreeNodeFlagsNoTreePushOnOpen = 1 << 3
	// TreeNodeFlagsNoAutoOpenOnLog doesn't automatically and temporarily open node when Logging is active
	// (by default logging will automatically open tree nodes).
	TreeNodeFlagsNoAutoOpenOnLog = 1 << 4
	// TreeNodeFlagsDefaultOpen defaults node to be open.
	TreeNodeFlagsDefaultOpen = 1 << 5
	// TreeNodeFlagsOpenOnDoubleClick needs double-click to open node.
	TreeNodeFlagsOpenOnDoubleClick = 1 << 6
	// TreeNodeFlagsOpenOnArrow opens only when clicking on the arrow part.
	// If TreeNodeFlagsOpenOnDoubleClick is also set, single-click arrow or double-click all box to open.
	TreeNodeFlagsOpenOnArrow = 1 << 7
	// TreeNodeFlagsLeaf allows no collapsing, no arrow (use as a convenience for leaf nodes).
	TreeNodeFlagsLeaf = 1 << 8
	// TreeNodeFlagsBullet displays a bullet instead of an arrow.
	TreeNodeFlagsBullet = 1 << 9
	// TreeNodeFlagsFramePadding uses FramePadding (even for an unframed text node) to
	// vertically align text baseline to regular widget height. Equivalent to calling AlignTextToFramePadding().
	TreeNodeFlagsFramePadding = 1 << 10
	// TreeNodeFlagsSpanAvailWidth extends hit box to the right-most edge, even if not framed.
	// This is not the default in order to allow adding other items on the same line.
	// In the future we may refactor the hit system to be front-to-back, allowing natural overlaps
	// and then this can become the default.
	TreeNodeFlagsSpanAvailWidth = 1 << 11
	// TreeNodeFlagsSpanFullWidth extends hit box to the left-most and right-most edges (bypass the indented area).
	TreeNodeFlagsSpanFullWidth = 1 << 12
	// TreeNodeFlagsNavLeftJumpsBackHere (WIP) Nav: left direction may move to this TreeNode() from any of its child
	// (items submitted between TreeNode and TreePop)
	TreeNodeFlagsNavLeftJumpsBackHere = 1 << 13
	// TreeNodeFlagsCollapsingHeader combines TreeNodeFlagsFramed and TreeNodeFlagsNoAutoOpenOnLog.
	TreeNodeFlagsCollapsingHeader = TreeNodeFlagsFramed | TreeNodeFlagsNoTreePushOnOpen | TreeNodeFlagsNoAutoOpenOnLog
)

const (
	// WindowFlagsNone default = 0
	WindowFlagsNone = 0
	// WindowFlagsNoTitleBar disables title-bar.
	WindowFlagsNoTitleBar = 1 << 0
	// WindowFlagsNoResize disables user resizing with the lower-right grip.
	WindowFlagsNoResize = 1 << 1
	// WindowFlagsNoMove disables user moving the window.
	WindowFlagsNoMove = 1 << 2
	// WindowFlagsNoScrollbar disables scrollbars. Window can still scroll with mouse or programmatically.
	WindowFlagsNoScrollbar = 1 << 3
	// WindowFlagsNoScrollWithMouse disables user vertically scrolling with mouse wheel. On child window, mouse wheel
	// will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoScrollWithMouse = 1 << 4
	// WindowFlagsNoCollapse disables user collapsing window by double-clicking on it.
	WindowFlagsNoCollapse = 1 << 5
	// WindowFlagsAlwaysAutoResize resizes every window to its content every frame.
	WindowFlagsAlwaysAutoResize = 1 << 6
	// WindowFlagsNoBackground disables drawing background color (WindowBg, etc.) and outside border. Similar as using
	// SetNextWindowBgAlpha(0.0f).
	WindowFlagsNoBackground = 1 << 7
	// WindowFlagsNoSavedSettings will never load/save settings in .ini file.
	WindowFlagsNoSavedSettings = 1 << 8
	// WindowFlagsNoMouseInputs disables catching mouse, hovering test with pass through.
	WindowFlagsNoMouseInputs = 1 << 9
	// WindowFlagsMenuBar has a menu-bar.
	WindowFlagsMenuBar = 1 << 10
	// WindowFlagsHorizontalScrollbar allows horizontal scrollbar to appear (off by default). You may use
	// SetNextWindowContentSize(ImVec2(width,0.0f)); prior to calling Begin() to specify width. Read code in imgui_demo
	// in the "Horizontal Scrolling" section.
	WindowFlagsHorizontalScrollbar = 1 << 11
	// WindowFlagsNoFocusOnAppearing disables taking focus when transitioning from hidden to visible state.
	WindowFlagsNoFocusOnAppearing = 1 << 12
	// WindowFlagsNoBringToFrontOnFocus disables bringing window to front when taking focus. e.g. clicking on it or
	// programmatically giving it focus.
	WindowFlagsNoBringToFrontOnFocus = 1 << 13
	// WindowFlagsAlwaysVerticalScrollbar always shows vertical scrollbar, even if ContentSize.y < Size.y .
	WindowFlagsAlwaysVerticalScrollbar = 1 << 14
	// WindowFlagsAlwaysHorizontalScrollbar always shows horizontal scrollbar, even if ContentSize.x < Size.x .
	WindowFlagsAlwaysHorizontalScrollbar = 1 << 15
	// WindowFlagsAlwaysUseWindowPadding ensures child windows without border uses style.WindowPadding (ignored by
	// default for non-bordered child windows, because more convenient).
	WindowFlagsAlwaysUseWindowPadding = 1 << 16
	// WindowFlagsNoNavInputs has no gamepad/keyboard navigation within the window.
	WindowFlagsNoNavInputs = 1 << 18
	// WindowFlagsNoNavFocus has no focusing toward this window with gamepad/keyboard navigation
	// (e.g. skipped by CTRL+TAB)
	WindowFlagsNoNavFocus = 1 << 19
	// WindowFlagsUnsavedDocument appends '*' to title without affecting the ID, as a convenience to avoid using the
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
