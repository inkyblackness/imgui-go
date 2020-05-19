#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggContext iggCreateContext(IggFontAtlas sharedFontAtlas);
extern void iggDestroyContext(IggContext context);
extern IggContext iggGetCurrentContext();
extern void iggSetCurrentContext(IggContext context);

#ifdef __cplusplus
}
#endif