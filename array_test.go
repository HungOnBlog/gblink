package gblink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray_Len(t *testing.T) {
	assert := assert.New(t)

	a := Array[int]{1, 2, 3}

	assert.Equal(3, a.Len())
}

func TestArray_Append(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.Append(1)
	a.Append(2)
	a.Append(3)

	assert.Equal(3, a.Len())
}

func TestArray_AppendAll(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	assert.Equal(3, a.Len())
}

func TestArray_AppendArray(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendArray(&Array[int]{1, 2, 3})

	assert.Equal(3, a.Len())
}

func TestArray_AppendArrayAll(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendArrayAll(&Array[int]{1, 2, 3}, &Array[int]{4, 5, 6})

	assert.Equal(6, a.Len())
}

func TestArray_Clear(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)
	a.Clear()

	assert.Equal(0, a.Len())
}

func TestArray_Contains(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	assert.True(a.Contains(2))

	assert.False(a.Contains(4))
}

func TestArray_ContainsAll(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	assert.True(a.ContainsAll(1, 2, 3))

	assert.False(a.ContainsAll(1, 2, 4))
}

func TestArray_Some(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	assert.True(a.Some(func(i int) bool {
		return i == 2
	}))

	assert.False(a.Some(func(i int) bool {
		return i == 4
	}))
}

func TestArray_Every(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	assert.True(a.Every(func(i int) bool {
		return i > 0
	}))

	assert.False(a.Every(func(i int) bool {
		return i > 1
	}))

}

func TestArray_Filter(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	b := a.Filter(func(i int) bool {
		return i > 1
	})

	assert.Equal(2, b.Len())
}

func TestArray_Find(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	b := a.Find(func(i int) bool {
		return i > 1
	})

	assert.Equal(2, b)

	c := a.Find(func(i int) bool {
		return i > 3
	})

	assert.Equal(0, c)
}

func TestArray_FindIndex(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	b := a.FindIndex(func(i int) bool {
		return i > 1
	})

	assert.Equal(1, b)

	c := a.FindIndex(func(i int) bool {
		return i > 3
	})

	assert.Equal(-1, c)
}

func TestArray_FindLast(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	b := a.FindLast(func(i int) bool {
		return i > 1
	})

	assert.Equal(3, b)

	c := a.FindLast(func(i int) bool {
		return i > 3
	})

	assert.Equal(0, c)
}

func TestArray_FindLastIndex(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	b := a.FindLastIndex(func(i int) bool {
		return i > 1
	})

	assert.Equal(2, b)

	c := a.FindLastIndex(func(i int) bool {
		return i > 3
	})

	assert.Equal(-1, c)
}

func TestArray_Each(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	a.Each(func(i int, v int) {
		assert.Equal(i+1, v)
	})
}

func TestArray_Copy(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	b := a.Copy()

	assert.Equal(3, b.Len())
}

func TestArray_EachIndex(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	a.EachIndex(func(i int) {
		assert.Equal(i, i)
	})
}

func TestArray_EachValue(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	a.EachValue(func(v int) {
		assert.Equal(v, v)
	})
}

func TestArray_Map(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	b := a.Map(func(i int) int {
		return i + 1
	})

	assert.Equal(3, b.Len())
}

func TestArray_SortBy(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 2, 3)

	a.SortBy(func(i, j int) int {
		if i > j {
			return 1
		}

		if i < j {
			return -1
		}

		return 0
	})

	assert.Equal(1, a[0])
	assert.Equal(2, a[1])
	assert.Equal(3, a[2])
}

func TestArray_Sort(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(1, 1, 5, 6, 8, 2, 3, 5, 9, 3, 5, 6, 7, 8, 9, 0, 4, 3, 2, 1)

	a.Sort(true)

	assert.Equal(0, a[0])
	assert.Equal(1, a[1])
	assert.Equal(1, a[2])

	a.Sort(false)

	assert.Equal(9, a[0])
	assert.Equal(9, a[1])
	assert.Equal(8, a[2])
}

func TestArray_Reverse(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	a.AppendAll(3, 1, 2)

	a.Reverse()

	assert.Equal(2, a[0])
	assert.Equal(1, a[1])
	assert.Equal(3, a[2])
}

func TestArray_IndexOf(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]

	a.AppendAll(1, 2, 3)

	assert.Equal(1, a.IndexOf(2))

	assert.Equal(-1, a.IndexOf(4))
}

func TestArray_Insert(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]

	a.AppendAll(1, 2, 3)

	a.Insert(1, 4)

	assert.Equal(4, a[1])

	// Error
	err := a.Insert(10, 4)
	assert.NotNil(err)
}

func TestArray_Remove(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]

	a.AppendAll(1, 2, 3)

	a.Remove(1)

	assert.Equal(3, a[1])

	// Error
	err := a.Remove(10)
	assert.NotNil(err)
}

func TestArray_RemoveRange(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]

	a.AppendAll(1, 2, 3, 4, 5)

	a.RemoveRange(1, 3)

	assert.Equal(4, a[1])
}

func TestArray_Slice(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]

	a.AppendAll(1, 2, 3, 4, 5)

	b := a.Slice(1, 3)

	assert.Equal(2, b[0])
}

func TestArray_InsertArray(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]
	var b Array[int]

	a.AppendAll(1, 2, 3)
	b.AppendAll(4, 5, 6)

	a.InsertArray(1, b)

	assert.Equal(4, a[1])

	// Error
	err := a.InsertArray(10, b)
	assert.NotNil(err)
}

func TestArray_Reduce(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]

	a.AppendAll(1, 2, 3, 4, 5)

	b := a.Reduce(func(i, j int) int {
		return i + j
	}, 0)

	assert.Equal(15, b)

	var c Array[string]

	c.AppendAll("a", "b", "c", "d", "e")

	d := c.Reduce(func(i, j string) string {
		return i + j
	}, "")

	assert.Equal("abcde", d)
}

func TestArray_ReduceIf(t *testing.T) {
	assert := assert.New(t)

	var a Array[int]

	a.AppendAll(1, 2, 3, 4, 5)

	b := a.ReduceIf(func(i, j int) int {
		return i + j
	}, func(i int) bool {
		return i%2 == 0
	}, 0)

	assert.Equal(6, b)
}
