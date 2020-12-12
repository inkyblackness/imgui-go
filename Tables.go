package imgui

// #include "wrapper/Tables.h"
import "C"

// Tables
// [BETA API] API may evolve!
// - Full-featured replacement for old Columns API.
// - See Demo->Tables for details.
// - See TableFlags and TableColumnFlags enums for a description of available flags.
// The typical call flow is:
// - 1. Call BeginTable()
// - 2. Optionally call TableSetupColumn() to submit column name/flags/defaults
// - 3. Optionally call TableSetupScrollFreeze() to request scroll freezing of columns/rows
// - 4. Optionally call TableHeadersRow() to submit a header row (names will be pulled from data submitted to TableSetupColumns)
// - 5. Populate contents
//    - In most situations you can use TableNextRow() + TableSetColumnIndex(N) to start appending into a column.
//    - If you are using tables as a sort of grid, where every columns is holding the same type of contents,
//      you may prefer using TableNextColumn() instead of TableNextRow() + TableSetColumnIndex().
//      TableNextColumn() will automatically wrap-around into the next row if needed.
//    - IMPORTANT: Comparatively to the old Columns() API, we need to call TableNextColumn() for the first column!
//    - Both TableSetColumnIndex() and TableNextColumn() return true when the column is visible or performing
//      width measurements. Otherwise, you may skip submitting the contents of a cell/column, BUT ONLY if you know
//      it is not going to contribute to row height.
//      In many situations, you may skip submitting contents for every columns but one (e.g. the first one).
//    - Summary of possible call flow:
//      ----------------------------------------------------------------------------------------------------------
//       TableNextRow() -> TableSetColumnIndex(0) -> Text("Hello 0") -> TableSetColumnIndex(1) -> Text("Hello 1")  // OK
//       TableNextRow() -> TableNextColumn()      -> Text("Hello 0") -> TableNextColumn()      -> Text("Hello 1")  // OK
//                         TableNextColumn()      -> Text("Hello 0") -> TableNextColumn()      -> Text("Hello 1")  // OK: TableNextColumn() automatically gets to next row!
//       TableNextRow()                           -> Text("Hello 0")                                               // Not OK! Missing TableSetColumnIndex() or TableNextColumn()! Text will not appear!
//      ----------------------------------------------------------------------------------------------------------
// - 5. Call EndTable()

