package imgui

// #include "imguiWrapperTypes.h"
import "C"

func wrapBool(goValue *bool) (wrapped *C.IggBool, finisher func()) {
	if goValue != nil {
		var cValue C.IggBool
		if *goValue {
			cValue = 1
		}
		wrapped = &cValue
		finisher = func() {
			*goValue = cValue != 0
		}
	} else {
		finisher = func() {}
	}
	return
}
