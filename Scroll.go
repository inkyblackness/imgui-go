package imgui

// #include "wrapper/Scroll.h"
import "C"

// GetScrollX returns the horizontal scrolling amount [0..GetScrollMaxX()].
func GetScrollX() float32 {
	return float32(C.iggGetScrollX())
}

// GetScrollY returns the vertical scrolling amount [0..GetScrollMaxY()].
func GetScrollY() float32 {
	return float32(C.iggGetScrollY())
}

// GetScrollMaxX returns the maximum horizontal scrolling amount: ContentSize.X
// - WindowSize.X .
func GetScrollMaxX() float32 {
	return float32(C.iggGetScrollMaxX())
}

// GetScrollMaxY returns the maximum vertical scrolling amount: ContentSize.Y -
// WindowSize.Y .
func GetScrollMaxY() float32 {
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
