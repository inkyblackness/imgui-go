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

void iggTableSetupColumn(char const *label, int flags, float init_width_or_weight, unsigned int user_id);
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