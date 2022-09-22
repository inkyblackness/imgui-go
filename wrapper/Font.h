#pragma once

#include "Types.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void iggPushFont(IggFont handle);
extern void iggPopFont(void);
extern void iggCalcTextSize(const char *text, int length, IggBool hide_text_after_double_hash, float wrap_width, IggVec2 *value);
extern float iggGetFontSize();
extern float iggFontFontSize(IggFont handle);
extern IggFontGlyph iggFindGlyph(IggFont font, int ch);
extern int iggFontGlyphColored(IggFontGlyph glyph);
extern int iggFontGlyphVisible(IggFontGlyph glyph);
extern int iggFontGlyphCodepoint(IggFontGlyph glyph);
extern float iggFontGlyphAdvanceX(IggFontGlyph glyph);
extern float iggFontGlyphX0(IggFontGlyph glyph);
extern float iggFontGlyphY0(IggFontGlyph glyph);
extern float iggFontGlyphX1(IggFontGlyph glyph);
extern float iggFontGlyphY1(IggFontGlyph glyph);
extern float iggFontGlyphU0(IggFontGlyph glyph);
extern float iggFontGlyphV0(IggFontGlyph glyph);
extern float iggFontGlyphU1(IggFontGlyph glyph);
extern float iggFontGlyphV1(IggFontGlyph glyph);

#ifdef __cplusplus
}
#endif
