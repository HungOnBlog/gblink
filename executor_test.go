package gblink

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecutor_Execute(t *testing.T) {
	assert := assert.New(t)

	executor := Executor[int]{}

	executor.Execute(func() (int, error) {
		return 1, nil
	}, func(value int) {
		assert.Equal(1, value)
	}, func(err error) {

	})

	executor.Execute(func() (int, error) {
		return -1, fmt.Errorf("error")
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
	})
}

func TestExecutor_ExecuteWithTimeout(t *testing.T) {
	assert := assert.New(t)

	executor := Executor[int]{}

	executor.ExecuteWithTimeout(func() (int, error) {
		return 1, nil
	}, func(value int) {
		assert.Equal(1, value)
	}, func(err error) {
	}, 1)

	executor.ExecuteWithTimeout(func() (int, error) {
		return -1, fmt.Errorf("error")
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
	}, 1)

	// Timeout
	executor.ExecuteWithTimeout(func() (int, error) {
		i := 1000000
		for i > 0 {
			i--
		}
		return 1, nil
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
		assert.Contains(err.Error(), "timeout")
	}, 1)

}

func TestExecutor_ExecuteWithTimeoutAndRetry(t *testing.T) {
	assert := assert.New(t)

	executor := Executor[int]{}

	executor.ExecuteWithTimeoutAndRetry(func() (int, error) {
		return 1, nil
	}, func(value int) {
		assert.Equal(1, value)
	}, func(err error) {
	}, 1, 1)

	executor.ExecuteWithTimeoutAndRetry(func() (int, error) {
		return -1, fmt.Errorf("error")
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
	}, 1, 1)

	// Timeout
	executor.ExecuteWithTimeoutAndRetry(func() (int, error) {
		i := 1000000
		for i > 0 {
			i--
		}
		return 1, nil
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
		assert.Contains(err.Error(), "timeout")
	}, 1, 1)

	// Retry
	executor.ExecuteWithTimeoutAndRetry(func() (int, error) {
		return -1, fmt.Errorf("error")
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
		fmt.Print(err.Error())
	}, 100000, 2)
}

func TestExecutor_ExecuteRetryBackOff(t *testing.T) {
	assert := assert.New(t)

	executor := Executor[int]{}

	executor.ExecuteRetryBackOff(func() (int, error) {
		return 1, nil
	}, func(value int) {
		assert.Equal(1, value)
	}, func(err error) {
	}, 1, 1, 1)

	executor.ExecuteRetryBackOff(func() (int, error) {
		return -1, fmt.Errorf("error")
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
	}, 1, 1, 1)

	// Retry
	executor.ExecuteRetryBackOff(func() (int, error) {
		return -1, fmt.Errorf("error")
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
	}, time.Microsecond*2, 2, 1)

	// Retry
	executor.ExecuteRetryBackOff(func() (int, error) {
		return -1, fmt.Errorf("error")
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
	}, time.Microsecond*2, 2, time.Millisecond*2)
}

func TestExecutor_ExecuteWithTimeoutAndRetryBackOff(t *testing.T) {
	assert := assert.New(t)

	executor := Executor[int]{}

	executor.ExecuteWithTimeoutAndRetryBackOff(func() (int, error) {
		return 1, nil
	}, func(value int) {
		assert.Equal(1, value)
	}, func(err error) {
	}, 1, 1, 1)

	executor.ExecuteWithTimeoutAndRetryBackOff(func() (int, error) {
		return -1, fmt.Errorf("error")
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
	}, 1, 1, 1)

	// Timeout
	executor.ExecuteWithTimeoutAndRetryBackOff(func() (int, error) {
		i := 1000000
		for i > 0 {
			i--
		}
		return 1, nil
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
		assert.Contains(err.Error(), "timeout")
	}, 1, 1, 1)

	// Retry
	executor.ExecuteWithTimeoutAndRetryBackOff(func() (int, error) {
		return -1, fmt.Errorf("error")
	}, func(value int) {
	}, func(err error) {
		assert.NotNil(err)
	}, time.Microsecond*2, 2, 1)
}
