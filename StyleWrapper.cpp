#include "imguiWrappedHeader.h"
#include "StyleWrapper.h"

void iggStyleScaleAllSizes(IggGuiStyle handle, float scale)
{
   ImGuiStyle *style = reinterpret_cast<ImGuiStyle *>(handle);
   style->ScaleAllSizes(scale);
}
