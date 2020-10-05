package imgui

// #include "wrapper/Tables.h"
import "C"

// Tables
// [ALPHA API] API may evolve!
// - Full-featured replacement for old Columns API
// - See Demo->Tables for details.
// - See TableFlags and TableColumnsFlags enums for a description of available flags.
// The typical call flow is:
// - 1. Call BeginTable()
// - 2. Optionally call TableSetupColumn() to submit column name/flags/defaults
// - 3. Optionally call TableSetupScrollFreeze() to request scroll freezing of columns/rows
// - 4. Optionally call TableHeadersRow() to submit a header row (names will be pulled from data submitted to TableSetupColumns)
// - 4. Populate contents
//     - In most situations you can use TableNextRow() + TableSetColumnIndex(xx) to start appending into a column.
//     - If you are using tables as a sort of grid, where every columns is holding the same type of contents,
//       you may prefer using TableNextColumn() instead of TableNextRow() + TableSetColumnIndex().
//       TableNextColumn() will automatically wrap-around into the next row if needed.
//     - IMPORTANT: Comparatively to the old Columns() API, we need to call TableNextColumn() for the first column!
//     - Here's a summary of possible call flow:
//       ----------------------------------------------------------------------------------------------------------------
//         TableNextRow() -> TableSetColumnIndex(0) -> Button("Hello 0") -> TableSetColumnIndex(1) -> Button("Hello 1")   // OK
//         TableNextRow() -> TableNextColumn()         Button("Hello 0") -> TableNextColumn()      -> Button("Hello 1")   // OK
//                           TableNextColumn()         Button("Hello 0") -> TableNextColumn()      -> Button("Hello 1")   // OK: TableNextColumn() automatically gets to next row!
//         TableNextRow()                              Button("Hello 0") ................................................ // Not OK! Missing TableSetColumnIndex() or TableNextColumn()!
//       ----------------------------------------------------------------------------------------------------------------
// - 5. Call EndTable()

// Flags for ImGui::BeginTable()
// - Important! Sizing policies have particularly complex and subtle side effects, more so than you would expect.
//   Read comments/demos carefully + experiment with live demos to get acquainted with them.
// - The default sizing policy for columns depends on whether the ScrollX flag is set on the table:
//   When ScrollX is off:
//    - Table defaults to ImGuiTableFlags_SizingPolicyStretchX -> all Columns defaults to ImGuiTableColumnFlags_WidthStretch.
//    - Columns sizing policy allowed: Fixed/Auto or Stretch.
//    - Stretch Columns will share the width available in table.
//    - Fixed Columns will generally obtain their requested width unless the Table cannot fit them all.
//   When ScrollX is on:
//    - Table defaults to ImGuiTableFlags_SizingPolicyFixedX -> all Columns defaults to ImGuiTableColumnFlags_WidthFixed.
//    - Columns sizing policy allowed: Fixed/Auto mostly! Using Stretch columns OFTEN DOES NOT MAKE SENSE if ScrollX is on, UNLESS you have specified a value for 'inner_width' in BeginTable().
//    - Fixed Columns can be enlarged as needed. Table will show an horizontal scrollbar if needed.
//    - Stretch Columns, if any, will calculate their width using inner_width, assuming no scrolling (it really doesn't make sense to do otherwise).
// - Mixing up columns with different sizing policy is possible BUT can be tricky and has some side-effects and restrictions.
//   (their visible order and the scrolling state have subtle but necessary effects on how they can be manually resized).
//   The typical use of mixing sizing policies is to have ScrollX disabled, one or two Stretch Column and many Fixed Columns.

