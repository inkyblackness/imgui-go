#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern float iggGetScrollX();
extern float iggGetScrollY();
extern float iggGetScrollMaxX();
extern float iggGetScrollMaxY();
extern void iggSetScrollHereX(float centerXRatio);
extern void iggSetScrollHereY(float centerYRatio);
extern void iggSetScrollX(float scrollX);
extern void iggSetScrollY(float scrollY);

#ifdef __cplusplus
}
#endif
