#include "ConfiguredImGui.h"

#include "Color.h"
#include "WrapperConverter.h"

IggBool iggColorEdit3(char const *label, float *col, int flags)
{
   return ImGui::ColorEdit3(label, col, flags) ? 1 : 0;
}

IggBool iggColorEdit4(char const *label, float *col, int flags)
{
   return ImGui::ColorEdit4(label, col, flags) ? 1 : 0;
}

IggBool iggColorButton(char const *label, IggVec4 const *col, int flags, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   Vec4Wrapper colArg(col);
   bool ret = ImGui::ColorButton(label, *colArg, flags, *sizeArg) ? 1 : 0;
   return ret;
}

IggBool iggColorPicker3(char const *label, float *col, int flags)
{
   return ImGui::ColorPicker3(label, col, flags) ? 1 : 0;
}

IggBool iggColorPicker4(char const *label, float *col, int flags)
{
   return ImGui::ColorPicker4(label, col, flags) ? 1 : 0;
}