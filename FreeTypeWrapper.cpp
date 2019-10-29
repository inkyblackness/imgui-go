#include "imguiWrappedHeader.h"
#include "imgui_freetype.h"
#include "FreeTypeWrapper.h"

int iggImGuiFreeTypeBuildFontAtlas(IggFontAtlas handle, unsigned int flags)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   return ImGuiFreeType::BuildFontAtlas(fontAtlas, flags);
}
