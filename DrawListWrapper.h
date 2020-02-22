#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern int iggDrawListGetCommandCount(IggDrawList handle);
extern IggDrawCmd iggDrawListGetCommand(IggDrawList handle, int index);
extern void iggDrawListGetRawIndexBuffer(IggDrawList handle, void **data, int *byteSize);
extern void iggDrawListGetRawVertexBuffer(IggDrawList handle, void **data, int *byteSize);

extern void iggGetIndexBufferLayout(size_t *entrySize);
extern void iggGetVertexBufferLayout(size_t *entrySize, size_t *posOffset, size_t *uvOffset, size_t *colOffset);

extern void iggAddRect(IggDrawList handle, IggVec2 const *min, IggVec2 const *max, ImU32 col, float rounding, int flags, float thickness);
extern void iggAddRectFilled(IggDrawList handle, IggVec2 const *min, IggVec2 const *max, ImU32 col, float rounding, int flags);
extern void iggAddCircle(IggDrawList handle, IggVec2 const *center, float radius, ImU32 col, int numSegments, float thickness);
extern void iggAddCircleFilled(IggDrawList handle, IggVec2 const *center, float radius, ImU32 col, int numSegments);
extern void iggAddTriangle(IggDrawList handle, IggVec2 *p1, IggVec2 *p2, IggVec2 *p3, ImU32 col, float thickness);
extern void iggAddTriangleFilled(IggDrawList handle, IggVec2 *p1, IggVec2 *p2, IggVec2 *p3, ImU32 col);

extern IggDrawList iggGetWindowDrawList();

#ifdef __cplusplus
}
#endif
