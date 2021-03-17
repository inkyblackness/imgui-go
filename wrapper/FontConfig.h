#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern IggFontConfig iggNewFontConfig();
extern void iggFontConfigDelete(IggFontConfig handle);

extern void iggFontConfigSetSize(IggFontConfig handle, float sizePixels);
extern void iggFontConfigSetOversampleH(IggFontConfig handle, int value);
extern void iggFontConfigSetOversampleV(IggFontConfig handle, int value);
extern void iggFontConfigSetPixelSnapH(IggFontConfig handle, IggBool value);
extern void iggFontConfigSetGlyphMinAdvanceX(IggFontConfig handle, float value);
extern void iggFontConfigSetGlyphMaxAdvanceX(IggFontConfig handle, float value);
extern void iggFontConfigSetGlyphOffsetX(IggFontConfig handle, float value);
extern void iggFontConfigSetGlyphOffsetY(IggFontConfig handle, float value);
extern void iggFontConfigSetMergeMode(IggFontConfig handle, IggBool value);
extern void iggFontConfigSetName(IggFontConfig handle, char const *value);
extern int iggFontConfigGetFontDataOwnedByAtlas(IggFontConfig handle);

extern unsigned int iggFontConfigGetFontBuilderFlags(IggFontConfig handle);
extern void         iggFontConfigSetFontBuilderFlags(IggFontConfig handle, unsigned int flags);

#ifdef __cplusplus
}
#endif
