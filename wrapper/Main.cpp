#include "ConfiguredImGui.h"

#include "Main.h"

char const *iggGetVersion()
{
   return ImGui::GetVersion();
}

double iggGetTime()
{
   return ImGui::GetTime();
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
