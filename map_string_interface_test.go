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

func TestMapStringInterface_Delete(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	m.Delete("one")
	assert.Equal(2, m.Len())
	assert.Equal(2, m["two"])
	assert.Equal(3, m["three"])
}

func TestMapStringInterface_DeleteIf(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	m.DeleteIf(func(key string, value interface{}) bool {
		return value.(int) > 1
	})

	assert.Equal(1, m.Len())
	assert.Equal(1, m["one"])
}

func TestMapStringInterface_Each(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	m.Each(func(key string, value interface{}) {
		assert.Equal(m[key], value)
	})
}

func TestMapStringInterface_Clone(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	m2 := m.Clone()

	assert.Equal(m, m2)
}

func TestMapStringInterface_Clear(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)

	m.Clear()
	assert.Equal(0, m.Len())
}

func TestMapStringInterface_JsonString(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	assert := assert.New(t)
	jsonString, err := m.JsonString()
	assert.Nil(err)
	assert.Equal("{\"one\":1,\"three\":3,\"two\":2}", jsonString)
}

func TestMapStringInterface_GetDeep(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)

	_, err := m.GetDeep("")
	assert.NotNil(err)

	v, err := m.GetDeep("four.five")
	assert.Nil(err)
	assert.Equal(5, v)

	_, err = m.GetDeep("four.seven")
	assert.NotNil(err)

	v, err = m.GetDeep("four")
	assert.Nil(err)
	assert.Equal(MapStringInterface{"five": 5, "six": 6}, v)
}

func TestMapStringInterface_getDeep(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)

	_, err := m.getDeep([]string{})
	assert.NotNil(err)

	v, err := m.getDeep([]string{"one"})
	assert.Nil(err)
	assert.Equal(1, v)

	v, err = m.getDeep([]string{"four", "five"})
	assert.Nil(err)
	assert.Equal(5, v)

	_, err = m.getDeep([]string{"four", "seven"})
	assert.NotNil(err)
}

func TestMapStringInterface_SetDeep(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)
	m.SetDeep("", 0)
	assert.Equal(5, m.Len())

	m.SetDeep("four.five", 7)
	assert.Equal(7, m["four"].(MapStringInterface)["five"])

	m.SetDeep("four.seven", 8)
	assert.Equal(8, m["four"].(MapStringInterface)["seven"])

	m.SetDeep("four", 9)
	assert.Equal(9, m["four"])
}

func TestMapStringInterface_setDeep(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)

	m.setDeep([]string{}, 0)
	assert.Equal(4, m.Len())

	m.setDeep([]string{"four", "five"}, 7)
	assert.Equal(7, m["four"].(MapStringInterface)["five"])

	m.setDeep([]string{"four", "seven"}, 8)
	assert.Equal(8, m["four"].(MapStringInterface)["seven"])

	m.setDeep([]string{"four"}, 9)
	assert.Equal(9, m["four"])
}

func TestMapStringInterface_DeleteDeep(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)

	m.DeleteDeep("four.five")
	assert.Equal(1, m["four"].(MapStringInterface).Len())
	assert.Equal(6, m["four"].(MapStringInterface)["six"])
}

func TestMapStringInterface_deleteDeep(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)

	m.deleteDeep([]string{})
	assert.Equal(4, m.Len())

	m.deleteDeep([]string{"one"})
	assert.Equal(3, m.Len())

	m.deleteDeep([]string{"four", "five"})
	assert.Equal(1, m["four"].(MapStringInterface).Len())
}

func TestMapStringInterface_HasDeep(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)

	assert.False(m.HasDeep(""))
	assert.True(m.HasDeep("one"))
	assert.True(m.HasDeep("four.five"))
	assert.False(m.HasDeep("four.seven"))
}

func TestMapStringInterface_hasDeep(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)

	assert.False(m.hasDeep([]string{}))
	assert.True(m.hasDeep([]string{"one"}))
	assert.True(m.hasDeep([]string{"four", "five"}))
	assert.False(m.hasDeep([]string{"four", "seven"}))
}

func TestMapStringInterface_Clean(t *testing.T) {
	m := MapStringInterface{
		"one":   nil,
		"two":   nil,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)

	cleanedMap := m.Clean(nil)
	assert.Equal(2, cleanedMap.Len())
	assert.Equal(3, cleanedMap["three"])
}

func TestMapStringInterface_CleanIf(t *testing.T) {
	m := MapStringInterface{
		"one":   nil,
		"two":   nil,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)

	cleanedMap := m.CleanIf(func(key string, value interface{}) bool {
		return value == nil
	})
	assert.Equal(2, cleanedMap.Len())
	assert.Equal(3, cleanedMap["three"])
}

func TestMapStringInterface_CleanDeep(t *testing.T) {
	m := MapStringInterface{
		"one":   nil,
		"two":   nil,
		"three": 3,
		"four": MapStringInterface{
			"five": nil,
			"six":  6,
		},
	}
	assert := assert.New(t)

	cleanedMap := m.CleanDeep(nil)
	assert.Equal(2, cleanedMap.Len())
	assert.Equal(3, cleanedMap["three"])
	assert.Equal(1, cleanedMap["four"].(MapStringInterface).Len())
}

func TestMapStringInterface_CleanDeepIf(t *testing.T) {
	m := MapStringInterface{
		"one":   nil,
		"two":   nil,
		"three": 3,
		"four": MapStringInterface{
			"five": nil,
			"six":  6,
		},
	}
	assert := assert.New(t)

	cleanedMap := m.CleanDeepIf(func(key string, value interface{}) bool {
		return value == nil
	})
	assert.Equal(2, cleanedMap.Len())
	assert.Equal(3, cleanedMap["three"])
	assert.Equal(1, cleanedMap["four"].(MapStringInterface).Len())
}

func TestMapStringInterface_MergeDeep(t *testing.T) {
	m := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
		"four": MapStringInterface{
			"five": 5,
			"six":  6,
		},
	}
	assert := assert.New(t)

	m2 := MapStringInterface{
		"one":   1,
		"two":   2,
		"three": 3,
		"four": MapStringInterface{
			"five":  5,
			"seven": 7,
			"eight": 8,
		},
	}

	m3 := m.MergeDeep(m2)
	assert.Equal(4, m3.Len())
	assert.Equal(3, m3["three"])
	m3four := m3["four"].(MapStringInterface)
	assert.Equal(4, m3four.Len())
	assert.Equal(5, m3four["five"])
	assert.Equal(6, m3four["six"])
	assert.Equal(7, m3four["seven"])
	assert.Equal(8, m3four["eight"])
}
