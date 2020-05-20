#pragma once

// Overrides as per standard imconfig.h

extern "C" void iggAssert(int result, char const *expression, char const *file, int line);
#define IM_ASSERT(_EXPR) iggAssert((_EXPR) != 0, #_EXPR, __FILE__, __LINE__)

#define IMGUI_DISABLE_OBSOLETE_FUNCTIONS
