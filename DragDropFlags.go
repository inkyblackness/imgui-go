package imgui

const (
	// BeginDragDropSource() flags
	// DragDropFlagsSourceNoPreviewTooltip By default, a successful call to BeginDragDropSource opens a tooltip so you can display a preview or description of the source contents. This flag disable this behavior.
	DragDropFlagsSourceNoPreviewTooltip = 1 << 0
	// DragDropFlagsSourceNoDisableHover By default, when dragging we clear data so that IsItemHovered() will return true, to avoid subsequent user code submitting tooltips. This flag disable this behavior so you can still call IsItemHovered() on the source item.
	DragDropFlagsSourceNoDisableHover = 1 << 1
	// DragDropFlagsSourceNoHoldToOpenOthers Disable the behavior that allows to open tree nodes and collapsing header by holding over them while dragging a source item.
	DragDropFlagsSourceNoHoldToOpenOthers = 1 << 2
	// DragDropFlagsSourceAllowNullID Allow items such as Text(), Image() that have no unique identifier to be used as drag source, by manufacturing a temporary identifier based on their window-relative position. This is extremely unusual within the dear  ecosystem and so we made it explicit.
	DragDropFlagsSourceAllowNullID = 1 << 3
	// DragDropFlagsSourceExtern External source (from outside of ), won't attempt to read current item/window info. Will always return true. Only one Extern source can be active simultaneously.
	DragDropFlagsSourceExtern = 1 << 4
	// AcceptDragDropPayload() flags
	// DragDropFlagsAcceptBeforeDelivery AcceptDragDropPayload() will returns true even before the mouse button is released. You can then call IsDelivery() to test if the payload needs to be delivered.
	DragDropFlagsAcceptBeforeDelivery = 1 << 10
	// DragDropFlagsAcceptNoDrawDefaultRect Do not draw the default highlight rectangle when hovering over target.
	DragDropFlagsAcceptNoDrawDefaultRect = 1 << 11
	// DragDropFlagsAcceptPeekOnly For peeking ahead and inspecting the payload before delivery.
	DragDropFlagsAcceptPeekOnly = DragDropFlagsAcceptBeforeDelivery | DragDropFlagsAcceptNoDrawDefaultRect
)
