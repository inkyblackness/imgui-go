#define IMGUI_USER_CONFIG "wrapper/imconfig.h"

// imgui code
// imgui-<version> is added to include path in all.go
#include "imgui.cpp"
#include "imgui_draw.cpp"
#include "imgui_demo.cpp"
#include "imgui_widgets.cpp"
#ifdef IMGUI_FREETYPE_ENABLED
#include "misc/freetype/imgui_freetype.cpp"
#endif

// imgui-go code
#include "wrapper/DragDropWrapper.cpp"
#include "wrapper/DrawCommandWrapper.cpp"
#include "wrapper/DrawDataWrapper.cpp"
#include "wrapper/DrawListWrapper.cpp"
#include "wrapper/FontAtlasWrapper.cpp"
#include "wrapper/FontConfigWrapper.cpp"
#include "wrapper/imguiWrapper.cpp"
#include "wrapper/InputTextCallbackDataWrapper.cpp"
#include "wrapper/IOWrapper.cpp"
#include "wrapper/ListClipper.cpp"
#include "wrapper/StyleWrapper.cpp"
#include "wrapper/UtilsWrapper.cpp"
#include "wrapper/WrapperConverter.cpp"
#ifdef IMGUI_FREETYPE_ENABLED
#include "wrapper/FreeTypeWrapper.cpp"
#endif
