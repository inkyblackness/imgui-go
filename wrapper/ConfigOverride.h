#pragma once

// Overrides as per standard imconfig.h

extern "C" void iggAssert(char const *expression, char const *file, int line);
#define IM_ASSERT(_EXPR) do { if ((_EXPR) == 0) { iggAssert(#_EXPR, __FILE__, __LINE__); } } while(false)

#define IMGUI_DISABLE_OBSOLETE_FUNCTIONS
