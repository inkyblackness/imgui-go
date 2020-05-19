#include "ConfiguredImGui.h"

#include "Window.h"
#include "WrapperConverter.h"

void iggShowDemoWindow(IggBool *open)
{
   BoolWrapper openArg(open);

   ImGui::ShowDemoWindow(openArg);
}

void iggShowUserGuide(void)
{
   ImGui::ShowUserGuide();
}

IggBool iggBegin(char const *id, IggBool *open, int flags)
{
   BoolWrapper openArg(open);
   return ImGui::Begin(id, openArg, flags) ? 1 : 0;
}

void iggEnd(void)
{
   ImGui::End();
}

IggBool iggBeginChild(char const *id, IggVec2 const *size, IggBool border, int flags)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::BeginChild(id, *sizeArg, border, flags) ? 1 : 0;
}

void iggEndChild(void)
{
   ImGui::EndChild();
}

void iggWindowPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetWindowPos());
}

void iggWindowSize(IggVec2 *size)
{
   exportValue(*size, ImGui::GetWindowSize());
}

float iggWindowWidth(void)
{
   return ImGui::GetWindowWidth();
}

float iggWindowHeight(void)
{
   return ImGui::GetWindowHeight();
}

void iggContentRegionAvail(IggVec2 *size)
{
   exportValue(*size, ImGui::GetContentRegionAvail());
}

void iggGetContentRegionMax(IggVec2 *out)
{
   ImVec2 im_out = ImGui::GetContentRegionMax();
   exportValue(*out, im_out);
}

void iggSetNextWindowPos(IggVec2 const *pos, int cond, IggVec2 const *pivot)
{
   Vec2Wrapper posArg(pos);
   Vec2Wrapper pivotArg(pivot);
   ImGui::SetNextWindowPos(*posArg, cond, *pivotArg);
}

void iggSetNextWindowCollapsed(IggBool collapsed, int cond)
{
   ImGui::SetNextWindowCollapsed(collapsed, cond);
}

void iggSetNextWindowSize(IggVec2 const *size, int cond)
{
   Vec2Wrapper sizeArg(size);
   ImGui::SetNextWindowSize(*sizeArg, cond);
}

void iggSetNextWindowSizeConstraints(const IggVec2 *size_min, const IggVec2 *size_max)
{
   Vec2Wrapper sizeMinArg(size_min);
   Vec2Wrapper sizeMaxArg(size_max);
   ImGui::SetNextWindowSizeConstraints(*sizeMinArg, *sizeMaxArg);
}

void iggSetNextWindowContentSize(IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   ImGui::SetNextWindowContentSize(*sizeArg);
}

void iggSetNextWindowFocus(void)
{
   ImGui::SetNextWindowFocus();
}

void iggSetNextWindowBgAlpha(float value)
{
   ImGui::SetNextWindowBgAlpha(value);
}

void iggPushItemWidth(float width)
{
   ImGui::PushItemWidth(width);
}

void iggPopItemWidth(void)
{
   ImGui::PopItemWidth();
}

float iggCalcItemWidth(void)
{
   return ImGui::CalcItemWidth();
}

void iggPushTextWrapPos(float wrapPosX)
{
   ImGui::PushTextWrapPos(wrapPosX);
}

void iggPopTextWrapPos(void)
{
   ImGui::PopTextWrapPos();
}
