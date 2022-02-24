// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"github.com/lasthyphen/avalanchego-1.7.3/codec"
	"github.com/lasthyphen/avalanchego-1.7.3/codec/linearcodec"
	"github.com/lasthyphen/avalanchego-1.7.3/codec/reflectcodec"
	"github.com/lasthyphen/avalanchego-1.7.3/utils/units"
	"github.com/lasthyphen/avalanchego-1.7.3/utils/wrappers"
)

const (
	codecVersion   uint16 = 0
	maxMessageSize        = 512 * units.KiB
	maxSliceLen           = maxMessageSize
)

// Codec does serialization and deserialization
var c codec.Manager

func init() {
	c = codec.NewManager(maxMessageSize)
	lc := linearcodec.New(reflectcodec.DefaultTagName, maxSliceLen)

	errs := wrappers.Errs{}
	errs.Add(
		lc.RegisterType(&Tx{}),
		c.RegisterCodec(codecVersion, lc),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}
