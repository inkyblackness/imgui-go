package imgui

// Condition for SetWindow***(), SetNextWindow***(), SetNextTreeNode***() functions.
// Important: Treat as a regular enum! Do NOT combine multiple values using binary operators!
// All the functions above treat 0 as a shortcut to ConditionAlways.
type Condition int

const (
	// ConditionNone no condition (always set the variable), same as ConditionAlways
	ConditionNone Condition = 0
	// ConditionAlways no condition (always set the variable)
	ConditionAlways = 1 << 0
	// ConditionOnce sets the variable once per runtime session (only the first call will succeed)
	ConditionOnce = 1 << 1
	// ConditionFirstUseEver sets the variable if the object/window has no persistently saved data (no entry in .ini file).
	ConditionFirstUseEver = 1 << 2
	// ConditionAppearing sets the variable if the object/window is appearing after being hidden/inactive (or the first time).
	ConditionAppearing = 1 << 3
)
