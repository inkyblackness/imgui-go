#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct tagIggTableColumnSortSpecs
{
	unsigned int ColumnUserID;
	short        ColumnIndex;
	short        SortOrder;
	int          SortDirection;
} IggTableColumnSortSpecs;

extern IggBool     iggBeginTable(char const *str_id, int columns_count, int flags, IggVec2 const *outer_size, float inner_width);
extern void        iggEndTable();
extern void        iggTableNextRow(int row_flags, float min_row_height);
extern IggBool     iggTableNextColumn();
extern IggBool     iggTableSetColumnIndex(int column_n);
extern int         iggTableGetColumnIndex();
extern int         iggTableGetRowIndex();
extern void        iggTableSetupColumn(char const *label, int flags, float init_width_or_weight, unsigned int user_id);
extern void        iggTableSetupScrollFreeze(int cols, int rows);
extern void        iggTableHeadersRow();
extern void        iggTableHeader(char const *label);
extern int         iggTableGetColumnCount();
extern char const *iggTableGetColumnName(int column_n);
extern int         iggTableGetColumnFlags(int column_n);
extern void        iggTableSetBgColor(int target, IggVec4 const *color, int column_n);

extern IggTableSortSpecs iggTableGetSortSpecs();
extern void iggTableSortSpecsGetSpec(IggTableSortSpecs handle, int index, IggTableColumnSortSpecs *out);
extern int iggTableSortSpecsGetSpecsCount(IggTableSortSpecs handle);
extern IggBool iggTableSortSpecsGetSpecsDirty(IggTableSortSpecs handle);
extern void iggTableSortSpecsSetSpecsDirty(IggTableSortSpecs handle, IggBool value);

#ifdef __cplusplus
}
#endif