const (
	// Features
	TableFlagsNone              = 0
	TableFlagsResizable         = 1 << 0 // Allow resizing columns.
	TableFlagsReorderable       = 1 << 1 // Allow reordering columns (need calling TableSetupColumn() + TableHeadersRow() to display headers)
	TableFlagsHideable          = 1 << 2 // Allow hiding columns (with right-click on header) (FIXME-TABLE: allow without headers).
	TableFlagsSortable          = 1 << 3 // Allow sorting on one column (sort_specs_count will always be == 1). Call TableGetSortSpecs() to obtain sort specs.
	TableFlagsMultiSortable     = 1 << 4 // Allow sorting on multiple columns by holding Shift (sort_specs_count may be > 1). Call TableGetSortSpecs() to obtain sort specs.
	TableFlagsNoSavedSettings   = 1 << 5 // Disable persisting columns order, width and sort settings in the .ini file.
	TableFlagsContextMenuInBody = 1 << 6 // Right-click on columns body/contents will display table context menu. By default it is available in TableHeadersRow().
	// Decoration
	TableFlagsRowBg              = 1 << 7                                            // Set each RowBg color with StyleColorTableRowBg or StyleColorTableRowBgAlt (equivalent to calling TableSetBgColor with TableBgFlagsRowBg0 on each row manually)
	TableFlagsBordersInnerH      = 1 << 8                                            // Draw horizontal borders between rows.
	TableFlagsBordersOuterH      = 1 << 9                                            // Draw horizontal borders at the top and bottom.
	TableFlagsBordersInnerV      = 1 << 10                                           // Draw vertical borders between columns.
	TableFlagsBordersOuterV      = 1 << 11                                           // Draw vertical borders on the left and right sides.
	TableFlagsBordersH           = TableFlagsBordersInnerH | TableFlagsBordersOuterH // Draw horizontal borders.
	TableFlagsBordersV           = TableFlagsBordersInnerV | TableFlagsBordersOuterV // Draw vertical borders.
	TableFlagsBordersInner       = TableFlagsBordersInnerV | TableFlagsBordersInnerH // Draw inner borders.
	TableFlagsBordersOuter       = TableFlagsBordersOuterV | TableFlagsBordersOuterH // Draw outer borders.
	TableFlagsBorders            = TableFlagsBordersInner | TableFlagsBordersOuter   // Draw all borders.
	TableFlagsBordersFullHeightV = 1 << 12                                           // Borders covers all rows even when Headers are being used. Allow resizing from any rows.
	// Padding, Sizing
	TableFlagsSizingPolicyFixedX   = 1 << 13 // Default if ScrollX is on. Columns will default to use _WidthFixed or _WidthAlwaysAutoResize policy. Read description above for more details.
	TableFlagsSizingPolicyStretchX = 1 << 14 // Default if ScrollX is off. Columns will default to use _WidthStretch policy. Read description above for more details.
	TableFlagsNoHeadersWidth       = 1 << 15 // Disable header width contribution to automatic width calculation.
	TableFlagsNoHostExtendY        = 1 << 16 // (FIXME-TABLE: Reword as SizingPolicy?) Disable extending past the limit set by outer_size.y, only meaningful when neither of ScrollX|ScrollY are set (data below the limit will be clipped and not visible)
	TableFlagsNoKeepColumnsVisible = 1 << 17 // (FIXME-TABLE) Disable code that keeps column always minimally visible when table width gets too small and horizontal scrolling is off.
	TableFlagsNoClip               = 1 << 18 // Disable clipping rectangle for every individual columns (reduce draw command count, items will be able to overflow into other columns). Generally incompatible with ScrollFreeze options.
	// Scrolling
	TableFlagsScrollX = 1 << 19 // Enable horizontal scrolling. Require 'outer_size' parameter of BeginTable() to specify the container size. Because this create a child window, ScrollY is currently generally recommended when using ScrollX.
	TableFlagsScrollY = 1 << 20 // Enable vertical scrolling. Require 'outer_size' parameter of BeginTable() to specify the container size.
	TableFlagsScroll  = TableFlagsScrollX | TableFlagsScrollY
)

