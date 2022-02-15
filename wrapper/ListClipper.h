#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct tagIggListClipper
{
   int DisplayStart;
   int DisplayEnd;
   void* imguiListClipper;
} IggListClipper;

extern IggBool iggListClipperStep(IggListClipper *clipper);
extern void iggListClipperBegin(IggListClipper *clipper, int items_count, float items_height);
extern void iggListClipperEnd(IggListClipper *clipper);

#ifdef __cplusplus
}
#endif
