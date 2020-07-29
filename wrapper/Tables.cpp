#include "ConfiguredImGui.h"

#include "Tables.h"
#include "WrapperConverter.h"


IggBool iggBeginTable(char const *str_id, int columns_count, int flags, IggVec2 const *outer_size, float inner_width)
{
	Vec2Wrapper outerSizeArg(outer_size);
	return ImGui::BeginTable(str_id, columns_count, flags, *outerSizeArg, inner_width) ? 1 : 0;
}

void iggEndTable(void)
{
	ImGui::EndTable();
}

void iggTableNextRow(int row_flags, float min_row_height)
{
	ImGui::TableNextRow(row_flags, min_row_height);
}

IggBool iggTableNextCell(void)
{
	return ImGui::TableNextCell() ? 1 : 0;
}

IggBool iggTableSetColumnIndex(int column_n)
{
	return ImGui::TableSetColumnIndex(column_n) ? 1 : 0;
}

int iggTableGetColumnIndex(void)
{
	return ImGui::TableGetColumnIndex();
}

char const *iggTableGetColumnName(int column_n)
{
	return ImGui::TableGetColumnName(column_n);
}

IggBool iggTableGetColumnIsVisible(int column_n)
{
	return ImGui::TableGetColumnIsVisible(column_n) ? 1 : 0;
}

IggBool iggTableGetColumnIsSorted(int column_n)
{
	return ImGui::TableGetColumnIsSorted(column_n) ? 1 : 0;
}

int iggTableGetHoveredColumn(void)
{
	return ImGui::TableGetHoveredColumn();	
}

void iggTableSetBgColor(int bg_target, IggVec4 const *color, int column_n)
{
	Vec4Wrapper colorArg(color);
	auto col = ImGui::GetColorU32(*colorArg);
	ImGui::TableSetBgColor(bg_target, col, column_n);
}

void iggTableSetupColumn(char const *label, int flags, float init_width_or_weight, unsigned int user_id)
{
	ImGui::TableSetupColumn(label, flags, init_width_or_weight, user_id);
}

void iggTableAutoHeaders(void)
{
	ImGui::TableAutoHeaders();
}

void iggTableHeader(char const *label)
{
	ImGui::TableHeader(label);
}

static void exportTableSortSpecs(IggTableSortSpecs &out, ImGuiTableSortSpecs const &in)
{
	for (int n = 0; n < in.SpecsCount; n++)
	{
		out.Specs[n].ColumnUserID = in.Specs[n].ColumnUserID;
		out.Specs[n].ColumnIndex = in.Specs[n].ColumnIndex;
		out.Specs[n].SortOrder = in.Specs[n].SortOrder;
		out.Specs[n].SortDirection = in.Specs[n].SortDirection;
	}
	out.SpecsCount = in.SpecsCount;
	out.SpecsChanged = in.SpecsChanged;
	out.ColumnsMask = in.ColumnsMask;
}

IggBool iggTableGetSortSpecs(IggTableSortSpecs *sort_specs)
{
	const ImGuiTableSortSpecs *ret = ImGui::TableGetSortSpecs();
	if (ret == NULL) return 0;

	exportTableSortSpecs(*sort_specs, *ret); return 1;
}