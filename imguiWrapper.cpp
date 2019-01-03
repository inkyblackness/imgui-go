#include "imguiWrappedHeader.h"
#include "imguiWrapper.h"
#include "WrapperConverter.h"

IggContext iggCreateContext(IggFontAtlas sharedFontAtlas)
{
   ImGuiContext *context = ImGui::CreateContext(reinterpret_cast<ImFontAtlas *>(sharedFontAtlas));
   return reinterpret_cast<IggContext>(context);
}

void iggDestroyContext(IggContext context)
{
   ImGui::DestroyContext(reinterpret_cast<ImGuiContext *>(context));
}

IggContext iggGetCurrentContext()
{
   return reinterpret_cast<IggContext>(ImGui::GetCurrentContext());
}

void iggSetCurrentContext(IggContext context)
{
   ImGui::SetCurrentContext(reinterpret_cast<ImGuiContext *>(context));
}

IggIO iggGetCurrentIO()
{
   return reinterpret_cast<IggIO>(&ImGui::GetIO());
}

IggGuiStyle iggGetCurrentStyle()
{
   return reinterpret_cast<IggGuiStyle>(&ImGui::GetStyle());
}

void iggNewFrame()
{
   ImGui::NewFrame();
}

void iggRender()
{
   ImGui::Render();
}

IggDrawData iggGetDrawData()
{
   return reinterpret_cast<IggDrawData>(ImGui::GetDrawData());
}

void iggEndFrame()
{
   ImGui::EndFrame();
}

char const *iggGetVersion()
{
   return ImGui::GetVersion();
}

void iggShowDemoWindow(IggBool *open)
{
   BoolWrapper openArg(open);

   ImGui::ShowDemoWindow(openArg);
}

void iggShowUserGuide(void)
{
   ImGui::ShowUserGuide();
}

IggBool iggBegin(char const *id, IggBool *open, int flags)
{
   BoolWrapper openArg(open);
   return ImGui::Begin(id, openArg, flags) ? 1 : 0;
}

void iggEnd(void)
{
   ImGui::End();
}

IggBool iggBeginChild(char const *id, IggVec2 const *size, IggBool border, int flags)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::BeginChild(id, *sizeArg, border, flags) ? 1 : 0;
}

void iggEndChild(void)
{
   ImGui::EndChild();
}

void iggSetNextWindowPos(IggVec2 const *pos, int cond, IggVec2 const *pivot)
{
   Vec2Wrapper posArg(pos);
   Vec2Wrapper pivotArg(pivot);
   ImGui::SetNextWindowPos(*posArg, cond, *pivotArg);
}

void iggSetNextWindowSize(IggVec2 const *size, int cond)
{
   Vec2Wrapper sizeArg(size);
   ImGui::SetNextWindowSize(*sizeArg, cond);
}

void iggSetNextWindowFocus(void)
{
   ImGui::SetNextWindowFocus();
}

void iggSetNextWindowBgAlpha(float value)
{
   ImGui::SetNextWindowBgAlpha(value);
}

void iggPushFont(IggFont handle)
{
   ImFont *font = reinterpret_cast<ImFont *>(handle);
   ImGui::PushFont(font);
}

void iggPopFont(void)
{
   ImGui::PopFont();
}

void iggPushStyleColor(int index, IggVec4 const *col)
{
   Vec4Wrapper colArg(col);
   ImGui::PushStyleColor(index, *colArg);
}

void iggPopStyleColor(int count)
{
   ImGui::PopStyleColor(count);
}

void iggPushStyleVarFloat(int index, float value)
{
   ImGui::PushStyleVar(index, value);
}

void iggPushStyleVarVec2(int index, IggVec2 const *value)
{
   Vec2Wrapper valueArg(value);
   ImGui::PushStyleVar(index, *valueArg);
}

void iggPopStyleVar(int count)
{
   ImGui::PopStyleVar(count);
}

void iggPushItemWidth(float width)
{
   ImGui::PushItemWidth(width);
}

void iggPopItemWidth(void)
{
   ImGui::PopItemWidth();
}

void iggPushTextWrapPos(float wrapPosX)
{
   ImGui::PushTextWrapPos(wrapPosX);
}

void iggPopTextWrapPos(void)
{
   ImGui::PopTextWrapPos();
}

void iggPushID(char const *id)
{
   ImGui::PushID(id);
}
void iggPopID(void)
{
   ImGui::PopID();
}

void iggTextUnformatted(char const *text)
{
   ImGui::TextUnformatted(text);
}

void iggLabelText(char const *label, char const *text)
{
   ImGui::LabelText(label, "%s", text);
}

IggBool iggButton(char const *label, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::Button(label, *sizeArg) ? 1 : 0;
}

void iggImage(IggTextureID textureID,
   IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
   IggVec4 const *tintCol, IggVec4 const *borderCol)
{
   Vec2Wrapper sizeArg(size);
   Vec2Wrapper uv0Arg(uv0);
   Vec2Wrapper uv1Arg(uv1);
   Vec4Wrapper tintColArg(tintCol);
   Vec4Wrapper borderColArg(borderCol);
   ImGui::Image(static_cast<ImTextureID>(textureID), *sizeArg, *uv0Arg, *uv1Arg, *tintColArg, *borderColArg);
}

