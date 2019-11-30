package imgui

//#include "imguiWrapperTypes.h"
//#include "imguiWrapper.h"
import "C"

type ListClipper struct {
	StartPosY    float32
	ItemsHeight  float32
	ItemsCount   int
	StepNo       int
	DisplayStart int
	DisplayEnd   int
}

// wrapped return C struct and func for setting the values when done
func (clipper *ListClipper) wrapped() (out *C.IggListClipper, finisher func()) {
	if clipper != nil {
		out = &C.IggListClipper{
			StartPosY:    C.float(clipper.StartPosY),
			ItemsHeight:  C.float(clipper.ItemsHeight),
			ItemsCount:   C.int(clipper.ItemsCount),
			StepNo:       C.int(clipper.StepNo),
			DisplayStart: C.int(clipper.DisplayStart),
			DisplayEnd:   C.int(clipper.DisplayEnd),
		}
		finisher = func() {
			clipper.StartPosY = float32(out.StartPosY)     // nolint: gotype
			clipper.ItemsHeight = float32(out.ItemsHeight) // nolint: gotype
			clipper.ItemsCount = int(out.ItemsCount)       // nolint: gotype
			clipper.StepNo = int(out.StepNo)               // nolint: gotype
			clipper.DisplayStart = int(out.DisplayStart)   // nolint: gotype
			clipper.DisplayEnd = int(out.DisplayEnd)       // nolint: gotype
		}
	} else {
		finisher = func() {}
	}
	return
}

// Step Call until it returns false. The DisplayStart/DisplayEnd fields will be set and you can process/draw those items.
func (clipper *ListClipper) Step() bool {
	arg, finnishFunc := clipper.wrapped()
	defer finnishFunc()
	return C.iggListClipperStep(arg) != 0
}

// Begin Called before stepping
func (clipper *ListClipper) Begin(itemsCount int, itemsHeight float32) {
	arg, finnishFunc := clipper.wrapped()
	defer finnishFunc()
	C.iggListClipperBegin(arg, C.int(itemsCount), C.float(itemsHeight))
}

// End Automatically called on the last call of Step() that returns false.
func (clipper *ListClipper) End() {
	arg, finnishFunc := clipper.wrapped()
	defer finnishFunc()
	C.iggListClipperEnd(arg)
}