// Flags for BeginTable()
// - Important! Sizing policies have particularly complex and subtle side effects, more so than you would expect.
//   Read comments/demos carefully + experiment with live demos to get acquainted with them.
// - The default sizing policy for columns depends on whether the ScrollX flag is set on the table:
//   When ScrollX is off:
//    - Table defaults to TableFlagsColumnsWidthStretch -> all Columns defaults to TableColumnFlagsWidthStretch.
//    - Columns sizing policy allowed: Stretch (default) or Fixed/Auto.
//    - Stretch Columns will share the width available in table.
//    - Fixed Columns will generally obtain their requested width unless the Table cannot fit them all.
//   When ScrollX is on:
//    - Table defaults to TableFlagsColumnsWidthFixed -> all Columns defaults to TableColumnFlagsWidthFixed.
//    - Columns sizing policy allowed: Fixed/Auto mostly!
//    - Fixed Columns can be enlarged as needed. Table will show an horizontal scrollbar if needed.
//    - Using Stretch columns OFTEN DOES NOT MAKE SENSE if ScrollX is on, UNLESS you have specified a value for 'inner_width' in BeginTable().
//    - Stretch Columns, if any, will calculate their width using inner_width, assuming no scrolling (it really doesn't make sense to do otherwise).
// - Mixing up columns with different sizing policy is possible BUT can be tricky and has some side-effects and restrictions.
//   (their visible order and the scrolling state have subtle but necessary effects on how they can be manually resized).
//   The typical use of mixing sizing policies is to have ScrollX disabled, one or two Stretch Column and many Fixed Columns.
const (
	// Features
	TableFlagsNone              = 0
	TableFlagsResizable         = 1 << 0 // Allow resizing columns.
	TableFlagsReorderable       = 1 << 1 // Allow reordering columns in header row (need calling TableSetupColumn() + TableHeadersRow() to display headers)
	TableFlagsHideable          = 1 << 2 // Allow hiding/disabling columns in context menu.
	TableFlagsSortable          = 1 << 3 // Allow sorting on one column (sort_specs_count will always be == 1). Call TableGetSortSpecs() to obtain sort specs.
	TableFlagsNoSavedSettings   = 1 << 4 // Disable persisting columns order, width and sort settings in the .ini file.
	TableFlagsContextMenuInBody = 1 << 5 // Right-click on columns body/contents will display table context menu. By default it is available in TableHeadersRow().
	// Decorations
	TableFlagsRowBg                      = 1 << 6                                            // Set each RowBg color with StyleColorTableRowBg or StyleColorTableRowBgAlt (equivalent to calling TableSetBgColor with TableBgFlagsRowBg0 on each row manually)
	TableFlagsBordersInnerH              = 1 << 7                                            // Draw horizontal borders between rows.
	TableFlagsBordersOuterH              = 1 << 8                                            // Draw horizontal borders at the top and bottom.
	TableFlagsBordersInnerV              = 1 << 9                                            // Draw vertical borders between columns.
	TableFlagsBordersOuterV              = 1 << 10                                           // Draw vertical borders on the left and right sides.
	TableFlagsBordersH                   = TableFlagsBordersInnerH | TableFlagsBordersOuterH // Draw horizontal borders.
	TableFlagsBordersV                   = TableFlagsBordersInnerV | TableFlagsBordersOuterV // Draw vertical borders.
	TableFlagsBordersInner               = TableFlagsBordersInnerV | TableFlagsBordersInnerH // Draw inner borders.
	TableFlagsBordersOuter               = TableFlagsBordersOuterV | TableFlagsBordersOuterH // Draw outer borders.
	TableFlagsBorders                    = TableFlagsBordersInner | TableFlagsBordersOuter   // Draw all borders.
	TableFlagsNoBordersInBody            = 1 << 11                                           // Disable vertical borders in columns Body (borders will always appears in Headers).
	TableFlagsNoBordersInBodyUntilResize = 1 << 12                                           // Disable vertical borders in columns Body until hovered for resize (borders will always appears in Headers).
	// Sizing
	TableFlagsColumnsWidthStretch  = 1 << 13 // Default if ScrollX is off. Columns will default to use _WidthStretch. Read description above for more details.
	TableFlagsColumnsWidthFixed    = 1 << 14 // Default if ScrollX is on. Columns will default to use _WidthFixed or _WidthAutoResize policy (if Resizable is off). Read description above for more details.
	TableFlagsSameWidths           = 1 << 15 // Make all columns the same widths which is useful with Fixed columns policy (but granted by default with Stretch policy + no resize). Implicitly enable TableFlagsNoKeepColumnsVisible and disable TableFlagsResizable.
	TableFlagsNoHeadersWidth       = 1 << 16 // Disable headers' contribution to automatic width calculation.
	TableFlagsNoHostExtendY        = 1 << 17 // Disable extending past the limit set by outer_size.y, only meaningful when neither of ScrollX|ScrollY are set (data below the limit will be clipped and not visible)
	TableFlagsNoKeepColumnsVisible = 1 << 18 // Disable keeping column always minimally visible when ScrollX is off and table gets too small.
	TableFlagsPreciseWidths        = 1 << 19 // Disable distributing remainder width to stretched columns (width allocation on a 100-wide table with 3 columns: Without this flag: 33,33,34. With this flag: 33,33,33). With larger number of columns, resizing will appear to be less smooth.
	TableFlagsNoClip               = 1 << 20 // Disable clipping rectangle for every individual columns (reduce draw command count, items will be able to overflow into other columns). Generally incompatible with TableSetupScrollFreeze().
	// Padding
	TableFlagsPadOuterX   = 1 << 21 // Default if BordersOuterV is on. Enable outer-most padding.
	TableFlagsNoPadOuterX = 1 << 22 // Default if BordersOuterV is off. Disable outer-most padding.
	TableFlagsNoPadInnerX = 1 << 23 // Disable inner padding between columns (double inner padding if BordersOuterV is on, single inner padding if BordersOuterV is off).
	// Scrolling
	TableFlagsScrollX = 1 << 24 // Enable horizontal scrolling. Require 'outer_size' parameter of BeginTable() to specify the container size. Changes default sizing policy. Because this create a child window, ScrollY is currently generally recommended when using ScrollX.
	TableFlagsScrollY = 1 << 25 // Enable vertical scrolling. Require 'outer_size' parameter of BeginTable() to specify the container size.
	// Sorting
	TableFlagsSortMulti    = 1 << 26 // Hold shift when clicking headers to sort on multiple column. TableGetSortSpecs() may return specs where (SpecsCount > 1).
	TableFlagsSortTristate = 1 << 27 // Allow no sorting, disable default sorting. TableGetSortSpecs() may return specs where (SpecsCount == 0).
)

