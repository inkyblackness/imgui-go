#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern IggGlyphRanges iggGetGlyphRangesDefault(IggFontAtlas handle);
extern IggGlyphRanges iggGetGlyphRangesKorean(IggFontAtlas handle);
extern IggGlyphRanges iggGetGlyphRangesJapanese(IggFontAtlas handle);
extern IggGlyphRanges iggGetGlyphRangesChinese(IggFontAtlas handle);
extern IggGlyphRanges iggGetGlyphRangesCyrillic(IggFontAtlas handle);
extern IggGlyphRanges iggGetGlyphRangesThai(IggFontAtlas handle);

extern IggFont iggAddFontDefault(IggFontAtlas handle);
extern IggFont iggAddFontFromFileTTF(IggFontAtlas handle, char const *filename, float sizePixels,
		IggFontConfig config, IggGlyphRanges glyphRanges);

extern void iggFontAtlasSetTexDesiredWidth(IggFontAtlas handle, int value);

extern void iggFontAtlasGetTexDataAsAlpha8(IggFontAtlas handle, unsigned char **pixels,
      int *width, int *height, int *bytesPerPixel);
extern void iggFontAtlasSetTextureID(IggFontAtlas handle, IggTextureID id);

#ifdef __cplusplus
}
#endif
