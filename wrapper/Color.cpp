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

IggBool iggColorButton(char const *label, float *col, int flags, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   ImVec4 v4col;
   v4col.x = col[0];
   v4col.y = col[1];
   v4col.z = col[2];
   v4col.w = col[3];
   bool ret = ImGui::ColorButton(label, v4col, flags, *sizeArg) ? 1 : 0;
   col[0] = v4col.x;
   col[1] = v4col.y;
   col[2] = v4col.z;
   col[3] = v4col.w;
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