package gblink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloomFilter_CalculateBitSet(t *testing.T) {
	assert := assert.New(t)

	// test cases
	testCases := []struct {
		numItems           uint
		falsePositiveRate  float64
		expectedBitSetSize uint
	}{
		{100, 0.01, 18446744073709550658},
		{100, 0.001, 18446744073709550179},
		{100, 0.0001, 18446744073709549699},
		{1000000000, 0.01, 18446744064124493239},
	}

	// run test cases
	for _, tc := range testCases {
		bitSetSize := CalculateBloomFilterBitSetSize(tc.numItems, tc.falsePositiveRate)
		assert.Equal(tc.expectedBitSetSize, bitSetSize)
	}
}

func TestBloomFilter_CalculateNumberOfHashFunctions(t *testing.T) {
	assert := assert.New(t)

	// test cases
	testCases := []struct {
		bitSetSize uint
		numItems   uint
		expectedK  uint
	}{
		{645, 100, 4},
		{1290, 100, 8},
		{2580, 100, 17},
		{5160, 100, 35},
		{10320, 100, 71},
		{20640, 100, 143},
		{41280, 100, 286},
		{82560, 100, 572},
		{18446744064124493239, 1000000000, 12786308638},
	}

	// run test cases
	for _, tc := range testCases {
		k := CalculateBloomFilterNumHashFunctions(tc.bitSetSize, tc.numItems)
		assert.Equal(tc.expectedK, k)
	}
}

func TestBloomFilter_Add(t *testing.T) {
	assert := assert.New(t)

	// create Bloom filter
	bf := NewBloomFilter(100, 4)

	// add items
	bf.Add("foo")
	bf.Add("bar")
	bf.Add("baz")

	// check if items are in Bloom filter
	assert.True(bf.Contains("foo"))
	assert.True(bf.Contains("bar"))
	assert.True(bf.Contains("baz"))
}

func TestBloomFilter_Contains(t *testing.T) {
	assert := assert.New(t)

	// create Bloom filter
	bf := NewBloomFilter(100, 4)

	// add items
	bf.Add("foo")
	bf.Add("bar")
	bf.Add("baz")

	// check if items are in Bloom filter
	assert.True(bf.Contains("foo"))
	assert.True(bf.Contains("bar"))
	assert.True(bf.Contains("baz"))
}
