#include "ConfiguredImGui.h"

#include "Color.h"

IggBool iggColorEdit3(char const *label, float *col, int flags)
{
   return ImGui::ColorEdit3(label, col, flags) ? 1 : 0;
}

IggBool iggColorEdit4(char const *label, float *col, int flags)
{
   return ImGui::ColorEdit4(label, col, flags) ? 1 : 0;
}

IggBool iggColorPicker3(char const *label, float *col, int flags)
{
   return ImGui::ColorPicker3(label, col, flags) ? 1 : 0;
}

IggBool iggColorPicker4(char const *label, float *col, int flags)
{
   return ImGui::ColorPicker4(label, col, flags) ? 1 : 0;
}