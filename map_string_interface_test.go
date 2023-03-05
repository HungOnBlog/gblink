package gblink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapStringInterface_Get(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	v, err := m.Get("two")
	assert.Nil(err)
	assert.Equal(2, v)

	_, err = m.Get("four")
	assert.NotNil(err)
}

func TestMapStringInterface_Set(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	m.Set("four", 4)
	assert.Equal(4, m["four"])
}

func TestMapStringInterface_String(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	assert.Equal("map[one:1 three:3 two:2]", m.String())
}

func TestMapStringInterface_Len(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	assert.Equal(3, m.Len())
}

func TestMapStringInterface_IsEmpty(t *testing.T) {
	m := MapStringInterface{}
	assert := assert.New(t)

	assert.True(m.IsEmpty())
}

func TestMapStringInterface_Contains(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	assert.True(m.Contains("one"))
	assert.False(m.Contains("four"))
}

func TestMapStringInterface_Keys(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)
	keys := m.Keys()

	assert.Equal(3, len(keys))
	assert.Contains(keys, "one")
	assert.Contains(keys, "two")
	assert.Contains(keys, "three")
}

func TestMapStringInterface_Values(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)
	values := m.Values()

	assert.Equal(3, len(values))
	assert.Contains(values, 1)
	assert.Contains(values, 2)
	assert.Contains(values, 3)
}

func TestMapStringInterface_Pairs(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)
	pairs := m.Pairs()

	assert.Equal(3, len(pairs))
	assert.Contains(pairs, [2]interface{}{"one", 1})
	assert.Contains(pairs, [2]interface{}{"two", 2})
	assert.Contains(pairs, [2]interface{}{"three", 3})
}

func TestMapStringInterface_Merge(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	m2 := MapStringInterface{
		"four": 4,
		"five": 5,
	}
	mm := m.Merge(m2)

	assert.Equal(5, mm.Len())
	assert.Equal(1, mm["one"])
}

func TestMapStringInterface_Filter(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	mm := m.Filter(func(key string, value interface{}) bool {
		return value.(int) > 1
	})

	assert.Equal(2, mm.Len())
	assert.Equal(2, mm["two"])
	assert.Equal(3, mm["three"])
}
