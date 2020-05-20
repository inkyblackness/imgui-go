#include "ConfiguredImGui.h"

#include "Focus.h"

void iggSetItemDefaultFocus()
{
   ImGui::SetItemDefaultFocus();
}

IggBool iggIsItemFocused()
{
   return ImGui::IsItemFocused();
}

IggBool iggIsAnyItemFocused()
{
   return ImGui::IsAnyItemFocused();
}

void iggSetKeyboardFocusHere(int offset)
{
   ImGui::SetKeyboardFocusHere(offset);
}
