#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

/*******************************/
/* Context creation and access */
/*******************************/
extern IggContext iggCreateContext(IggFontAtlas sharedFontAtlas);
extern void iggDestroyContext(IggContext context);
extern IggContext iggGetCurrentContext();
extern void iggSetCurrentContext(IggContext context);
//DebugCheckVersionAndDataLayout

/*************************/
/* Main - v1.67 complete */
/*************************/
extern IggIO iggGetIO(void);
extern IggGuiStyle iggGetStyle(void);
extern void iggNewFrame(void);
extern void iggEndFrame(void);
extern void iggRender(void);
extern IggDrawData iggGetDrawData(void);

// here because we don't want to break with previous versions of this wrapper
extern IggIO iggGetCurrentIO(void);
// here because we don't want to break with previous versions of this wrapper
extern IggGuiStyle iggGetCurrentStyle(void);

/****************************/
/* Demo, Debug, Information */
/****************************/
extern void iggShowDemoWindow(IggBool *open);
// ShowAboutWindow
// ShowMetricsWindow
// ShowStyleEditor
// ShowStyleSelector
// ShowFontSelector
extern void iggShowUserGuide(void);
extern char const *iggGetVersion(void);

/********************/
/* Styles - MISSING */
/********************/
// StyleColorsDark
// StyleColorsClassic
// StyleColorsLight

/****************************/
/* Windows - v1.67 complete */
/****************************/
extern IggBool iggBegin(char const *id, IggBool *open, int flags);
extern void iggEnd(void);

/*****************/
/* Child Windows */
/*****************/
extern IggBool iggBeginChild(char const *id, IggVec2 const *size, IggBool border, int flags);
// BeginChild - intId
extern void iggEndChild(void);

/*********************/
/* Windows Utilities */
/*********************/
// IsWindowAppearing
// IsWindowCollapsed
// IsWindowFocused
// IsWindowHovered
// GetWindowDrawList
extern void iggWindowPos(IggVec2 *pos);
extern void iggWindowSize(IggVec2 *size);
extern float iggWindowWidth(void);
extern float iggWindowHeight(void);
// GetContentRegionMax
extern void iggContentRegionAvail(IggVec2 *size);
// GetContentRegionAvailWidth
// GetWindowContentRegionMin
// GetWindowContentRegionMax
// GetWindowContentRegionWidth
extern void iggSetNextWindowPos(IggVec2 const *pos, int cond, IggVec2 const *pivot);
extern void iggSetNextWindowSize(IggVec2 const *size, int cond);
// SetNextWindowSizeConstraints
// SetNextWindowContentSize
// SetNextWindowCollapsed
extern void iggSetNextWindowFocus(void);
extern void iggSetNextWindowBgAlpha(float value);
// SetWindowPos
// SetWindowSize
// SetWindowCollapsed
// SetWindowFocus
// SetWindowFontScale
// SetWindowPos - named
// SetWindowSize - named
// SetWindowCollapsed - named
// SetWindowFocus - named

/*********************/
/* Windows Scrolling */
/*********************/
// GetScrollX
// GetScrollY
// GetScrollMaxX
// GetScrollMaxY
// SetScrollX
// SetScrollY
extern void iggSetScrollHereY(float centerYRatio);
// SetScrollFromPosY

/******************************/
/* Parameters stacks (shared) */
/******************************/
extern void iggPushFont(IggFont handle);
extern void iggPopFont(void);
// PushStyleColor - ImU32
extern void iggPushStyleColor(int index, IggVec4 const *col);
extern void iggPopStyleColor(int count);
extern void iggPushStyleVarFloat(int index, float value);
extern void iggPushStyleVarVec2(int index, IggVec2 const *value);
extern void iggPopStyleVar(int count);
// GetStyleColorVec4
// GetFont
extern float iggGetFontSize();
// GetFontTexUvWhitePixel
// GetColorU32
// GetColorU32Vec4
// GetColorU32 - ImU32
extern void iggPushItemWidth(float width);
extern void iggPopItemWidth(void);
extern float iggCalcItemWidth(void);
extern void iggPushTextWrapPos(float wrapPosX);
extern void iggPopTextWrapPos(void);
// PushAllowKeyboardFocus
// PopAllowKeyboardFocus
// PushButtonRepeat
// PopButtonRepeat

/*******************/
/* Cursor / Layout */
/*******************/
extern void iggSeparator(void);
extern void iggSameLine(float posX, float spacingW);
// NewLine
extern void iggSpacing(void);
extern void iggDummy(IggVec2 const *size);
// Indent
// Unindent
extern void iggBeginGroup(void);
extern void iggEndGroup(void);
extern IggVec2 iggGetCursorPos(void);
extern float iggGetCursorPosX(void);
extern float iggGetCursorPosY(void);
extern void iggSetCursorPos(IggVec2 const *localPos);
// SetCursorPosX
// SetCursorPosY
extern IggVec2 iggGetCursorStartPos(void);
extern IggVec2 iggGetCursorScreenPos(void);
extern void iggSetCursorScreenPos(IggVec2 const *absPos);
extern void iggAlignTextToFramePadding();
extern float iggGetTextLineHeight(void);
extern float iggGetTextLineHeightWithSpacing(void);
// GetFrameHeight
// GetFrameHeightWithSpacing

