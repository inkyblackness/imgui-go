#include "ConfiguredImGui.h"

#include "Font.h"

void iggPushFont(IggFont handle)
{
   ImFont *font = reinterpret_cast<ImFont *>(handle);
   ImGui::PushFont(font);
}

void iggPopFont(void)
{
   ImGui::PopFont();
}

float iggGetFontSize()
{
   return ImGui::GetFontSize();
}

void iggCalcTextSize(const char *text, int length, IggBool hide_text_after_double_hash, float wrap_width, IggVec2 *value)
{
   exportValue(*value, ImGui::CalcTextSize(text, text + length, hide_text_after_double_hash, wrap_width));
}

IggFontGlyph iggFindGlyph(IggFont handle, int ch)
{
   ImFont *font = reinterpret_cast<ImFont *>(handle);
   return (IggFontGlyph)font->FindGlyph(ch);
}

int iggFontGlyphColored(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->Colored;
}

int iggFontGlyphVisible(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->Visible;
}

int iggFontGlyphCodepoint(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->Codepoint;
}

float iggFontGlyphAdvanceX(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->AdvanceX;
}

float iggFontGlyphX0(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->X0;
}

float iggFontGlyphY0(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->Y0;
}

float iggFontGlyphX1(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->X1;
}

float iggFontGlyphY1(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->Y1;
}

float iggFontGlyphU0(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->U0;
}

float iggFontGlyphV0(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->V0;
}

float iggFontGlyphU1(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->U1;
}

float iggFontGlyphV1(IggFontGlyph handle)
{
    ImFontGlyph *glyph = reinterpret_cast<ImFontGlyph *>(handle);
    return glyph->V1;
}
