package gblink

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens        uint64        // Current number of tokens in the bucket.
	capacity      uint64        // Maximum number of tokens that the bucket can hold.
	rate          time.Duration // Rate at which tokens are added to the bucket.
	mu            sync.Mutex    // Mutex to synchronize access to the bucket.
	lastTokenTime time.Time     // Last time a token was added to the bucket.
}

// NewTokenBucket creates a new Token Bucket with the specified capacity and refill rate.
func NewTokenBucket(capacity uint64, rate time.Duration) *TokenBucket {
	return &TokenBucket{
		tokens:        capacity,
		capacity:      capacity,
		rate:          rate,
		lastTokenTime: time.Now(),
	}
}

// TakeToken attempts to take a token from the bucket.
func (tb *TokenBucket) TakeToken() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	// Calculate the number of tokens that should have been added since the last token was added.
	elapsedTime := time.Since(tb.lastTokenTime)
	numTokensToAdd := uint64(elapsedTime.Nanoseconds() / tb.rate.Nanoseconds())

	// Add the calculated tokens to the bucket, up to the capacity of the bucket.
	tb.tokens += numTokensToAdd
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	// Attempt to take a token from the bucket.
	if tb.tokens > 0 {
		tb.tokens--
		tb.lastTokenTime = time.Now()
		return true
	}
	return false
}

// Example of a token bucket.
// Limit the rate of incoming requests to 100 requests per second.
func ExampleTokenBucket() {
	// Create a new token bucket with a capacity of 100 tokens and a fill rate of 100 tokens per second.
	tb := NewTokenBucket(100, 100)

	// Simulate 1000 incoming requests over the course of 11 seconds.
	start := time.Now()
	for i := 0; i < 1000; i++ {
		// Wait for the token bucket to allow the request to proceed.
		for !tb.TakeToken() {
			time.Sleep(time.Millisecond * 10)
		}

		// Process the request.
		processRequest()

		// Throttle the request rate to exactly 100 requests per second.
		if i%10 == 9 {
			time.Sleep(time.Second / 100)
		}
	}
	elapsed := time.Since(start)

	// Print some statistics about the simulation.
	fmt.Printf("Processed %d requests in %s (%.2f requests per second)\n", 1000, elapsed, float64(1000)/elapsed.Seconds())
}

func processRequest() {
	// Simulate some work.
	time.Sleep(time.Millisecond * 50)
}
