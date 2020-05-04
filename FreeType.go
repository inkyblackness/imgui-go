package imgui

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
