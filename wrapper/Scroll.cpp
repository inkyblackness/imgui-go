#include "ConfiguredImGui.h"

#include "Scroll.h"

float iggGetScrollX()
{
   return ImGui::GetScrollX();
}

float iggGetScrollY()
{
   return ImGui::GetScrollY();
}

float iggGetScrollMaxX()
{
   return ImGui::GetScrollMaxX();
}

float iggGetScrollMaxY()
{
   return ImGui::GetScrollMaxY();
}

void iggSetScrollHereX(float centerXRatio)
{
   ImGui::SetScrollHereX(centerXRatio);
}

void iggSetScrollHereY(float centerYRatio)
{
   ImGui::SetScrollHereY(centerYRatio);
}

void iggSetScrollX(float scrollX)
{
   ImGui::SetScrollX(scrollX);
}

void iggSetScrollY(float scrollY)
{
   ImGui::SetScrollY(scrollY);
}