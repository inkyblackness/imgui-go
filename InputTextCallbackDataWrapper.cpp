#include "imguiWrappedHeader.h"
#include "InputTextCallbackDataWrapper.h"

int iggInputTextCallbackDataGetEventFlags(IggInputTextCallbackData handle)
{
   ImGuiInputTextCallbackData *data = reinterpret_cast<ImGuiInputTextCallbackData *>(handle);
   return data->EventFlag;
}
