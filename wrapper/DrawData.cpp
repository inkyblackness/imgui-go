#include "ConfiguredImGui.h"

#include "DrawData.h"
#include "WrapperConverter.h"

IggDrawData iggGetDrawData()
{
   return reinterpret_cast<IggDrawData>(ImGui::GetDrawData());
}

IggBool iggDrawDataValid(IggDrawData handle)
{
   ImDrawData *drawData = reinterpret_cast<ImDrawData *>(handle);
   IggBool result = 0;
   exportValue(result, drawData->Valid);
   return result;
}

void iggDrawDataGetCommandLists(IggDrawData handle, void **handles, int *count)
{
   ImDrawData *drawData = reinterpret_cast<ImDrawData *>(handle);
   *handles = reinterpret_cast<void **>(drawData->CmdLists);
   *count = drawData->CmdListsCount;
}

void iggDrawDataDisplayPos(IggDrawData handle, IggVec2 *value)
{
   ImDrawData *drawData = reinterpret_cast<ImDrawData *>(handle);
   exportValue(*value, drawData->DisplayPos);
}

void iggDrawDataDisplaySize(IggDrawData handle, IggVec2 *value)
{
   ImDrawData *drawData = reinterpret_cast<ImDrawData *>(handle);
   exportValue(*value, drawData->DisplaySize);
}

void iggDrawDataFrameBufferScale(IggDrawData handle, IggVec2 *value)
{
   ImDrawData *drawData = reinterpret_cast<ImDrawData *>(handle);
   exportValue(*value, drawData->FramebufferScale);
}

void iggDrawDataScaleClipRects(IggDrawData handle, IggVec2 const *scale)
{
   ImDrawData *drawData = reinterpret_cast<ImDrawData *>(handle);
   Vec2Wrapper wrappedScale(scale);
   drawData->ScaleClipRects(*wrappedScale);
}
