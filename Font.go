package imgui

// #include "wrapper/Font.h"
// #include "wrapper/Types.h"
import "C"

// Font describes one loaded font in an atlas.
type Font uintptr

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

	C.iggCalcTextSize((*C.char)(CString.ptr), C.int(CString.size), castBool(hideTextAfterDoubleHash), C.float(wrapWidth), valueArg)
	returnFunc()

	return vec2
}

func (font Font) handle() C.IggFont {
	return C.IggFont(font)
}
