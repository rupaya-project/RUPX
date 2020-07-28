// Copyright 2019 The Rupaya Authors
// This file is part of the Core Rupaya infrastructure
// https://rupaya.com
// Package rupexDAO provides an interface to work with rupex database, including leveldb for masternode and mongodb for SDK node
package rupexDAO

import (
	"github.com/rupaya-project/rupx/common"
	"github.com/rupaya-project/rupx/ethdb"
)

const defaultCacheLimit = 1024

type RupeXDAO interface {
	// for both leveldb and mongodb
	IsEmptyKey(key []byte) bool
	Close()

	// mongodb methods
	HasObject(hash common.Hash, val interface{}) (bool, error)
	GetObject(hash common.Hash, val interface{}) (interface{}, error)
	PutObject(hash common.Hash, val interface{}) error
	DeleteObject(hash common.Hash, val interface{}) error // won't return error if key not found
	GetListItemByTxHash(txhash common.Hash, val interface{}) interface{}
	GetListItemByHashes(hashes []string, val interface{}) interface{}
	DeleteItemByTxHash(txhash common.Hash, val interface{})

		// basic rupex
		InitBulk()
		CommitBulk() error

		// rupex lending
		InitLendingBulk()
		CommitLendingBulk() error

	// leveldb methods
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error
	NewBatch() ethdb.Batch
}

// use alloc to prevent reference manipulation
func EmptyKey() []byte {
	key := make([]byte, common.HashLength)
	return key
}