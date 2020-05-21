package imgui

// #include "wrapper/Popup.h"
import "C"

// OpenPopup marks popup as open (don't call every frame!).
// Popups are closed when user click outside, or if CloseCurrentPopup() is called within a BeginPopup()/EndPopup() block.
// By default, Selectable()/MenuItem() are calling CloseCurrentPopup().
// Popup identifiers are relative to the current ID-stack (so OpenPopup and BeginPopup needs to be at the same level).
func OpenPopup(id string) {
	idArg, idFin := wrapString(id)
	defer idFin()
	C.iggOpenPopup(idArg)
}

// BeginPopupV returns true if the popup is open, and you can start outputting to it.
// Only call EndPopup() if BeginPopup() returns true.
func BeginPopupV(name string, flags int) bool {
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
func BeginPopupModalV(name string, open *bool, flags int) bool {
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

// BeginPopupContextItemV returns true if the identified mouse button was pressed
// while hovering over the last item.
func BeginPopupContextItemV(label string, mouseButton int) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	return C.iggBeginPopupContextItem(labelArg, C.int(mouseButton)) != 0
}

// BeginPopupContextItem calls BeginPopupContextItemV("", 1).
func BeginPopupContextItem() bool {
	return BeginPopupContextItemV("", 1)
}

// EndPopup finshes a popup. Only call EndPopup() if BeginPopupXXX() returns true!
func EndPopup() {
	C.iggEndPopup()
}

// CloseCurrentPopup closes the popup we have begin-ed into.
// Clicking on a MenuItem or Selectable automatically close the current popup.
func CloseCurrentPopup() {
	C.iggCloseCurrentPopup()
}
