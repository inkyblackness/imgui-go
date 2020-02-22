package imgui

import "image/color"

// PackedColor is a 32-bit RGBA color value, with 8 bits per color channel.
// The bytes are assigned as 0xAARRGGBB.
type PackedColor uint32

const (
	packedRedShift   = 16
	packedGreenShift = 8
	packedBlueShift  = 0
	packedAlphaShift = 24
)

// RGBA implements the color.Color interface.
func (clr PackedColor) RGBA() (r, g, b, a uint32) {
	return color.NRGBA{
		R: uint8(uint32(clr) >> packedRedShift),
		G: uint8(uint32(clr) >> packedGreenShift),
		B: uint8(uint32(clr) >> packedBlueShift),
		A: uint8(uint32(clr) >> packedAlphaShift),
	}.RGBA()
}

// PackedColorModel converts colors to PackedColor instances.
var PackedColorModel = color.ModelFunc(func(in color.Color) color.Color { return Packed(in) })

// Packed converts the given color to a PackedColor instance.
func Packed(c color.Color) PackedColor {
	nrgba := color.NRGBAModel.Convert(c).(color.NRGBA)
	return PackedColor(0 |
		uint32(nrgba.R)<<packedRedShift |
		uint32(nrgba.G)<<packedGreenShift |
		uint32(nrgba.B)<<packedBlueShift |
		uint32(nrgba.A)<<packedAlphaShift)
}
