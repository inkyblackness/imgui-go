#include "ConfiguredImGui.h"

#include "Popup.h"
#include "WrapperConverter.h"

IggBool iggBeginPopup(const char *name, int flags)
{
   return ImGui::BeginPopup(name, flags) ? 1 : 0;
}

IggBool iggBeginPopupModal(const char *name, IggBool *open, int flags)
{
   BoolWrapper openArg(open);
   return ImGui::BeginPopupModal(name, openArg, flags) ? 1 : 0;
}

void iggEndPopup(void)
{
   ImGui::EndPopup();
}

void iggOpenPopup(const char *id, int flags)
{
   ImGui::OpenPopup(id, flags);
}

void iggOpenPopupOnItemClick(const char *id, int flags)
{
   ImGui::OpenPopupOnItemClick(id, flags);
}

void iggCloseCurrentPopup(void)
{
   ImGui::CloseCurrentPopup();
}

IggBool iggBeginPopupContextItem(const char *id, int flags)
{
   return ImGui::BeginPopupContextItem(id, flags) ? 1 : 0;
}

IggBool iggBeginPopupContextWindow(const char *id, int flags)
{
   return ImGui::BeginPopupContextWindow(id, flags) ? 1 : 0;
}

IggBool iggBeginPopupContextVoid(const char *id, int flags)
{
   return ImGui::BeginPopupContextVoid(id, flags) ? 1 : 0;
}

IggBool iggIsPopupOpen(const char *id, int flags)
{
   return ImGui::IsPopupOpen(id, flags) ? 1 : 0;
}
