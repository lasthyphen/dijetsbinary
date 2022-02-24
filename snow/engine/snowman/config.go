// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/lasthyphen/avalanchego-1.7.3/snow/consensus/snowball"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/consensus/snowman"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/engine/snowman/bootstrap"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	bootstrap.Config

	Params    snowball.Parameters
	Consensus snowman.Consensus
}
