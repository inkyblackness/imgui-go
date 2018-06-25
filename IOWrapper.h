#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern IggBool iggWantCaptureMouse(IggIO handle);

extern IggFontAtlas iggIoGetFonts(IggIO handle);

extern void iggIoSetDisplaySize(IggIO handle, IggVec2 const *value);
extern void iggIoSetMousePosition(IggIO handle, IggVec2 const *value);
extern void iggIoSetMouseButtonDown(IggIO handle, int index, IggBool value);
extern void iggIoAddMouseWheelDelta(IggIO handle, float x, float y);
extern void iggIoSetDeltaTime(IggIO handle, float value);
extern void iggIoSetFontGlobalScale(IggIO handle, float value);

#ifdef __cplusplus
}
#endif
