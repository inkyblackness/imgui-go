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

type ListClipper struct {
	handle       C.IggListClipper
	DisplayStart int
	DisplayEnd   int
}

func ListClipperInit(count int) ListClipper {
	var list_clipper ListClipper
	list_clipper.handle = C.iggListClipperInit(C.int(count))
	return list_clipper
}

func (clipper *ListCLipper) Step() bool {
	var display_start C.int
	var display_end C.int
	step := C.iggListClipperStep(clipper.handle, &display_start, &display_end)
	clipper.DisplayStart = int(*display_start)
	clipper.DisplayEnd = int(*display_end)
	return step != 0
}
