// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package trie

import (
	"github.com/medibloc/go-medibloc/storage"
)

// Action represents operation types in BatchTrie
type Action int

// Action constants
const (
	Insert Action = iota
	Update
	Delete
)

// Entry entry for not committed changes
type Entry struct {
	action Action
	key    []byte
	old    []byte
	update []byte
}

// Batch batch implementation for trie
type Batch struct {
	batching    bool
	changeLogs  []*Entry
	oldRootHash []byte
	trie        *Trie
	storage     storage.Storage
}

// NewBatch return new Batch instance
func NewBatch(rootHash []byte, stor storage.Storage) (*Batch, error) {
	t, err := NewTrie(rootHash, stor)
	if err != nil {
		return nil, err
	}
	return &Batch{
		trie:    t,
		storage: stor,
	}, nil
}

// BeginBatch begin batch
func (t *Batch) BeginBatch() error {
	if t.batching {
		return ErrBeginAgainInBatch
	}
	t.oldRootHash = t.trie.RootHash()
	t.batching = true
	return nil
}

// Clone clone Batch
func (t *Batch) Clone() (*Batch, error) {
	if t.batching {
		return nil, ErrCannotCloneOnBatching
	}
	return NewBatch(t.trie.RootHash(), t.storage)
}

// Commit commit batch WARNING: not thread-safe
func (t *Batch) Commit() error {
	if !t.batching {
		return ErrNotBatching
	}
	t.changeLogs = t.changeLogs[:0]
	t.batching = false
	return nil
}

// Delete delete in trie
func (t *Batch) Delete(key []byte) error {
	if !t.batching {
		return ErrNotBatching
	}
	entry := &Entry{action: Delete, key: key, old: nil}
	old, getErr := t.trie.Get(key)
	if getErr == nil {
		entry.old = old
	}
	t.changeLogs = append(t.changeLogs, entry)
	return t.trie.Delete(key)
}

// Get get from trie
func (t *Batch) Get(key []byte) ([]byte, error) {
	return t.trie.Get(key)
}

// Iterator iterates trie.
func (t *Batch) Iterator(prefix []byte) (*Iterator, error) {
	return t.trie.Iterator(prefix)
}

// Put put to trie
func (t *Batch) Put(key []byte, value []byte) error {
	if !t.batching {
		return ErrNotBatching
	}
	entry := &Entry{action: Insert, key: key, old: nil, update: value}
	old, getErr := t.trie.Get(key)
	if getErr == nil {
		entry.action = Update
		entry.old = old
	}
	t.changeLogs = append(t.changeLogs, entry)
	return t.trie.Put(key, value)
}

// RollBack rollback batch WARNING: not thread-safe
func (t *Batch) RollBack() error {
	if !t.batching {
		return ErrNotBatching
	}
	// TODO rollback with changelogs
	t.changeLogs = t.changeLogs[:0]
	t.trie.SetRootHash(t.oldRootHash)
	t.batching = false
	return nil
}

// RootHash getter for rootHash
func (t *Batch) RootHash() []byte {
	return t.trie.RootHash()
}
