#include "ConfiguredImGui.h"

#include "DragDrop.h"

IggBool iggBeginDragDropSource(int flags)
{
   return ImGui::BeginDragDropSource(flags) ? 1 : 0;
}

IggBool iggSetDragDropPayload(const char *type, const void *data, int size, int cond)
{
   return ImGui::SetDragDropPayload(type, data, size, cond) ? 1 : 0;
}

void iggEndDragDropSource()
{
   ImGui::EndDragDropSource();
}

IggBool iggBeginDragDropTarget()
{
   return ImGui::BeginDragDropTarget() ? 1 : 0;
}

void *iggPayloadData(IggPayload payload)
{
   const ImGuiPayload *p = reinterpret_cast<const ImGuiPayload *>(payload);
   return p->Data;
}

int iggPayloadDataSize(const IggPayload payload)
{
   const ImGuiPayload *p = reinterpret_cast<const ImGuiPayload *>(payload);
   return p->DataSize;
}

const IggPayload iggAcceptDragDropPayload(const char *type, int flags)
{
   const ImGuiPayload *payload = ImGui::AcceptDragDropPayload(type, flags);
   const IggPayload res = reinterpret_cast<IggPayload>(const_cast<ImGuiPayload *>(payload));
   return res;
}

void iggEndDragDropTarget()
{
   ImGui::EndDragDropTarget();
}