package imgui

// #include "wrapper/Scroll.h"
import "C"

// GetScrollX is deprecated and will be removed in v3.0.0 .
// Deprecated: Use ScrollX() instead.
func GetScrollX() float32 {
	return ScrollX()
}

// ScrollX returns the horizontal scrolling amount [0..GetScrollMaxX()].
func ScrollX() float32 {
	return float32(C.iggGetScrollX())
}

// GetScrollY is deprecated and will be removed in v3.0.0 .
// Deprecated: Use ScrollY() instead.
func GetScrollY() float32 {
	return ScrollY()
}

// ScrollY returns the vertical scrolling amount [0..GetScrollMaxY()].
func ScrollY() float32 {
	return float32(C.iggGetScrollY())
}

// GetScrollMaxX is deprecated and will be removed in v3.0.0 .
// Deprecated: Use ScrollMaxX() instead.
func GetScrollMaxX() float32 {
	return ScrollMaxX()
}

// ScrollMaxX returns the maximum horizontal scrolling amount: ContentSize.X - WindowSize.X .
func ScrollMaxX() float32 {
	return float32(C.iggGetScrollMaxX())
}

// GetScrollMaxY is deprecated and will be removed in v3.0.0 .
// Deprecated: Use ScrollMaxY() instead.
func GetScrollMaxY() float32 {
	return ScrollMaxY()
}

// ScrollMaxY returns the maximum vertical scrolling amount: ContentSize.Y - WindowSize.Y .
func ScrollMaxY() float32 {
	return float32(C.iggGetScrollMaxY())
}

// SetScrollHereX adjusts horizontal scrolling amount to make current cursor
// position visible.  ratio=0.0: left, 0.5: center, 1.0: right.  When using to
// make a "default/current item" visible, consider using SetItemDefaultFocus()
// instead.
func SetScrollHereX(ratio float32) {
	C.iggSetScrollHereX(C.float(ratio))
}

// SetScrollHereY adjusts vertical scrolling amount to make current cursor
// position visible.  ratio=0.0: top, 0.5: center, 1.0: bottom.  When using to
// make a "default/current item" visible, consider using SetItemDefaultFocus()
// instead.
func SetScrollHereY(ratio float32) {
	C.iggSetScrollHereY(C.float(ratio))
}

// SetScrollX sets horizontal scrolling amount [0..GetScrollMaxX()].
func SetScrollX(scrollX float32) {
	C.iggSetScrollX(C.float(scrollX))
}

// SetScrollY sets vertical scrolling amount [0..GetScrollMaxY()].
func SetScrollY(scrollY float32) {
	C.iggSetScrollY(C.float(scrollY))
}
