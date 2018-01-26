package state

import (
	"github.com/vechain/thor/kv"
	"github.com/vechain/thor/thor"
)

// Creator state creator to cut-off kv dependency.
type Creator struct {
	kv kv.GetPutter
}

// NewCreator create a new state creator.
func NewCreator(kv kv.GetPutter) *Creator {
	return &Creator{kv}
}

// NewState create a new state object.
func (c *Creator) NewState(root thor.Hash) (*State, error) {
	return New(root, c.kv)
}