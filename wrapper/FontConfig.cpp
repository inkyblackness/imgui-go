#include "ConfiguredImGui.h"

#include "FontConfig.h"

IggFontConfig iggNewFontConfig()
{
   ImFontConfig *fontConfig = new ImFontConfig();
   return static_cast<IggFontConfig>(fontConfig);
}

void iggFontConfigDelete(IggFontConfig handle)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   delete fontConfig;
}

void iggFontConfigSetSize(IggFontConfig handle, float sizePixels)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   fontConfig->SizePixels = sizePixels;
}

void iggFontConfigSetOversampleH(IggFontConfig handle, int value)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   fontConfig->OversampleH = value;
}

void iggFontConfigSetOversampleV(IggFontConfig handle, int value)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   fontConfig->OversampleV = value;
}

void iggFontConfigSetPixelSnapH(IggFontConfig handle, IggBool value)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   fontConfig->PixelSnapH = value;
}

void iggFontConfigSetGlyphMinAdvanceX(IggFontConfig handle, float value)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   fontConfig->GlyphMinAdvanceX = value;
}

void iggFontConfigSetGlyphMaxAdvanceX(IggFontConfig handle, float value)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   fontConfig->GlyphMaxAdvanceX = value;
}

void iggFontConfigSetGlyphOffsetX(IggFontConfig handle, float value)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   fontConfig->GlyphOffset.x = value;
}

void iggFontConfigSetGlyphOffsetY(IggFontConfig handle, float value)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   fontConfig->GlyphOffset.y = value;
}

void iggFontConfigSetMergeMode(IggFontConfig handle, IggBool value)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   fontConfig->MergeMode = value;
}

void iggFontConfigSetName(IggFontConfig handle, char const *value)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   const size_t bufSize = sizeof(fontConfig->Name);
   strncpy(fontConfig->Name, value, bufSize - 1);
   fontConfig->Name[bufSize - 1] = '\0';
}

int iggFontConfigGetFontDataOwnedByAtlas(IggFontConfig handle)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   return fontConfig->FontDataOwnedByAtlas;
}

unsigned int iggFontConfigGetFontBuilderFlags(IggFontConfig handle)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   return fontConfig->FontBuilderFlags;
}

void iggFontConfigSetFontBuilderFlags(IggFontConfig handle, unsigned int flags)
{
   ImFontConfig *fontConfig = reinterpret_cast<ImFontConfig *>(handle);
   fontConfig->FontBuilderFlags = flags;
}
