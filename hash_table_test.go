package gblink

import (
	"hash/fnv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTable_Set(t *testing.T) {
	assert := assert.New(t)

	table := NewHashTable[int, string](fnv.New64())
	table.Set(1, "one")
	table.Set(2, "two")
	table.Set(3, "three")

	assert.Equal(3, table.Len())
}

func TestHashTable_Get(t *testing.T) {
	assert := assert.New(t)

	table := NewHashTable[string, int](fnv.New64a())

	table.Set("one", 1)
	table.Set("two", 2)
	table.Set("three", 3)

	v, err := table.Get("two")
	assert.Nil(err)
	assert.Equal(2, v)

	v, err = table.Get("three")
	assert.Nil(err)
	assert.Equal(3, v)

	v, err = table.Get("four")
	assert.NotNil(err)
	assert.Equal(0, v)
}

func TestHashTable_GetError(t *testing.T) {
	assert := assert.New(t)

	table := NewHashTable[int, string](fnv.New64())
	table.Set(1, "one")
	table.Set(2, "two")
	table.Set(3, "three")

	_, err := table.Get(4)
	assert.NotNil(err)
}

func TestHashTable_Len(t *testing.T) {
	assert := assert.New(t)

	table := NewHashTable[int, string](fnv.New64())
	table.Set(1, "one")
	table.Set(2, "two")
	table.Set(3, "three")
	table.Set(4, "four")
	table.Set(5, "five")

	assert.Equal(5, table.Len())
}

func TestHashTable_Delete(t *testing.T) {
	assert := assert.New(t)

	table := NewHashTable[int, string](fnv.New64())
	table.Set(1, "one")
	table.Set(2, "two")
	table.Set(3, "three")
	table.Set(4, "four")
	table.Set(5, "five")

	table.Delete(1)
	table.Delete(2)
	table.Delete(3)

	assert.Equal(2, table.Len())
}

func TestHashTable_Clear(t *testing.T) {
	assert := assert.New(t)

	table := NewHashTable[int, string](fnv.New64())
	table.Set(1, "one")
	table.Set(2, "two")
	table.Set(3, "three")
	table.Set(4, "four")
	table.Set(5, "five")

	table.Clear()

	assert.Equal(0, table.Len())
}
