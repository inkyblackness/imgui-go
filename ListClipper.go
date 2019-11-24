package imgui

// #include "imguiWrapper.h"
import "C"

// items_count:  Use -1 to ignore (you can call Begin later). Use INT_MAX if you don't know how many items you have (in which case the cursor won't be advanced in the final step).
// items_height: Use -1.0f to be calculated automatically on first step. Otherwise pass in the distance between your items, typically GetTextLineHeightWithSpacing() or GetFrameHeightWithSpacing().
// If you don't specify an items_height, you NEED to call Step(). If you specify items_height you may call the old Begin()/End() api directly, but prefer calling Step().
type ListClipper struct {
	CListClipper C.IggListClipper
}

func New(itemsCount int, itemsHeight float32) ListClipper {
	clipper := ListClipper{}
	clipper.CListClipper = C.iggListClipperInit(C.int(itemsCount), C.float(itemsHeight))
	return clipper
}

// Step Call until it returns false. The DisplayStart/DisplayEnd fields will be set and you can process/draw those items.
func (clipper *ListClipper) Step() bool {
	return C.iggListClipperStep(&clipper.CListClipper) != 0
}

// Begin Automatically called by constructor if you passed 'items_count' or by Step() in Step 1.
func (clipper *ListClipper) Begin(itemsCount int, itemsHeight float32) {
	C.iggListClipperBegin(&clipper.CListClipper, C.int(itemsCount), C.float(itemsHeight))
}

// End Automatically called on the last call of Step() that returns false.
func (clipper *ListClipper) End() {
	C.iggListClipperEnd(&clipper.CListClipper)
}

// GetDisplayStart returns DisplayStart
func (clipper *ListClipper) GetDisplayStart() int {
	return int(C.iggListClipperDisplayStart(&clipper.CListClipper))
}

// GetDisplayEnd returns DisplayEnd
func (clipper *ListClipper) GetDisplayEnd() int {
	return int(C.iggListClipperDisplayEnd(&clipper.CListClipper))
}

// GetStepNo returns StepNo
func (clipper *ListClipper) GetStepNo() int {
	return int(C.iggListClipperStepNo(&clipper.CListClipper))
}

// GetItemsCount returns ItemsCount
func (clipper *ListClipper) GetItemsCount() int {
	return int(C.iggListClipperItemsCount(&clipper.CListClipper))
}

// GetStartPosY returns StartPosY
func (clipper *ListClipper) GetStartPosY() float32 {
	return float32(C.iggListClipperStartPosY(&clipper.CListClipper))
}

// GetItemsHeight returns ItemsHeight
func (clipper *ListClipper) GetItemsHeight() float32 {
	return float32(C.iggListClipperItemsHeight(&clipper.CListClipper))
}

// SetDisplayStart sets DisplayStart
func (clipper *ListClipper) SetDisplayStart(value int) {
	C.iggSetListClipperDisplayStart(&clipper.CListClipper, C.int(value))
}

// SetDisplayEnd sets DisplayEnd
func (clipper *ListClipper) SetDisplayEnd(value int) {
	C.iggSetListClipperDisplayEnd(&clipper.CListClipper, C.int(value))
}

// SetStepNo sets StepNo
func (clipper *ListClipper) SetStepNo(value int) {
	C.iggSetListClipperStepNo(&clipper.CListClipper, C.int(value))
}

// SetItemsCount sets ItemsCount
func (clipper *ListClipper) SetItemsCount(value int) {
	C.iggSetListClipperItemsCount(&clipper.CListClipper, C.int(value))
}

// SetStartPosY sets StartPosY
func (clipper *ListClipper) SetStartPosY(value float32) {
	C.iggSetListClipperStartPosY(&clipper.CListClipper, C.float(value))
}

// SetItemsHeight sets ItemsHeight
func (clipper *ListClipper) SetItemsHeight(value float32) {
	C.iggSetListClipperItemsHeight(&clipper.CListClipper, C.float(value))
}
