// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/lasthyphen/avalanchego-1.7.3/database/memdb"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/consensus/snowball"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/consensus/snowman"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/engine/common"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/engine/common/queue"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/engine/snowman/block"
	"github.com/lasthyphen/avalanchego-1.7.3/snow/engine/snowman/bootstrap"
)

func DefaultConfig() Config {
	blocked, _ := queue.NewWithMissing(memdb.New(), "", prometheus.NewRegistry())
	return Config{
		Config: bootstrap.Config{
			Config:  common.DefaultConfigTest(),
			Blocked: blocked,
			VM:      &block.TestVM{},
		},
		Params: snowball.Parameters{
			K:                     1,
			Alpha:                 1,
			BetaVirtuous:          1,
			BetaRogue:             2,
			ConcurrentRepolls:     1,
			OptimalProcessing:     100,
			MaxOutstandingItems:   1,
			MaxItemProcessingTime: 1,
		},
		Consensus: &snowman.Topological{},
	}
}
