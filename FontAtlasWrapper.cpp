#include "imguiWrappedHeader.h"
#include "FontAtlasWrapper.h"
#include "WrapperConverter.h"

void iggFontAtlasGetTexDataAsAlpha8(IggFontAtlas handle, unsigned char **pixels,
                                    int *width, int *height, int *bytesPerPixel)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   fontAtlas->GetTexDataAsAlpha8(pixels, width, height, bytesPerPixel);
}

void iggFontAtlasSetTextureID(IggFontAtlas handle, IggTextureID id)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   fontAtlas->SetTexID(id);
}
