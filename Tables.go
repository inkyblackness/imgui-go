package imgui

// #include "wrapper/Tables.h"
import "C"

// Tables
// [ALPHA API] API will evolve! (FIXME-TABLE)
// - Full-featured replacement for old Columns API
// - In most situations you can use TableNextRow() + TableSetColumnIndex() to populate a table.
// - If you are using tables as a sort of grid, populating every columns with the same type of contents,
//   you may prefer using TableNextCell() instead of TableNextRow() + TableSetColumnIndex().
// - See Demo->Tables for details.
// - See TableFlags and TableColumnsFlags enums for a description of available flags.

// Flags for ImGui::BeginTable()
// - Columns can either varying resizing policy: "Fixed", "Stretch" or "AlwaysAutoResize". Toggling ScrollX needs to alter default sizing policy.
// - Sizing policy have many subtle side effects which may be hard to fully comprehend at first.. They'll eventually make sense.
//   - with SizingPolicyFixedX (default is ScrollX is on):     Columns can be enlarged as needed. Enable scrollbar if ScrollX is enabled, otherwise extend parent window's contents rect. Only Fixed columns allowed. Weighted columns will calculate their width assuming no scrolling.
//   - with SizingPolicyStretchX (default is ScrollX is off):  Fit all columns within available table width (so it doesn't make sense to use ScrollX with Stretch columns!). Fixed and Weighted columns allowed.

const (
	// Features
	TableFlagsNone            = 0
	TableFlagsResizable       = 1 << 0 // Allow resizing columns.
	TableFlagsReorderable     = 1 << 1 // Allow reordering columns (need calling TableSetupColumn() + TableAutoHeaders() or TableHeaders() to display headers)
	TableFlagsHideable        = 1 << 2 // Allow hiding columns (with right-click on header) (FIXME-TABLE: allow without headers).
	TableFlagsSortable        = 1 << 3 // Allow sorting on one column (sort_specs_count will always be == 1). Call TableGetSortSpecs() to obtain sort specs.
	TableFlagsMultiSortable   = 1 << 4 // Allow sorting on multiple columns by holding Shift (sort_specs_count may be > 1). Call TableGetSortSpecs() to obtain sort specs.
	TableFlagsNoSavedSettings = 1 << 5 // Disable persisting columns order, width and sort settings in the .ini file.
	// Decoration
	TableFlagsRowBg              = 1 << 6                                            // Use ImGuiCol_TableRowBg and ImGuiCol_TableRowBgAlt colors behind each rows.
	TableFlagsBordersHInner      = 1 << 7                                            // Draw horizontal borders between rows.
	TableFlagsBordersHOuter      = 1 << 8                                            // Draw horizontal borders at the top and bottom.
	TableFlagsBordersVInner      = 1 << 9                                            // Draw vertical borders between columns.
	TableFlagsBordersVOuter      = 1 << 10                                           // Draw vertical borders on the left and right sides.
	TableFlagsBordersH           = TableFlagsBordersHInner | TableFlagsBordersHOuter // Draw horizontal borders.
	TableFlagsBordersV           = TableFlagsBordersVInner | TableFlagsBordersVOuter // Draw vertical borders.
	TableFlagsBordersInner       = TableFlagsBordersVInner | TableFlagsBordersHInner // Draw inner borders.
	TableFlagsBordersOuter       = TableFlagsBordersVOuter | TableFlagsBordersHOuter // Draw outer borders.
	TableFlagsBorders            = TableFlagsBordersInner | TableFlagsBordersOuter   // Draw all borders.
	TableFlagsBordersVFullHeight = 1 << 11                                           // Borders covers all rows even when Headers are being used. Allow resizing from any rows.
	// Padding, Sizing
	TableFlagsNoClipX              = 1 << 12 // Disable pushing clipping rectangle for every individual columns (reduce draw command count, items will be able to overflow)
	TableFlagsSizingPolicyFixedX   = 1 << 13 // Default if ScrollX is on. Columns will default to use _WidthFixed or _WidthAlwaysAutoResize policy. Read description above for more details.
	TableFlagsSizingPolicyStretchX = 1 << 14 // Default if ScrollX is off. Columns will default to use _WidthStretch policy. Read description above for more details.
	TableFlagsNoHeadersWidth       = 1 << 15 // Disable header width contribution to automatic width calculation.
	TableFlagsNoHostExtendY        = 1 << 16 // (FIXME-TABLE: Reword as SizingPolicy?) Disable extending past the limit set by outer_size.y, only meaningful when neither of ScrollX|ScrollY are set (data below the limit will be clipped and not visible)
	TableFlagsNoKeepColumnsVisible = 1 << 17 // (FIXME-TABLE) Disable code that keeps column always minimally visible when table width gets too small.
	// Scrolling
	TableFlagsScrollX                = 1 << 18 // Enable horizontal scrolling. Require 'outer_size' parameter of BeginTable() to specify the container size. Because this create a child window, ScrollY is currently generally recommended when using ScrollX.
	TableFlagsScrollY                = 1 << 19 // Enable vertical scrolling. Require 'outer_size' parameter of BeginTable() to specify the container size.
	TableFlagsScroll                 = TableFlagsScrollX | TableFlagsScrollY
	TableFlagsScrollFreezeTopRow     = 1 << 20 // We can lock 1 to 3 rows (starting from the top). Use with ScrollY enabled.
	TableFlagsScrollFreeze2Rows      = 2 << 20
	TableFlagsScrollFreeze3Rows      = 3 << 20
	TableFlagsScrollFreezeLeftColumn = 1 << 22 // We can lock 1 to 3 columns (starting from the left). Use with ScrollX enabled.
	TableFlagsScrollFreeze2Columns   = 2 << 22
	TableFlagsScrollFreeze3Columns   = 3 << 22
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

// TableNextCell appends into the next column (next column, or next row if currently in last column).
// Return true if column is visible.
func TableNextCell() bool {
	return C.iggTableNextCell() != 0
}

// TableSetColumnIndex appends into the specified column. Return true if column is visible.
func TableSetColumnIndex(columnN int) bool {
	return C.iggTableSetColumnIndex(C.int(columnN)) != 0
}

// TableGetColumnIndex return current column index.
func TableGetColumnIndex() int {
	return int(C.iggTableGetColumnIndex())
}

// TableGetColumnNameV returns NULL if column didn't have a name declared by TableSetupColumn(). Pass -1 to use current column.
func TableGetColumnNameV(columnN int) string {
	return C.GoString(C.iggTableGetColumnName(C.int(columnN)))
}

// TableGetColumnName calls TableGetColumnNameV(-1)
func TableGetColumnName() string {
	return TableGetColumnNameV(-1)
}

// TableGetColumnIsVisibleV returns true if column is visible. Same value is also returned by TableNextCell() and TableSetColumnIndex(). Pass -1 to use current column.
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

// Tables: Headers & Columns declaration
// - Use TableSetupColumn() to specify label, resizing policy, default width, id, various other flags etc.
// - The name passed to TableSetupColumn() is used by TableAutoHeaders() and by the context-menu
// - Use TableAutoHeaders() to submit the whole header row, otherwise you may treat the header row as a regular row, manually call TableHeader() and other widgets.
// - Headers are required to perform some interactions: reordering, sorting, context menu (FIXME-TABLE: context menu should work without!)
func TableSetupColumnV(label string, flags int, initWidthOrHeight float32, userID uint) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	C.iggTableSetupColumn(labelArg, C.int(flags), C.float(initWidthOrHeight), C.uint(userID))
}

