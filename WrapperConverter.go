package imgui

// #include "wrapper/Types.h"
// #include <memory.h>
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

func wrapBytes(value []byte) (wrapped unsafe.Pointer, finisher func()) {
	wrapped = C.CBytes(value)
	finisher = func() { C.free(wrapped) } // nolint: gas
	return
}

type stringBuffer struct {
	ptr  unsafe.Pointer
	size int
}

func newStringBuffer(initialValue string) *stringBuffer {
	rawText := []byte(initialValue)
	bufSize := len(rawText) + 1
	newPtr := C.malloc(C.size_t(bufSize))
	zeroOffset := bufSize - 1
	buf := ptrToByteSlice(newPtr)
	copy(buf[:zeroOffset], rawText)
	buf[zeroOffset] = 0

	return &stringBuffer{ptr: newPtr, size: bufSize}
}

func (buf *stringBuffer) free() {
	C.free(buf.ptr)
	buf.size = 0
}

func (buf *stringBuffer) resizeTo(requestedSize int) {
	bufSize := requestedSize
	if bufSize < 1 {
		bufSize = 1
	}
	newPtr := C.malloc(C.size_t(bufSize))
	copySize := bufSize
	if copySize > buf.size {
		copySize = buf.size
	}
	if copySize > 0 {
		C.memcpy(newPtr, buf.ptr, C.size_t(copySize))
	}
	ptrToByteSlice(newPtr)[bufSize-1] = 0
	C.free(buf.ptr)
	buf.ptr = newPtr
	buf.size = bufSize
}

func (buf stringBuffer) toGo() string {
	if (buf.ptr == nil) || (buf.size < 1) {
		return ""
	}
	ptrToByteSlice(buf.ptr)[buf.size-1] = 0
	return C.GoString((*C.char)(buf.ptr))
}

// unrealisticLargePointer is used to cast an arbitrary native pointer to a slice.
// Its value is chosen to fit into a 32bit architecture, and still be large
// enough to cover "any" data blob. Note that this value is in bytes.
// Should an array of larger primitives be addressed, be sure to divide the value
// by the size of the elements.
const unrealisticLargePointer = 1 << 30

func ptrToByteSlice(p unsafe.Pointer) []byte {
	return (*[unrealisticLargePointer]byte)(p)[:]
}

func ptrToUint16Slice(p unsafe.Pointer) []uint16 {
	return (*[unrealisticLargePointer / 2]uint16)(p)[:]
}
