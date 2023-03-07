package gblink

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHyperLogLog(t *testing.T) {
	hll, _ := NewHyperLogLog(4, &DefaultHasher{})
	assert := assert.New(t)

	// Add 1000 items to the HLL. with only 6 distinct items, the HLL should estimate the number of distinct items to be 6.
	for i := 0; i < 1000; i++ {
		hll.Add([]byte{'a'})
		hll.Add([]byte{'b'})
		hll.Add([]byte{'c'})
		hll.Add([]byte{'d'})
		hll.Add([]byte{'e'})
		hll.Add([]byte{'f'})
		hll.Add([]byte{'g'})
		hll.Add([]byte{'h'})
		hll.Add([]byte{'i'})
		hll.Add([]byte{'j'})
		hll.Add([]byte{'k'})
		hll.Add([]byte{'l'})
		hll.Add([]byte{'m'})
	}

	// The HLL should estimate the number of distinct items to be 6.
	count := hll.Count()
	fmt.Printf("count: %d\n", count)
	assert.InDelta(13, count, 10)
}
