#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern void iggFontAtlasGetTexDataAsAlpha8(IggFontAtlas handle, unsigned char **pixels,
      int *width, int *height, int *bytesPerPixel);
extern void iggFontAtlasSetTextureID(IggFontAtlas handle, IggTextureID id);

#ifdef __cplusplus
}
#endif
