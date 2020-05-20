#include "ConfiguredImGui.h"

#include "Context.h"

IggContext iggCreateContext(IggFontAtlas sharedFontAtlas)
{
   ImGuiContext *context = ImGui::CreateContext(reinterpret_cast<ImFontAtlas *>(sharedFontAtlas));
   return reinterpret_cast<IggContext>(context);
}

void iggDestroyContext(IggContext context)
{
   ImGui::DestroyContext(reinterpret_cast<ImGuiContext *>(context));
}

IggContext iggGetCurrentContext()
{
   return reinterpret_cast<IggContext>(ImGui::GetCurrentContext());
}

void iggSetCurrentContext(IggContext context)
{
   ImGui::SetCurrentContext(reinterpret_cast<ImGuiContext *>(context));
}
