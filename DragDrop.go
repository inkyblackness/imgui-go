package imgui

// #include "wrapper/DragDrop.h"
import "C"

// This is a list of BeginDragDropSource flags.
const (
	// DragDropFlagsNone specifies the default behaviour.
	DragDropFlagsNone = 0
	// DragDropFlagsSourceNoPreviewTooltip hides the tooltip that is open so you can display a preview or description of the source contents.
	DragDropFlagsSourceNoPreviewTooltip = 1 << 0
	// DragDropFlagsSourceNoDisableHover preserves the behaviour of IsItemHovered. By default, when dragging we clear data so that IsItemHovered() will return true, to avoid subsequent user code submitting tooltips.
	DragDropFlagsSourceNoDisableHover = 1 << 1
	// DragDropFlagsSourceNoHoldToOpenOthers disables the behavior that allows to open tree nodes and collapsing header by holding over them while dragging a source item.
	DragDropFlagsSourceNoHoldToOpenOthers = 1 << 2
	// DragDropFlagsSourceAllowNullID allows items such as Text(), Image() that have no unique identifier to be used as drag source, by manufacturing a temporary identifier based on their window-relative position. This is extremely unusual within the dear ecosystem and so we made it explicit.
	DragDropFlagsSourceAllowNullID = 1 << 3
	// DragDropFlagsSourceExtern specifies external source (from outside of), won't attempt to read current item/window info. Will always return true. Only one Extern source can be active simultaneously.
	DragDropFlagsSourceExtern = 1 << 4
)

// This is a list of AcceptDragDropPayload flags.
const (
	// DragDropFlagsAcceptBeforeDelivery makes AcceptDragDropPayload() return true even before the mouse button is released. You can then call IsDelivery() to test if the payload needs to be delivered.
	DragDropFlagsAcceptBeforeDelivery = 1 << 10
	// DragDropFlagsAcceptNoDrawDefaultRect does not draw the default highlight rectangle when hovering over target.
	DragDropFlagsAcceptNoDrawDefaultRect = 1 << 11
	// DragDropFlagsAcceptPeekOnly is for peeking ahead and inspecting the payload before delivery.
	DragDropFlagsAcceptPeekOnly = DragDropFlagsAcceptBeforeDelivery | DragDropFlagsAcceptNoDrawDefaultRect
)

// BeginDragDropSource registers the currently active item as drag'n'drop source.
// When this returns true you need to:
// a) call SetDragDropPayload() exactly once,
// b) you may render the payload visual/description,
// c) call EndDragDropSource().
func BeginDragDropSource(flags int) bool {
	return C.iggBeginDragDropSource(C.int(flags)) != 0
}

// SetDragDropPayload sets the payload for current draw and drop source.
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
