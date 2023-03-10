package gblink

import (
	"encoding/json"
	"fmt"
	"strings"
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

// Delete the key/value pair with the key k.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	m.Delete("two")
//	fmt.Println(m) // map[one:1 three:3]
func (m MapStringInterface) Delete(k string) {
	delete(m, k)
}

// Delete a random key/value pair from the map that satisfies the callback
// function f.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	m.DeleteFunc(func(k string, v interface{}) bool {
//	    return v.(int) > 1
//	})
//	fmt.Println(m) // map[one:1]
func (m MapStringInterface) DeleteIf(f func(string, interface{}) bool) {
	for k, v := range m {
		if f(k, v) {
			delete(m, k)
		}
	}
}

// Call the callback function f for each key/value pair in the map.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	m.Each(func(k string, v interface{}) {
//	    fmt.Println(k, v)
//	})
func (m MapStringInterface) Each(f func(string, interface{})) {
	for k, v := range m {
		f(k, v)
	}
}

// Clone the map.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	m2 := m.Clone()
//	fmt.Println(m2) // map[one:1 two:2 three:3]
func (m MapStringInterface) Clone() MapStringInterface {
	clone := MapStringInterface{}
	for k, v := range m {
		clone[k] = v
	}
	return clone
}

// Clear the map.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	m.Clear()
//	fmt.Println(m) // map[]
func (m MapStringInterface) Clear() {
	for k := range m {
		delete(m, k)
	}
}

// Json stringifies the map.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "one": 1,
//	    "two": 2,
//	    "three": 3,
//	}
//	s, err := m.JsonString()
//	fmt.Println(s) // {"one":1,"two":2,"three":3}
func (m MapStringInterface) JsonString() (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", MapError{fmt.Errorf("MapError: Cannot marshal %e", err)}
	}
	return string(b), err
}

// Return a value with nested keys. If the key does not exist, return the error
// MapError.
//
// Example:
//
//		m := gblink.MapStringInterface{
//		    "a": 1,
//		    "b": &gblink.MapStringInterface{
//		        "c": 2,
//		        "d": &gblink.MapStringInterface{
//		            "e": 3,
//		        },
//		    },
//		}
//		v, err := m.GetDeep("a.b.c")
//	 if err != nil {
//	     fmt.Println(err) // MapError: Key "a.b.c" does not exist
//	 }
//		v , err := m.GetDeep("b.d.e")
//		fmt.Println(v) // 3
func (m MapStringInterface) GetDeep(keys string) (interface{}, error) {
	keysSlice := strings.Split(keys, ".")
	return m.getDeep(keysSlice)
}

func (m MapStringInterface) getDeep(keys []string) (interface{}, error) {
	if len(keys) == 0 {
		return nil, MapError{fmt.Errorf("MapError: No keys provided")}
	}

	if len(keys) == 1 {
		v, ok := m[keys[0]]
		if !ok {
			return nil, MapError{fmt.Errorf("MapError: Key %q does not exist", keys[0])}
		}
		return v, nil
	}

	// Length of keys is greater than 1
	v, ok := m[keys[0]]
	if !ok {
		return nil, MapError{fmt.Errorf("MapError: Key %q does not exist", keys[0])}
	}

	m2, ok := v.(MapStringInterface)
	if !ok {
		return nil, MapError{fmt.Errorf("MapError: Key %q is not a map", keys[0])}
	}

	return m2.getDeep(keys[1:])
}

// Set a value with nested keys.
//
// Example:
//
//	m := gblink.MapStringInterface{}
//	m.SetDeep("a.b.c", 1)
//	fmt.Println(m) // map[a:map[b:map[c:1]]]
func (m MapStringInterface) SetDeep(keys string, value interface{}) {
	keysSlice := strings.Split(keys, ".")
	m.setDeep(keysSlice, value)
}

func (m MapStringInterface) setDeep(keys []string, value interface{}) {
	if len(keys) == 0 {
		return
	}

	if len(keys) == 1 {
		m[keys[0]] = value
		return
	}

	// Length of keys is greater than 1
	v, ok := m[keys[0]]
	if !ok {
		m[keys[0]] = MapStringInterface{}
		v = m[keys[0]]
	}

	m2, ok := v.(MapStringInterface)
	if !ok {
		m2 = MapStringInterface{}
		m[keys[0]] = m2
	}

	m2.setDeep(keys[1:], value)
}

// Delete a value with nested keys.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "a": 1,
//	    "b": &gblink.MapStringInterface{
//	        "c": 2,
//	        "d": &gblink.MapStringInterface{
//	            "e": 3,
//	        },
//	    },
//	}
//	m.DeleteDeep("b.d.e")
//	fmt.Println(m) // map[a:1 b:map[c:2 d:map[]]]
func (m MapStringInterface) DeleteDeep(keys string) {
	keysSlice := strings.Split(keys, ".")
	m.deleteDeep(keysSlice)
}

func (m MapStringInterface) deleteDeep(keys []string) {
	if len(keys) == 0 {
		return
	}

	if len(keys) == 1 {
		delete(m, keys[0])
		return
	}

	// Length of keys is greater than 1
	v, ok := m[keys[0]]
	if !ok {
		return
	}

	m2, ok := v.(MapStringInterface)
	if !ok {
		return
	}

	m2.deleteDeep(keys[1:])
}

