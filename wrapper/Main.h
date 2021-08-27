#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern char const *iggGetVersion(void);
extern double iggGetTime(void);

extern void iggNewFrame(void);
extern void iggRender(void);
extern void iggEndFrame(void);

#ifdef __cplusplus
}
#endif
