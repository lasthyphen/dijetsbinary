// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

import (
	"github.com/lasthyphen/avalanchego-1.7.3/snow/consensus/avalanche"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/engine/avalanche/bootstrap"
)

// Config wraps all the parameters needed for an avalanche engine
type Config struct {
	bootstrap.Config

	Params    avalanche.Parameters
	Consensus avalanche.Consensus
}
