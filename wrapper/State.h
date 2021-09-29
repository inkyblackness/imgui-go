#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggClearActiveID(void);

extern IggBool iggIsItemClicked();
extern IggBool iggIsItemHovered(int flags);
extern IggBool iggIsItemActive();
extern IggBool iggIsItemEdited();
extern IggBool iggIsItemActivated();
extern IggBool iggIsItemDeactivated();
extern IggBool iggIsItemDeactivatedAfterEdit();
extern IggBool iggIsItemToggledOpen();

extern void iggSetItemAllowOverlap();

extern IggBool iggIsAnyItemActive();
extern IggBool iggIsItemVisible();

extern IggBool iggIsWindowAppearing();
extern IggBool iggIsWindowCollapsed();
extern IggBool iggIsWindowFocused(int flags);
extern IggBool iggIsWindowHovered(int flags);

extern IggBool iggIsKeyDown(int key);
extern IggBool iggIsKeyPressed(int key, IggBool repeat);
extern IggBool iggIsKeyReleased(int key);
extern IggBool iggIsMouseDown(int button);
extern IggBool iggIsAnyMouseDown();
extern IggBool iggIsMouseClicked(int button, IggBool repeat);
extern IggBool iggIsMouseReleased(int button);
extern IggBool iggIsMouseDoubleClicked(int button);
extern IggBool iggIsMouseDragging(int button, float threshold);
extern void iggGetMouseDragDelta(IggVec2 *value, int button, float lock_threshold);
extern void iggResetMouseDragDelta(int button);
extern void iggMousePos(IggVec2 *pos);
extern int iggGetMouseCursor();
extern void iggSetMouseCursor(int cursor);

extern void iggGetItemRectMax(IggVec2 *pos);
extern void iggGetItemRectMin(IggVec2 *pos);

#ifdef __cplusplus
}
#endif
