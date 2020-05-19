#ifdef IMGUI_FREETYPE_ENABLED

#include "ConfiguredImGui.h"

#include "misc/freetype/imgui_freetype.h"

#include "FreeType.h"

int iggFreeTypeBuildFontAtlas(IggFontAtlas handle, unsigned int flags)
{
   ImFontAtlas *fontAtlas = reinterpret_cast<ImFontAtlas *>(handle);
   return ImGuiFreeType::BuildFontAtlas(fontAtlas, flags);
}

#endif