// Flags for TableSetupColumn()
const (
	// Input configuration flags
	TableColumnFlagsNone                 = 0
	TableColumnFlagsDefaultHide          = 1 << 0  // Default as a hidden/disabled column.
	TableColumnFlagsDefaultSort          = 1 << 1  // Default as a sorting column.
	TableColumnFlagsWidthStretch         = 1 << 2  // Column will stretch. Preferable with horizontal scrolling disabled (default if table sizing policy is _ColumnsWidthStretch).
	TableColumnFlagsWidthFixed           = 1 << 3  // Column will not stretch. Preferable with horizontal scrolling enabled (default if table sizing policy is _ColumnsWidthFixed and table is resizable).
	TableColumnFlagsWidthAutoResize      = 1 << 4  // Column will not stretch and keep resizing based on submitted contents (default if table sizing policy is _ColumnsWidthFixed and table is not resizable).
	TableColumnFlagsNoResize             = 1 << 5  // Disable manual resizing.
	TableColumnFlagsNoReorder            = 1 << 6  // Disable manual reordering this column, this will also prevent other columns from crossing over this column.
	TableColumnFlagsNoHide               = 1 << 7  // Disable ability to hide/disable this column.
	TableColumnFlagsNoClip               = 1 << 8  // Disable clipping for this column (all NoClip columns will render in a same draw command).
	TableColumnFlagsNoSort               = 1 << 9  // Disable ability to sort on this field (even if TableFlagsSortable is set on the table).
	TableColumnFlagsNoSortAscending      = 1 << 10 // Disable ability to sort in the ascending direction.
	TableColumnFlagsNoSortDescending     = 1 << 11 // Disable ability to sort in the descending direction.
	TableColumnFlagsNoHeaderWidth        = 1 << 12 // Disable header text width contribution to automatic column width.
	TableColumnFlagsPreferSortAscending  = 1 << 13 // Make the initial sort direction Ascending when first sorting on this column (default).
	TableColumnFlagsPreferSortDescending = 1 << 14 // Make the initial sort direction Descending when first sorting on this column.
	TableColumnFlagsIndentEnable         = 1 << 15 // Use current Indent value when entering cell (default for column 0).
	TableColumnFlagsIndentDisable        = 1 << 16 // Ignore current Indent value when entering cell (default for columns > 0). Indentation changes _within_ the cell will still be honored.

	// Output status flags, read-only via TableGetColumnFlags()
	TableColumnFlagsIsEnabled = 1 << 20 // Status: is enabled == not hidden by user/api (referred to as "Hide" in _DefaultHide and _NoHide) flags.
	TableColumnFlagsIsVisible = 1 << 21 // Status: is visible == is enabled AND not clipped by scrolling.
	TableColumnFlagsIsSorted  = 1 << 22 // Status: is currently part of the sort specs
	TableColumnFlagsIsHovered = 1 << 23 // Status: is hovered by mouse
)

