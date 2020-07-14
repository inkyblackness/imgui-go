#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggBool     iggBeginTable(char const *str_id, int columns_count, ImGuiTableFlags flags, IggVec2 const *outer_size, float inner_width);
extern void        iggEndTable();
extern void        iggTableNextRow(ImGuiTableRowFlags row_flags, float min_row_height);
extern IggBool     iggTableNextCell();
extern IggBool     iggTableSetColumnIndex(int column_n);
extern int         iggTableGetColumnIndex();
extern char const *iggTableGetColumnName(int column_n);
extern IggBool     iggTableGetColumnIsVisible(int column_n);
extern IggBool     iggTableGetColumnIsSorted(int column_n);
extern int         iggTableGetHoveredColumn();
extern void        iggTableSetupColumn(char const *label, ImGuiTableColumnFlags flags, float init_width_or_weight, unsigned int user_id);
extern void        iggTableAutoHeaders();
extern void        iggTableHeader(char const *label);
// extern const ImGuiTableSortSpecs* iggTableGetSortSpecs();

#ifdef __cplusplus
}
#endif
