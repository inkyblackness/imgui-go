package imgui

// #include "FontAtlasWrapper.h"
import "C"
import "unsafe"

// Alpha8Image represents a imgui backed 8-bit alpha value image.
type Alpha8Image struct {
	Width, Height int
	Pixels        unsafe.Pointer
}

// FontAtlas contains runtime data for multiple fonts,
// bake multiple fonts into a single texture, TTF/OTF font loader
type FontAtlas uintptr

func (atlas FontAtlas) handle() C.IggFontAtlas {
	return C.IggFontAtlas(atlas)
}

// AddFontDefault adds the default font to the atlas. This is done by default if you do not call any
// of the AddFont* methods before retrieving the texture data.
func (atlas FontAtlas) AddFontDefault() Font {
	fontHandle := C.iggAddFontDefault(atlas.handle())
	return Font(fontHandle)
}

// AddFontFromFileTTF attempts to load a font from given TTF file.
func (atlas FontAtlas) AddFontFromFileTTF(filename string, sizePixels float32) Font {
	filenameArg, filenameFin := wrapString(filename)
	defer filenameFin()
	fontHandle := C.iggAddFontFromFileTTF(atlas.handle(), filenameArg, C.float(sizePixels))
	return Font(fontHandle)
}

// TextureDataAlpha8 returns the image in 8-bit alpha values for the font atlas.
// The returned image is valid as long as the font atlas is.
func (atlas FontAtlas) TextureDataAlpha8() *Alpha8Image {
	var pixels *C.uchar
	var width C.int
	var height C.int
	var bytesPerPixel C.int
	C.iggFontAtlasGetTexDataAsAlpha8(atlas.handle(), &pixels, &width, &height, &bytesPerPixel)

	return &Alpha8Image{
		Width:  int(width),
		Height: int(height),
		Pixels: unsafe.Pointer(pixels), // nolint: gas
	}
}

// SetTextureID sets user data to refer to the texture once it has been uploaded to user's graphic systems.
// It is passed back to you during rendering via the DrawCommand.
func (atlas FontAtlas) SetTextureID(id TextureID) {
	C.iggFontAtlasSetTextureID(atlas.handle(), id.handle())
}
