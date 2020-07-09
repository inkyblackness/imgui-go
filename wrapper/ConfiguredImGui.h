#pragma once

// As recommended in imgui/imconfig.h, the IMGUI_USER_CONFIG must be specified
// whenever imgui is used. As such, the wrapper code should never include imgui.h directly.

#define IMGUI_USER_CONFIG "wrapper/ConfigOverride.h"
#include "imgui.h"
#include "imgui_internal.h"
