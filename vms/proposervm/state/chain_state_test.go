// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lasthyphen/avalanchego-1.7.3/database"
	"github.com/lasthyphen/avalanchego-1.7.3/database/memdb"
	"github.com/lasthyphen/avalanchego-1.7.3/ids"
)

func testChainState(a *assert.Assertions, cs ChainState) {
	lastAccepted := ids.GenerateTestID()

	_, err := cs.GetLastAccepted()
	a.Equal(database.ErrNotFound, err)

	err = cs.SetLastAccepted(lastAccepted)
	a.NoError(err)

	err = cs.SetLastAccepted(lastAccepted)
	a.NoError(err)

	cs.clearCache()

	fetchedLastAccepted, err := cs.GetLastAccepted()
	a.NoError(err)
	a.Equal(lastAccepted, fetchedLastAccepted)

	fetchedLastAccepted, err = cs.GetLastAccepted()
	a.NoError(err)
	a.Equal(lastAccepted, fetchedLastAccepted)

	err = cs.DeleteLastAccepted()
	a.NoError(err)

	_, err = cs.GetLastAccepted()
	a.Equal(database.ErrNotFound, err)
}

func TestChainState(t *testing.T) {
	a := assert.New(t)

	db := memdb.New()
	cs := NewChainState(db)

	testChainState(a, cs)
}
