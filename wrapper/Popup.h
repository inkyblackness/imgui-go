#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggBool iggBeginPopup(const char *name, int flags);
extern IggBool iggBeginPopupModal(const char *name, IggBool *open, int flags);
extern void iggEndPopup(void);
extern void iggOpenPopup(const char *id, int flags);
extern void iggOpenPopupOnItemClick(const char *id, int flags);
extern void iggCloseCurrentPopup(void);
extern IggBool iggBeginPopupContextItem(const char *id, int flags);
extern IggBool iggBeginPopupContextWindow(const char *id, int flags);
extern IggBool iggBeginPopupContextVoid(const char *id, int flags);
extern IggBool iggIsPopupOpen(const char *id, int flags);

#ifdef __cplusplus
}
#endif
