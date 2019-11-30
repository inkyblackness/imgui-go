#include "imguiWrappedHeader.h"
#include "imguiWrapperTypes.h"
#include "WrapperConverter.h"

void importValue(bool &out, IggBool const &in)
{
   out = in != 0;
}

void exportValue(IggBool &out, bool const &in)
{
   out = in ? 1 : 0;
}

void importValue(ImVec2 &out, IggVec2 const &in)
{
   out.x = in.x;
   out.y = in.y;
}

void exportValue(IggVec2 &out, ImVec2 const &in)
{
   out.x = in.x;
   out.y = in.y;
}

void importValue(ImVec4 &out, IggVec4 const &in)
{
   out.x = in.x;
   out.y = in.y;
   out.z = in.z;
   out.w = in.w;
}

void exportValue(IggVec4 &out, ImVec4 const &in)
{
   out.x = in.x;
   out.y = in.y;
   out.z = in.z;
   out.w = in.w;
}

void importValue(ImGuiListClipper &out, IggListClipper const &in)
{
    out.StartPosY = in.StartPosY;
    out.ItemsHeight = in.ItemsHeight;
    out.ItemsCount = in.ItemsCount;
    out.StepNo = in.StepNo;
    out.DisplayStart = in.DisplayStart;
    out.DisplayEnd = in.DisplayEnd;
}

void exportValue(IggListClipper &out, ImGuiListClipper const &in)
{
    out.StartPosY = in.StartPosY;
    out.ItemsHeight = in.ItemsHeight;
    out.ItemsCount = in.ItemsCount;
    out.StepNo = in.StepNo;
    out.DisplayStart = in.DisplayStart;
    out.DisplayEnd = in.DisplayEnd;
}
