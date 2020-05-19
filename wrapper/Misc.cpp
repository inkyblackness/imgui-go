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

void iggPushID(char const *id)
{
   ImGui::PushID(id);
}

void iggPushIDInt(int id)
{
   ImGui::PushID(id);
}

void iggPopID(void)
{
   ImGui::PopID();
}

int iggGetMouseCursor()
{
   return ImGui::GetMouseCursor();
}

void iggSetMouseCursor(int cursor)
{
   ImGui::SetMouseCursor(cursor);
}

