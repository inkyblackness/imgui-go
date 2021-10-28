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

void iggStyleSetItemInnerSpacing(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->ItemInnerSpacing, *value);
}

void iggStyleSetItemSpacing(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->ItemSpacing, *value);
}

void iggStyleSetFramePadding(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->FramePadding, *value);
}

void iggStyleSetWindowPadding(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->WindowPadding, *value);
}

void iggStyleSetCellPadding(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->CellPadding, *value);
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

void iggGetTouchExtraPadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->TouchExtraPadding);
}

void iggSetTouchExtraPadding(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->TouchExtraPadding, *value);
}

float iggGetAlpha(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->Alpha;
}

void iggSetAlpha(IggGuiStyle handle, float alpha)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->Alpha = alpha;
}

float iggGetDisabledAlpha(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->DisabledAlpha;
}

void iggSetDisabledAlpha(IggGuiStyle handle, float disabledAlpha)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->DisabledAlpha = disabledAlpha;
}

float iggGetWindowRounding(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->WindowRounding;
}

void iggSetWindowRounding(IggGuiStyle handle, float windowRounding)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->WindowRounding = windowRounding;
}

float iggGetWindowBorderSize(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->WindowBorderSize;
}

void iggSetWindowBorderSize(IggGuiStyle handle, float windowBorderSize)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->WindowBorderSize = windowBorderSize;
}

float iggGetChildRounding(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->ChildRounding;
}

void iggSetChildRounding(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->ChildRounding = v;
}

float iggGetChildBorderSize(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->ChildBorderSize;
}

void iggSetChildBorderSize(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->ChildBorderSize = v;
}

float iggGetPopupRounding(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->PopupRounding;
}

void iggSetPopupRounding(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->PopupRounding = v;
}

float iggGetPopupBorderSize(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->PopupBorderSize;
}

void iggSetPopupBorderSize(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->PopupBorderSize = v;
}

float iggGetFrameRounding(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->FrameRounding;
}

void iggSetFrameRounding(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->FrameRounding = v;
}

float iggGetFrameBorderSize(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->FrameBorderSize;
}

void iggSetFrameBorderSize(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->FrameBorderSize = v;
}

float iggGetIndentSpacing(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->IndentSpacing;
}

void iggSetIndentSpacing(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->IndentSpacing = v;
}

float iggGetColumnsMinSpacing(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->ColumnsMinSpacing;
}

void iggSetColumnsMinSpacing(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->ColumnsMinSpacing = v;
}

float iggGetScrollbarSize(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->ScrollbarSize;
}

void iggSetScrollbarSize(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->ScrollbarSize = v;
}

float iggGetScrollbarRounding(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->ScrollbarRounding;
}

void iggSetScrollbarRounding(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->ScrollbarRounding = v;
}

float iggGetGrabMinSize(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->GrabMinSize;
}

void iggSetGrabMinSize(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->GrabMinSize = v;
}

float iggGetGrabRounding(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->GrabRounding;
}

void iggSetGrabRounding(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->GrabRounding = v;
}

float iggGetLogSliderDeadzone(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->LogSliderDeadzone;
}

void iggSetLogSliderDeadzone(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->LogSliderDeadzone = v;
}

float iggGetTabRounding(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->TabRounding;
}

void iggSetTabRounding(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->TabRounding = v;
}

float iggGetTabBorderSize(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->TabBorderSize;
}

void iggSetTabBorderSize(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->TabBorderSize = v;
}

float iggGetTabMinWidthForCloseButton(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->TabMinWidthForCloseButton;
}

void iggSetTabMinWidthForCloseButton(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->TabMinWidthForCloseButton = v;
}

float iggGetCurveTessellationTol(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->CurveTessellationTol;
}

void iggSetCurveTessellationTol(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->CurveTessellationTol = v;
}

float iggGetCircleTessellationMaxError(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->CircleTessellationMaxError;
}

void iggSetCircleTessellationMaxError(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->CircleTessellationMaxError = v;
}

float iggGetMouseCursorScale(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->MouseCursorScale;
}

void iggSetMouseCursorScale(IggGuiStyle handle, float v)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->MouseCursorScale = v;
}

void iggStyleGetWindowMinSize(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->WindowMinSize);
}

void iggSetWindowMinSize(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->WindowMinSize, *value);
}

void iggStyleGetWindowTitleAlign(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->WindowTitleAlign);
}

void iggSetWindowTitleAlign(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->WindowTitleAlign, *value);
}

void iggStyleGetButtonTextAlign(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->ButtonTextAlign);
}

void iggSetButtonTextAlign(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->ButtonTextAlign, *value);
}

void iggStyleGetSelectableTextAlign(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->SelectableTextAlign);
}

void iggSetSelectableTextAlign(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->SelectableTextAlign, *value);
}

void iggStyleGetDisplayWindowPadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->DisplayWindowPadding);
}

void iggSetDisplayWindowPadding(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->DisplayWindowPadding, *value);
}

void iggStyleGetDisplaySafeAreaPadding(IggGuiStyle handle, IggVec2 *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   exportValue(*value, style->DisplaySafeAreaPadding);
}

void iggSetDisplaySafeAreaPadding(IggGuiStyle handle, IggVec2 const *value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   importValue(style->DisplaySafeAreaPadding, *value);
}

IggBool iggStyleGetAntiAliasedLines(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->AntiAliasedLines ? 1 : 0;
}

void iggStyleSetAntiAliasedLines(IggGuiStyle handle, IggBool value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->AntiAliasedLines = (value == 0) ? false : true;
}

IggBool iggStyleGetAntiAliasedLinesUseTex(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->AntiAliasedLinesUseTex ? 1 : 0;
}

void iggStyleSetAntiAliasedLinesUseTex(IggGuiStyle handle, IggBool value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->AntiAliasedLinesUseTex = (value == 0) ? false : true;
}

IggBool iggStyleGetAntiAliasedFill(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->AntiAliasedFill ? 1 : 0;
}

void iggStyleSetAntiAliasedFill(IggGuiStyle handle, IggBool value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->AntiAliasedFill = (value == 0) ? false : true;
}

IggDir iggStyleGetWindowMenuButtonPosition(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->WindowMenuButtonPosition;
}

void iggStyleSetWindowMenuButtonPosition(IggGuiStyle handle, IggDir value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->WindowMenuButtonPosition = value;
}

IggDir iggStyleGetColorButtonPosition(IggGuiStyle handle)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   return style->ColorButtonPosition;
}

void iggStyleSetColorButtonPosition(IggGuiStyle handle, IggDir value)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->ColorButtonPosition = value;
}
