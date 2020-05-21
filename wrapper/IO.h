#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggIO iggGetCurrentIO(void);

extern IggBool iggWantCaptureMouse(IggIO handle);
extern IggBool iggWantCaptureKeyboard(IggIO handle);
extern IggBool iggWantTextInput(IggIO handle);
extern float iggFramerate(IggIO handle);
extern int iggMetricsRenderVertices(IggIO handle);
extern int iggMetricsRenderIndices(IggIO handle);
extern int iggMetricsRenderWindows(IggIO handle);
extern int iggMetricsActiveWindows(IggIO handle);
extern int iggMetricsActiveAllocations(IggIO handle);
extern void iggMouseDelta(IggIO handle, IggVec2 *value);

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
extern void iggIoKeyCtrl(IggIO handle, int leftCtrl, int rightCtrl);
extern void iggIoKeyShift(IggIO handle, int leftShift, int rightShift);
extern void iggIoKeyAlt(IggIO handle, int leftAlt, int rightAlt);
extern void iggIoKeySuper(IggIO handle, int leftSuper, int rightSuper);
extern void iggIoAddInputCharactersUTF8(IggIO handle, char const *utf8Chars);
extern void iggIoSetIniFilename(IggIO handle, char const *value);
extern void iggIoSetConfigFlags(IggIO handle, int flags);
extern void iggIoSetBackendFlags(IggIO handle, int flags);

extern void iggIoRegisterClipboardFunctions(IggIO handle);
extern void iggIoClearClipboardFunctions(IggIO handle);

#ifdef __cplusplus
}
#endif