// Return true if the nested key exists.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "a": 1,
//	    "b": &gblink.MapStringInterface{
//	        "c": 2,
//	        "d": &gblink.MapStringInterface{
//	            "e": 3,
//	        },
//	    },
//	}
//	fmt.Println(m.HasDeep("a")) // true
//	fmt.Println(m.HasDeep("b.d.e")) // true
//	fmt.Println(m.HasDeep("b.d.f")) // false
func (m MapStringInterface) HasDeep(keys string) bool {
	keysSlice := strings.Split(keys, ".")
	return m.hasDeep(keysSlice)
}

func (m MapStringInterface) hasDeep(keys []string) bool {
	if len(keys) == 0 {
		return false
	}

	if len(keys) == 1 {
		_, ok := m[keys[0]]
		return ok
	}

	// Length of keys is greater than 1
	v, ok := m[keys[0]]
	if !ok {
		return false
	}

	m2, ok := v.(MapStringInterface)
	if !ok {
		return false
	}

	return m2.hasDeep(keys[1:])
}

// Clean a map by removing all keys with specified values.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "a": 1,
//	    "b": 2,
//	    "c": 3,
//	}
//
// mCleaned := m.Clean(nil)
// fmt.Println(mCleaned) // map[a:1 b:2 c:3]
func (m MapStringInterface) Clean(value interface{}) MapStringInterface {
	cleanedMap := MapStringInterface{}
	for k, v := range m {
		if v != value {
			cleanedMap[k] = v
		}
	}

	return cleanedMap
}

// Clean a map by removing all keys if the callback returns true.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "a": 1,
//	    "b": 2,
//	    "c": 3,
//	}
//
//	mCleaned := m.CleanIf(func(key string, value interface{}) bool {
//	    return value == 2
//	})
//
// fmt.Println(mCleaned) // map[a:1 c:3]
func (m MapStringInterface) CleanIf(callback func(key string, value interface{}) bool) MapStringInterface {
	cleanedMap := MapStringInterface{}
	for k, v := range m {
		if !callback(k, v) {
			cleanedMap[k] = v
		}
	}

	return cleanedMap
}

// Clean deep a map by removing all keys with specified values.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "a": 1,
//	    "b": &gblink.MapStringInterface{
//	        "c": 2,
//	        "d": &gblink.MapStringInterface{
//	            "e": 3,
//	        },
//	    },
//	}
//
// mCleaned := m.CleanDeep(nil)
// fmt.Println(mCleaned) // map[a:1 b:map[c:2 d:map[e:3]]]
func (m MapStringInterface) CleanDeep(value interface{}) MapStringInterface {
	cleanedMap := MapStringInterface{}
	for k, v := range m {
		if v != value {
			if m2, ok := v.(MapStringInterface); ok {
				cleanedMap[k] = m2.CleanDeep(value)
			} else {
				cleanedMap[k] = v
			}
		}
	}

	return cleanedMap
}

// Clean deep a map by removing all keys if the callback returns true.
//
// Example:
//
//	m := gblink.MapStringInterface{
//	    "a": 1,
//	    "b": &gblink.MapStringInterface{
//	        "c": 2,
//	        "d": &gblink.MapStringInterface{
//	            "e": 3,
//	        },
//	    },
//	}
//
//	mCleaned := m.CleanDeepIf(func(key string, value interface{}) bool {
//	    return value == 2
//	})
//
// fmt.Println(mCleaned) // map[a:1 b:map[d:map[e:3]]]
func (m MapStringInterface) CleanDeepIf(callback func(key string, value interface{}) bool) MapStringInterface {
	cleanedMap := MapStringInterface{}
	for k, v := range m {
		if !callback(k, v) {
			if m2, ok := v.(MapStringInterface); ok {
				cleanedMap[k] = m2.CleanDeepIf(callback)
			} else {
				cleanedMap[k] = v
			}
		}
	}

	return cleanedMap
}

// Deep merge maps.
//
// Example:
//
//	m1 := gblink.MapStringInterface{
//	    "a": 1,
//	    "b": &gblink.MapStringInterface{
//	        "c": 2,
//	        "d": &gblink.MapStringInterface{
//	            "e": 3,
//	        },
//	    },
//	}
//
//	m2 := gblink.MapStringInterface{
//	    "a": 4,
//	    "b": &gblink.MapStringInterface{
//	        "d": &gblink.MapStringInterface{
//	            "f": 5,
//	        },
//	    },
//	}
//
//	m3 := m1.MergeDeep(m2)
//	fmt.Println(m3) // map[a:4 b:map[c:2 d:map[e:3 f:5]]]
func (m MapStringInterface) MergeDeep(m2 MapStringInterface) MapStringInterface {
	mergedMap := MapStringInterface{}
	for k, v := range m {
		mergedMap[k] = v
	}

	for k, v := range m2 {
		if m2v, ok := v.(MapStringInterface); ok {
			if m1v, ok := mergedMap[k].(MapStringInterface); ok {
				mergedMap[k] = m1v.MergeDeep(m2v)
			} else {
				mergedMap[k] = m2v
			}
		} else {
			mergedMap[k] = v
		}
	}

	return mergedMap
}
