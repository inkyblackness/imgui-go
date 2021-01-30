package imgui

// #include "wrapper/ListClipper.h"
import "C"

// ListClipper is a helper to manually clip large list of items.
// If you are submitting lots of evenly spaced items and you have a random access to the list, you can perform coarse
// clipping based on visibility to save yourself from processing those items at all.
// The clipper calculates the range of visible items and advance the cursor to compensate for the non-visible items we have skipped.
// (Dear ImGui already clip items based on their bounds but it needs to measure text size to do so, whereas manual coarse clipping before submission makes this cost and your own data fetching/submission cost almost null)
// Usage:
//   var clipper imgui.ListClipper
//   clipper.Begin(1000) // We have 1000 elements, evenly spaced.
//   for clipper.Step() {
//       for i := clipper.DisplayStart; i < clipper.DisplayEnd; i += 1 {
//           imgui.Text(fmt.Sprintf("line number %d", i))
//       }
//   }
// Generally what happens is:
// - Clipper lets you process the first element (DisplayStart = 0, DisplayEnd = 1) regardless of it being visible or not.
// - User code submit one element.
// - Clipper can measure the height of the first element
// - Clipper calculate the actual range of elements to display based on the current clipping rectangle, position the cursor before the first visible element.
// - User code submit visible elements.
type ListClipper struct {
	DisplayStart int
	DisplayEnd   int

	// [Internal]
	ItemsCount  int
	StepNo      int
	ItemsFrozen int
	ItemsHeight float32
	StartPosY   float32
}

// wrapped returns C struct and func for setting the values when done.
func (clipper *ListClipper) wrapped() (out *C.IggListClipper, finisher func()) {
	if clipper == nil {
		return nil, func() {}
	}
	out = &C.IggListClipper{
		DisplayStart: C.int(clipper.DisplayStart),
		DisplayEnd:   C.int(clipper.DisplayEnd),

		ItemsCount:  C.int(clipper.ItemsCount),
		StepNo:      C.int(clipper.StepNo),
		ItemsFrozen: C.int(clipper.ItemsFrozen),
		ItemsHeight: C.float(clipper.ItemsHeight),
		StartPosY:   C.float(clipper.StartPosY),
	}
	finisher = func() {
		clipper.DisplayStart = int(out.DisplayStart)
		clipper.DisplayEnd = int(out.DisplayEnd)

		clipper.ItemsCount = int(out.ItemsCount)
		clipper.StepNo = int(out.StepNo)
		clipper.ItemsFrozen = int(out.ItemsFrozen)
		clipper.ItemsHeight = float32(out.ItemsHeight)
		clipper.StartPosY = float32(out.StartPosY)
	}
	return
}

// Step must be called in a loop until it returns false.
// The DisplayStart/DisplayEnd fields will be set and you can process/draw those items.
func (clipper *ListClipper) Step() bool {
	arg, finishFunc := clipper.wrapped()
	defer finishFunc()
	return C.iggListClipperStep(arg) != 0
}

// Begin calls BeginV(itemsCount, -1.0) .
func (clipper *ListClipper) Begin(itemsCount int) {
	clipper.BeginV(itemsCount, -1.0)
}

// BeginV must be called before stepping.
// Use an itemCount of math.MaxInt if you don't know how many items you have.
// In this case the cursor won't be advanced in the final step.
//
// For itemsHeight, use -1.0 to be calculated automatically on first step.
// Otherwise pass in the distance between your items, typically
// GetTextLineHeightWithSpacing() or GetFrameHeightWithSpacing().
func (clipper *ListClipper) BeginV(itemsCount int, itemsHeight float32) {
	arg, finishFunc := clipper.wrapped()
	defer finishFunc()
	C.iggListClipperBegin(arg, C.int(itemsCount), C.float(itemsHeight))
}

// End resets the clipper. This function is automatically called on the last call of Step() that returns false.
func (clipper *ListClipper) End() {
	arg, finishFunc := clipper.wrapped()
	defer finishFunc()
	C.iggListClipperEnd(arg)
}
