package disabled

import (
	"github.com/kalyan3104/k-core-storage-go/common"
)

type persister struct{}

// NewPersister returns a new instance of this disabled persister
func NewPersister() *persister {
	return &persister{}
}

// Put returns nil
func (p *persister) Put(_, _ []byte) error {
	return nil
}

// Get returns nil and ErrKeyNotFound
func (p *persister) Get(_ []byte) ([]byte, error) {
	return nil, common.ErrKeyNotFound
}

// Has returns ErrKeyNotFound
func (p *persister) Has(_ []byte) error {
	return common.ErrKeyNotFound
}

// Close returns nil
func (p *persister) Close() error {
	return nil
}

// Remove returns nil
func (p *persister) Remove(_ []byte) error {
	return nil
}

// Destroy returns nil
func (p *persister) Destroy() error {
	return nil
}

// DestroyClosed returns nil
func (p *persister) DestroyClosed() error {
	return nil
}

// RangeKeys does nothing
func (p *persister) RangeKeys(_ func(key []byte, val []byte) bool) {}

// IsInterfaceNil returns true if there is no value under the interface
func (p *persister) IsInterfaceNil() bool {
	return p == nil
}