// here because we don't want to break with previous versions of this wrapper
extern void iggCursorPos(IggVec2 *pos);
// here because we don't want to break with previous versions of this wrapper
extern float iggCursorPosX(void);
// here because we don't want to break with previous versions of this wrapper
extern float iggCursorPosY(void);
// here because we don't want to break with previous versions of this wrapper
extern void iggCursorStartPos(IggVec2 *pos);
// here because we don't want to break with previous versions of this wrapper
extern void iggCursorScreenPos(IggVec2 *pos);

/*******************/
/* ID stack/scopes */
/*******************/
extern void iggPushID(char const *id);
// PushID - str_id_end
// PushID - ptr_id
// PushID - int
extern void iggPopID(void);
// GetID
// GetID - str_id_end
// GetID - ptr_id

/*****************/
/* Widgets: Text */
/*****************/
extern void iggTextUnformatted(char const *text);
// Text
// TextV
// TextColored
// TextColoredV
// TextDisabled
// TextDisabledV
// TextWrapped
// TextWrappedV
extern void iggLabelText(char const *label, char const *text);
// LabelTextV
// BulletText
// BulletTextV

/**********************************/
/* Widgets: Main - v1.67 complete */
/**********************************/
extern IggBool iggButton(char const *label, IggVec2 const *size);
extern IggBool iggSmallButton(char const *label);
extern IggBool iggInvisibleButton(char const *str_id, IggVec2 const *size);
extern IggBool iggArrowButton(char const *str_id, int dir);
extern void iggImage(IggTextureID textureID,
	IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
	IggVec4 const *tintCol, IggVec4 const *borderCol);
extern IggBool iggImageButton(IggTextureID textureID,
	IggVec2 const *size, IggVec2 const *uv0, IggVec2 const *uv1,
	int framePadding, IggVec4 const *bgCol,
	IggVec4 const *tintCol);
extern IggBool iggCheckbox(char const *label, IggBool *selected);
extern IggBool iggCheckboxFlags(char const *label, unsigned int *flags, unsigned int flags_value);
extern IggBool iggRadioButton(char const *label, IggBool active);
extern IggBool iggRadioButtonInt(char const *label, int *v, int v_button);
extern void iggProgressBar(float fraction, IggVec2 const *size, char const *overlay);
extern void iggBullet();

/**********************/
/* Widgets: Combo Box */
/**********************/
extern IggBool iggBeginCombo(char const *label, char const *previewValue, int flags);
extern void iggEndCombo(void);
// Combo
// ComboStr
// ComboPtr

/******************/
/* Widgets: Drags */
/******************/
extern IggBool iggDragFloat(char const *label, float *value, float speed, float min, float max, char const *format, float power);
// DragFloat2
// DragFloat3
// DragFloat4
// DragFloatRange2
extern IggBool iggDragInt(char const *label, int *value, float speed, int min, int max, char const *format);
// DragInt2
// DragInt3
// DragInt4
// DragIntRange2
// DragScalar
// DragScalarN

/********************/
/* Widgets: Sliders */
/********************/
extern IggBool iggSliderFloat(char const *label, float *value, float minValue, float maxValue, char const *format, float power);
// SliderFloat2
// SliderFloat3
// SliderFloat4
// SliderAngle
extern IggBool iggSliderInt(char const *label, int *value, int minValue, int maxValue, char const *format);
// SliderInt2
// SliderInt3
// SliderInt4
// SliderScalar
// SliderScalarN
// VSliderFloat
// VSliderInt
// VSliderScalar

// here because we don't want to break with previous versions of this wrapper
extern IggBool iggSliderFloatN(char const *label, float *value, int n, float minValue, float maxValue, char const *format, float power);

/*******************************/
/* Widgets: Input with Keyboard */
/*******************************/
extern IggBool iggInputText(char const* label, char* buf, unsigned int bufSize, int flags, int callbackKey);
extern IggBool iggInputTextMultiline(char const* label, char* buf, unsigned int bufSize, IggVec2 const *size, int flags, int callbackKey);
// InputFloat
// InputFloat2
// InputFloat3
// InputFloat4
// InputInt
// InputInt2
// InputInt3
// InputInt4
// InputDouble
// InputScalar
// InputScalarN

/******************************************/
/* Widgets: Color Editor/Picker - MISSING */
/******************************************/
// ColorEdit3
// ColorEdit4
// ColorPicker3
// ColorPicker4
// ColorButton
// SetColorEditOptions

