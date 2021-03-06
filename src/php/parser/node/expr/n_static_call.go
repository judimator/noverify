package expr

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/php/parser/walker"
)

// StaticCall node
type StaticCall struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Class        node.Node
	Call         node.Node
	ArgumentList *node.ArgumentList
}

// NewStaticCall node constructor
func NewStaticCall(Class node.Node, Call node.Node, ArgumentList *node.ArgumentList) *StaticCall {
	return &StaticCall{
		FreeFloating: nil,
		Class:        Class,
		Call:         Call,
		ArgumentList: ArgumentList,
	}
}

// SetPosition sets node position
func (n *StaticCall) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *StaticCall) GetPosition() *position.Position {
	return n.Position
}

func (n *StaticCall) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *StaticCall) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Class != nil {
		n.Class.Walk(v)
	}

	if n.Call != nil {
		n.Call.Walk(v)
	}

	if n.ArgumentList != nil {
		n.ArgumentList.Walk(v)
	}

	v.LeaveNode(n)
}
