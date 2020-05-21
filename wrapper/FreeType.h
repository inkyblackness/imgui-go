#pragma once

#include "Types.h"

#ifdef IMGUI_FREETYPE_ENABLED

#ifdef __cplusplus
extern "C" {
#endif

extern int iggFreeTypeBuildFontAtlas(IggFontAtlas handle, unsigned int flags);

#ifdef __cplusplus
}
#endif

#endif
