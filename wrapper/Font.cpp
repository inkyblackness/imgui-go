#include "ConfiguredImGui.h"

#include "Font.h"

void iggPushFont(IggFont handle)
{
   ImFont *font = reinterpret_cast<ImFont *>(handle);
   ImGui::PushFont(font);
}

void iggPopFont(void)
{
   ImGui::PopFont();
}

float iggGetFontSize()
{
   return ImGui::GetFontSize();
}

void iggCalcTextSize(const char *text, int length, IggBool hide_text_after_double_hash, float wrap_width, IggVec2 *value)
{
   exportValue(*value, ImGui::CalcTextSize(text, text + length, hide_text_after_double_hash, wrap_width));
}
