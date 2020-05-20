#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggOpenPopup(char const *id);
extern IggBool iggBeginPopup(char const *name, int flags);
extern IggBool iggBeginPopupModal(char const *name, IggBool *open, int flags);
extern IggBool iggBeginPopupContextItem(char const *label, int mouseButton);
extern void iggEndPopup(void);
extern void iggCloseCurrentPopup(void);

#ifdef __cplusplus
}
#endif
