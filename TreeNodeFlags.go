package imgui

const (
	// TreeNodeFlagsSelected draws as selected.
	TreeNodeFlagsSelected             = 1 << iota
	// TreeNodeFlagsFramed draws full colored frame (e.g. for CollapsingHeader).
	TreeNodeFlagsFramed
	// TreeNodeFlagsAllowItemOverlap hit testing to allow subsequent widgets to overlap this one.
	TreeNodeFlagsAllowItemOverlap
	// TreeNodeFlagsNoTreePushOnOpen doesn't do a TreePush() when open
	// (e.g. for CollapsingHeader) = no extra indent nor pushing on ID stack.
	TreeNodeFlagsNoTreePushOnOpen
	// TreeNodeFlagsNoAutoOpenOnLog doesn't automatically and temporarily open node when Logging is active
	// (by default logging will automatically open tree nodes).
	TreeNodeFlagsNoAutoOpenOnLog
	// TreeNodeFlagsDefaultOpen defaults node to be open.
	TreeNodeFlagsDefaultOpen
	// TreeNodeFlagsOpenOnDoubleClick needs double-click to open node.
	TreeNodeFlagsOpenOnDoubleClick
	// TreeNodeFlagsOpenOnArrow opens only when clicking on the arrow part.
	// If TreeNodeFlagsOpenOnDoubleClick is also set, single-click arrow or double-click all box to open.
	TreeNodeFlagsOpenOnArrow
	// TreeNodeFlagsLeaf allows no collapsing, no arrow (use as a convenience for leaf nodes).
	TreeNodeFlagsLeaf
	// TreeNodeFlagsBullet displays a bullet instead of arrow.
	TreeNodeFlagsBullet
	// TreeNodeFlagsFramePadding use FramePadding (even for an unframed text node) to
	// vertically align text baseline to regular widget height. Equivalent to calling AlignTextToFramePadding().
	TreeNodeFlagsFramePadding
	_
	_
	// TreeNodeFlagsNavLeftJumpsBackHere (WIP) Nav: left direction may move to this TreeNode()
	// from any of its child (items submitted between TreeNode and TreePop).
	TreeNodeFlagsNavLeftJumpsBackHere
	// TreeNodeFlagsCollapsingHeader combines TreeNodeFlagsFramed and TreeNodeFlagsNoAutoOpenOnLog.
	TreeNodeFlagsCollapsingHeader     = TreeNodeFlagsFramed | TreeNodeFlagsNoAutoOpenOnLog
)