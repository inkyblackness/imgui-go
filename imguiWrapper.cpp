#include "imguiWrappedHeader.h"
#include "imguiWrapper.h"
#include "WrapperConverter.h"

/*******************************/
/* Context creation and access */
/*******************************/

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

/*************************/
/* Main - v1.67 complete */
/*************************/

IggIO iggGetIO()
{
   return reinterpret_cast<IggIO>(&ImGui::GetIO());
}

IggGuiStyle iggGetStyle()
{
   return reinterpret_cast<IggGuiStyle>(&ImGui::GetStyle());
}

void iggNewFrame()
{
   ImGui::NewFrame();
}

void iggEndFrame()
{
   ImGui::EndFrame();
}

void iggRender()
{
   ImGui::Render();
}

IggDrawData iggGetDrawData()
{
   return reinterpret_cast<IggDrawData>(ImGui::GetDrawData());
}

// here because we don't want to break with previous versions of this wrapper
IggIO iggGetCurrentIO()
{
   return iggGetIO();
}

// here because we don't want to break with previous versions of this wrapper
IggGuiStyle iggGetCurrentStyle()
{
   return iggGetStyle();
}

/****************************/
/* Demo, Debug, Information */
/****************************/

void iggShowDemoWindow(IggBool *open)
{
   BoolWrapper openArg(open);

   ImGui::ShowDemoWindow(openArg);
}

void iggShowUserGuide(void)
{
   ImGui::ShowUserGuide();
}

char const *iggGetVersion()
{
   return ImGui::GetVersion();
}

/********************/
/* Styles - MISSING */
/********************/

/****************************/
/* Windows - v1.67 complete */
/****************************/

IggBool iggBegin(char const *id, IggBool *open, int flags)
{
   BoolWrapper openArg(open);
   return ImGui::Begin(id, openArg, flags) ? 1 : 0;
}

void iggEnd(void)
{
   ImGui::End();
}

/*****************/
/* Child Windows */
/*****************/

IggBool iggBeginChild(char const *id, IggVec2 const *size, IggBool border, int flags)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::BeginChild(id, *sizeArg, border, flags) ? 1 : 0;
}

void iggEndChild(void)
{
   ImGui::EndChild();
}

/*********************/
/* Windows Utilities */
/*********************/

void iggWindowPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetWindowPos());
}

void iggWindowSize(IggVec2 *size)
{
   exportValue(*size, ImGui::GetWindowSize());
}

float iggWindowWidth(void)
{
   return ImGui::GetWindowWidth();
}

float iggWindowHeight(void)
{
   return ImGui::GetWindowHeight();
}

void iggContentRegionAvail(IggVec2 *size)
{
   exportValue(*size, ImGui::GetContentRegionAvail());
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

/*********************/
/* Windows Scrolling */
/*********************/

void iggSetScrollHereY(float centerYRatio)
{
   ImGui::SetScrollHereY(centerYRatio);
}

/******************************/
/* Parameters stacks (shared) */
/******************************/

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

float iggGetFontSize()
{
   return ImGui::GetFontSize();
}

void iggPushItemWidth(float width)
{
   ImGui::PushItemWidth(width);
}

void iggPopItemWidth(void)
{
   ImGui::PopItemWidth();
}

float iggCalcItemWidth(void)
{
   return ImGui::CalcItemWidth();
}

void iggPushTextWrapPos(float wrapPosX)
{
   ImGui::PushTextWrapPos(wrapPosX);
}

void iggPopTextWrapPos(void)
{
   ImGui::PopTextWrapPos();
}

/*******************/
/* Cursor / Layout */
/*******************/

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

IggVec2 iggGetCursorPos()
{
    ImVec2 r = ImGui::GetCursorPos();
    return IggVec2{x: r.x, y:r.y};
}

float iggGetCursorPosX(void)
{
   return ImGui::GetCursorPosX();
}

float iggGetCursorPosY(void)
{
   return ImGui::GetCursorPosY();
}

void iggSetCursorPos(IggVec2 const *localPos)
{
   Vec2Wrapper localPosArg(localPos);
   ImGui::SetCursorPos(*localPosArg);
}

IggVec2 iggGetCursorStartPos(void)
{
    ImVec2 r = ImGui::GetCursorStartPos();
    return IggVec2{x: r.x, y:r.y};
}

IggVec2 iggCursorScreenPos(void)
{
    ImVec2 r = ImGui::GetCursorScreenPos();
    return IggVec2{x: r.x, y:r.y};
}

void iggSetCursorScreenPos(IggVec2 const *absPos)
{
   Vec2Wrapper absPosArg(absPos);
   ImGui::SetCursorScreenPos(*absPosArg);
}

void iggAlignTextToFramePadding()
{
   ImGui::AlignTextToFramePadding();
}

float iggGetTextLineHeight(void)
{
   return ImGui::GetTextLineHeight();
}

float iggGetTextLineHeightWithSpacing(void)
{
   return ImGui::GetTextLineHeightWithSpacing();
}

// here because we don't want to break with previous versions of this wrapper
void iggCursorPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetCursorPos());
}

