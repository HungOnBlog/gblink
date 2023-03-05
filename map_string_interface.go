package gblink

import (
	"fmt"
)

type MapStringInterface map[string]interface{}

// Returns the value associated with the key k.
// If the key is not found, it returns an error.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	v, err := m.Get("two")
//	if err != nil {
//	    fmt.Println(err) // MapError: key two not found
//	} else {
//	    fmt.Println(v) // 2
//	}
func (m MapStringInterface) Get(k string) (interface{}, error) {
	v, ok := m[k]
	if !ok {
		return m[k], &MapError{fmt.Errorf("MapError: key %v not found", k)}
	}
	return v, nil
}

// Set the value v associated with the key k.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	m.Set("four", 4)
//	fmt.Println(m) // map[one:1 two:2 three:3 four:4]
func (m MapStringInterface) Set(k string, v interface{}) {
	m[k] = v
}

// Stringify the map.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	fmt.Println(m) // map[one:1 two:2 three:3]
func (m MapStringInterface) String() string {
	return fmt.Sprintf("%v", map[string]interface{}(m))
}

// Return the length of the map.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	fmt.Println(m.Len()) // 3
func (m MapStringInterface) Len() int {
	return len(m)
}

// Return true if the map is empty.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	fmt.Println(m.IsEmpty()) // false
func (m MapStringInterface) IsEmpty() bool {
	return len(m) == 0
}

// Return true if the map contains the key k.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	fmt.Println(m.Contains("two")) // true
func (m MapStringInterface) Contains(k string) bool {
	_, ok := m[k]
	return ok
}

// Return a slice of keys.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	fmt.Println(m.Keys()) // [one two three]
func (m MapStringInterface) Keys() []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// Return a slice of values.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	fmt.Println(m.Values()) // [1 2 3]
func (m MapStringInterface) Values() []interface{} {
	values := make([]interface{}, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

// Return a slice of key/value pairs.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	fmt.Println(m.Pairs()) // [[one 1] [two 2] [three 3]]
func (m MapStringInterface) Pairs() [][2]interface{} {
	pairs := make([][2]interface{}, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, [2]interface{}{k, v})
	}
	return pairs
}

// Return a new map with the merged key/value pairs of the current map and the
// map passed as argument.
//
// Example:
//
//	m1 := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	m2 := gblink.MapStringInterface{
//	    "four": 4,
//	    "five": 5,
//	    "six": 6,
//	}
//	m3 := gblink.MapStringInterface{
//	    "seven": 7,
//	    "eight": 8,
//	    "nine": 9,
//	}
//	m4 := m1.Merge(m2, m3)
//
// fmt.Println(m4) // map["one":1 "two":2 "three":3 "four":4 "five":5 "six":6 "seven":7 "eight":8 "nine":9]
func (m MapStringInterface) Merge(maps ...MapStringInterface) MapStringInterface {
	merged := MapStringInterface{}
	for k, v := range m {
		merged[k] = v
	}
	for _, mm := range maps {
		for k, v := range mm {
			merged[k] = v
		}
	}
	return merged
}

// Return a new map with the key/value pairs of the current map where the
// callback function returns true.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	m2 := m.Filter(func(k string, v interface{}) bool {
//	    return v.(int) > 1
//	})
//	fmt.Println(m2) // map[two:2 three:3]
func (m MapStringInterface) Filter(f func(string, interface{}) bool) MapStringInterface {
	filtered := MapStringInterface{}
	for k, v := range m {
		if f(k, v) {
			filtered[k] = v
		}
	}
	return filtered
}
