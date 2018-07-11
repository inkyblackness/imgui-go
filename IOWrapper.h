#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern IggBool iggWantCaptureMouse(IggIO handle);
extern IggBool iggWantCaptureKeyboard(IggIO handle);
extern IggBool iggWantTextInput(IggIO handle);

extern IggFontAtlas iggIoGetFonts(IggIO handle);

extern void iggIoSetDisplaySize(IggIO handle, IggVec2 const *value);
extern void iggIoSetMousePosition(IggIO handle, IggVec2 const *value);
extern void iggIoSetMouseButtonDown(IggIO handle, int index, IggBool value);
extern void iggIoAddMouseWheelDelta(IggIO handle, float x, float y);
extern void iggIoSetDeltaTime(IggIO handle, float value);
extern void iggIoSetFontGlobalScale(IggIO handle, float value);

extern void iggIoKeyPress(IggIO handle, int key);
extern void iggIoKeyRelease(IggIO handle, int key);
extern void iggIoKeyMap(IggIO handle, int imguiKey, int nativeKey);
extern void iggIoKeyCtrl(IggIO handle, int leftCtrl, int rigthCtrl);
extern void iggIoKeyShift(IggIO handle, int leftShift, int rightShift);
extern void iggIoKeyAlt(IggIO handle, int leftAlt, int rightAlt);
extern void iggIoKeySuper(IggIO handle, int leftSuper, int rightSuper);
extern void iggIoAddInputCharactersUTF8(IggIO handle, const char *utf8_char);

#ifdef __cplusplus
}
#endif
