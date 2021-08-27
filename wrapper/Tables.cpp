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

IggBool iggTableNextColumn(void)
{
	return ImGui::TableNextColumn() ? 1 : 0;
}

IggBool iggTableSetColumnIndex(int column_n)
{
	return ImGui::TableSetColumnIndex(column_n) ? 1 : 0;
}

int iggTableGetColumnIndex(void)
{
	return ImGui::TableGetColumnIndex();
}

int iggTableGetRowIndex(void)
{
	return ImGui::TableGetRowIndex();
}

void iggTableSetupColumn(char const *label, int flags, float init_width_or_weight, unsigned int user_id)
{
	ImGui::TableSetupColumn(label, flags, init_width_or_weight, user_id);
}

void iggTableSetupScrollFreeze(int cols, int rows)
{
	ImGui::TableSetupScrollFreeze(cols, rows);
}

void iggTableHeadersRow(void)
{
	ImGui::TableHeadersRow();
}

void iggTableHeader(char const *label)
{
	ImGui::TableHeader(label);
}

int iggTableGetColumnCount(void)
{
	return ImGui::TableGetColumnCount();
}

char const *iggTableGetColumnName(int column_n)
{
	return ImGui::TableGetColumnName(column_n);
}

int iggTableGetColumnFlags(int column_n)
{
	return ImGui::TableGetColumnFlags(column_n);
}

void iggTableSetBgColor(int target, IggVec4 const *color, int column_n)
{
	Vec4Wrapper colorArg(color);
	auto col = ImGui::GetColorU32(*colorArg);
	ImGui::TableSetBgColor(target, col, column_n);
}

IggTableSortSpecs iggTableGetSortSpecs()
{
	return static_cast<IggTableSortSpecs>(ImGui::TableGetSortSpecs());
}

void iggTableSortSpecsGetSpec(IggTableSortSpecs handle, int index, IggTableColumnSortSpecs *out)
{
	ImGuiTableSortSpecs *sort_specs = reinterpret_cast<ImGuiTableSortSpecs *>(handle);
	ImGuiTableColumnSortSpecs column_spec = sort_specs->Specs[index];

	out->ColumnUserID = column_spec.ColumnUserID;
	out->ColumnIndex = column_spec.ColumnIndex;
	out->SortOrder = column_spec.SortOrder;
	out->SortDirection = column_spec.SortDirection;
}

int iggTableSortSpecsGetSpecsCount(IggTableSortSpecs handle)
{
	ImGuiTableSortSpecs *sort_specs = reinterpret_cast<ImGuiTableSortSpecs *>(handle);
	return sort_specs->SpecsCount;
}

IggBool iggTableSortSpecsGetSpecsDirty(IggTableSortSpecs handle)
{
	ImGuiTableSortSpecs *sort_specs = reinterpret_cast<ImGuiTableSortSpecs *>(handle);
	return sort_specs->SpecsDirty ? 1 : 0;
}

void iggTableSortSpecsSetSpecsDirty(IggTableSortSpecs handle, IggBool value)
{
	ImGuiTableSortSpecs *sort_specs = reinterpret_cast<ImGuiTableSortSpecs *>(handle);
	sort_specs->SpecsDirty = value != 0;
}
