package gblink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	assert := assert.New(t)

	st := NewStack[int]()

	st.Push(1)
	st.Push(2)
	st.Push(3)

	assert.Equal(3, st.Len())

	v, err := st.Pop()
	assert.Nil(err)
	assert.Equal(3, v)
}

func TestStack_Pop(t *testing.T) {
	assert := assert.New(t)

	st := NewStack[int]()

	st.Push(1)
	st.Push(2)
	st.Push(3)

	assert.Equal(3, st.Len())

	v, err := st.Pop()
	assert.Nil(err)
	assert.Equal(3, v)

	v, err = st.Pop()
	assert.Nil(err)
	assert.Equal(2, v)

	v, err = st.Pop()
	assert.Nil(err)
	assert.Equal(1, v)

	_, err = st.Pop()
	assert.NotNil(err)
}

func TestStack_Peek(t *testing.T) {
	assert := assert.New(t)

	st := NewStack[int]()

	st.Push(1)
	st.Push(2)
	st.Push(3)

	assert.Equal(3, st.Len())

	v, err := st.Peek()
	assert.Nil(err)
	assert.Equal(3, v)

	v, err = st.Peek()
	assert.Nil(err)
	assert.Equal(3, v)

	v, err = st.Pop()
	assert.Nil(err)
	assert.Equal(3, v)

	v, err = st.Peek()
	assert.Nil(err)
	assert.Equal(2, v)

	v, err = st.Pop()
	assert.Nil(err)
	assert.Equal(2, v)

	v, err = st.Peek()
	assert.Nil(err)
	assert.Equal(1, v)

	v, err = st.Pop()
	assert.Nil(err)
	assert.Equal(1, v)

	_, err = st.Peek()
	assert.NotNil(err)
}

func TestStack_Len(t *testing.T) {
	assert := assert.New(t)

	st := NewStack[int]()

	st.Push(1)
	st.Push(2)
	st.Push(3)

	assert.Equal(3, st.Len())

	st.Pop()
	assert.Equal(2, st.Len())

	st.Pop()
	assert.Equal(1, st.Len())

	st.Pop()
	assert.Equal(0, st.Len())
}

func TestStack_IsEmpty(t *testing.T) {
	assert := assert.New(t)

	st := NewStack[int]()

	assert.True(st.IsEmpty())

	st.Push(1)
	assert.False(st.IsEmpty())

	st.Pop()
	assert.True(st.IsEmpty())
}
