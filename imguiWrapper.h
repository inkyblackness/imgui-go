#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern IggContext iggCreateContext(IggFontAtlas sharedFontAtlas);
extern void iggDestroyContext(IggContext context);
extern IggContext iggGetCurrentContext();
extern void iggSetCurrentContext(IggContext context);

extern IggIO iggGetCurrentIO(void);
extern IggGuiStyle iggGetCurrentStyle(void);
extern void iggNewFrame(void);
extern void iggRender(void);
extern IggDrawData iggGetDrawData(void);
extern void iggEndFrame(void);

extern char const *iggGetVersion(void);
extern void iggShowDemoWindow(IggBool *open);
extern void iggShowUserGuide(void);

extern IggBool iggBegin(char const *id, IggBool *open, int flags);
extern void iggEnd(void);
extern IggBool iggBeginChild(char const *id, IggVec2 const *size, IggBool border, int flags);
extern void iggEndChild(void);

extern void iggSetNextWindowPos(IggVec2 const *pos, int cond, IggVec2 const *pivot);
extern void iggSetNextWindowSize(IggVec2 const *size, int cond);

extern void iggTextUnformatted(char const *text);

extern IggBool iggButton(char const *label, IggVec2 const *size);
extern IggBool iggCheckbox(char const *label, IggBool *selected);

extern void iggSeparator(void);
extern void iggSameLine(float posX, float spacingW);
extern void iggBeginGroup(void);
extern void iggEndGroup(void);

extern IggBool iggSelectable(char const *label, IggBool selected, int flags, IggVec2 const *size);

extern IggBool iggBeginMainMenuBar(void);
extern void iggEndMainMenuBar(void);
extern IggBool iggBeginMenuBar(void);
extern void iggEndMenuBar(void);
extern IggBool iggBeginMenu(char const *label, IggBool enabled);
extern void iggEndMenu(void);
extern IggBool iggMenuItem(char const *label, char const *shortcut, IggBool selected, IggBool enabled);

#ifdef __cplusplus
}
#endif
