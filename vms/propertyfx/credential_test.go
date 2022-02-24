// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package propertyfx

import (
	"testing"

	"github.com/lasthyphen/dijetsnetgo1.2/vms/components/verify"
)

func TestCredentialState(t *testing.T) {
	intf := interface{}(&Credential{})
	if _, ok := intf.(verify.State); ok {
		t.Fatalf("shouldn't be marked as state")
	}
}
