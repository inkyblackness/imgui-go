package imgui

import (
	"image/color"
	"math"
)

// PackedColor is a 32-bit RGBA color value, with 8 bits per color channel.
// The bytes are assigned as 0xAABBGGRR.
type PackedColor uint32

const (
	packedRedShift   = 0
	packedGreenShift = 8
	packedBlueShift  = 16
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

// PackedColorFromVec4 converts the given four-dimensional vector into a packed color.
func PackedColorFromVec4(vec Vec4) PackedColor {
	convert := func(f float32) uint32 {
		scaled := (f * math.MaxUint8) + 0.5 // nolint: gomnd
		switch {
		case scaled <= 0:
			return 0
		case scaled >= math.MaxUint8:
			return math.MaxUint8
		default:
			return uint32(scaled)
		}
	}
	return PackedColor(0 |
		convert(vec.X)<<packedRedShift |
		convert(vec.Y)<<packedGreenShift |
		convert(vec.Z)<<packedBlueShift |
		convert(vec.W)<<packedAlphaShift)
}
