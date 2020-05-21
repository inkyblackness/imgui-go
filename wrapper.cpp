#include "wrapper/ConfiguredImGui.h"

// imgui code
// imgui/ is added to include path in wrapper.go
#include "imgui.cpp"
#include "imgui_draw.cpp"
#include "imgui_demo.cpp"
#include "imgui_widgets.cpp"
#ifdef IMGUI_FREETYPE_ENABLED
#include "misc/freetype/imgui_freetype.cpp"
#endif

// imgui-go code
#include "wrapper/Context.cpp"
#include "wrapper/Focus.cpp"
#include "wrapper/DragDrop.cpp"
#include "wrapper/DrawCommand.cpp"
#include "wrapper/DrawData.cpp"
#include "wrapper/DrawList.cpp"
#include "wrapper/Font.cpp"
#include "wrapper/FontAtlas.cpp"
#include "wrapper/FontConfig.cpp"
#include "wrapper/InputTextCallbackData.cpp"
#include "wrapper/IO.cpp"
#include "wrapper/Layout.cpp"
#include "wrapper/ListClipper.cpp"
#include "wrapper/Main.cpp"
#include "wrapper/Popup.cpp"
#include "wrapper/Scroll.cpp"
#include "wrapper/State.cpp"
#include "wrapper/Style.cpp"
#include "wrapper/Widgets.cpp"
#include "wrapper/Window.cpp"
#include "wrapper/WrapperConverter.cpp"
#ifdef IMGUI_FREETYPE_ENABLED
#include "wrapper/FreeType.cpp"
#endif
