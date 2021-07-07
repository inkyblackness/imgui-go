#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggDrawData iggGetDrawData(void);
extern IggBool iggDrawDataValid(IggDrawData handle);
extern void iggDrawDataGetCommandLists(IggDrawData handle, void **handles, int *count);
extern void iggDrawDataDisplayPos(IggDrawData handle, IggVec2 *value);
extern void iggDrawDataDisplaySize(IggDrawData handle, IggVec2 *value);
extern void iggDrawDataFrameBufferScale(IggDrawData handle, IggVec2 *value);
extern void iggDrawDataScaleClipRects(IggDrawData handle, IggVec2 const *scale);

#ifdef __cplusplus
}
#endif