// Flags for TableNextRow()
const (
	TableRowFlagsNone    = 0
	TableRowFlagsHeaders = 1 << 0 // Identify header row (set default background color + width of its contents accounted different for auto column width)
)

// Enum for TableSetBgColor()
// Background colors are rendering in 3 layers:
//  - Layer 0: draw with RowBg0 color if set, otherwise draw with ColumnBg0 if set.
//  - Layer 1: draw with RowBg1 color if set, otherwise draw with ColumnBg1 if set.
//  - Layer 2: draw with CellBg color if set.
// The purpose of the two row/columns layers is to let you decide if a background color changes should override or blend with the existing color.
// When using TableFlagsRowBg on the table, each row has the RowBg0 color automatically set for odd/even rows.
// If you set the color of RowBg0 target, your color will override the existing RowBg0 color.
// If you set the color of RowBg1 or ColumnBg1 target, your color will blend over the RowBg0 color.
const (
	TableBgTargetNone   = 0
	TableBgTargetRowBg0 = 1 // Set row background color 0 (generally used for background, automatically set when TableFlagsRowBg is used)
	TableBgTargetRowBg1 = 2 // Set row background color 1 (generally used for selection marking)
	TableBgTargetCellBg = 3 // Set cell background color (top-most color)
)

// A sorting direction
const (
	SortDirectionNone       = 0
	SortDirectionAscending  = 1 // Ascending = 0->9, A->Z etc.
	SortDirectionDescending = 2 // Descending = 9->0, Z->A etc.
)

// (Read carefully because this is subtle but it does make sense!)
// About 'outerSize', its meaning needs to differ slightly depending of if we are using ScrollX/ScrollY flags:
//   X:
//   - outerSize.x < 0.0f  ->  right align from window/work-rect maximum x edge.
//   - outerSize.x = 0.0f  ->  auto enlarge, use all available space.
//   - outerSize.x > 0.0f  ->  fixed width
//   Y with ScrollX/ScrollY: using a child window for scrolling:
//   - outerSize.y < 0.0f  ->  bottom align
//   - outerSize.y = 0.0f  ->  bottom align, consistent with BeginChild(). not recommended unless table is last item in parent window.
//   - outerSize.y > 0.0f  ->  fixed child height. recommended when using Scrolling on any axis.
//   Y without scrolling, we output table directly in parent window:
//   - outerSize.y < 0.0f  ->  bottom align (will auto extend, unless NoHostExtendV is set)
//   - outerSize.y = 0.0f  ->  zero minimum height (will auto extend, unless NoHostExtendV is set)
//   - outerSize.y > 0.0f  ->  minimum height (will auto extend, unless NoHostExtendV is set)
// About 'innerWidth':
//   With ScrollX:
//   - innerWidth  < 0.0f  ->  *illegal* fit in known width (right align from outerSize.x) <-- weird
//   - innerWidth  = 0.0f  ->  fit in outer_width: Fixed size columns will take space they need (if avail, otherwise shrink down), Stretch columns becomes Fixed columns.
//   - innerWidth  > 0.0f  ->  override scrolling width, generally to be larger than outerSize.x. Fixed column take space they need (if avail, otherwise shrink down), Stretch columns share remaining space!
//   Without ScrollX:
//   - innerWidth          ->  *ignored*
// Details:
// - If you want to use Stretch columns with ScrollX, you generally need to specify 'innerWidth' otherwise the concept
//   of "available space" doesn't make sense.
// - Even if not really useful, we allow 'innerWidth < outerSize.x' for consistency and to facilitate understanding
//   of what the value does.
func BeginTableV(id string, columnsCount int, flags int, outerSize Vec2, innerWidth float32) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	outerSizeArg, _ := outerSize.wrapped()
	return C.iggBeginTable(idArg, C.int(columnsCount), C.int(flags), outerSizeArg, C.float(innerWidth)) != 0
}

// BeginTable calls BeginTableV(id, columnsCount, 0, imgui.Vec2{}, 0.0).
func BeginTable(id string, columnsCount int) bool {
	return BeginTableV(id, columnsCount, 0, Vec2{}, 0.0)
}

