// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/lasthyphen/dijetsnetgo1.2/snow/consensus/snowball"
	"github.com/lasthyphen/dijetsnetgo1.2/snow/consensus/snowman"
	"github.com/lasthyphen/dijetsnetgo1.2/snow/engine/snowman/bootstrap"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	bootstrap.Config

	Params    snowball.Parameters
	Consensus snowman.Consensus
}
