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

func wrapFloat(goValue *float32) (wrapped *C.float, finisher func()) {
	if goValue != nil {
		cValue := C.float(*goValue)
		wrapped = &cValue
		finisher = func() {
			*goValue = float32(cValue)
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

func wrapStringBuffer(value *string, bufSize uint32) (wrapped *C.char, finisher func()) {
	buf := C.malloc(C.size_t(bufSize))
	if buf == nil {
		panic("out-of-memory allocating buffer")
	}

	copy(((*[1 << 30]byte)(buf))[:bufSize-1], []byte(*value))
	textLen := len(*value)
	if uint32(textLen) < bufSize {
		((*[1 << 30]byte)(buf))[textLen] = 0
	}

	return (*C.char)(buf), func() {
		*value = C.GoString((*C.char)(buf))
		C.free(unsafe.Pointer(buf))
	}
}