// TableSetupColumn calls TableSetupColumnV(label, 0, -1.0, 0)
func TableSetupColumn(label string) {
	TableSetupColumnV(label, 0, -1.0, 0)
}

// TableAutoHeaders submits all headers cells based on data provided to TableSetupColumn() + submit context menu
func TableAutoHeaders() {
	C.iggTableAutoHeaders()
}

// TableHeader submits one header cell manually
func TableHeader(label string) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	C.iggTableHeader(labelArg)
}

// Tables: Sorting
// - Call TableGetSortSpecs() to retrieve latest sort specs for the table. Return value will be nil if no sorting.
// - You can sort your data again when 'SpecsChanged == true'. It will be true with sorting specs have changed since last call, or the first time.
// - Lifetime: don't hold on this pointer over multiple frames or past any subsequent call to BeginTable()!

// Sorting specification for one column of a table (sizeof == 8 bytes)
type ImGuiTableSortSpecsColumn struct {
	ColumnUserID  uint  // User id of the column (if specified by a TableSetupColumn() call)
	ColumnIndex   uint8 // Index of the column
	SortOrder     uint8 // Index within parent ImGuiTableSortSpecs (always stored in order starting from 0, tables sorted on a single criteria will always have a 0 here)
	SortDirection int   // ImGuiSortDirection_Ascending or ImGuiSortDirection_Descending (you can use this or SortSign, whichever is more convenient for your sort function)

	// ImGuiTableSortSpecsColumn() { ColumnUserID = 0; ColumnIndex = 0; SortOrder = 0; SortDirection = ImGuiSortDirection_Ascending; }
}

// Sorting specifications for a table (often handling sort specs for a single column, occasionally more)
// Obtained by calling TableGetSortSpecs()
type ImGuiTableSortSpecs struct {
	ImGuiTableSortSpecsColumn *Specs // Pointer to sort spec array.
	SpecsCount                int    // Sort spec count. Most often 1 unless e.g. ImGuiTableFlags_MultiSortable is enabled.
	SpecsChanged              bool   // Set to true by TableGetSortSpecs() call if the specs have changed since the previous call. Use this to sort again!
	ColumnsMask               uint64 // Set to the mask of column indexes included in the Specs array. e.g. (1 << N) when column N is sorted.

	// ImGuiTableSortSpecs()       { Specs = nil; SpecsCount = 0; SpecsChanged = false; ColumnsMask = 0x00; }
}


// TableGetSortSpecs gets latest sort specs for the table (nil if not sorting).
func ImGuiTableSortSpecs* TableGetSortSpecs() {

}
