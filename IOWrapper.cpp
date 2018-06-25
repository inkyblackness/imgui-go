#include "imguiWrappedHeader.h"
#include "IOWrapper.h"
#include "WrapperConverter.h"

IggBool iggWantCaptureMouse(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return io->WantCaptureMouse ? 1 : 0;
}

IggFontAtlas iggIoGetFonts(IggIO handle)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   return reinterpret_cast<IggFontAtlas>(io->Fonts);
}

void iggIoSetDisplaySize(IggIO handle, IggVec2 const *value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   importValue(io->DisplaySize, *value);
}

void iggIoSetMousePosition(IggIO handle, IggVec2 const *value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   importValue(io->MousePos, *value);
}

void iggIoSetMouseButtonDown(IggIO handle, int index, IggBool value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   io->MouseDown[index] = value != 0;
}

void iggIoAddMouseWheelDelta(IggIO handle, float horizontal, float vertical)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   io->MouseWheelH += horizontal;
   io->MouseWheel += vertical;
}

void iggIoSetDeltaTime(IggIO handle, float value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   io->DeltaTime = value;
}

void iggIoSetFontGlobalScale(IggIO handle, float value)
{
   ImGuiIO *io = reinterpret_cast<ImGuiIO *>(handle);
   io->FontGlobalScale = value;
}
