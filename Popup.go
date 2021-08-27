package imgui

// #include "wrapper/Popup.h"
import "C"

// PopupFlags Flags for OpenPopup*(), BeginPopupContext*(), IsPopupOpen() functions.
// - To be backward compatible with older API which took an 'int mouse_button = 1' argument, we need to treat
//   small flags values as a mouse button index, so we encode the mouse button in the first few bits of the flags.
//   It is therefore guaranteed to be legal to pass a mouse button index in ImGuiPopupFlags.
// - For the same reason, we exceptionally default the ImGuiPopupFlags argument of BeginPopupContextXXX functions to 1 instead of 0.
//   IMPORTANT: because the default parameter is 1 (==ImGuiPopupFlags_MouseButtonRight), if you rely on the default parameter
//   and want to another another flag, you need to pass in the ImGuiPopupFlags_MouseButtonRight flag.
// - Multiple buttons currently cannot be combined/or-ed in those functions (we could allow it later).
type PopupFlags int

const (
	// PopupFlagsNone no popup flags apply.
	PopupFlagsNone PopupFlags = 0
	// PopupFlagsMouseButtonLeft For BeginPopupContext*(): open on Left Mouse release. Guaranteed to always be == 0 (same as ImGuiMouseButton_Left).
	PopupFlagsMouseButtonLeft PopupFlags = 0
	// PopupFlagsMouseButtonRight For BeginPopupContext*(): open on Right Mouse release. Guaranteed to always be == 1 (same as ImGuiMouseButton_Right).
	PopupFlagsMouseButtonRight PopupFlags = 1
	// PopupFlagsMouseButtonMiddle For BeginPopupContext*(): open on Middle Mouse release. Guaranteed to always be == 2 (same as ImGuiMouseButton_Middle).
	PopupFlagsMouseButtonMiddle PopupFlags = 2
	// PopupFlagsNoOpenOverExistingPopup For OpenPopup*(), BeginPopupContext*(): don't open if there's already a popup at the same level of the popup stack.
	PopupFlagsNoOpenOverExistingPopup PopupFlags = 1 << 5
	// PopupFlagsNoOpenOverItems For BeginPopupContextWindow(): don't return true when hovering items, only when hovering empty space.
	PopupFlagsNoOpenOverItems PopupFlags = 1 << 6
	// PopupFlagsAnyPopupID For IsPopupOpen(): ignore the ImGuiID parameter and test for any popup.
	PopupFlagsAnyPopupID PopupFlags = 1 << 7
	// PopupFlagsAnyPopupLevel For IsPopupOpen(): search/test at any level of the popup stack (default test in the current level).
	PopupFlagsAnyPopupLevel PopupFlags = 1 << 8
	// PopupFlagsAnyPopup for any usage.
	PopupFlagsAnyPopup = PopupFlagsAnyPopupID | PopupFlagsAnyPopupLevel
)

// BeginPopupV returns true if the popup is open, and you can start outputting to it.
// Only call EndPopup() if BeginPopup() returns true.
// WindowFlags are forwarded to the window.
func BeginPopupV(name string, flags WindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()
	return C.iggBeginPopup(nameArg, C.int(flags)) != 0
}

// BeginPopup calls BeginPopupV(name, nil, 0).
func BeginPopup(name string) bool {
	return BeginPopupV(name, 0)
}

// BeginPopupModalV creates modal dialog (regular window with title bar, block interactions behind the modal window,
// can't close the modal window by clicking outside).
// WindowFlags are forwarded to the window.
func BeginPopupModalV(name string, open *bool, flags WindowFlags) bool {
	nameArg, nameFin := wrapString(name)
	defer nameFin()
	openArg, openFin := wrapBool(open)
	defer openFin()
	return C.iggBeginPopupModal(nameArg, openArg, C.int(flags)) != 0
}

// BeginPopupModal calls BeginPopupModalV(name, nil, 0).
func BeginPopupModal(name string) bool {
	return BeginPopupModalV(name, nil, 0)
}

