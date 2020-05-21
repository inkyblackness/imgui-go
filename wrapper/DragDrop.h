#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggBool iggBeginDragDropSource(int flags);
extern IggBool iggSetDragDropPayload(const char *type, const void *data, int size, int cond);
extern void iggEndDragDropSource();
extern IggBool iggBeginDragDropTarget();
extern const IggPayload iggAcceptDragDropPayload(const char *type, int flags);
extern void iggEndDragDropTarget();

extern void *iggPayloadData(const IggPayload payload);
extern int iggPayloadDataSize(const IggPayload payload);

#ifdef __cplusplus
}
#endif
