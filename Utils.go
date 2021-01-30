package imgui

// #include "wrapper/Utils.h"
import "C"

// BufferingBar shows a loading bar
func BufferingBar(label string, value float32, size Vec2, fg_color Vec4, bg_color Vec4) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	sizeArg, _ := size.wrapped()
	fgColorArg, _ := fg_color.wrapped()
	bgColorArg, _ := bg_color.wrapped()
	C.iggBufferingBar(labelArg, C.float(value), sizeArg, fgColorArg, bgColorArg)
}

func LoadingIndicatorCircle(label string, indicator_radius float32, circle_count int, speed float32, fg_color Vec4, bg_color Vec4) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	fgColorArg, _ := fg_color.wrapped()
	bgColorArg, _ := bg_color.wrapped()
	C.iggLoadingIndicatorCircle(labelArg, C.float(indicator_radius), C.int(circle_count), C.float(speed), fgColorArg, bgColorArg)
}

// Spinner draws a rotating spinner
func Spinner(label string, radius float32, thickness int, color Vec4) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	colorArg, _ := color.wrapped()
	C.iggSpinner(labelArg, C.float(radius), C.int(thickness), colorArg)
}

// ToggleButton creates a toggle button in the selected state.
// The return value indicates if the selected state has changed.
func ToggleButton(id string, selected *bool) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	selectedArg, selectedFin := wrapBool(selected)
	defer selectedFin()
	return C.iggToggleButton(idArg, selectedArg) != 0
}
