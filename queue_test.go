package gblink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Push(t *testing.T) {

	assert := assert.New(t)

	queue := NewQueue[int]()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	assert.Equal(3, queue.Len())

	v, err := queue.Pop()
	assert.Nil(err)
	assert.Equal(1, v)
}

func TestQueue_Pop(t *testing.T) {

	assert := assert.New(t)

	queue := NewQueue[int]()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	assert.Equal(3, queue.Len())

	v, err := queue.Pop()
	assert.Nil(err)
	assert.Equal(1, v)

	v, err = queue.Pop()
	assert.Nil(err)
	assert.Equal(2, v)

	v, err = queue.Pop()
	assert.Nil(err)
	assert.Equal(3, v)

	_, err = queue.Pop()
	assert.NotNil(err)
}

func TestQueue_Peek(t *testing.T) {

	assert := assert.New(t)

	queue := NewQueue[int]()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	assert.Equal(3, queue.Len())

	v, err := queue.Peek()
	assert.Nil(err)
	assert.Equal(1, v)

	v, err = queue.Peek()
	assert.Nil(err)
	assert.Equal(1, v)

	v, err = queue.Peek()
	assert.Nil(err)
	assert.Equal(1, v)
}

func TestQueue_Len(t *testing.T) {

	assert := assert.New(t)

	queue := NewQueue[int]()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	assert.Equal(3, queue.Len())

	v, err := queue.Pop()
	assert.Nil(err)
	assert.Equal(1, v)

	assert.Equal(2, queue.Len())

	v, err = queue.Pop()
	assert.Nil(err)
	assert.Equal(2, v)

	assert.Equal(1, queue.Len())

	v, err = queue.Pop()
	assert.Nil(err)
	assert.Equal(3, v)

	assert.Equal(0, queue.Len())

	_, err = queue.Pop()
	assert.NotNil(err)
}

func TestQueue_IsEmpty(t *testing.T) {

	assert := assert.New(t)

	queue := NewQueue[int]()

	assert.True(queue.IsEmpty())

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	assert.False(queue.IsEmpty())

	v, err := queue.Pop()
	assert.Nil(err)
	assert.Equal(1, v)

	assert.False(queue.IsEmpty())

	v, err = queue.Pop()
	assert.Nil(err)
	assert.Equal(2, v)

	assert.False(queue.IsEmpty())

	v, err = queue.Pop()
	assert.Nil(err)
	assert.Equal(3, v)

	assert.True(queue.IsEmpty())

	_, err = queue.Pop()
	assert.NotNil(err)
}
