#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern void iggBufferingBar(const char* label, float value, IggVec2 const *size_arg, IggVec4 const *fg_color_arg, IggVec4 const *bg_color_arg);
extern void iggLoadingIndicatorCircle(const char* label, float indicator_radius, int circle_count, float speed, IggVec4 const *fg_color_arg, IggVec4 const *bg_color_arg);
extern void iggSpinner(const char* label, float radius, int thickness, IggVec4 const *color_arg);
extern IggBool iggToggleButton(const char* str_id, IggBool* selected);

#ifdef __cplusplus
}
#endif
