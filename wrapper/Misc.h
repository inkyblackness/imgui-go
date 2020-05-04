#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern char const *iggGetVersion(void);

extern void iggNewFrame(void);
extern void iggRender(void);
extern void iggEndFrame(void);

extern void iggPushID(char const *id);
extern void iggPushIDInt(int id);
extern void iggPopID(void);

extern int iggGetMouseCursor();
extern void iggSetMouseCursor(int cursor);

#ifdef __cplusplus
}
#endif
