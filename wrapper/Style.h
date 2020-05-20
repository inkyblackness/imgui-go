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

extern void iggStyleGetFramePadding(IggGuiStyle handle, IggVec2 *value);

extern void iggStyleSetColor(IggGuiStyle handle, int index, IggVec4 const *color);

extern void iggStyleScaleAllSizes(IggGuiStyle handle, float scale);

#ifdef __cplusplus
}
#endif
