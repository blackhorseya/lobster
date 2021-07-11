package idgen

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
)

// New serve caller to create a snowflake.Node
func New(nodeID int64) (*snowflake.Node, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, err
	}

	return node, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(New)
