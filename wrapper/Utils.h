#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern void iggSpinner(const char* label, float radius, int thickness);
extern IggBool iggSplitter(IggBool split_vertically, float thickness, float *size1, float *size2);
extern IggListClipper iggListClipperInit(int count);
extern IggBool iggListClipperStep(IggListClipper handle, int *display_start, int *display_end);

#ifdef __cplusplus
}
#endif
