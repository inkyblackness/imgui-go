package imgui

// #cgo pkg-config: freetype2
// #include "FreeTypeWrapper.h"
import "C"
import (
	"errors"
)

const (
	// By default, hinting is enabled and the font's native hinter is preferred over the auto-hinter.
	FreeTypeRasterizerFlagsNoHinting     = 1 << 0 // Disable hinting. This generally generates 'blurrier' bitmap glyphs when the glyph are rendered in any of the anti-aliased modes.
	FreeTypeRasterizerFlagsNoAutoHint    = 1 << 1 // Disable auto-hinter.
	FreeTypeRasterizerFlagsForceAutoHint = 1 << 2 // Indicates that the auto-hinter is preferred over the font's native hinter.
	FreeTypeRasterizerFlagsLightHinting  = 1 << 3 // A lighter hinting algorithm for gray-level modes. Many generated glyphs are fuzzier but better resemble their original shape. This is achieved by snapping glyphs to the pixel grid only vertically (Y-axis), as is done by Microsoft's ClearType and Adobe's proprietary font renderer. This preserves inter-glyph spacing in horizontal text.
	FreeTypeRasterizerFlagsMonoHinting   = 1 << 4 // Strong hinting algorithm that should only be used for monochrome output.
	FreeTypeRasterizerFlagsBold          = 1 << 5 // Styling: Should we artificially embolden the font?
	FreeTypeRasterizerFlagsOblique       = 1 << 6 // Styling: Should we slant the font, emulating italic style?
	FreeTypeRasterizerFlagsMonochrome    = 1 << 7 // Disable anti-aliasing. Combine this with MonoHinting for best results!
)

// BuildFontAtlasFreeType builds the FontAtlas using FreeType instead of the
// default rasterizer. FreeType renders small fonts better. Make sure to call
// `FontAtlas.BuildFontAtlasFreeType()` *BEFORE* calling
// `FontAtlas.GetTexDataAsRGBA32()` or `FontAtlas.Build()` (so normal Build()
// won't be called)
func (atlas FontAtlas) BuildFontAtlasFreeTypeV(flags int) error {
	if C.iggImGuiFreeTypeBuildFontAtlas(atlas.handle(), C.uint(flags)) == 0 {
		return errors.New("Failed to build FreeType FontAtlas")
	} else {
		return nil
	}
}

// BuildFontAtlasFreeType calls BuildFontAtlasFreeTypeV(0)
func (atlas FontAtlas) BuildFontAtlasFreeType() error {
	return atlas.BuildFontAtlasFreeTypeV(0)
}
