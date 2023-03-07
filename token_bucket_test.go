package gblink

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokenBucket(t *testing.T) {
	// Create a new Token Bucket with a capacity of 5 and a refill rate of 1 token per second.
	tb := NewTokenBucket(5, time.Second)

	// Attempt to take 5 tokens from the bucket, which should all be successful.
	for i := 0; i < 5; i++ {
		assert.True(t, tb.TakeToken())
	}

	// Attempt to take another token from the bucket, which should fail because the bucket is empty.
	assert.False(t, tb.TakeToken())

	// Wait for 1 second to allow the bucket to refill.
	time.Sleep(time.Second)

	// Attempt to take another token from the bucket, which should be successful because the bucket has refilled.
	assert.True(t, tb.TakeToken())
}
