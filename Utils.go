package imgui

// #include "wrapper/Utils.h"
import "C"

// Spinner draws a rotating spinner
func Spinner(label string, radius float32, thickness int) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	C.iggSpinner(labelArg, C.float(radius), C.int(thickness))
}

// Splitter adds a movable splitter between to childs
func Splitter(splitVertically bool, thickness float32, size1 *float32, size2 *float32) bool {
	size1Arg, size1Fin := wrapFloat(size1)
	defer size1Fin()
	size2Arg, size2Fin := wrapFloat(size2)
	defer size2Fin()
	return C.iggSplitter(castBool(splitVertically), C.float(thickness), size1Arg, size2Arg) != 0
}
