package imgui 

// Condition for SetWindow***(), SetNextWindow***(), SetNextTreeNode***() functions.
// Important: Treat as a regular enum! Do NOT combine multiple values using binary operators!
// All the functions above treat 0 as a shortcut to ConditionAlways.
type Condition int
