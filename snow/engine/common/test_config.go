// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"github.com/lasthyphen/dijetsnetgo1.2/ids"
	"github.com/lasthyphen/dijetsnetgo1.2/snow"
	"github.com/lasthyphen/dijetsnetgo1.2/snow/validators"
)

// DefaultConfigTest returns a test configuration
func DefaultConfigTest() Config {
	isBootstrapped := false
	subnet := &SubnetTest{
		IsBootstrappedF: func() bool { return isBootstrapped },
		BootstrappedF:   func(ids.ID) { isBootstrapped = true },
	}

	return Config{
		Ctx:                           snow.DefaultConsensusContextTest(),
		Validators:                    validators.NewSet(),
		Beacons:                       validators.NewSet(),
		Sender:                        &SenderTest{},
		Bootstrapable:                 &BootstrapableTest{},
		Subnet:                        subnet,
		Timer:                         &TimerTest{},
		MultiputMaxContainersSent:     2000,
		MultiputMaxContainersReceived: 2000,
	}
}
