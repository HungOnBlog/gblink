package gblink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLikedList_Append(t *testing.T) {
	assert := assert.New(t)

	list := NewLikedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(3, list.Len())
}

func TestLinkedList_Prepend(t *testing.T) {
	assert := assert.New(t)

	list := NewLikedList[int]()
	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)
	assert.Equal(3, list.Len())
	assert.Equal(3, list.Head.Value)
}

func TestLinkedList_Insert(t *testing.T) {
	assert := assert.New(t)

	list := NewLikedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Insert(2, 4)
	assert.Equal(4, list.Len())
	assert.Equal(4, list.Head.Next.Next.Value)
}

func TestLinkedList_Remove(t *testing.T) {
	assert := assert.New(t)

	list := NewLikedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(2)
	assert.Equal(2, list.Len())
	assert.Equal(2, list.Head.Next.Value)
}

func TestLinkedList_Get(t *testing.T) {
	assert := assert.New(t)

	list := NewLikedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	value, err := list.Get(2)
	assert.Nil(err)
	assert.Equal(3, value)

	value, err = list.Get(3)
	assert.NotNil(err)
	assert.Equal(0, value)
}

func TestLinkedList_IndexOf(t *testing.T) {
	assert := assert.New(t)

	list := NewLikedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	index := list.IndexOf(3)
	assert.Equal(2, index)

	index = list.IndexOf(4)
	assert.Equal(-1, index)
}

func TestLikedList_Contains(t *testing.T) {
	assert := assert.New(t)

	list := NewLikedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	assert.True(list.Contains(3))
	assert.False(list.Contains(4))
}

func TestLikedList_Clear(t *testing.T) {
	assert := assert.New(t)

	list := NewLikedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Clear()
	assert.Equal(0, list.Len())
}