// here because we don't want to break with previous versions of this wrapper
float iggCursorPosX(void)
{
   return iggGetCursorPosX();
}

// here because we don't want to break with previous versions of this wrapper
float iggCursorPosY(void)
{
   return iggGetCursorPosY();
}

// here because we don't want to break with previous versions of this wrapper
void iggCursorStartPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetCursorStartPos());
}

// here because we don't want to break with previous versions of this wrapper
void iggCursorScreenPos(IggVec2 *pos)
{
   exportValue(*pos, ImGui::GetCursorScreenPos());
}

/*******************/
/* ID stack/scopes */
/*******************/

void iggPushID(char const *id)
{
   ImGui::PushID(id);
}
void iggPopID(void)
{
   ImGui::PopID();
}

/*****************/
/* Widgets: Text */
/*****************/

void iggTextUnformatted(char const *text)
{
   ImGui::TextUnformatted(text);
}

void iggLabelText(char const *label, char const *text)
{
   ImGui::LabelText(label, "%s", text);
}

/**********************************/
/* Widgets: Main - v1.67 complete */
/**********************************/

IggBool iggButton(char const *label, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::Button(label, *sizeArg) ? 1 : 0;
}

IggBool iggSmallButton (char const *label) {
    return ImGui::SmallButton(label) ? 1 : 0;
}

IggBool iggInvisibleButton(char const *str_id, IggVec2 const *size) {
    Vec2Wrapper sizeArg(size);
    return ImGui::InvisibleButton(str_id, *sizeArg) ? 1 : 0;
}

IggBool iggArrowButton (char const *label, int dir) {
    return ImGui::ArrowButton(label, dir) ? 1 : 0;
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

IggBool iggCheckboxFlags(char const *label, unsigned int *flags, unsigned int flags_value)
{
   return ImGui::CheckboxFlags(label, flags, flags_value) ? 1 : 0;
}

IggBool iggRadioButton(char const *label, IggBool *active)
{
   return ImGui::RadioButton(label, active) ? 1 : 0;
}

IggBool iggRadioButtonInt(char const *label, int *v, int v_button)
{
   return ImGui::RadioButton(label, v, v_button) ? 1 : 0;
}

void iggProgressBar(float fraction, IggVec2 const *size, char const *overlay)
{
   Vec2Wrapper sizeArg(size);
   ImGui::ProgressBar(fraction, *sizeArg, overlay);
}

void iggBullet() {
    ImGui::Bullet();
}

/**********************/
/* Widgets: Combo Box */
/**********************/

IggBool iggBeginCombo(char const *label, char const *previewValue, int flags)
{
   return ImGui::BeginCombo(label, previewValue, flags) ? 1 : 0;
}

void iggEndCombo(void)
{
   ImGui::EndCombo();
}

/******************/
/* Widgets: Drags */
/******************/

IggBool iggDragFloat(char const *label, float *value, float speed, float min, float max, char const *format, float power)
{
   return ImGui::DragFloat(label, value, speed, min, max, format, power) ? 1 : 0;
}

IggBool iggDragInt(char const *label, int *value, float speed, int min, int max, char const *format)
{
   return ImGui::DragInt(label, value, speed, min, max, format) ? 1 : 0;
}

/********************/
/* Widgets: Sliders */
/********************/

IggBool iggSliderFloat(char const *label, float *value, float minValue, float maxValue, char const *format, float power)
{
   return ImGui::SliderFloat(label, value, minValue, maxValue, format, power) ? 1 : 0;
}

IggBool iggSliderInt(char const *label, int *value, int minValue, int maxValue, char const *format)
{
   return ImGui::SliderInt(label, value, minValue, maxValue, format) ? 1 : 0;
}

// here because we don't want to break with previous versions of this wrapper
IggBool iggSliderFloatN(char const *label, float *value, int n, float minValue, float maxValue, char const *format, float power)
{
   return ImGui::SliderScalarN(label, ImGuiDataType_Float, (void *)value, n, &minValue, &maxValue, format, power) ? 1 : 0;
}

/*******************************/
/* Widgets: Input with Keyboard */
/*******************************/

extern "C" int iggInputTextCallback(IggInputTextCallbackData data, int key);

static int iggInputTextCallbackWrapper(ImGuiInputTextCallbackData *data)
{
   return iggInputTextCallback(reinterpret_cast<IggInputTextCallbackData>(data), static_cast<int>(reinterpret_cast<size_t>(data->UserData)));
}

IggBool iggInputText(char const *label, char* buf, unsigned int bufSize, int flags, int callbackKey)
{
   return ImGui::InputText(label, buf, static_cast<size_t>(bufSize), flags,
                           iggInputTextCallbackWrapper, reinterpret_cast<void *>(callbackKey)) ? 1 : 0;
}

IggBool iggInputTextMultiline(char const *label, char* buf, unsigned int bufSize, IggVec2 const *size, int flags, int callbackKey)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::InputTextMultiline(label, buf, static_cast<size_t>(bufSize), *sizeArg, flags,
                                    iggInputTextCallbackWrapper, reinterpret_cast<void *>(callbackKey)) ? 1 : 0;
}

/******************************************/
/* Widgets: Color Editor/Picker - MISSING */
/******************************************/