// EndTable closes the scope for the previously opened table.
// only call EndTable() if BeginTable() returns true!
func EndTable() {
	C.iggEndTable()
}

// TableNextRowV appends into the first cell of a new row
func TableNextRowV(flags int, minRowHeight float32) {
	C.iggTableNextRow(C.int(flags), C.float(minRowHeight))
}

// TableNextRow calls TableNextRowV(0, 0.0)
func TableNextRow() {
	TableNextRowV(0, 0.0)
}

// TableNextColumn appends into the next column (or first column of next row if currently in last column)
// Return true when column is visible.
func TableNextColumn() bool {
	return C.iggTableNextColumn() != 0
}

// TableSetColumnIndex appends into the specified column. Return true when column is visible.
func TableSetColumnIndex(columnN int) bool {
	return C.iggTableSetColumnIndex(C.int(columnN)) != 0
}

// TableGetColumnIndex return current column index.
func TableGetColumnIndex() int {
	return int(C.iggTableGetColumnIndex())
}

// TableGetRowIndex return current row index.
func TableGetRowIndex() int {
	return int(C.iggTableGetRowIndex())
}

// Tables: Headers & Columns declaration
// - Use TableSetupColumn() to specify label, resizing policy, default width/weight, id, various other flags etc.
//   Important: this will not display anything! The name passed to TableSetupColumn() is used by TableHeadersRow() and context-menus.
// - Use TableHeadersRow() to create a row and automatically submit a TableHeader() for each column.
//   Headers are required to perform: reordering, sorting, and opening the context menu (but context menu can also be available in columns body using TableFlagsContextMenuInBody).
// - You may manually submit headers using TableNextRow() + TableHeader() calls, but this is only useful in some advanced cases (e.g. adding custom widgets in header row).
// - Use TableSetupScrollFreeze() to lock columns (from the right) or rows (from the top) so they stay visible when scrolled.
func TableSetupColumnV(label string, flags int, initWidthOrHeight float32, userID uint) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	C.iggTableSetupColumn(labelArg, C.int(flags), C.float(initWidthOrHeight), C.uint(userID))
}

// TableSetupColumn calls TableSetupColumnV(label, 0, -1.0, 0)
func TableSetupColumn(label string) {
	TableSetupColumnV(label, 0, -1.0, 0)
}

// TableSetupScrollFreeze locks columns/rows so they stay visible when scrolled
func TableSetupScrollFreeze(cols int, rows int) {
	C.iggTableSetupScrollFreeze(C.int(cols), C.int(rows))
}

// TableHeadersRow submits all headers cells based on data provided to TableSetupColumn() + submit context menu
func TableHeadersRow() {
	C.iggTableHeadersRow()
}

// TableHeader submits one header cell manually (rarely used)
func TableHeader(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	C.iggTableHeader(labelArg)
}

// Tables: Miscellaneous functions
// - Most functions taking 'int column_n' treat the default value of -1 as the same as passing the current column index

// TableGetColumnCount returns number of columns (value passed to BeginTable)
func TableGetColumnCount() int {
	return int(C.iggTableGetColumnCount())
}

// TableGetColumnNameV returns "" if column didn't have a name declared by TableSetupColumn(). Pass -1 to use current column.
func TableGetColumnNameV(columnN int) string {
	return C.GoString(C.iggTableGetColumnName(C.int(columnN)))
}

// TableGetColumnName calls TableGetColumnNameV(-1)
func TableGetColumnName() string {
	return TableGetColumnNameV(-1)
}

// TableGetColumnFlags return column flags so you can query their Enabled/Visible/Sorted/Hovered status flags.
func TableGetColumnFlagsV(columnN int) int {
	return int(C.iggTableGetColumnFlags(C.int(columnN)))
}

// TableGetColumnFlags calls TableGetColumnFlagsV(-1)
func TableGetColumnFlags() int {
	return TableGetColumnFlagsV(-1)
}

