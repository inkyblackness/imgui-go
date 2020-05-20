#include "ConfiguredImGui.h"

#include "Popup.h"
#include "WrapperConverter.h"

void iggOpenPopup(char const *id)
{
   ImGui::OpenPopup(id);
}

IggBool iggBeginPopup(char const *name, int flags)
{
   return ImGui::BeginPopup(name, flags) ? 1 : 0;
}

IggBool iggBeginPopupModal(char const *name, IggBool *open, int flags)
{
   BoolWrapper openArg(open);
   return ImGui::BeginPopupModal(name, openArg, flags) ? 1 : 0;
}

IggBool iggBeginPopupContextItem(char const *label, int mouseButton)
{
   return ImGui::BeginPopupContextItem(label, mouseButton) ? 1 : 0;
}

void iggEndPopup(void)
{
   ImGui::EndPopup();
}

void iggCloseCurrentPopup(void)
{
   ImGui::CloseCurrentPopup();
}