// EndPopup finishes a popup. Only call EndPopup() if BeginPopupXXX() returns true!
func EndPopup() {
	C.iggEndPopup()
}

// OpenPopupV marks popup as open (don't call every frame!).
// Popups are closed when user click outside, or if CloseCurrentPopup() is called within a BeginPopup()/EndPopup() block.
// By default, Selectable()/MenuItem() are calling CloseCurrentPopup().
// Popup identifiers are relative to the current ID-stack (so OpenPopup and BeginPopup needs to be at the same level).
func OpenPopupV(id string, flags PopupFlags) {
	idArg, idFin := wrapString(id)
	defer idFin()
	C.iggOpenPopup(idArg, C.int(flags))
}

// OpenPopup calls OpenPopupV(id, 0).
func OpenPopup(id string) {
	OpenPopupV(id, 0)
}

// OpenPopupOnItemClickV helper to open popup when clicked on last item. return true when just opened. (note: actually triggers on the mouse _released_ event to be consistent with popup behaviors).
func OpenPopupOnItemClickV(id string, flags PopupFlags) {
	idArg, idFin := wrapString(id)
	defer idFin()
	C.iggOpenPopupOnItemClick(idArg, C.int(flags))
}

// OpenPopupOnItemClick calls OpenPopupOnItemClickV("", PopupFlagsMouseButtonRight).
func OpenPopupOnItemClick() {
	OpenPopupOnItemClickV("", PopupFlagsMouseButtonRight)
}

// CloseCurrentPopup closes the popup we have begin-ed into.
// Clicking on a MenuItem or Selectable automatically close the current popup.
func CloseCurrentPopup() {
	C.iggCloseCurrentPopup()
}

// BeginPopupContextItemV returns true if the identified mouse button was pressed
// while hovering over the last item.
func BeginPopupContextItemV(id string, flags PopupFlags) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	return C.iggBeginPopupContextItem(idArg, C.int(flags)) != 0
}

// BeginPopupContextItem calls BeginPopupContextItemV("", PopupFlagsMouseButtonRight).
func BeginPopupContextItem() bool {
	return BeginPopupContextItemV("", PopupFlagsMouseButtonRight)
}

// BeginPopupContextWindowV open+begin popup when clicked on current window.
func BeginPopupContextWindowV(id string, flags PopupFlags) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	return C.iggBeginPopupContextWindow(idArg, C.int(flags)) != 0
}

// BeginPopupContextWindow calls BeginPopupContextWindowV("", PopupFlagsMouseButtonRight).
func BeginPopupContextWindow() bool {
	return BeginPopupContextWindowV("", PopupFlagsMouseButtonRight)
}

// BeginPopupContextVoidV open+begin popup when clicked in void (where there are no windows).
func BeginPopupContextVoidV(id string, flags PopupFlags) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	return C.iggBeginPopupContextVoid(idArg, C.int(flags)) != 0
}

// BeginPopupContextVoid calls BeginPopupContextVoidV("", PopupFlagsMouseButtonRight).
func BeginPopupContextVoid() bool {
	return BeginPopupContextVoidV("", PopupFlagsMouseButtonRight)
}

// IsPopupOpenV return true if the popup is open.
// IsPopupOpenV(id, PopupFlagsNone): return true if the popup is open at the current BeginPopup() level of the popup stack.
// IsPopupOpenV(id, PopupFlagsAnyPopupID: return true if any popup is open at the current BeginPopup() level of the popup stack.
// IsPopupOpenV(id, PopupFlagsAnyPopup): return true if any popup is open.
func IsPopupOpenV(id string, flags PopupFlags) bool {
	idArg, idFin := wrapString(id)
	defer idFin()
	return C.iggIsPopupOpen(idArg, C.int(flags)) != 0
}

// IsPopupOpen calls IsPopupOpenV(id, PopupFlagsNone).
func IsPopupOpen(id string) bool {
	return IsPopupOpenV(id, PopupFlagsNone)
}
