package imgui

// #include "wrapper/Font.h"
// #include "wrapper/Types.h"
import "C"

// Font describes one loaded font in an atlas.
type Font uintptr

// FontGlyph represents a single font glyph from a font atlas.
type FontGlyph uintptr

// DefaultFont can be used to refer to the default font of the current font atlas without
// having the actual font reference.
const DefaultFont Font = 0

// PushFont adds the given font on the stack. Use DefaultFont to refer to the default font.
func PushFont(font Font) {
	C.iggPushFont(font.handle())
}

// PopFont removes the previously pushed font from the stack.
func PopFont() {
	C.iggPopFont()
}

// FontSize returns the current font size (= height in pixels) of the current font with the current scale applied.
func FontSize() float32 {
	return float32(C.iggGetFontSize())
}

// CalcTextSize calculates the size of the text.
func CalcTextSize(text string, hideTextAfterDoubleHash bool, wrapWidth float32) Vec2 {
	CString := newStringBuffer(text)
	defer CString.free()

	var vec2 Vec2
	valueArg, returnFunc := vec2.wrapped()

	C.iggCalcTextSize((*C.char)(CString.ptr), C.int(CString.size)-1, castBool(hideTextAfterDoubleHash), C.float(wrapWidth), valueArg)
	returnFunc()

	return vec2
}

func (font Font) handle() C.IggFont {
	return C.IggFont(font)
}

// FindGlyph returns the FontGlyph corresponding to the given rune.
func (font Font) FindGlyph(ch rune) FontGlyph {
	return FontGlyph(C.iggFindGlyph(font.handle(), C.int(ch)))
}

func (glyph FontGlyph) handle() C.IggFontGlyph {
	return C.IggFontGlyph(glyph)
}

// Colored returns whether the glyph is colored.
func (glyph FontGlyph) Colored() bool {
	return C.iggFontGlyphColored(glyph.handle()) != 0
}

// Visible indicates whether the glyph is visible; it will be false for a space character, for example.
func (glyph FontGlyph) Visible() bool {
	return C.iggFontGlyphVisible(glyph.handle()) != 0
}

// Codepoint returns the codepoint associated with the glyph.
func (glyph FontGlyph) Codepoint() int {
	return int(C.iggFontGlyphCodepoint(glyph.handle()))
}

// AdvanceX returns the horizontal difference to the next character after the glyph is drawn.
func (glyph FontGlyph) AdvanceX() float32 {
	return float32(C.iggFontGlyphAdvanceX(glyph.handle()))
}

// X0 returns the lower x coordinate of the glyph.
func (glyph FontGlyph) X0() float32 {
	return float32(C.iggFontGlyphX0(glyph.handle()))
}

// Y0 returns the lower y coordinate of the glyph.
func (glyph FontGlyph) Y0() float32 {
	return float32(C.iggFontGlyphY0(glyph.handle()))
}

// X1 returns the upper x coordinate of the glyph.
func (glyph FontGlyph) X1() float32 {
	return float32(C.iggFontGlyphX1(glyph.handle()))
}

// Y1 returns the upper y coordinate of the glyph.
func (glyph FontGlyph) Y1() float32 {
	return float32(C.iggFontGlyphY1(glyph.handle()))
}

// U0 returns the u texture coordinate associated with the X0() vertex of the glyph.
func (glyph FontGlyph) U0() float32 {
	return float32(C.iggFontGlyphU0(glyph.handle()))
}

// V0 returns the v texture coordinate associated with the Y0() vertex of the glyph.
func (glyph FontGlyph) V0() float32 {
	return float32(C.iggFontGlyphV0(glyph.handle()))
}

// U1 returns the u texture coordinate associated with the X1() vertex of the glyph.
func (glyph FontGlyph) U1() float32 {
	return float32(C.iggFontGlyphU1(glyph.handle()))
}

// V1 returns the v texture coordinate associated with the Y1() vertex of the glyph.
func (glyph FontGlyph) V1() float32 {
	return float32(C.iggFontGlyphV1(glyph.handle()))
}
