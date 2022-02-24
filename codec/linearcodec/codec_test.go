// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package linearcodec

import (
	"testing"

	"github.com/lasthyphen/avalanchego-1.7.3/codec"
)

func TestVectors(t *testing.T) {
	for _, test := range codec.Tests {
		c := NewDefault()
		test(c, t)
	}
}