/******************/
/* Widgets: Trees */
/******************/

IggBool iggTreeNode(char const *label, int flags)
{
   return ImGui::TreeNodeEx(label, flags) ? 1 : 0;
}

void iggTreePop(void)
{
   ImGui::TreePop();
}

float iggGetTreeNodeToLabelSpacing(void)
{
   return ImGui::GetTreeNodeToLabelSpacing();
}

void iggSetNextTreeNodeOpen(IggBool open, int cond)
{
   ImGui::SetNextTreeNodeOpen(open != 0, cond);
}

/************************/
/* Widgets: Selectables */
/************************/

IggBool iggSelectable(char const *label, IggBool selected, int flags, IggVec2 const *size)
{
   Vec2Wrapper sizeArg(size);
   return ImGui::Selectable(label, selected != 0, flags, *sizeArg) ? 1 : 0;
}

/***********************/
/* Widgets: List Boxes */
/***********************/

IggBool iggListBox(char const *label, int *currentItem, char const *const items[], int itemsCount, int heightItems)
{
   return ImGui::ListBox(label, currentItem, items, itemsCount, heightItems) ? 1 : 0;
}

// here because we don't want to break with previous versions of this wrapper
IggBool iggListBoxV(char const *label, int *currentItem, char const *const items[], int itemsCount, int heightItems)
{
   return iggListBox(label, currentItem, items, itemsCount, heightItems);
}

/**************************/
/* Widgets: Data Plotting */
/**************************/

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

/***************************************/
/* Widgets: Value() Helpers. - MISSING */
/***************************************/

/******************/
/* Widgets: Menus */
/******************/

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

/************/
/* Tooltips */
/************/

void iggBeginTooltip(void)
{
   ImGui::BeginTooltip();
}

void iggEndTooltip(void)
{
   ImGui::EndTooltip();
}

void iggSetTooltip(char const *text)
{
   ImGui::SetTooltip("%s", text);
}

/******************/
/* Popups, Modals */
/******************/

void iggOpenPopup(char const *id)
{
   ImGui::OpenPopup(id);
}

IggBool iggBeginPopupContextItem(char const *label, int mouseButton)
{
   return ImGui::BeginPopupContextItem(label, mouseButton) ? 1 : 0;
}

IggBool iggBeginPopupModal(char const *name, IggBool *open, int flags)
{
   BoolWrapper openArg(open);
   return ImGui::BeginPopupModal(name, openArg, flags) ? 1 : 0;
}

void iggEndPopup(void)
{
   ImGui::EndPopup();
}

void iggCloseCurrentPopup(void)
{
   ImGui::CloseCurrentPopup();
}

/***********/
/* Columns */
/***********/

void iggColumns(int count, char const *label, int flags)
{
   ImGui::Columns(count, label, flags);
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

// here because we don't want to break with previous versions of this wrapper
void iggBeginColumns(int count, char const *label, int flags)
{
   iggColumns(count, label, flags);
}

/***********************************/
/* Tab Bars, Tabs - v1.67 complete */
/***********************************/

IggBool iggBeginTabBar(char const *str_id, int flags) {
    return ImGui::BeginTabBar(str_id, flags) ? 1 : 0;
}

void iggEndTabBar() {
    ImGui::EndTabBar();
}

IggBool iggBeginTabItem(char const *label, IggBool *p_open, int flags) {
    BoolWrapper openArg(p_open);
    return ImGui::BeginTabItem(label, openArg, flags) ? 1 : 0;
}

void iggEndTabItem() {
    ImGui::EndTabItem();
}

void iggSetTabItemClosed(char const * tab_or_docked_window_label) {
    ImGui::SetTabItemClosed(tab_or_docked_window_label);
}

/*****************************/
/* Logging/Capture - MISSING */
/*****************************/

/***************************/
/* Drag and Drop - MISSING */
/***************************/

/**********************/
/* Clipping - MISSING */
/**********************/

/*********************/
/* Focus, Activation */
/*********************/
void iggSetItemDefaultFocus()
{
   ImGui::SetItemDefaultFocus();
}

/**************************/
/* Item/Widgets Utilities */
/**************************/
IggBool iggIsItemHovered(int flags)
{
   return ImGui::IsItemHovered(flags) ? 1 : 0;
}

IggBool iggIsItemFocused()
{
   return ImGui::IsItemFocused();
}

IggBool iggIsAnyItemFocused()
{
   return ImGui::IsAnyItemFocused();
}


/*************************************/
/* Miscellaneous Utilities - MISSING */
/*************************************/

/*****************************/
/* Color Utilities - MISSING */
/*****************************/

/********************/
/* Inputs Utilities */
/*******************/
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

int iggGetMouseCursor()
{
   return ImGui::GetMouseCursor();
}

void iggSetMouseCursor(int cursor)
{
   ImGui::SetMouseCursor(cursor);
}

/*********************************/
/* Clipboard Utilities - MISSING */
/*********************************/

/*************************************/
/* Settings/.Ini Utilities - MISSING */
/*************************************/

/******************************/
/* Memory Utilities - MISSING */
/******************************/
