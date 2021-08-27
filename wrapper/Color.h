#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggBool iggColorEdit3(char const *label, float *col, int flags);
extern IggBool iggColorEdit4(char const *label, float *col, int flags);
extern IggBool iggColorPicker3(char const *label, float *col, int flags);
extern IggBool iggColorPicker4(char const *label, float *col, int flags);

#ifdef __cplusplus
}
#endif
