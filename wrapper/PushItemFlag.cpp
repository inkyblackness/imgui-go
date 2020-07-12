#include "ConfiguredImGui.h"

#include "PushItemFlag.h"

#include "imgui_internal.h"

void iggPushItemFlag(int option)
{
   ImGui::PushItemFlag(option, true);
}

void iggPopItemFlag()
{
   ImGui::PopItemFlag();
}

