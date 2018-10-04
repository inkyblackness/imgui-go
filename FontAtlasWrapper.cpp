#include "imguiWrappedHeader.h"
#include "FontAtlasWrapper.h"
#include "WrapperConverter.h"

IggGlyphRanges iggGetGlyphRangesDefault(IggFontAtlas handle)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   return static_cast<IggGlyphRanges>(const_cast<ImWchar *>(fontAtlas->GetGlyphRangesDefault()));
}

IggGlyphRanges iggGetGlyphRangesKorean(IggFontAtlas handle)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   return static_cast<IggGlyphRanges>(const_cast<ImWchar *>(fontAtlas->GetGlyphRangesKorean()));
}

IggGlyphRanges iggGetGlyphRangesJapanese(IggFontAtlas handle)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   return static_cast<IggGlyphRanges>(const_cast<ImWchar *>(fontAtlas->GetGlyphRangesJapanese()));
}

IggGlyphRanges iggGetGlyphRangesChinese(IggFontAtlas handle)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   return static_cast<IggGlyphRanges>(const_cast<ImWchar *>(fontAtlas->GetGlyphRangesChinese()));
}

IggGlyphRanges iggGetGlyphRangesCyrillic(IggFontAtlas handle)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   return static_cast<IggGlyphRanges>(const_cast<ImWchar *>(fontAtlas->GetGlyphRangesCyrillic()));
}

IggGlyphRanges iggGetGlyphRangesThai(IggFontAtlas handle)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   return static_cast<IggGlyphRanges>(const_cast<ImWchar *>(fontAtlas->GetGlyphRangesThai()));
}

IggFont iggAddFontDefault(IggFontAtlas handle)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   ImFont *font = fontAtlas->AddFontDefault();
   return static_cast<IggFont>(font);
}

IggFont iggAddFontFromFileTTF(IggFontAtlas handle, char const *filename, float sizePixels,
      IggFontConfig config, IggGlyphRanges glyphRanges)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(config);
   ImWchar *glyphChars = reinterpret_cast<ImWchar *>(glyphRanges);
   ImFont *font = fontAtlas->AddFontFromFileTTF(filename, sizePixels, fontConfig, glyphChars);
   return static_cast<IggFont>(font);
}

void iggFontAtlasSetTexDesiredWidth(IggFontAtlas handle, int value)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   fontAtlas->TexDesiredWidth = value;
}

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
