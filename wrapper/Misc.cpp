#include "ConfiguredImGui.h"

#include "Misc.h"

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


