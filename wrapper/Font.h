#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggPushFont(IggFont handle);
extern void iggPopFont(void);
extern void iggCalcTextSize(const char *text, int length, IggBool hide_text_after_double_hash, float wrap_width, IggVec2 *value);
extern float iggGetFontSize();

#ifdef __cplusplus
}
#endif
