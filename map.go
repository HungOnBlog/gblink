package gblink

import "fmt"

type Map[K comparable, V any] map[K]V

type MapError struct {
	error
}

// Returns the value associated with the key k.
// If the key is not found, it returns an error.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	v, err := m.Get(2)
//	if err != nil {
//	    fmt.Println(err) // MapError: key 2 not found
//	} else {
//	    fmt.Println(v) // two
//	}
func (m Map[K, V]) Get(k K) (V, error) {
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
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	m.Set(4, "four")
//	fmt.Println(m) // map[1:one 2:two 3:three 4:four]
func (m Map[K, V]) Set(k K, v V) {
	m[k] = v
}

// Stringify the map.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	fmt.Println(m) // map[1:one 2:two 3:three]
func (m Map[K, V]) String() string {
	return fmt.Sprintf("%v", map[K]V(m))
}

// Returns the number of elements in the map.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	fmt.Println(m.Len()) // 3
func (m Map[K, V]) Len() int {
	return len(m)
}

// Returns true if the map is empty.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	fmt.Println(m.IsEmpty()) // false
func (m Map[K, V]) IsEmpty() bool {
	return len(m) == 0
}

// Returns true if the map contains the key k.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	fmt.Println(m.Contains(2)) // true
func (m Map[K, V]) Contains(k K) bool {
	_, ok := m[k]
	return ok
}

// Returns a slice of the keys in the map.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	fmt.Println(m.Keys()) // [1 2 3]
func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Returns a slice of the values in the map.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	fmt.Println(m.Values()) // [one two three]
func (m Map[K, V]) Values() []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Returns a slice of the key-value pairs in the map.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	fmt.Println(m.Pairs()) // [[1 one] [2 two] [3 three]]
func (m Map[K, V]) Pairs() [][2]interface{} {
	pairs := make([][2]interface{}, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, [2]interface{}{k, v})
	}
	return pairs
}

// Return a new map with the merged key-value pairs from the maps. Key conflicts are resolved by the last map.
// That mean if the same key is present in multiple maps, the value from the last map will be used.
//
// Example:
//
//	m1 := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	m2 := gblink.Map[int, string]{
//	    4: "four",
//	    5: "five",
//	    6: "six",
//	}
//	m3 := gblink.Map[int, string]{
//	    7: "seven",
//	    8: "eight",
//	    9: "nine",
//	}
//	m := m1.Merge(m2, m3)
//	fmt.Println(m) // map[1:one 2:two 3:three 4:four 5:five 6:six 7:seven 8:eight 9:nine]
func (m Map[K, V]) Merge(maps ...Map[K, V]) Map[K, V] {
	merged := Map[K, V]{}
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

// Return a new map with the key-value pairs from the map that satisfy the predicate.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	m = m.Filter(func(k int, v string) bool {
//	    return k > 1
//	})
//	fmt.Println(m) // map[2:two 3:three]
func (m Map[K, V]) Filter(predicate func(K, V) bool) Map[K, V] {
	filtered := Map[K, V]{}
	for k, v := range m {
		if predicate(k, v) {
			filtered[k] = v
		}
	}
	return filtered
}

// Delete the key k from the map.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	m.Delete(2)
//	fmt.Println(m) // map[1:one 3:three]
func (m Map[K, V]) Delete(k K) {
	delete(m, k)
}

// Delete the key-value pairs from the map that satisfy the predicate.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	m.DeleteIf(func(k int, v string) bool {
//	    return k > 1
//	})
//	fmt.Println(m) // map[1:one]
func (m Map[K, V]) DeleteIf(predicate func(K, V) bool) {
	for k, v := range m {
		if predicate(k, v) {
			delete(m, k)
		}
	}
}

// Run callback for each key-value pair in the map.
//
// Example:
//
//	m := gblink.Map[int, string]{
//	    1: "one",
//	    2: "two",
//	    3: "three",
//	}
//	m.Each(func(k int, v string) {
//	    fmt.Println(k, v)
//	}) // 1 one 2 two 3 three
func (m Map[K, V]) Each(callback func(K, V)) {
	for k, v := range m {
		callback(k, v)
	}
}
