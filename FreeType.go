package imgui

// FreeTypeError describes a problem with FreeType font rendering.
type FreeTypeError string

// Error returns the readable text presentation of the error.
func (err FreeTypeError) Error() string {
	return string(err)
}

const (
	// ErrFreeTypeNotAvailable is used if the implementation of freetype is not available in this build.
	ErrFreeTypeNotAvailable = FreeTypeError("Not available for this build")
	// ErrFreeTypeFailed is used if building a font atlas was not possible.
	ErrFreeTypeFailed = FreeTypeError("Failed to build FontAtlas with FreeType")
)

// BuildWithFreeTypeV builds the FontAtlas using FreeType instead of the default rasterizer.
// FreeType renders small fonts better.
// Call this function instead of FontAtlas.Build() . As with FontAtlas.Build(), this function
// needs to be called before retrieving the texture data.
//
// FreeType support must be enabled with the build tag "imguifreetype".
func (atlas FontAtlas) BuildWithFreeTypeV(flags int) error {
	return atlas.buildWithFreeType(flags)
}

// BuildWithFreeType calls BuildWithFreeTypeV(0).
func (atlas FontAtlas) BuildWithFreeType() error {
	return atlas.BuildWithFreeTypeV(0)
}
