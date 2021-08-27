#include "ConfiguredImGui.h"

#include "Style.h"
#include "WrapperConverter.h"

IggGuiStyle iggGetCurrentStyle()
{
   return reinterpret_cast<IggGuiStyle>(&ImGui::GetStyle());
}

void iggStyleColorsDark()
{
   ImGui::StyleColorsDark();
}

void iggStyleColorsClassic()
{
   ImGui::StyleColorsClassic();
}

void iggStyleColorsLight()
{
   ImGui::StyleColorsLight();
}

void iggPushStyleColor(int index, IggVec4 const *col)
{
   Vec4Wrapper colArg(col);
   ImGui::PushStyleColor(index, *colArg);
}

void iggPopStyleColor(int count)
{
   ImGui::PopStyleColor(count);
}

void iggPushStyleVarFloat(int index, float value)
{
   ImGui::PushStyleVar(index, value);
}

void iggPushStyleVarVec2(int index, IggVec2 const *value)
{
   Vec2Wrapper valueArg(value);
   ImGui::PushStyleVar(index, *valueArg);
}

void iggPopStyleVar(int count)
{
   ImGui::PopStyleVar(count);
}

void iggStyleGetItemInnerSpacing(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->ItemInnerSpacing);
}

void iggStyleGetItemSpacing(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->ItemSpacing);
}

void iggStyleGetFramePadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->FramePadding);
}

void iggStyleGetWindowPadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->WindowPadding);
}

void iggStyleGetCellPadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->CellPadding);
}

void iggStyleSetColor(IggGuiStyle handle, int colorID, IggVec4 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   if ((colorID >= 0) && (colorID < ImGuiCol_COUNT))
   {
      importValue(style->Colors[colorID], *value);
   }
}

void iggStyleGetColor(IggGuiStyle handle, int colorID, IggVec4 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   if ((colorID >= 0) && (colorID < ImGuiCol_COUNT))
   {
      exportValue(*value, style->Colors[colorID]);
   }
}

void iggStyleScaleAllSizes(IggGuiStyle handle, float scale)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->ScaleAllSizes(scale);
}

void iggSetTouchExtraPadding(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->TouchExtraPadding, *value);
}
