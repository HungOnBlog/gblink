package gblink

import (
	"hash/fnv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCuckooFilter_Add(t *testing.T) {
	assert := assert.New(t)

	cf := NewCuckooFilter(1000, fnv.New64a())

	canAdd := cf.Add("one")
	assert.True(canAdd)

	canAdd = cf.Add("two")
	assert.True(canAdd)

	canAdd = cf.Add("three")
	assert.True(canAdd)

	canAdd = cf.Add("four")
	assert.True(canAdd)
}

func TestCuckooFilter_Contains(t *testing.T) {
	assert := assert.New(t)

	cf := NewCuckooFilter(1000, fnv.New64a())

	cf.Add("one")
	cf.Add("two")
	cf.Add("three")
	cf.Add("four")

	assert.True(cf.Contains("one"))
	assert.True(cf.Contains("two"))
	assert.True(cf.Contains("three"))
	assert.True(cf.Contains("four"))

	assert.False(cf.Contains("five"))
}
