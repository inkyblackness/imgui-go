package imgui

// #include "imguiWrapperTypes.h"
import "C"

// FontConfig describes properties of a single font.
type FontConfig uintptr

// DefaultFontConfig lets ImGui take default properties as per implementation.
const DefaultFontConfig FontConfig = 0

func (config FontConfig) handle() C.IggFontConfig {
	return C.IggFontConfig(config)
}
