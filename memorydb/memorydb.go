package memorydb

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sync"

	"github.com/kalyan3104/k-core-storage-go/types"
)

var _ types.Persister = (*DB)(nil)

// DB represents the memory database storage. It holds a map of key value pairs
// and a mutex to handle concurrent accesses to the map
type DB struct {
	db   map[string][]byte
	mutx sync.RWMutex
}

// New creates a new memorydb object
func New() *DB {
	return &DB{
		db:   make(map[string][]byte),
		mutx: sync.RWMutex{},
	}
}

// Put adds the value to the (key, val) storage medium
func (s *DB) Put(key, val []byte) error {
	s.mutx.Lock()
	defer s.mutx.Unlock()

	s.db[string(key)] = val

	return nil
}

// Get gets the value associated to the key, or reports an error
func (s *DB) Get(key []byte) ([]byte, error) {
	s.mutx.RLock()
	defer s.mutx.RUnlock()

	val, ok := s.db[string(key)]

	if !ok {
		return nil, fmt.Errorf("key: %s not found", base64.StdEncoding.EncodeToString(key))
	}

	return val, nil
}

// Has returns true if the given key is present in the persistence medium, false otherwise
func (s *DB) Has(key []byte) error {
	s.mutx.RLock()
	defer s.mutx.RUnlock()

	_, ok := s.db[string(key)]

	if !ok {
		return errors.New("key not found")
	}
	return nil
}

// Close closes the files/resources associated to the storage medium
func (s *DB) Close() error {
	// nothing to do
	return nil
}

// Remove removes the data associated to the given key
func (s *DB) Remove(key []byte) error {
	s.mutx.Lock()
	defer s.mutx.Unlock()

	delete(s.db, string(key))

	return nil
}

// Destroy removes the storage medium stored data
func (s *DB) Destroy() error {
	s.mutx.Lock()
	defer s.mutx.Unlock()

	s.db = make(map[string][]byte)

	return nil
}

// RangeKeys will iterate over all contained (key, value) pairs calling the provided handler
func (s *DB) RangeKeys(handler func(key []byte, value []byte) bool) {
	if handler == nil {
		return
	}

	s.mutx.RLock()
	defer s.mutx.RUnlock()

	for k, v := range s.db {
		shouldContinue := handler([]byte(k), v)
		if !shouldContinue {
			return
		}
	}
}

// DestroyClosed removes the storage medium stored data
func (s *DB) DestroyClosed() error {
	return s.Destroy()
}

// IsInterfaceNil returns true if there is no value under the interface
func (s *DB) IsInterfaceNil() bool {
	return s == nil
}