// Flags for ImGui::TableSetupColumn()
// FIXME-TABLE: Rename to ImGuiColumns_*, stick old columns api flags in there under an obsolete api block
const (
	TableColumnFlagsNone                  = 0
	TableColumnFlagsDefaultHide           = 1 << 0  // Default as a hidden column.
	TableColumnFlagsDefaultSort           = 1 << 1  // Default as a sorting column.
	TableColumnFlagsWidthFixed            = 1 << 2  // Column will keep a fixed size, preferable with horizontal scrolling enabled (default if table sizing policy is SizingPolicyFixedX and table is resizable).
	TableColumnFlagsWidthStretch          = 1 << 3  // Column will stretch, preferable with horizontal scrolling disabled (default if table sizing policy is SizingPolicyStretchX).
	TableColumnFlagsWidthAlwaysAutoResize = 1 << 4  // Column will keep resizing based on submitted contents (with a one frame delay) == Fixed with auto resize (default if table sizing policy is SizingPolicyFixedX and table is not resizable).
	TableColumnFlagsNoResize              = 1 << 5  // Disable manual resizing.
	TableColumnFlagsNoClipX               = 1 << 6  // Disable clipping for this column (all NoClipX columns will render in a same draw command).
	TableColumnFlagsNoSort                = 1 << 7  // Disable ability to sort on this field (even if ImGuiTableFlags_Sortable is set on the table).
	TableColumnFlagsNoSortAscending       = 1 << 8  // Disable ability to sort in the ascending direction.
	TableColumnFlagsNoSortDescending      = 1 << 9  // Disable ability to sort in the descending direction.
	TableColumnFlagsNoHide                = 1 << 10 // Disable hiding this column.
	TableColumnFlagsNoHeaderWidth         = 1 << 11 // Header width don't contribute to automatic column width.
	TableColumnFlagsPreferSortAscending   = 1 << 12 // Make the initial sort direction Ascending when first sorting on this column (default).
	TableColumnFlagsPreferSortDescending  = 1 << 13 // Make the initial sort direction Descending when first sorting on this column.
	TableColumnFlagsIndentEnable          = 1 << 14 // Use current Indent value when entering cell (default for 1st column).
	TableColumnFlagsIndentDisable         = 1 << 15 // Ignore current Indent value when entering cell (default for columns after the 1st one). Indentation changes _within_ the cell will still be honored.
	TableColumnFlagsNoReorder             = 1 << 16 // Disable reordering this column, this will also prevent other columns from crossing over this column.
)

