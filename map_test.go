package gblink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Get(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	v, err := m.Get(2)
	assert.Nil(err)
	assert.Equal("two", v)

	_, err = m.Get(4)
	assert.NotNil(err)
}

func TestMap_Set(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	m.Set(4, "four")
	assert.Equal("four", m[4])
}

func TestMap_String(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	assert.Equal("map[1:one 2:two 3:three]", m.String())
}

func TestMap_Len(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	assert.Equal(3, m.Len())
}

func TestMap_IsEmpty(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	assert.False(m.IsEmpty())
}

func TestMap_Contains(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	assert.True(m.Contains(2))
	assert.False(m.Contains(4))
}

func TestMap_Keys(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	keys := m.Keys()
	assert.Equal(3, len(keys))
	assert.Contains(keys, 1)
	assert.Contains(keys, 2)
	assert.Contains(keys, 3)
}

func TestMap_Values(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	values := m.Values()
	assert.Equal(3, len(values))
	assert.Contains(values, "one")
	assert.Contains(values, "two")
	assert.Contains(values, "three")
}

func TestMap_Pairs(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	pairs := m.Pairs()
	assert.Equal(3, len(pairs))
	assert.Contains(pairs, [2]interface{}{1, "one"})
	assert.Contains(pairs, [2]interface{}{2, "two"})
	assert.Contains(pairs, [2]interface{}{3, "three"})
}

func TestMap_Merge(t *testing.T) {
	assert := assert.New(t)

	m1 := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	m2 := Map[int, string]{
		4: "four",
		5: "five",
		6: "six",
	}

	m3 := Map[int, string]{
		7: "seven",
		8: "eight",
		9: "nine",
	}

	m := m1.Merge(m2, m3)
	assert.Equal(9, len(m))
	assert.Equal("one", m[1])
	assert.Equal("two", m[2])
	assert.Equal("three", m[3])
	assert.Equal("four", m[4])
	assert.Equal("five", m[5])
	assert.Equal("six", m[6])
	assert.Equal("seven", m[7])
	assert.Equal("eight", m[8])
	assert.Equal("nine", m[9])
}

func TestMap_Filter(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	m2 := m.Filter(func(k int, v string) bool {
		return k%2 == 0
	})

	assert.Equal(1, len(m2))
	assert.Equal("two", m2[2])
}

func TestMap_Delete(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	m.Delete(2)
	assert.Equal(2, len(m))
	assert.Equal("one", m[1])
	assert.Equal("three", m[3])
}

func TestMap_DeleteIf(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	m.DeleteIf(func(k int, v string) bool {
		return k%2 == 0
	})

	assert.Equal(2, len(m))
	assert.Equal("one", m[1])
	assert.Equal("three", m[3])
}

func TestMap_Each(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	m.Each(func(k int, v string) {
		m[k] = v + "!"
	})

	assert.Equal("one!", m[1])
	assert.Equal("two!", m[2])
	assert.Equal("three!", m[3])
}

func TestMap_Clone(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	m2 := m.Clone()

	assert.Equal(m, m2)
}

func TestMap_Clear(t *testing.T) {
	assert := assert.New(t)

	m := Map[int, string]{
		1: "one",
		2: "two",
		3: "three",
	}

	m.Clear()

	assert.Equal(0, len(m))
}
