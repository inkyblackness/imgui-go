package imgui

// #include "imguiWrapperTypes.h"
// #include <stdlib.h>
import "C"
import "unsafe"

func castBool(value bool) (cast C.IggBool) {
	if value {
		cast = 1
	}
	return
}

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

func wrapInt32(goValue *int32) (wrapped *C.int, finisher func()) {
	if goValue != nil {
		cValue := C.int(*goValue)
		wrapped = &cValue
		finisher = func() {
			*goValue = int32(cValue)
		}
	} else {
		finisher = func() {}
	}
	return
}

func wrapString(value string) (wrapped *C.char, finisher func()) {
	wrapped = C.CString(value)
	finisher = func() { C.free(unsafe.Pointer(wrapped)) } // nolint: gas
	return
}
