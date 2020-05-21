#include "ConfiguredImGui.h"

#include "Widgets.h"
#include "WrapperConverter.h"

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

IggBool iggInvisibleButton(char const *label, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::InvisibleButton(label, *sizeArg) ? 1 : 0;
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
   ImGui::Image(reinterpret_cast<ImTextureID>(textureID), *sizeArg, *uv0Arg, *uv1Arg, *tintColArg, *borderColArg);
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
   return ImGui::ImageButton(reinterpret_cast<ImTextureID>(textureID), *sizeArg, *uv0Arg, *uv1Arg, framePadding, *bgColArg, *tintColArg) ? 1 : 0;
}

IggBool iggCheckbox(char const *label, IggBool *selected)
{
   BoolWrapper selectedArg(selected);
   return ImGui::Checkbox(label, selectedArg) ? 1 : 0;
}

IggBool iggRadioButton(char const *label, IggBool active)
{
   return ImGui::RadioButton(label, active != 0) ? 1 : 0;
}

void iggBullet()
{
   ImGui::Bullet();
}

void iggProgressBar(float fraction, IggVec2 const *size, char const *overlay)
{
   Vec2Wrapper sizeArg(size);
   ImGui::ProgressBar(fraction, *sizeArg, overlay);
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
   return ImGui::SliderScalarN(label, ImGuiDataType_Float, (void *)value, n, &minValue, &maxValue, format, power) ? 1 : 0;
}

IggBool iggSliderInt(char const *label, int *value, int minValue, int maxValue, char const *format)
{
   return ImGui::SliderInt(label, value, minValue, maxValue, format) ? 1 : 0;
}

IggBool iggVSliderFloat(char const *label, IggVec2 const *size, float *value, float minValue, float maxValue, char const *format, float power)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::VSliderFloat(label, *sizeArg, value, minValue, maxValue, format, power) ? 1 : 0;
}

IggBool iggVSliderInt(char const *label, IggVec2 const *size, int *value, int minValue, int maxValue, char const *format)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::VSliderInt(label, *sizeArg, value, minValue, maxValue, format) ? 1 : 0;
}

extern "C" int iggInputTextCallback(IggInputTextCallbackData data, int key);

static int iggInputTextCallbackWrapper(ImGuiInputTextCallbackData *data)
{
   return iggInputTextCallback(reinterpret_cast<IggInputTextCallbackData>(data), static_cast<int>(reinterpret_cast<size_t>(data->UserData)));
}

IggBool iggInputTextSingleline(char const *label, char const *hint, char *buf, unsigned int bufSize, int flags, int callbackKey)
{
   return ImGui::InputTextWithHint(label, hint, buf, static_cast<size_t>(bufSize), flags,
             iggInputTextCallbackWrapper, reinterpret_cast<void *>(callbackKey))
      ? 1
      : 0;
}

IggBool iggInputTextMultiline(char const *label, char *buf, unsigned int bufSize, IggVec2 const *size, int flags, int callbackKey)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::InputTextMultiline(label, buf, static_cast<size_t>(bufSize), *sizeArg, flags,
             iggInputTextCallbackWrapper, reinterpret_cast<void *>(callbackKey))
      ? 1
      : 0;
}

IggBool iggInputInt(char const *label, int *value, int step, int step_fast, int flags)
{
   return ImGui::InputInt(label, value, step, step_fast, flags) ? 1 : 0;
}

IggBool iggColorEdit3(char const *label, float *col, int flags)
{
   return ImGui::ColorEdit3(label, col, flags) ? 1 : 0;
}

IggBool iggColorEdit4(char const *label, float *col, int flags)
{
   return ImGui::ColorEdit4(label, col, flags) ? 1 : 0;
}

IggBool iggColorPicker3(char const *label, float *col, int flags)
{
   return ImGui::ColorPicker3(label, col, flags) ? 1 : 0;
}

IggBool iggColorPicker4(char const *label, float *col, int flags)
{
   return ImGui::ColorPicker4(label, col, flags) ? 1 : 0;
}

IggBool iggCollapsingHeader(const char *label)
{
   return ImGui::CollapsingHeader(label) ? 1 : 0;
}

IggBool iggTreeNode(char const *label, int flags)
{
   return ImGui::TreeNodeEx(label, flags) ? 1 : 0;
}

void iggTreePop(void)
{
   ImGui::TreePop();
}

void iggSetNextItemOpen(IggBool open, int cond)
{
   ImGui::SetNextItemOpen(open != 0, cond);
}

float iggGetTreeNodeToLabelSpacing(void)
{
   return ImGui::GetTreeNodeToLabelSpacing();
}

IggBool iggSelectable(char const *label, IggBool selected, int flags, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::Selectable(label, selected != 0, flags, *sizeArg) ? 1 : 0;
}

IggBool iggListBoxV(char const *label, int *currentItem, char const *const items[], int itemsCount, int heightItems)
{
   return ImGui::ListBox(label, currentItem, items, itemsCount, heightItems) ? 1 : 0;
}

void iggPlotLines(char const *label, float const *values, int valuesCount, int valuesOffset, char const *overlayText, float scaleMin, float scaleMax, IggVec2 const *graphSize)
{
   Vec2Wrapper graphSizeArg(graphSize);
   ImGui::PlotLines(label, values, valuesCount, valuesOffset, overlayText, scaleMin, scaleMax, *graphSizeArg);
}

void iggPlotHistogram(char const *label, float const *values, int valuesCount, int valuesOffset, char const *overlayText, float scaleMin, float scaleMax, IggVec2 const *graphSize)
{
   Vec2Wrapper graphSizeArg(graphSize);
   ImGui::PlotHistogram(label, values, valuesCount, valuesOffset, overlayText, scaleMin, scaleMax, *graphSizeArg);
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

void iggColumns(int count, char const *label, IggBool border)
{
   ImGui::Columns(count, label, border);
}

void iggNextColumn()
{
   ImGui::NextColumn();
}

int iggGetColumnIndex()
{
   return ImGui::GetColumnIndex();
}

int iggGetColumnWidth(int index)
{
   return ImGui::GetColumnWidth(index);
}

void iggSetColumnWidth(int index, float width)
{
   ImGui::SetColumnWidth(index, width);
}

float iggGetColumnOffset(int index)
{
   return ImGui::GetColumnOffset(index);
}

void iggSetColumnOffset(int index, float offsetX)
{
   ImGui::SetColumnOffset(index, offsetX);
}

int iggGetColumnsCount()
{
   return ImGui::GetColumnsCount();
}

IggBool iggBeginTabBar(char const *str_id, int flags)
{
   return ImGui::BeginTabBar(str_id, flags) ? 1 : 0;
}

void iggEndTabBar()
{
   ImGui::EndTabBar();
}

IggBool iggBeginTabItem(char const *label, IggBool *p_open, int flags)
{
   BoolWrapper openArg(p_open);
   return ImGui::BeginTabItem(label, openArg, flags) ? 1 : 0;
}

void iggEndTabItem()
{
   ImGui::EndTabItem();
}

void iggSetTabItemClosed(char const *tab_or_docked_window_label)
{
   ImGui::SetTabItemClosed(tab_or_docked_window_label);
}
