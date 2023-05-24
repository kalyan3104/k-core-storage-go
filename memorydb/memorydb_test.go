package memorydb_test

import (
	"testing"

	"github.com/kalyan3104/k-core-storage-go/memorydb"
	"github.com/stretchr/testify/assert"
)

func TestPutNoError(t *testing.T) {
	key, val := []byte("key"), []byte("value")
	mdb := memorydb.New()

	err := mdb.Put(key, val)
	assert.Nil(t, err, "error saving in db")
}

func TestGetPresent(t *testing.T) {
	key, val := []byte("key1"), []byte("value1")
	mdb := memorydb.New()

	err := mdb.Put(key, val)
	assert.Nil(t, err, "error saving in db")

	v, err := mdb.Get(key)
	assert.Nil(t, err, "error not expected but got %s", err)
	assert.Equal(t, val, v, "expected %s but got %s", val, v)
}

func TestGetNotPresent(t *testing.T) {
	key := []byte("key2")
	mdb := memorydb.New()

	v, err := mdb.Get(key)
	assert.NotNil(t, err, "error expected but got nil, value %s", v)
}

func TestHasPresent(t *testing.T) {
	key, val := []byte("key3"), []byte("value3")
	mdb := memorydb.New()

	err := mdb.Put(key, val)
	assert.Nil(t, err, "error saving in db")

	err = mdb.Has(key)
	assert.Nil(t, err, "error not expected but got %s", err)
}

func TestHasNotPresent(t *testing.T) {
	key := []byte("key4")
	mdb := memorydb.New()

	err := mdb.Has(key)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "key not found")
}

func TestDeletePresent(t *testing.T) {
	key, val := []byte("key5"), []byte("value5")
	mdb := memorydb.New()

	err := mdb.Put(key, val)
	assert.Nil(t, err, "error saving in db")

	err = mdb.Remove(key)
	assert.Nil(t, err, "no error expected but got %s", err)

	err = mdb.Has(key)
	assert.NotNil(t, err, "element not expected as already deleted")
	assert.Contains(t, err.Error(), "key not found")
}

func TestDeleteNotPresent(t *testing.T) {
	key := []byte("key6")
	mdb := memorydb.New()

	err := mdb.Remove(key)
	assert.Nil(t, err, "no error expected but got %s", err)
}

func TestClose(t *testing.T) {
	mdb := memorydb.New()

	err := mdb.Close()
	assert.Nil(t, err, "no error expected but got %s", err)
}

func TestDestroy(t *testing.T) {
	mdb := memorydb.New()

	err := mdb.Destroy()
	assert.Nil(t, err, "no error expected but got %s", err)
}

func Test_RangeKeys(t *testing.T) {
	t.Parallel()

	mdb := memorydb.New()

	keysVals := map[string][]byte{
		"key1": []byte("value1"),
		"key2": []byte("value2"),
		"key3": []byte("value3"),
		"key4": []byte("value4"),
		"key5": []byte("value5"),
		"key6": []byte("value6"),
		"key7": []byte("value7"),
	}

	for key, val := range keysVals {
		_ = mdb.Put([]byte(key), val)
	}

	recovered := make(map[string][]byte)

	mdb.RangeKeys(func(key []byte, value []byte) bool {
		recovered[string(key)] = value
		return true
	})

	assert.Equal(t, keysVals, recovered)
}
