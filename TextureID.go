package imgui

// #include "imguiWrapperTypes.h"
import "C"

// TextureID is a user data to identify a texture.
type TextureID uintptr

func (id TextureID) handle() C.IggTextureID {
	return C.IggTextureID(id)
}
