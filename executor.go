package gblink

import (
	"fmt"
	"time"
)

type Executor[V any] struct{}

type ExecutorError struct {
	error
}

// Execute function which return a value of type V and the error
// If the error is not nil, run onError function which pass the error as parameter
// If the error is nil, run onSuccess function which pass the value as parameter
func (e *Executor[V]) Execute(fn func() (V, error), onSuccess func(V), onError func(error)) {
	value, err := fn()
	if err != nil {
		onError(err)
	}
	onSuccess(value)
}

// ExecuteWithTimeout function which return a value of type V and the error
// If the error is not nil, run onError function which pass the error as parameter
// If the error is nil, run onSuccess function which pass the value as parameter
func (e *Executor[V]) ExecuteWithTimeout(fn func() (V, error), onSuccess func(V), onError func(error), duration time.Duration) {
	done := make(chan bool)
	go func() {
		value, err := fn()
		if err != nil {
			onError(err)
		}
		onSuccess(value)
		done <- true
	}()
	select {
	case <-done:
		return
	case <-time.After(duration):
		onError(ExecutorError{fmt.Errorf("ExecutorError: timeout")})
	}
}

// ExecuteWithTimeoutAndRetry function which return a value of type V and the error
// If the error is not nil, wait for the duration and run the function again util reach the maxRetry
// If the error is nil, run onSuccess function which pass the value as parameter
func (e *Executor[V]) ExecuteWithTimeoutAndRetry(fn func() (V, error), onSuccess func(V), onError func(error), duration time.Duration, maxRetry int) {
	done := make(chan bool)
	go func() {
		value, err := fn()
		if err != nil {
			if maxRetry > 0 {
				time.Sleep(duration)
				e.ExecuteWithTimeoutAndRetry(fn, onSuccess, onError, duration, maxRetry-1)
			} else {
				onError(err)
			}
		}
		onSuccess(value)
		done <- true
	}()
	select {
	case <-done:
		return
	case <-time.After(duration):
		onError(ExecutorError{fmt.Errorf("ExecutorError: timeout")})
	}
}

// ExecuteWithTimeoutAndRetryBackOff function which return a value of type V and the error
// If the error is not nil, wait for the duration and run the function again util reach the maxRetry
// If the error is nil, run onSuccess function which pass the value as parameter
// The duration is recalculated by adding the duration to the backOffDuration
func (e *Executor[V]) ExecuteWithTimeoutAndRetryBackOff(fn func() (V, error), onSuccess func(V), onError func(error), duration time.Duration, maxRetry int, backOffDuration time.Duration) {
	done := make(chan bool)
	go func() {
		value, err := fn()
		if err != nil {
			if maxRetry > 0 {
				time.Sleep(duration)
				e.ExecuteWithTimeoutAndRetryBackOff(fn, onSuccess, onError, duration+backOffDuration, maxRetry-1, backOffDuration)
			} else {
				onError(err)
			}
		}
		onSuccess(value)
		done <- true
	}()
	select {
	case <-done:
		return
	case <-time.After(duration):
		onError(ExecutorError{fmt.Errorf("ExecutorError: timeout")})
	}
}

// ExecuteRetry function which return a value of type V and the error
// If the error is not nil, wait for the duration and run the function again util reach the maxRetry
// If the error is nil, run onSuccess function which pass the value as parameter
func (e *Executor[V]) ExecuteRetry(fn func() (V, error), onSuccess func(V), onError func(error), duration time.Duration, maxRetry int) {
	value, err := fn()
	if err != nil {
		if maxRetry > 0 {
			time.Sleep(duration)
			e.ExecuteRetry(fn, onSuccess, onError, duration, maxRetry-1)
		} else {
			onError(err)
		}
	}
	onSuccess(value)
}

// ExecuteRetryBackOff function which return a value of type V and the error
// If the error is not nil, wait for the duration and run the function again util reach the maxRetry
// If the error is nil, run onSuccess function which pass the value as parameter
// The duration is recalculated by adding the duration to the backOffDuration
func (e *Executor[V]) ExecuteRetryBackOff(fn func() (V, error), onSuccess func(V), onError func(error), duration time.Duration, maxRetry int, backOffDuration time.Duration) {
	value, err := fn()
	if err != nil {
		if maxRetry > 0 {
			time.Sleep(duration)
			e.ExecuteRetryBackOff(fn, onSuccess, onError, duration+backOffDuration, maxRetry-1, backOffDuration)
		} else {
			onError(err)
		}
	}
	onSuccess(value)
}
