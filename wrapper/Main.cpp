#include "ConfiguredImGui.h"

#include "Main.h"

char const *iggGetVersion()
{
   return ImGui::GetVersion();
}

void iggNewFrame()
{
   ImGui::NewFrame();
}

void iggRender()
{
   ImGui::Render();
}

void iggEndFrame()
{
   ImGui::EndFrame();
}
