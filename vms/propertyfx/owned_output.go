// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package propertyfx

import (
	"github.com/lasthyphen/avalanchego-1.7.3/vms/secp256k1fx"
)

type OwnedOutput struct {
	secp256k1fx.OutputOwners `serialize:"true"`
}
