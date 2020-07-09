#include "ConfiguredImGui.h"

void iggPushItemFlag(int option)
{
   ImGui::PushItemFlag(option, true);
}

void iggPopItemFlag()
{
   ImGui::PopItemFlag();
}

