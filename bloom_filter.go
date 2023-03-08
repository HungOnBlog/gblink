package gblink

import (
	"fmt"
	"hash/fnv"
	"math"
)

// BloomFilter is a probabilistic data structure that can be used to test if an item is in a set.
// It is a space-efficient implementation of a set that returns false positives but never false negatives.
// The probability of a false positive can be controlled by the size of the bitset and the number of hash functions.
//
// The number of hash functions should be set to k = m/n * ln(2), where m is the size of the bitset and n is the number of items in the set.
//
// The size of the bitset should be set to m = -n * ln(p) / (ln(2))^2, where n is the number of items in the set and p is the desired probability of a false positive.
//
// More: https://en.wikipedia.org/wiki/Bloom_filter
type BloomFilter struct {
	bitset []bool // the bitset used to store the filter
	k      uint   // the number of hash functions used
}

// NewBloomFilter creates a new Bloom filter with the specified bitset size and number of hash functions.
func NewBloomFilter(m uint, k uint) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, m),
		k:      k,
	}
}

// Add adds an item to the Bloom filter by setting the corresponding bits in the bitset.
func (bf *BloomFilter) Add(item string) {
	for i := uint(0); i < bf.k; i++ {
		hash := bf.hash(item, i)
		bf.bitset[hash] = true
	}
}

// Contains checks if an item is in the Bloom filter by checking if all the corresponding bits in the bitset are set.
func (bf *BloomFilter) Contains(item string) bool {
	for i := uint(0); i < bf.k; i++ {
		hash := bf.hash(item, i)
		if !bf.bitset[hash] {
			return false
		}
	}
	return true
}

// hash computes the hash value for an item using the FNV-1a hash function and the specified seed value.
func (bf *BloomFilter) hash(item string, seed uint) uint {
	hash := fnv.New32a()                             // create a new 32-bit FNV-1a hash object
	hash.Write([]byte(item))                         // write the item to the hash object
	hash.Write([]byte{byte(seed)})                   // write the seed value to the hash object
	return uint(hash.Sum32()) % uint(len(bf.bitset)) // compute the hash value and return it
}

// CalculateBloomFilterBitSetSize calculates the size of the bitset for a Bloom filter with the specified number of items and false positive rate.
func CalculateBloomFilterBitSetSize(numItems uint, falsePositiveRate float64) uint {
	return uint(float64(numItems) * math.Log(falsePositiveRate) / math.Pow(math.Log(2), 2))
}

// CalculateBloomFilterNumHashFunctions calculates the number of hash functions for a Bloom filter with the specified bitset size and number of items.
func CalculateBloomFilterNumHashFunctions(bitSetSize uint, numItems uint) uint {
	return uint(float64(bitSetSize) / float64(numItems) * math.Log(2))
}

// ExampleBloomFilter shows how to use a Bloom filter.
func ExampleBloomFilter() {
	// create Bloom filter
	bf := NewBloomFilter(100, 4)

	// add items
	bf.Add("foo")
	bf.Add("bar")
	bf.Add("baz")

	// check if items are in Bloom filter
	fmt.Println(bf.Contains("foo"))
	fmt.Println(bf.Contains("bar"))
	fmt.Println(bf.Contains("baz"))
	fmt.Println(bf.Contains("qux"))

	// Output:
	// true
	// true
	// true
	// false
}
