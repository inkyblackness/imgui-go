#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggShowDemoWindow(IggBool *open);
extern void iggShowUserGuide(void);

extern IggBool iggBegin(char const *id, IggBool *open, int flags);
extern void iggEnd(void);
extern IggBool iggBeginChild(char const *id, IggVec2 const *size, IggBool border, int flags);
extern void iggEndChild(void);

extern void iggWindowPos(IggVec2 *pos);
extern void iggWindowSize(IggVec2 *size);
extern float iggWindowWidth(void);
extern float iggWindowHeight(void);
extern void iggContentRegionAvail(IggVec2 *size);
extern void iggGetContentRegionMax(IggVec2 *out);

extern void iggSetNextWindowPos(IggVec2 const *pos, int cond, IggVec2 const *pivot);
extern void iggSetNextWindowSize(IggVec2 const *size, int cond);
extern void iggSetNextWindowCollapsed(IggBool collapsed, int cond);
extern void iggSetNextWindowSizeConstraints(const IggVec2 *size_min, const IggVec2 *size_max);
extern void iggSetNextWindowContentSize(IggVec2 const *size);
extern void iggSetNextWindowFocus(void);
extern void iggSetNextWindowBgAlpha(float value);

extern void iggPushItemWidth(float width);
extern void iggPopItemWidth(void);
extern float iggCalcItemWidth(void);
extern void iggPushTextWrapPos(float wrapPosX);
extern void iggPopTextWrapPos(void);

#ifdef __cplusplus
}
#endif
