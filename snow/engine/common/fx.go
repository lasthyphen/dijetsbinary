// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"github.com/lasthyphen/avalanchego-1.7.3/ids"
)

// Fx wraps an instance of a feature extension
type Fx struct {
	ID ids.ID
	Fx interface{}
}