/******************/
/* Widgets: Trees */
/******************/
extern IggBool iggTreeNode(char const *label, int flags);
// TreeNodeFmt
// TreeNodePtr
// TreeNodeV
// TreeNodePtrV
// TreeNodeEx
// TreeNodeExFmt
// TreeNodeExPtr
// TreeNodeExV
// TreeNodeExPtrV
// TreePush
// TreePushPtr
extern void iggTreePop(void);
// TreeAdvanceToLabelPos
extern float iggGetTreeNodeToLabelSpacing(void);
extern void iggSetNextTreeNodeOpen(IggBool open, int cond);
// CollapsingHeader
// CollapsingHeaderV

/************************/
/* Widgets: Selectables */
/************************/
extern IggBool iggSelectable(char const *label, IggBool selected, int flags, IggVec2 const *size);
// SelectableV

/***********************/
/* Widgets: List Boxes */
/***********************/
extern IggBool iggListBox(char const *label, int *currentItem, char const *const items[], int itemCount, int heightItems);
// ListBox
// ListBoxHeader
// ListBoxHeader
// ListBoxFooter

// here because we don't want to break with previous versions of this wrapper
extern IggBool iggListBoxV(char const *label, int *currentItem, char const *const items[], int itemCount, int heightItems);

/**************************/
/* Widgets: Data Plotting */
/**************************/
extern void iggPlotLines(const char* label, const float* values, int valuesCount, int valuesOffset, const char* overlayText, float scaleMin, float scaleMax, IggVec2 const *graphSize);
// PlotLines
extern void iggPlotHistogram(const char* label, const float* values, int valuesCount, int valuesOffset, const char* overlayText, float scaleMin, float scaleMax, IggVec2 const *graphSize);
// PlotHistogram

/***************************************/
/* Widgets: Value() Helpers. - MISSING */
/***************************************/
// ValueBool
// ValueInt
// ValueUInt
// ValueFloat

/******************/
/* Widgets: Menus */
/******************/
extern IggBool iggBeginMainMenuBar(void);
extern void iggEndMainMenuBar(void);
extern IggBool iggBeginMenuBar(void);
extern void iggEndMenuBar(void);
extern IggBool iggBeginMenu(char const *label, IggBool enabled);
extern void iggEndMenu(void);
extern IggBool iggMenuItem(char const *label, char const *shortcut, IggBool selected, IggBool enabled);
// MenuItem

/************/
/* Tooltips */
/************/
extern void iggBeginTooltip(void);
extern void iggEndTooltip(void);
extern void iggSetTooltip(char const *text);
// SetTooltip ...
// SetTooltipV

/******************/
/* Popups, Modals */
/******************/
extern void iggOpenPopup(char const *id);
// BeginPopup
extern IggBool iggBeginPopupContextItem(char const *label, int mouseButton);
// BeginPopupContextWindow
// BeginPopupContextVoid
extern IggBool iggBeginPopupModal(char const *name, IggBool *open, int flags);
extern void iggEndPopup(void);
// OpenPopupOnItemClick
// IsPopupOpen
extern void iggCloseCurrentPopup(void);

/****************************/
/* Columns - v1.67 complete */
/****************************/
extern void iggColumns(int count, char const *label, int flags);
extern void iggNextColumn();
extern int iggGetColumnIndex();
extern int iggGetColumnWidth(int index);
extern void iggSetColumnWidth(int index, float width);
extern float iggGetColumnOffset(int index);
extern void iggSetColumnOffset(int index, float offsetX);
extern int iggGetColumnsCount();

// here because we don't want to break with previous versions of this wrapper
extern void iggBeginColumns(int count, char const *label, int flags);

/***********************************/
/* Tab Bars, Tabs - v1.67 complete */
/***********************************/
extern IggBool iggBeginTabBar(char const *str_id, int flags);
extern void iggEndTabBar();
extern IggBool iggBeginTabItem(char const *label, IggBool *p_open, int flags);
extern void iggEndTabItem();
extern void iggSetTabItemClosed(char const * tab_or_docked_window_label);

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
extern void iggSetItemDefaultFocus();

/**************************/
/* Item/Widgets Utilities */
/**************************/
extern IggBool iggIsItemHovered(int flags);
extern IggBool iggIsItemFocused();
extern IggBool iggIsAnyItemFocused();

/*************************************/
/* Miscellaneous Utilities - MISSING */
/*************************************/

/*****************************/
/* Color Utilities - MISSING */
/*****************************/

/********************/
/* Inputs Utilities */
/*******************/
extern IggBool iggIsKeyDown(int key);
extern IggBool iggIsKeyPressed(int key, IggBool repeat);
extern IggBool iggIsKeyReleased(int key);
extern IggBool iggIsMouseDown(int button);
extern IggBool iggIsAnyMouseDown();
extern IggBool iggIsMouseClicked(int button, IggBool repeat);
extern IggBool iggIsMouseReleased(int button);
extern IggBool iggIsMouseDoubleClicked(int button);
extern int iggGetMouseCursor();
extern void iggSetMouseCursor(int cursor);

/*********************************/
/* Clipboard Utilities - MISSING */
/*********************************/

/*************************************/
/* Settings/.Ini Utilities - MISSING */
/*************************************/

/******************************/
/* Memory Utilities - MISSING */
/******************************/

#ifdef __cplusplus
}
#endif