IggBool iggImageButton(IggTextureID textureID,
   IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
   int framePadding, IggVec4 const *bgCol,
   IggVec4 const *tintCol)
{
   Vec2Wrapper sizeArg(size);
   Vec2Wrapper uv0Arg(uv0);
   Vec2Wrapper uv1Arg(uv1);
   Vec4Wrapper bgColArg(bgCol);
   Vec4Wrapper tintColArg(tintCol);
   return ImGui::ImageButton(static_cast<ImTextureID>(textureID), *sizeArg, *uv0Arg, *uv1Arg, framePadding, *bgColArg, *tintColArg) ? 1 : 0;
}

IggBool iggCheckbox(char const *label, IggBool *selected)
{
   BoolWrapper selectedArg(selected);
   return ImGui::Checkbox(label, selectedArg) ? 1 : 0;
}

IggBool iggBeginCombo(char const *label, char const *previewValue, int flags)
{
   return ImGui::BeginCombo(label, previewValue, flags) ? 1 : 0;
}

void iggEndCombo(void)
{
   ImGui::EndCombo();
}

IggBool iggDragFloat(char const *label, float *value, float speed, float min, float max, char const *format, float power)
{
    return ImGui::DragFloat(label, value, speed, min, max, format, power) ? 1 : 0;
}

IggBool iggDragInt(char const *label, int *value, float speed, int min, int max, char const *format)
{
    return ImGui::DragInt(label, value, speed, min, max, format) ? 1 : 0;
}

IggBool iggSliderFloat(char const *label, float *value, float minValue, float maxValue, char const *format, float power)
{
   return ImGui::SliderFloat(label, value, minValue, maxValue, format, power) ? 1 : 0;
}

IggBool iggSliderFloatN(char const *label, float *value, int n, float minValue, float maxValue, char const *format, float power)
{
   return  ImGui::SliderScalarN(label, ImGuiDataType_Float, (void*)value, n, &minValue, &maxValue, format, power) ? 1 : 0;
}

IggBool iggSliderInt(char const *label, int *value, int minValue, int maxValue, char const *format)
{
   return ImGui::SliderInt(label, value, minValue, maxValue, format) ? 1 : 0;
}

void iggSeparator(void)
{
   ImGui::Separator();
}

void iggSameLine(float posX, float spacingW)
{
   ImGui::SameLine(posX, spacingW);
}

void iggSpacing(void)
{
   ImGui::Spacing();
}

void iggDummy(IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   ImGui::Dummy(*sizeArg);
}

void iggBeginGroup(void)
{
   ImGui::BeginGroup();
}

void iggEndGroup(void)
{
   ImGui::EndGroup();
}

void iggSetCursorPos(IggVec2 const *localPos)
{
   Vec2Wrapper localPosArg(localPos);
   ImGui::SetCursorPos(*localPosArg);
}

float iggGetTextLineHeight(void)
{
   return ImGui::GetTextLineHeight();
}

float iggGetTextLineHeightWithSpacing(void)
{
   return ImGui::GetTextLineHeightWithSpacing();
}

IggBool iggTreeNode(char const *label, int flags)
{
   return ImGui::TreeNodeEx(label, flags) ? 1 : 0;
}

void iggTreePop(void)
{
   ImGui::TreePop();
}

void iggSetNextTreeNodeOpen(IggBool open, int cond)
{
   ImGui::SetNextTreeNodeOpen(open != 0, cond);
}

IggBool iggSelectable(char const *label, IggBool selected, int flags, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::Selectable(label, selected != 0, flags, *sizeArg) ? 1 : 0;
}

void iggSetTooltip(char const *text)
{
   ImGui::SetTooltip("%s", text);
}

void iggBeginTooltip(void)
{
   ImGui::BeginTooltip();
}

void iggEndTooltip(void)
{
   ImGui::EndTooltip();
}

IggBool iggBeginMainMenuBar(void)
{
   return ImGui::BeginMainMenuBar() ? 1 : 0;
}

void iggEndMainMenuBar(void)
{
   ImGui::EndMainMenuBar();
}

IggBool iggBeginMenuBar(void)
{
   return ImGui::BeginMenuBar() ? 1 : 0;
}

void iggEndMenuBar(void)
{
   ImGui::EndMenuBar();
}

IggBool iggBeginMenu(char const *label, IggBool enabled)
{
   return ImGui::BeginMenu(label, enabled != 0) ? 1 : 0;
}

void iggEndMenu(void)
{
   ImGui::EndMenu();
}

IggBool iggMenuItem(char const *label, char const *shortcut, IggBool selected, IggBool enabled)
{
   return ImGui::MenuItem(label, shortcut, selected != 0, enabled != 0) ? 1 : 0;
}

void iggOpenPopup(char const *id)
{
   ImGui::OpenPopup(id);
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

IggBool iggIsItemHovered(int flags)
{
   return ImGui::IsItemHovered(flags) ? 1 : 0;
}

IggBool iggIsKeyPressed(int key)
{
   return ImGui::IsKeyPressed(key);
}
