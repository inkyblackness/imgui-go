#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern void iggSpinner(const char* label, float radius, int thickness);
extern IggBool iggSplitter(IggBool split_vertically, float thickness, float *size1, float *size2);

#ifdef __cplusplus
}
#endif
