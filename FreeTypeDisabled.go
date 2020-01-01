// +build !imguifreetype

package imgui

// #include "FreeTypeWrapper.h"
import "C"

func buildFontAtlas(handle C.IggFontAtlas, flags int) error {
	return ErrFreeTypeNotAvailable
}
