package gblink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeakyBucket(t *testing.T) {
	lb := NewLeakyBucket(1, 10)
	assert := assert.New(t)

	assert.Equal(0.0, lb.waterLevel)
	canAdd := lb.AddWater(1)
	assert.True(canAdd)
	assert.GreaterOrEqual(1.0, lb.waterLevel)

	canAdd = lb.AddWater(1)
	assert.True(canAdd)
	assert.GreaterOrEqual(2.0, lb.waterLevel)

	canAdd = lb.AddWater(100)
	assert.False(canAdd)
	assert.GreaterOrEqual(2.0, lb.waterLevel)
}
