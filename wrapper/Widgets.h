#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggTextUnformatted(char const *text);
extern void iggLabelText(char const *label, char const *text);

extern IggBool iggButton(char const *label, IggVec2 const *size);
extern IggBool iggInvisibleButton(char const *label, IggVec2 const *size);
extern void iggImage(IggTextureID textureID,
   IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
   IggVec4 const *tintCol, IggVec4 const *borderCol);
extern IggBool iggImageButton(IggTextureID textureID,
   IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
   int framePadding, IggVec4 const *bgCol,
   IggVec4 const *tintCol);
extern IggBool iggCheckbox(char const *label, IggBool *selected);
extern IggBool iggRadioButton(char const *label, IggBool active);
extern void iggBullet();
extern void iggProgressBar(float fraction, IggVec2 const *size, char const *overlay);

extern IggBool iggBeginCombo(char const *label, char const *previewValue, int flags);
extern void iggEndCombo(void);

extern IggBool iggDragFloat(char const *label, float *value, float speed, float min, float max, char const *format, float power);
extern IggBool iggDragInt(char const *label, int *value, float speed, int min, int max, char const *format);

extern IggBool iggSliderFloat(char const *label, float *value, float minValue, float maxValue, char const *format, float power);
extern IggBool iggSliderFloatN(char const *label, float *value, int n, float minValue, float maxValue, char const *format, float power);

extern IggBool iggSliderInt(char const *label, int *value, int minValue, int maxValue, char const *format);

extern IggBool iggVSliderFloat(char const *label, IggVec2 const *size, float *value, float minValue, float maxValue, char const *format, float power);
extern IggBool iggVSliderInt(char const *label, IggVec2 const *size, int *value, int minValue, int maxValue, char const *format);

extern IggBool iggInputTextSingleline(char const *label, char const *hint, char *buf, unsigned int bufSize, int flags, int callbackKey);
extern IggBool iggInputTextMultiline(char const *label, char *buf, unsigned int bufSize, IggVec2 const *size, int flags, int callbackKey);

extern IggBool iggInputInt(char const *label, int *value, int step, int step_fast, int flags);

extern IggBool iggColorEdit3(char const *label, float *col, int flags);
extern IggBool iggColorEdit4(char const *label, float *col, int flags);
extern IggBool iggColorPicker3(char const *label, float *col, int flags);
extern IggBool iggColorPicker4(char const *label, float *col, int flags);

extern IggBool iggTreeNode(char const *label, int flags);
extern void iggTreePop(void);
extern void iggSetNextItemOpen(IggBool open, int cond);
extern float iggGetTreeNodeToLabelSpacing(void);

extern IggBool iggCollapsingHeader(const char *label);

extern IggBool iggSelectable(char const *label, IggBool selected, int flags, IggVec2 const *size);
extern IggBool iggListBoxV(char const *label, int *currentItem, char const *const items[], int itemCount, int heightItems);

extern void iggPlotLines(const char *label, const float *values, int valuesCount, int valuesOffset, const char *overlayText, float scaleMin, float scaleMax, IggVec2 const *graphSize);
extern void iggPlotHistogram(const char *label, const float *values, int valuesCount, int valuesOffset, const char *overlayText, float scaleMin, float scaleMax, IggVec2 const *graphSize);

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

extern void iggColumns(int count, char const *label, IggBool border);
extern void iggNextColumn();
extern int iggGetColumnIndex();
extern int iggGetColumnWidth(int index);
extern void iggSetColumnWidth(int index, float width);
extern float iggGetColumnOffset(int index);
extern void iggSetColumnOffset(int index, float offsetX);
extern int iggGetColumnsCount();

extern IggBool iggBeginTabBar(char const *str_id, int flags);
extern void iggEndTabBar();
extern IggBool iggBeginTabItem(char const *label, IggBool *p_open, int flags);
extern void iggEndTabItem();
extern void iggSetTabItemClosed(char const *tab_or_docked_window_label);

#ifdef __cplusplus
}
#endif
