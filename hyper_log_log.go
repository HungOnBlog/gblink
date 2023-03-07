package gblink

import (
	"errors"
	"fmt"
	"math"

	"github.com/spaolacci/murmur3"
)

// HyperLogLog is a probabilistic data structure that can be used to estimate the number of distinct elements in a data stream.
//
// The HyperLogLog algorithm was invented by Philippe Flajolet, Éric Fusy, Olivier Gandouet and Frédéric Meunier in 2007.
//
// HyperLogLog is a linear time algorithm that uses a fixed amount of memory, making it suitable for use in a distributed system.
// The HyperLogLog algorithm is also very accurate, with an error rate of less than 1.04/sqrt(m) where m is the number of registers.
//
// For more information, see the following resources: http://algo.inria.fr/flajolet/Publications/FlFuGaMe07.pdf
// Hasher is an interface for a hash function that takes a byte slice and returns a 64-bit integer.
type Hasher interface {
	Sum64([]byte) uint64
}

// defaultHasher is a simple implementation of the Hasher interface that uses the Murmur3 hash function.
type defaultHasher struct {
}

func (h defaultHasher) Sum64(data []byte) uint64 {
	return murmur3.Sum64(data)
}

// HyperLogLog is a probabilistic data structure that approximates the cardinality of a set with high accuracy and low memory usage.
type HyperLogLog struct {
	m         uint32
	alphaM    float64
	registers []uint8
	hasher    Hasher
}

// NewHyperLogLog returns a new HyperLogLog with the specified number of registers.
func NewHyperLogLog(m uint32, hasher Hasher) (*HyperLogLog, error) {
	if m < 4 || m > 16 {
		return nil, errors.New("m must be between 4 and 16")
	}

	return &HyperLogLog{
		m:         m,
		alphaM:    getAlpha(m),
		registers: make([]uint8, 1<<m),
		hasher:    hasher,
	}, nil
}

// Add adds the specified item to the HyperLogLog.
func (h *HyperLogLog) Add(item []byte) {
	hashVal := h.hasher.Sum64(item)

	// Determine the register index
	index := hashVal & ((1 << h.m) - 1)

	// Determine the rank of the first 1 bit after the m least significant bits
	rank := getRank(hashVal>>h.m, 64-int(h.m))

	// Update the register if the rank is greater than the current value
	if rank > int(h.registers[index]) {
		h.registers[index] = uint8(rank)
	}
}

// Count returns an estimate of the number of distinct items that have been added to the HyperLogLog.
func (h *HyperLogLog) Count() uint64 {
	var sum float64 = 0

	for _, val := range h.registers {
		sum += math.Pow(2, -float64(val))
	}

	estimate := h.alphaM * math.Pow(float64(1)/sum, 2)

	if estimate <= float64(2.5)*float64(len(h.registers)) {
		var zeros uint64
		for _, val := range h.registers {
			if val == 0 {
				zeros++
			}
		}
		if zeros != 0 {
			estimate = float64(len(h.registers)) * math.Log(float64(len(h.registers))/float64(zeros))
		}
	} else if estimate > float64(1<<32)/float64(30) {
		estimate = -math.Pow(2, 64) * math.Log(1-estimate/math.Pow(2, 64))
	}

	return uint64(estimate)
}

// getAlpha returns the alpha constant for the specified number of registers.
func getAlpha(m uint32) float64 {
	switch m {
	case 4:
		return 0.673
	case 5:
		return 0.697
	case 6:
		return 0.709
	default:
		return 0.7213 / (1 + 1.079/float64(uint(1<<m)))
	}
}

func getRank(hashVal uint64, p int) int {
	var rank int = 1
	for (hashVal&1) == 0 && rank <= p {
		rank++
		hashVal >>= 1
	}
	return rank
}

type DefaultHasher struct{}

func (h DefaultHasher) Sum64(data []byte) uint64 {
	return murmur3.Sum64(data)
}

// ExampleHyperLogLog demonstrates how to use the HyperLogLog data structure.
func ExampleHyperLogLog() {

	h, _ := NewHyperLogLog(4, DefaultHasher{})

	// Add some elements to the data stream (100 els) with only 3 distinct values
	for i := 0; i < 100; i++ {
		h.Add([]byte("foo"))
		h.Add([]byte("bar"))
		h.Add([]byte("baz"))
	}

	fmt.Println(h.Count())
	// 3
}
