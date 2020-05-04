package imgui

// #include "wrapper/DragDrop.h"
import "C"

// BeginDragDropSource opens the scope for current draw and drop source.
// Call when current ID is active.
// When this returns true you need to: a) call SetDragDropPayload() exactly once, b) you may render the payload visual/description, c) call EndDragDropSource()
func BeginDragDropSource(flags int) bool {
	return C.iggBeginDragDropSource(C.int(flags)) != 0
}

// SetDragDropPayload sets the payload for current draw and drop source.
// type is a user defined string of maximum 32 characters.
// Strings starting with '_' are reserved for dear imgui internal types.
// Data is copied and held by imgui.
func SetDragDropPayload(dataType string, data []byte, cond Condition) bool {
	typeArg, typeFin := wrapString(dataType)
	defer typeFin()
	dataArg, dataFin := wrapBytes(data)
	defer dataFin()
	return C.iggSetDragDropPayload(typeArg, dataArg, C.int(len(data)), C.int(cond)) != 0
}

// EndDragDropSource closes the scope for current draw and drop source.
// Only call EndDragDropSource() if BeginDragDropSource() returns true.
func EndDragDropSource() {
	C.iggEndDragDropSource()
}

// BeginDragDropTarget must be called after submitting an item that may receive an item.
// If this returns true, you can call AcceptDragDropPayload() and EndDragDropTarget().
func BeginDragDropTarget() bool {
	return C.iggBeginDragDropTarget() != 0
}

// AcceptDragDropPayload accepts contents of a given type.
// If ImGuiDragDropFlags_AcceptBeforeDelivery is set you can peek into the payload before the mouse button is released.
func AcceptDragDropPayload(dataType string, flags int) []byte {
	typeArg, typeFin := wrapString(dataType)
	defer typeFin()

	payload := C.iggAcceptDragDropPayload(typeArg, C.int(flags))
	if payload == nil {
		return nil
	}

	data := C.iggPayloadData(payload)
	size := C.iggPayloadDataSize(payload)
	return C.GoBytes(data, size)
}

// EndDragDropTarget closed the scope for current drag and drop target.
// Only call EndDragDropTarget() if BeginDragDropTarget() returns true.
func EndDragDropTarget() {
	C.iggEndDragDropTarget()
}
