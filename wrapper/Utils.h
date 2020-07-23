#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern void iggBufferingBar(const char* label, float value, IggVec2 const *size_arg, IggVec4 const *fg_color_arg, IggVec4 const *bg_color_arg);
extern void iggLoadingIndicatorCircle(const char* label, float indicator_radius, int circle_count, float speed, IggVec4 const *fg_color_arg, IggVec4 const *bg_color_arg);
extern void iggSpinner(const char* label, float radius, int thickness, IggVec4 const *color_arg);
extern IggBool iggSplitter(IggBool split_vertically, float thickness, float *size1, float *size2);
extern IggBool iggSelectableInput(const char* label, char* buf, int buf_size, int callback_key);

// typedef struct tagIggComboFilterState {
//     int     activeIdx;         // Index of currently 'active' item by use of up/down keys
//     IggBool selectionChanged;  // Flag to help focus the correct item when selecting active item
// } IggComboFilterState;

// extern IggBool iggComboFilter(const char *label, char *buffer, int bufferlen, const char **hints, int num_hints, IggComboFilterState *state, int callback_key);

#ifdef __cplusplus
}
#endif