// Flags for ImGui::TableNextRow()
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
// When using ImGuiTableFlags_RowBg on the table, each row has the RowBg0 color automatically set for odd/even rows.
// If you set the color of RowBg0 target, your color will override the existing RowBg0 color.
// If you set the color of RowBg1 or ColumnBg1 target, your color will blend over the RowBg0 color.
const (
	TableBgTargetNone      = 0
	TableBgTargetColumnBg0 = 1 // FIXME-TABLE: Todo. Set column background color 0 (generally used for background
	TableBgTargetColumnBg1 = 2 // FIXME-TABLE: Todo. Set column background color 1 (generally used for selection marking)
	TableBgTargetRowBg0    = 3 // Set row background color 0 (generally used for background, automatically set when ImGuiTableFlags_RowBg is used)
	TableBgTargetRowBg1    = 4 // Set row background color 1 (generally used for selection marking)
	TableBgTargetCellBg    = 5 // Set cell background color (top-most color)
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
// Return true if column is visible.
func TableNextColumn() bool {
	return C.iggTableNextColumn() != 0
}

// TableSetColumnIndex appends into the specified column. Return true if column is visible.
func TableSetColumnIndex(columnN int) bool {
	return C.iggTableSetColumnIndex(C.int(columnN)) != 0
}

// TableGetColumnIndex return current column index.
func TableGetColumnIndex() int {
	return int(C.iggTableGetColumnIndex())
}

// Tables: Headers & Columns declaration
// - Use TableSetupColumn() to specify label, resizing policy, default width, id, various other flags etc.
//   Important: this will not display anything! The name passed to TableSetupColumn() is used by TableHeadersRow() and context-menus.
// - Use TableHeadersRow() to create a row and automatically submit a TableHeader() for each column.
//   Headers are required to perform some interactions: reordering, sorting, context menu (FIXME-TABLE: context menu should work without!)
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

// TableGetColumnIsVisibleV returns true if column is visible.
// Same value is also returned by TableNextColumn() and TableSetColumnIndex().
// Pass -1 to use current column.
func TableGetColumnIsVisibleV(columnN int) bool {
	return C.iggTableGetColumnIsVisible(C.int(columnN)) != 0
}

// TableGetColumnIsVisible calls TableGetColumnIsVisibleV(-1)
func TableGetColumnIsVisible() bool {
	return TableGetColumnIsVisibleV(-1)
}

// TableGetColumnIsSortedV returns true if column is included in the sort specs.
// Rarely used, can be useful to tell if a data change should trigger resort.
// Equivalent to test ImGuiTableSortSpecs's ->ColumnsMask & (1 << column_n).
// Pass -1 to use current column.
func TableGetColumnIsSortedV(columnN int) bool {
	return C.iggTableGetColumnIsSorted(C.int(columnN)) != 0
}

// TableGetColumnIsSorted calls TableGetColumnIsSortedV(-1)
func TableGetColumnIsSorted() bool {
	return TableGetColumnIsSortedV(-1)
}

// TableGetHoveredColumn returns hovered column.
// return -1 when table is not hovered.
// return columnsCount if the unused space at the right of visible columns is hovered.
func TableGetHoveredColumn() int {
	return int(C.iggTableGetHoveredColumn())
}

// TableSetBgColorV changes the color of a cell, row, or column. See ImGuiTableBgTarget flags for details.
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

// Sorting specification for one column of a table (sizeof == 8 bytes)
type TableSortSpecsColumn struct {
	ColumnUserID  uint  // User id of the column (if specified by a TableSetupColumn() call)
	ColumnIndex   uint8 // Index of the column
	SortOrder     uint8 // Index within parent ImGuiTableSortSpecs (always stored in order starting from 0, tables sorted on a single criteria will always have a 0 here)
	SortDirection int   // ImGuiSortDirection_Ascending or ImGuiSortDirection_Descending (you can use this or SortSign, whichever is more convenient for your sort function)
}

// Sorting specifications for a table (often handling sort specs for a single column, occasionally more)
// Obtained by calling TableGetSortSpecs()
// When 'SpecsDirty == true' you can sort your data. It will be true with sorting specs have changed since last call, or the first time.
// Make sure to set 'SpecsDirty = false' after sorting, else you may wastefully sort your data every frame!
type TableSortSpecs struct {
	Specs       [64]TableSortSpecsColumn // Sort spec array.
	SpecsCount  int                      // Sort spec count. Most often 1 unless e.g. ImGuiTableFlags_MultiSortable is enabled.
	SpecsDirty  bool                     // Set to true when specs have changed since last time! Use this to sort again, then clear the flag.
	ColumnsMask uint64                   // Set to the mask of column indexes included in the Specs array. e.g. (1 << N) when column N is sorted.

	// [Internal]
	specsDirty *C.char // Needed to clear the flag after the sort
}

// TableGetSortSpecs gets latest sort specs for the table (nil if not sorting).
func TableGetSortSpecs() *TableSortSpecs {
	// TableSortSpecs.SpecsCount is most often 1 unless e.g. ImGuiTableFlags_MultiSortable is enabled.
	// We use static array with 64 TableSortSpecsColumn values since that's maximum ammount of columns
	sort_specs := &C.IggTableSortSpecs{Specs: [64]C.IggTableSortSpecsColumn{}}
	if C.iggTableGetSortSpecs(sort_specs) == 0 {
		return nil
	}

	result := &TableSortSpecs{
		SpecsCount:  int(sort_specs.SpecsCount),
		SpecsDirty:  sort_specs.SpecsDirty != 0,
		ColumnsMask: uint64(sort_specs.ColumnsMask),

		specsDirty: sort_specs.SpecsDirtyInternal,
	}

	for n := 0; n < int(sort_specs.SpecsCount); n += 1 {
		result.Specs[n] = TableSortSpecsColumn{
			ColumnUserID:  uint(sort_specs.Specs[n].ColumnUserID),
			ColumnIndex:   uint8(sort_specs.Specs[n].ColumnIndex),
			SortOrder:     uint8(sort_specs.Specs[n].SortOrder),
			SortDirection: int(sort_specs.Specs[n].SortDirection),
		}
	}

	return result
}

func (specs *TableSortSpecs) Clear() {
	*specs.specsDirty = 0
}
