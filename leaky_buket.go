package gblink

import (
	"fmt"
	"time"
)

// LeakyBucket simulates a bucket with a hole that leaks water at a fixed rate.
type LeakyBucket struct {
	flowRate       float64       // The rate at which water flows into the bucket.
	bucketCapacity float64       // The maximum amount of water that the bucket can hold.
	waterLevel     float64       // The current amount of water in the bucket.
	lastLeak       time.Time     // The time when the bucket was last leaked.
	flowTicker     *time.Ticker  // The ticker that controls the flow of water into the bucket.
	stopChan       chan struct{} // The channel used to stop the flow of water into the bucket.
}

// NewLeakyBucket creates a new leaky bucket with the specified flow rate and bucket capacity.
func NewLeakyBucket(flowRate float64, bucketCapacity float64) *LeakyBucket {
	return &LeakyBucket{
		flowRate:       flowRate,
		bucketCapacity: bucketCapacity,
		waterLevel:     0,
		lastLeak:       time.Now(),
		flowTicker:     time.NewTicker(time.Second), // By default, the bucket leaks water once per second.
		stopChan:       make(chan struct{}),
	}
}

// AddWater adds a specified volume of water to the bucket.
func (lb *LeakyBucket) AddWater(volume float64) bool {
	// Calculate the time since the bucket was last leaked.
	elapsed := time.Since(lb.lastLeak)

	// Calculate the amount of water that should have leaked from the bucket during this time.
	leaked := elapsed.Seconds() * lb.flowRate

	// Update the current water level by subtracting the leaked water.
	lb.waterLevel -= leaked

	// Ensure that the water level does not exceed the bucket capacity.
	if lb.waterLevel+volume > lb.bucketCapacity {
		return false // The bucket is full.
	}

	// Add the new water volume to the water level.
	lb.waterLevel += volume

	// Update the last leak time.
	lb.lastLeak = time.Now()

	return true // The water has been added to the bucket.
}

// Start starts the flow of water into the bucket.
func (lb *LeakyBucket) Start() {
	go func() {
		for {
			select {
			case <-lb.stopChan:
				lb.flowTicker.Stop() // Stop the ticker when the flow is stopped.
				return
			case <-lb.flowTicker.C:
				// Add the flow rate amount of water to the bucket.
				lb.AddWater(lb.flowRate)
			}
		}
	}()
}

// Stop stops the flow of water into the bucket.
func (lb *LeakyBucket) Stop() {
	lb.stopChan <- struct{}{}
}

// Example of usage:
// Using the leaky bucket to limit the number of requests per second.
// We want to limit the number of requests to 100 requests per second.
func ExampleLeakyBucket() {
	// Create a new leaky bucket with a flow rate of 100 requests per second and a capacity of 100 requests.
	bucket := NewLeakyBucket(100, 100)

	// Start the flow of requests into the bucket.
	bucket.Start()

	// Simulate incoming requests.
	for i := 1; i <= 200; i++ {
		if bucket.AddWater(1) {
			fmt.Printf("Request %d allowed at %s\n", i, time.Now().Format(time.StampMilli))
		} else {
			fmt.Printf("Request %d blocked at %s\n", i, time.Now().Format(time.StampMilli))
		}
		time.Sleep(10 * time.Millisecond) // Wait for 10 milliseconds between each request.
	}
}
