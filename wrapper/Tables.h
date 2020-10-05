#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct tagIggTableSortSpecsColumn
{
	unsigned int ColumnUserID;
	unsigned char ColumnIndex;
	unsigned char SortOrder;
	int  SortDirection;
} IggTableSortSpecsColumn;

typedef struct tagIggTableSortSpecs
{
	IggTableSortSpecsColumn Specs[64];
	int                     SpecsCount;
	IggBool                 SpecsDirty;
	unsigned long long      ColumnsMask;

	char*                   SpecsDirtyInternal; // Needed to clear the flag after the sort
} IggTableSortSpecs;

extern IggBool     iggBeginTable(char const *str_id, int columns_count, int flags, IggVec2 const *outer_size, float inner_width);
extern void        iggEndTable();
extern void        iggTableNextRow(int row_flags, float min_row_height);
extern IggBool     iggTableNextColumn();
extern IggBool     iggTableSetColumnIndex(int column_n);
extern int         iggTableGetColumnIndex();
extern void        iggTableSetupColumn(char const *label, int flags, float init_width_or_weight, unsigned int user_id);
extern void        iggTableSetupScrollFreeze(int cols, int rows);
extern void        iggTableHeadersRow();
extern void        iggTableHeader(char const *label);
extern int         iggTableGetColumnCount();
extern char const *iggTableGetColumnName(int column_n);
extern IggBool     iggTableGetColumnIsVisible(int column_n);
extern IggBool     iggTableGetColumnIsSorted(int column_n);
extern int         iggTableGetHoveredColumn();
extern void        iggTableSetBgColor(int bg_target, IggVec4 const *color, int column_n);
extern IggBool     iggTableGetSortSpecs(IggTableSortSpecs *sort_specs);

#ifdef __cplusplus
}
#endif
