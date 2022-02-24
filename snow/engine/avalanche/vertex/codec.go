// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"github.com/lasthyphen/avalanchego-1.7.3/codec"
	"github.com/lasthyphen/avalanchego-1.7.3/codec/linearcodec"
	"github.com/lasthyphen/avalanchego-1.7.3/codec/reflectcodec"
	"github.com/lasthyphen/avalanchego-1.7.3/utils/units"
)

const (
	// maxSize is the maximum allowed vertex size. It is necessary to deter DoS
	maxSize = units.MiB

	// codecVersion is the only supported codec version
	codecVersion uint16 = 0
)

var c codec.Manager

func init() {
	lc := linearcodec.New(reflectcodec.DefaultTagName, maxSize)
	c = codec.NewManager(maxSize)

	if err := c.RegisterCodec(codecVersion, lc); err != nil {
		panic(err)
	}
}
