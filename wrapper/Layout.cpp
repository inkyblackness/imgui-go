#include "ConfiguredImGui.h"

#include "Layout.h"
#include "WrapperConverter.h"

void iggPushID(char const *id)
{
   ImGui::PushID(id);
}

void iggPushIDInt(int id)
{
   ImGui::PushID(id);
}

void iggPopID(void)
{
   ImGui::PopID();
}

void iggSeparator(void)
{
   ImGui::Separator();
}

void iggSameLine(float posX, float spacingW)
{
   ImGui::SameLine(posX, spacingW);
}

void iggSpacing(void)
{
   ImGui::Spacing();
}

void iggDummy(IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   ImGui::Dummy(*sizeArg);
}

void iggBeginGroup(void)
{
   ImGui::BeginGroup();
}

void iggEndGroup(void)
{
   ImGui::EndGroup();
}

void iggBeginDisabled(IggBool disabled)
{
   ImGui::BeginDisabled(disabled);
}

void iggEndDisabled(void)
{
   ImGui::EndDisabled();
}

void iggIndent(float indent_w)
{
   ImGui::Indent(indent_w);
}

void iggUnindent(float indent_w)
{
   ImGui::Unindent(indent_w);
}

void iggCursorPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetCursorPos());
}

float iggCursorPosX(void)
{
   return ImGui::GetCursorPosX();
}

float iggCursorPosY(void)
{
   return ImGui::GetCursorPosY();
}

void iggCursorStartPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetCursorStartPos());
}

void iggCursorScreenPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetCursorScreenPos());
}

void iggSetCursorPos(IggVec2 const *localPos)
{
   Vec2Wrapper localPosArg(localPos);
   ImGui::SetCursorPos(*localPosArg);
}

void iggSetCursorScreenPos(IggVec2 const *absPos)
{
   Vec2Wrapper absPosArg(absPos);
   ImGui::SetCursorScreenPos(*absPosArg);
}

void iggAlignTextToFramePadding()
{
   ImGui::AlignTextToFramePadding();
}

float iggGetTextLineHeight(void)
{
   return ImGui::GetTextLineHeight();
}

float iggGetTextLineHeightWithSpacing(void)
{
   return ImGui::GetTextLineHeightWithSpacing();
}

float iggGetFrameHeight(void)
{
   return ImGui::GetFrameHeight();
}

float iggGetFrameHeightWithSpacing(void)
{
   return ImGui::GetFrameHeightWithSpacing();
}