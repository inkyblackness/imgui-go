package imgui

// #include "wrapper/Utils.h"
import "C"

// BufferingBar shows a loading bar
func BufferingBar(label string, value float32, size Vec2, fg_color Vec4, bg_color Vec4) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	sizeArg, _ := size.wrapped()
	fgColorArg, _ := fg_color.wrapped()
	bgColorArg, _ := bg_color.wrapped()
	C.iggBufferingBar(labelArg, C.float(value), sizeArg, fgColorArg, bgColorArg)
}

func LoadingIndicatorCircle(label string, indicator_radius float32, circle_count int, speed float32, fg_color Vec4, bg_color Vec4) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	fgColorArg, _ := fg_color.wrapped()
	bgColorArg, _ := bg_color.wrapped()
	C.iggLoadingIndicatorCircle(labelArg, C.float(indicator_radius), C.int(circle_count), C.float(speed), fgColorArg, bgColorArg)
}

// Spinner draws a rotating spinner
func Spinner(label string, radius float32, thickness int, color Vec4) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()
	colorArg, _ := color.wrapped()
	C.iggSpinner(labelArg, C.float(radius), C.int(thickness), colorArg)
}

// Splitter adds a movable splitter between to childs
func Splitter(splitVertically bool, thickness float32, size1 *float32, size2 *float32) bool {
	size1Arg, size1Fin := wrapFloat(size1)
	defer size1Fin()
	size2Arg, size2Fin := wrapFloat(size2)
	defer size2Fin()
	return C.iggSplitter(castBool(splitVertically), C.float(thickness), size1Arg, size2Arg) != 0
}

// type ComboFilterState struct {
// 	ActiveIdx        int  // Index of currently 'active' item by use of up/down keys
// 	SelectionChanged bool // Flag to help focus the correct item when selecting active item
// }

// func ComboFilter(label string, text *string, hints []string, state *ComboFilterState) bool {
// 	if text == nil {
// 		panic("text can't be nil")
// 	}

// 	labelArg, labelFin := wrapString(label)
// 	defer labelFin()

// 	text_state := newInputTextState(*text, nil)
// 	defer func() {
// 		*text = text_state.buf.toGo()
// 		text_state.release()
// 	}()

// 	hints_buf := newStringsBuffer(hints)
// 	// defer func() {
// 	// 	hints_buf.free()
// 	// }()

// 	s := &C.IggComboFilterState{
// 		activeIdx:        C.int(state.ActiveIdx),
// 		selectionChanged: castBool(state.SelectionChanged),
// 	}

// 	result := C.iggComboFilter(labelArg, (*C.char)(text_state.buf.ptr), C.int(text_state.buf.size), (**C.char)(hints_buf.ptr), C.int(len(hints)), s, text_state.key) != 0

// 	state.ActiveIdx = int(s.activeIdx)
// 	state.SelectionChanged = s.selectionChanged != 0

// 	return result
// }

// SelectableInput provides a field for dynamic text input which is enabled with left-mouse double-clicking.
// Returns true if within text input enter was pressed or changes were made and focus was lost (mouse clicked elsewhere or enter/tab/esc key was pressed)
func SelectableInput(label string, text *string) bool {
	if text == nil {
		panic("text can't be nil")
	}

	labelArg, labelFin := wrapString(label)
	defer labelFin()

	text_state := newInputTextState(*text, nil)
	defer func() {
		*text = text_state.buf.toGo()
		text_state.release()
	}()

	return C.iggSelectableInput(labelArg, (*C.char)(text_state.buf.ptr), C.int(text_state.buf.size), text_state.key) != 0
}
