// +build imguifreetype

package imgui

// #cgo pkg-config: freetype2
// #cgo CXXFLAGS: -DIMGUI_ENABLE_FREETYPE
// #cgo CFLAGS: -DIMGUI_ENABLE_FREETYPE
// #cgo CPPFLAGS: -DIMGUI_ENABLE_FREETYPE
import "C"

// Hinting greatly impacts visuals (and glyph sizes).
// - By default, hinting is enabled and the font's native hinter is preferred over the auto-hinter.
// - When disabled, FreeType generates blurrier glyphs, more or less matches the stb_truetype.h
// - The Default hinting mode usually looks good, but may distort glyphs in an unusual way.
// - The Light hinting mode generates fuzzier glyphs but better matches Microsoft's rasterizer.
// You can set those flags globaly in FontAtlas.SetFontBuilderFlags(flags)
// You can set those flags on a per font basis in FontConfig.SetFontBuilderFlags(flags).
const (
	// FreeTypeBuilderFlagsNoHinting disables hinting.
	// This generally generates 'blurrier' bitmap glyphs when the glyph are rendered in any of the anti-aliased modes.
	FreeTypeBuilderFlagsNoHinting = 1 << 0
	// FreeTypeBuilderFlagsNoAutoHint disables auto-hinter.
	FreeTypeBuilderFlagsNoAutoHint = 1 << 1
	// FreeTypeBuilderFlagsForceAutoHint indicates that the auto-hinter is preferred over the font's native hinter.
	FreeTypeBuilderFlagsForceAutoHint = 1 << 2
	// FreeTypeBuilderFlagsLightHinting is a lighter hinting algorithm for gray-level modes.
	// Many generated glyphs are fuzzier but better resemble their original shape.
	// This is achieved by snapping glyphs to the pixel grid only vertically (Y-axis),
	// as is done by Microsoft's ClearType and Adobe's proprietary font renderer.
	// This preserves inter-glyph spacing in horizontal text.
	FreeTypeBuilderFlagsLightHinting = 1 << 3
	// FreeTypeBuilderFlagsMonoHinting is a strong hinting algorithm that should only be used for monochrome output.
	FreeTypeBuilderFlagsMonoHinting = 1 << 4
	// FreeTypeBuilderFlagsBold is for styling: Should we artificially embolden the font?
	FreeTypeBuilderFlagsBold = 1 << 5
	// FreeTypeBuilderFlagsOblique is for styling: Should we slant the font, emulating italic style?
	FreeTypeBuilderFlagsOblique = 1 << 6
	// FreeTypeBuilderFlagsMonochrome disables anti-aliasing. Combine this with MonoHinting for best results!
	FreeTypeBuilderFlagsMonochrome = 1 << 7
	// FreeTypeBuilderFlagsLoadColor enables FreeType color-layered glyphs.
	FreeTypeBuilderFlagsLoadColor = 1 << 8
	// FreeTypeBuilderFlagsBitmap enables FreeType bitmap glyphs
	FreeTypeBuilderFlagsBitmap = 1 << 9
)
