#include "ConfiguredImGui.h"

#include "State.h"
#include "WrapperConverter.h"

void iggClearActiveID(void)
{
   ImGui::ClearActiveID();
}

IggBool iggIsItemClicked()
{
   return ImGui::IsItemClicked() ? 1 : 0;
}

IggBool iggIsItemHovered(int flags)
{
   return ImGui::IsItemHovered(flags) ? 1 : 0;
}

IggBool iggIsItemActive()
{
   return ImGui::IsItemActive() ? 1 : 0;
}

IggBool iggIsAnyItemActive()
{
   return ImGui::IsAnyItemActive() ? 1 : 0;
}

IggBool iggIsItemVisible()
{
   return ImGui::IsItemVisible() ? 1 : 0;
}

IggBool iggIsItemEdited()
{
   return ImGui::IsItemEdited() ? 1 : 0;
}

IggBool iggIsItemActivated()
{
   return ImGui::IsItemActivated() ? 1 : 0;
}

IggBool iggIsItemDeactivated()
{
   return ImGui::IsItemDeactivated() ? 1 : 0;
}

IggBool iggIsItemDeactivatedAfterEdit()
{
   return ImGui::IsItemDeactivatedAfterEdit() ? 1 : 0;
}

IggBool iggIsItemToggledOpen()
{
   return ImGui::IsItemToggledOpen() ? 1 : 0;
}

void iggSetItemAllowOverlap()
{
   ImGui::SetItemAllowOverlap();
}

IggBool iggIsWindowAppearing()
{
   return ImGui::IsWindowAppearing() ? 1 : 0;
}

IggBool iggIsWindowCollapsed()
{
   return ImGui::IsWindowCollapsed() ? 1 : 0;
}

IggBool iggIsWindowFocused(int flags)
{
   return ImGui::IsWindowFocused(flags) ? 1 : 0;
}

IggBool iggIsWindowHovered(int flags)
{
   return ImGui::IsWindowHovered(flags) ? 1 : 0;
}

IggBool iggIsKeyDown(int key)
{
   return ImGui::IsKeyDown(key);
}

IggBool iggIsKeyPressed(int key, IggBool repeat)
{
   return ImGui::IsKeyPressed(key, repeat);
}

IggBool iggIsKeyReleased(int key)
{
   return ImGui::IsKeyReleased(key);
}

IggBool iggIsMouseDown(int button)
{
   return ImGui::IsMouseDown(button);
}

IggBool iggIsAnyMouseDown()
{
   return ImGui::IsAnyMouseDown();
}

IggBool iggIsMouseClicked(int button, IggBool repeat)
{
   return ImGui::IsMouseClicked(button, repeat);
}

IggBool iggIsMouseReleased(int button)
{
   return ImGui::IsMouseReleased(button);
}

IggBool iggIsMouseDoubleClicked(int button)
{
   return ImGui::IsMouseDoubleClicked(button);
}

IggBool iggIsMouseDragging(int button, float lock_threshold)
{
   return ImGui::IsMouseDragging(button, lock_threshold);
}

void iggGetMouseDragDelta(IggVec2 *value, int button, float lock_threshold)
{
   exportValue(*value, ImGui::GetMouseDragDelta(button, lock_threshold));
}

void iggResetMouseDragDelta(int button)
{
   ImGui::ResetMouseDragDelta(button);
}

void iggMousePos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetMousePos());
}

int iggGetMouseCursor()
{
   return ImGui::GetMouseCursor();
}

void iggSetMouseCursor(int cursor)
{
   ImGui::SetMouseCursor(cursor);
}

void iggGetItemRectMin(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetItemRectMin());
}

void iggGetItemRectMax(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetItemRectMax());
}
