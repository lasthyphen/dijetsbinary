// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package codec

import "github.com/lasthyphen/avalanchego-1.7.3/utils/wrappers"

// Codec marshals and unmarshals
type Codec interface {
	MarshalInto(interface{}, *wrappers.Packer) error
	Unmarshal([]byte, interface{}) error
}
