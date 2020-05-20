#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggBool iggIsItemClicked();
extern IggBool iggIsItemHovered(int flags);
extern IggBool iggIsItemActive();
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
extern void iggMousePos(IggVec2 *pos);
extern int iggGetMouseCursor();
extern void iggSetMouseCursor(int cursor);

#ifdef __cplusplus
}
#endif
