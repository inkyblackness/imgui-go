#include "ConfiguredImGui.h"

#include "State.h"
#include "WrapperConverter.h"

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
