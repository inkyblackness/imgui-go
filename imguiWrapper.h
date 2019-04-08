#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern IggContext iggCreateContext(IggFontAtlas sharedFontAtlas);
extern void iggDestroyContext(IggContext context);
extern IggContext iggGetCurrentContext();
extern void iggSetCurrentContext(IggContext context);

extern IggIO iggGetCurrentIO(void);
extern IggGuiStyle iggGetCurrentStyle(void);
extern void iggNewFrame(void);
extern void iggRender(void);
extern IggDrawData iggGetDrawData(void);
extern void iggEndFrame(void);

extern char const *iggGetVersion(void);
extern void iggShowDemoWindow(IggBool *open);
extern void iggShowUserGuide(void);

extern IggBool iggBegin(char const *id, IggBool *open, int flags);
extern void iggEnd(void);
extern IggBool iggBeginChild(char const *id, IggVec2 const *size, IggBool border, int flags);
extern void iggEndChild(void);

extern void iggSetNextWindowPos(IggVec2 const *pos, int cond, IggVec2 const *pivot);
extern void iggSetNextWindowSize(IggVec2 const *size, int cond);
extern void iggSetNextWindowFocus(void);
extern void iggSetNextWindowBgAlpha(float value);

extern void iggPushFont(IggFont handle);
extern void iggPopFont(void);
extern void iggPushStyleColor(int index, IggVec4 const *col);
extern void iggPopStyleColor(int count);
extern void iggPushStyleVarFloat(int index, float value);
extern void iggPushStyleVarVec2(int index, IggVec2 const *value);
extern void iggPopStyleVar(int count);

extern void iggPushItemWidth(float width);
extern void iggPopItemWidth(void);
extern void iggPushTextWrapPos(float wrapPosX);
extern void iggPopTextWrapPos(void);

extern void iggPushID(char const *id);
extern void iggPopID(void);

extern void iggTextUnformatted(char const *text);
extern void iggLabelText(char const *label, char const *text);

extern IggBool iggButton(char const *label, IggVec2 const *size);
extern void iggImage(IggTextureID textureID,
	IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
	IggVec4 const *tintCol, IggVec4 const *borderCol);
extern IggBool iggImageButton(IggTextureID textureID,
	IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
	int framePadding, IggVec4 const *bgCol,
	IggVec4 const *tintCol);
extern IggBool iggCheckbox(char const *label, IggBool *selected);
extern void iggProgressBar(float fraction, IggVec2 const *size, char const *overlay);

extern IggBool iggBeginCombo(char const *label, char const *previewValue, int flags);
extern void iggEndCombo(void);

extern IggBool iggDragFloat(char const *label, float *value, float speed, float min, float max, char const *format, float power);
extern IggBool iggDragInt(char const *label, int *value, float speed, int min, int max, char const *format);

extern IggBool iggSliderFloat(char const *label, float *value, float minValue, float maxValue, char const *format, float power);
extern IggBool iggSliderFloatN(char const *label, float *value, int n, float minValue, float maxValue, char const *format, float power);

extern IggBool iggSliderInt(char const *label, int *value, int minValue, int maxValue, char const *format);

extern IggBool iggInputText(char const* label, char* buf, unsigned int bufSize, int flags, int callbackKey);
extern IggBool iggInputTextMultiline(char const* label, char* buf, unsigned int bufSize, IggVec2 const *size, int flags, int callbackKey);

extern void iggSeparator(void);
extern void iggSameLine(float posX, float spacingW);
extern void iggSpacing(void);
extern void iggDummy(IggVec2 const *size);
extern void iggBeginGroup(void);
extern void iggEndGroup(void);
extern void iggSetCursorPos(IggVec2 const *localPos);
extern void iggAlignTextToFramePadding();
extern float iggGetTextLineHeight(void);
extern float iggGetTextLineHeightWithSpacing(void);

extern IggBool iggTreeNode(char const *label, int flags);
extern void iggTreePop(void);
extern void iggSetNextTreeNodeOpen(IggBool open, int cond);

extern IggBool iggSelectable(char const *label, IggBool selected, int flags, IggVec2 const *size);
extern IggBool iggListBoxV(char const *label, int *current_item, char const *const items[], int items_count, int height_items);

extern void iggPlotLines(const char* label, const float* values, int values_count, int values_offset, const char* overlay_text, float scale_min, float scale_max, IggVec2 const *graph_size);
extern void iggPlotHistogram(const char* label, const float* values, int values_count, int values_offset, const char* overlay_text, float scale_min, float scale_max, IggVec2 const *graph_size);

extern void iggSetTooltip(char const *text);
extern void iggBeginTooltip(void);
extern void iggEndTooltip(void);

extern IggBool iggBeginMainMenuBar(void);
extern void iggEndMainMenuBar(void);
extern IggBool iggBeginMenuBar(void);
extern void iggEndMenuBar(void);
extern IggBool iggBeginMenu(char const *label, IggBool enabled);
extern void iggEndMenu(void);
extern IggBool iggMenuItem(char const *label, char const *shortcut, IggBool selected, IggBool enabled);

extern void iggOpenPopup(char const *id);
extern IggBool iggBeginPopupModal(char const *name, IggBool *open, int flags);
extern IggBool iggBeginPopupContextItem(char const *label, int mouseButton);
extern void iggEndPopup(void);
extern void iggCloseCurrentPopup(void);

extern IggBool iggIsItemHovered(int flags);

extern IggBool iggIsKeyPressed(int key);

extern void iggBeginColumns(int count, char const *label, int flags);
extern void iggNextColumn();
extern int iggGetColumnIndex();
extern int iggGetColumnWidth(int index);
extern void iggSetColumnWidth(int index, float width);
extern float iggGetColumnOffset(int index);
extern void iggSetColumnOffset(int index, float offsetX);
extern int iggGetColumnsCount();
extern void iggSetScrollHereY(float ratio);

#ifdef __cplusplus
}
#endif