// TableSetBgColorV changes the color of a cell, row, or column. See TableBgTarget flags for details.
func TableSetBgColorV(bgTarget int, color Vec4, columnN int) {
	colorArg, _ := color.wrapped()
	C.iggTableSetBgColor(C.int(bgTarget), colorArg, C.int(columnN))
}

// TableSetBgColor calls TableSetBgColorV(bgTarget, color, -1)
func TableSetBgColor(bgTarget int, color Vec4) {
	TableSetBgColorV(bgTarget, color, -1)
}

// Tables: Sorting
//   Call TableGetSortSpecs() to retrieve latest sort specs for the table. Return value will be NULL if no sorting.
//   When 'SpecsDirty == true' you should sort your data. It will be true when sorting specs have changed since last call, or the first time.
//   Make sure to set 'SpecsDirty = false' after sorting, else you may wastefully sort your data every frame!
//   Lifetime: don't hold on this pointer over multiple frames or past any subsequent call to BeginTable().

// Sorting specification for one column of a table (sizeof == 12 bytes)
type TableColumnSortSpecs struct {
	ColumnUserID  uint  // User id of the column (if specified by a TableSetupColumn() call)
	ColumnIndex   int16 // Index of the column
	SortOrder     int16 // Index within parent TableSortSpecs (always stored in order starting from 0, tables sorted on a single criteria will always have a 0 here)
	SortDirection int   // SortDirectionAscending or SortDirectionDescending (you can use this or SortSign, whichever is more convenient for your sort function)
}

// Sorting specifications for a table (often handling sort specs for a single column, occasionally more)
// Obtained by calling TableGetSortSpecs()
// When 'SpecsDirty == true' you can sort your data. It will be true with sorting specs have changed since last call, or the first time.
// Make sure to set 'SpecsDirty = false' after sorting, else you may wastefully sort your data every frame!
type TableSortSpecs uintptr

// TableGetSortSpecs gets latest sort specs for the table (nil if not sorting).
func TableGetSortSpecs() TableSortSpecs {
	return TableSortSpecs(C.iggTableGetSortSpecs())
}

func (specs TableSortSpecs) handle() C.IggTableSortSpecs {
	return C.IggTableSortSpecs(specs)
}

// Specs returns columns sort spec array
func (specs TableSortSpecs) Specs() []TableColumnSortSpecs {
	count := specs.SpecsCount()
	column_specs := make([]TableColumnSortSpecs, count)
	for i := 0; i < count; i++ {
		out := &C.IggTableColumnSortSpecs{}
		C.iggTableSortSpecsGetSpec(specs.handle(), C.int(i), out)

		column_specs[i] = TableColumnSortSpecs{
			ColumnUserID:  uint(out.ColumnUserID),
			ColumnIndex:   int16(out.ColumnIndex),
			SortOrder:     int16(out.SortOrder),
			SortDirection: int(out.SortDirection),
		}
	}
	return column_specs
}

// SpecsCount returns sort spec count. Most often 1 unless e.g. TableFlagsMultiSortable is enabled.
func (specs TableSortSpecs) SpecsCount() int {
	if specs == 0 {
		return 0
	}
	return int(C.iggTableSortSpecsGetSpecsCount(specs.handle()))
}

// SpecsDirty returns if specs have changed since last time! Use this to sort again
func (specs TableSortSpecs) SpecsDirty() bool {
	if specs == 0 {
		return false
	}
	return C.iggTableSortSpecsGetSpecsDirty(specs.handle()) != 0
}

// SetSpecsDirty sets SpecsDirty value to a given value, usually used to clear the flag (set to false)
func (specs TableSortSpecs) SetSpecsDirty(value bool) {
	if specs == 0 {
		return
	}
	C.iggTableSortSpecsSetSpecsDirty(specs.handle(), castBool(value))
}

// ClearSpecsDirty calls specs.SetSpecsDirty(false)
func (specs TableSortSpecs) ClearSpecsDirty() {
	specs.SetSpecsDirty(false)
}
