#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggStyleColorsDark();
extern void iggStyleColorsClassic();
extern void iggStyleColorsLight();

extern IggGuiStyle iggGetCurrentStyle(void);

extern void iggPushStyleColor(int index, IggVec4 const *col);
extern void iggPopStyleColor(int count);
extern void iggPushStyleVarFloat(int index, float value);
extern void iggPushStyleVarVec2(int index, IggVec2 const *value);
extern void iggPopStyleVar(int count);

extern void iggStyleGetItemInnerSpacing(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetItemSpacing(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleSetItemInnerSpacing(IggGuiStyle handle, IggVec2 const *value);
extern void iggStyleSetItemSpacing(IggGuiStyle handle, IggVec2 const *value);

extern void iggStyleGetFramePadding(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetWindowPadding(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleGetCellPadding(IggGuiStyle handle, IggVec2 *value);
extern void iggStyleSetFramePadding(IggGuiStyle handle, IggVec2 const *value);
extern void iggStyleSetWindowPadding(IggGuiStyle handle, IggVec2 const *value);
extern void iggStyleSetCellPadding(IggGuiStyle handle, IggVec2 const *value);

extern void iggStyleSetColor(IggGuiStyle handle, int index, IggVec4 const *color);
extern void iggStyleGetColor(IggGuiStyle handle, int index, IggVec4 *value);

extern void iggStyleScaleAllSizes(IggGuiStyle handle, float scale);

extern void iggGetTouchExtraPadding(IggGuiStyle handle, IggVec2 *value);
extern void iggSetTouchExtraPadding(IggGuiStyle handle, IggVec2 const *value);

extern float iggGetAlpha(IggGuiStyle handle);
extern void iggSetAlpha(IggGuiStyle handle, float alpha);
extern float iggGetDisabledAlpha(IggGuiStyle handle);
extern void iggSetDisabledAlpha(IggGuiStyle handle, float disabledAlpha);
extern float iggGetWindowRounding(IggGuiStyle handle);
extern void iggSetWindowRounding(IggGuiStyle handle, float windowRounding);
extern float iggGetWindowBorderSize(IggGuiStyle handle);
extern void iggSetWindowBorderSize(IggGuiStyle handle, float windowBorderSize);
extern float iggGetChildRounding(IggGuiStyle handle);
extern void iggSetChildRounding(IggGuiStyle handle, float v);
extern float iggGetChildBorderSize(IggGuiStyle handle);
extern void iggSetChildBorderSize(IggGuiStyle handle, float v);
extern float iggGetPopupRounding(IggGuiStyle handle);
extern void iggSetPopupRounding(IggGuiStyle handle, float v);
extern float iggGetPopupBorderSize(IggGuiStyle handle);
extern void iggSetPopupBorderSize(IggGuiStyle handle, float v);
extern float iggGetFrameRounding(IggGuiStyle handle);
extern void iggSetFrameRounding(IggGuiStyle handle, float v);
extern float iggGetFrameBorderSize(IggGuiStyle handle);
extern void iggSetFrameBorderSize(IggGuiStyle handle, float v);
extern float iggGetIndentSpacing(IggGuiStyle handle);
extern void iggSetIndentSpacing(IggGuiStyle handle, float v);
extern float iggGetColumnsMinSpacing(IggGuiStyle handle);
extern void iggSetColumnsMinSpacing(IggGuiStyle handle, float v);
extern float iggGetScrollbarSize(IggGuiStyle handle);
extern void iggSetScrollbarSize(IggGuiStyle handle, float v);
extern float iggGetScrollbarRounding(IggGuiStyle handle);
extern void iggSetScrollbarRounding(IggGuiStyle handle, float v);
extern float iggGetGrabMinSize(IggGuiStyle handle);
extern void iggSetGrabMinSize(IggGuiStyle handle, float v);
extern float iggGetGrabRounding(IggGuiStyle handle);
extern void iggSetGrabRounding(IggGuiStyle handle, float v);
extern float iggGetLogSliderDeadzone(IggGuiStyle handle);
extern void iggSetLogSliderDeadzone(IggGuiStyle handle, float v);
extern float iggGetTabRounding(IggGuiStyle handle);
extern void iggSetTabRounding(IggGuiStyle handle, float v);
extern float iggGetTabBorderSize(IggGuiStyle handle);
extern void iggSetTabBorderSize(IggGuiStyle handle, float v);
extern float iggGetTabMinWidthForCloseButton(IggGuiStyle handle);
extern void iggSetTabMinWidthForCloseButton(IggGuiStyle handle, float v);
extern float iggGetCurveTessellationTol(IggGuiStyle handle);
extern void iggSetCurveTessellationTol(IggGuiStyle handle, float v);
extern float iggGetCircleTessellationMaxError(IggGuiStyle handle);
extern void iggSetCircleTessellationMaxError(IggGuiStyle handle, float v);
extern float iggGetMouseCursorScale(IggGuiStyle handle);
extern void iggSetMouseCursorScale(IggGuiStyle handle, float v);
extern void iggStyleGetWindowMinSize(IggGuiStyle handle, IggVec2 *value);
extern void iggSetWindowMinSize(IggGuiStyle handle, IggVec2 const *value);
extern void iggStyleGetWindowTitleAlign(IggGuiStyle handle, IggVec2 *value);
extern void iggSetWindowTitleAlign(IggGuiStyle handle, IggVec2 const *value);
extern void iggStyleGetButtonTextAlign(IggGuiStyle handle, IggVec2 *value);
extern void iggSetButtonTextAlign(IggGuiStyle handle, IggVec2 const *value);
extern void iggStyleGetSelectableTextAlign(IggGuiStyle handle, IggVec2 *value);
extern void iggSetSelectableTextAlign(IggGuiStyle handle, IggVec2 const *value);
extern void iggStyleGetDisplayWindowPadding(IggGuiStyle handle, IggVec2 *value);
extern void iggSetDisplayWindowPadding(IggGuiStyle handle, IggVec2 const *value);
extern void iggStyleGetDisplaySafeAreaPadding(IggGuiStyle handle, IggVec2 *value);
extern void iggSetDisplaySafeAreaPadding(IggGuiStyle handle, IggVec2 const *value);
extern IggBool iggStyleGetAntiAliasedLines(IggGuiStyle handle);
extern void iggStyleSetAntiAliasedLines(IggGuiStyle handle, IggBool value);
extern IggBool iggStyleGetAntiAliasedLinesUseTex(IggGuiStyle handle);
extern void iggStyleSetAntiAliasedLinesUseTex(IggGuiStyle handle, IggBool value);
extern IggBool iggStyleGetAntiAliasedFill(IggGuiStyle handle);
extern void iggStyleSetAntiAliasedFill(IggGuiStyle handle, IggBool value);
extern IggDir iggStyleGetWindowMenuButtonPosition(IggGuiStyle handle);
extern void iggStyleSetWindowMenuButtonPosition(IggGuiStyle handle, IggDir value);
extern IggDir iggStyleGetColorButtonPosition(IggGuiStyle handle);
extern void iggStyleSetColorButtonPosition(IggGuiStyle handle, IggDir value);

#ifdef __cplusplus
}
#endif
