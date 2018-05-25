#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern void iggStyleSetColor(IggGuiStyle handle, int index, IggVec4 const *color);

extern void iggStyleScaleAllSizes(IggGuiStyle handle, float scale);

#ifdef __cplusplus
}
#endif
