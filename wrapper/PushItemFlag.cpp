#include "ConfiguredImGui.h"

#include "PushItemFlag.h"

void iggPushItemFlag(int option)
{
   ImGui::PushItemFlag(option, true);
}

void iggPopItemFlag()
{
   ImGui::PopItemFlag();
}

