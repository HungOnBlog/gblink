package gblink

import "hash"

// CuckooFilter is a probabilistic data structure that can be used to test if an item is in a set.
// It is a space-efficient implementation of a set that returns false positives but never false negatives.
type Bucket struct {
	Fingerprint uint32
}

type CuckooFilter struct {
	Size      uint32
	HashFn    hash.Hash64
	MaxKicks  uint32
	BucketArr []*Bucket
}

const (
	MaxNumKicks = 500 // Maximum number of kicks before we give up on inserting an item
	FpSize      = 32  // Size of the fingerprint in bits
)

// NewCuckooFilter creates a new Cuckoo filter with the specified size and number of hash functions.
func NewCuckooFilter(size uint32, hashFn hash.Hash64) *CuckooFilter {
	return &CuckooFilter{
		Size:      size,
		HashFn:    hashFn,
		MaxKicks:  MaxNumKicks,
		BucketArr: make([]*Bucket, size),
	}
}

// Add adds an item to the Cuckoo filter by setting the corresponding bits in the bitset.
func (cf *CuckooFilter) Add(item string) bool {
	return cf.insert(item, true)
}

// Insert inserts an item into the Cuckoo filter by setting the corresponding bits in the bitset.
func (cf *CuckooFilter) insert(item string, isInsert bool) bool {
	// Compute the hash values for the item
	hash1 := cf.hash(item, 0)
	hash2 := cf.hash(item, hash1)

	// Check if the item is already in the filter
	if cf.contains(item, hash1, hash2) {
		return true
	}

	// Insert the item into the filter
	if isInsert {
		return cf.insertItem(item, hash1, hash2)
	}

	return false
}

// contains checks if an item is in the Cuckoo filter.
func (cf *CuckooFilter) contains(item string, hash1 uint32, hash2 uint32) bool {
	// Compute the fingerprint for the item
	fingerprint := cf.fingerprint(item)

	// Check if the item is in the filter
	if cf.BucketArr[hash1] != nil && cf.BucketArr[hash1].Fingerprint == fingerprint {
		return true
	}
	if cf.BucketArr[hash2] != nil && cf.BucketArr[hash2].Fingerprint == fingerprint {
		return true
	}

	return false
}

// hash computes the hash value for an item using the FNV-1a hash function and the specified seed value.
func (cf *CuckooFilter) hash(item string, seed uint32) uint32 {
	cf.HashFn.Reset() // reset the hash object
	cf.HashFn.Write([]byte(item))
	cf.HashFn.Write([]byte{byte(seed)})
	return uint32(cf.HashFn.Sum64()) % cf.Size
}

// insertItem inserts an item into the Cuckoo filter by setting the corresponding bits in the bitset.
func (cf *CuckooFilter) insertItem(item string, hash1 uint32, hash2 uint32) bool {
	// Compute the fingerprint for the item
	fingerprint := cf.fingerprint(item)

	// Insert the item into the filter
	for i := uint32(0); i < cf.MaxKicks; i++ {
		// Insert the item into the filter
		if cf.insertItemIntoBucket(fingerprint, hash1) {
			return true
		}
		if cf.insertItemIntoBucket(fingerprint, hash2) {
			return true
		}

		// Swap the fingerprint with a random bucket's fingerprint
		hash := cf.hash(item, hash1)
		fingerprint, cf.BucketArr[hash].Fingerprint = cf.BucketArr[hash].Fingerprint, fingerprint
	}

	return false
}

// fingerprint computes the fingerprint for an item.
func (cf *CuckooFilter) fingerprint(item string) uint32 {
	return cf.hash(item, 0) & ((1 << FpSize) - 1)
}

// insertItemIntoBucket inserts an item into the specified bucket.
func (cf *CuckooFilter) insertItemIntoBucket(fingerprint uint32, hash uint32) bool {
	// Check if the bucket is empty
	if cf.BucketArr[hash] == nil {
		cf.BucketArr[hash] = &Bucket{Fingerprint: fingerprint}
		return true
	}

	// Check if the bucket already contains the item
	if cf.BucketArr[hash].Fingerprint == fingerprint {
		return true
	}

	return false
}

// Contains checks if an item is in the Cuckoo filter.
func (cf *CuckooFilter) Contains(item string) bool {
	// Compute the hash values for the item
	hash1 := cf.hash(item, 0)
	hash2 := cf.hash(item, hash1)

	// Check if the item is in the filter
	return cf.contains(item, hash1, hash2)
}

// Delete deletes an item from the Cuckoo filter by clearing the corresponding bits in the bitset.
func (cf *CuckooFilter) Delete(item string) bool {
	// Compute the hash values for the item
	hash1 := cf.hash(item, 0)
	hash2 := cf.hash(item, hash1)

	// Compute the fingerprint for the item
	fingerprint := cf.fingerprint(item)

	// Delete the item from the filter
	if cf.BucketArr[hash1] != nil && cf.BucketArr[hash1].Fingerprint == fingerprint {
		cf.BucketArr[hash1] = nil
		return true
	}
	if cf.BucketArr[hash2] != nil && cf.BucketArr[hash2].Fingerprint == fingerprint {
		cf.BucketArr[hash2] = nil
		return true
	}

	return false
}
