#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggSetItemDefaultFocus();
extern IggBool iggIsItemFocused();
extern IggBool iggIsAnyItemFocused();
extern void iggSetKeyboardFocusHere(int offset);

#ifdef __cplusplus
}
#endif
