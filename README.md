# GBlink

GBlink is a simple utility data structures and algorithms library for Golang.

Go version >= 1.18 is required.

## Installation

```bash
go get github.com/HungOnBlog/gblink
```

If you don't want to install the package, you can also use it directly in your code

## Table of Contents

- [GBlink](#gblink)
  - [Installation](#installation)
  - [Table of Contents](#table-of-contents)
  - [Data Structures](#data-structures)
    - [Map (Generic)](#map-generic)
    - [MapStringInterface](#mapstringinterface)
    - [Array (Slice, Generic)](#array-slice-generic)
    - [Bloom Filter](#bloom-filter)
    - [Stack](#stack)
    - [Queue](#queue)
    - [Linked List](#linked-list)
    - [Hash Table](#hash-table)
  - [Algorithms](#algorithms)
    - [Rate Limiter](#rate-limiter)
      - [Leaky Bucket](#leaky-bucket)
      - [Token Bucket](#token-bucket)

## Data Structures

### Map (Generic)

This type of map is suitable for flat data structures. It is not suitable for nested data structures.

```go
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
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// v, err := m.Get(2)
// if err != nil {
//     fmt.Println(err) // MapError: key 2 not found
// } else {
//     fmt.Println(v) // two
// }
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
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// m.Set(4, "four")
// fmt.Println(m) // map[1:one 2:two 3:three 4:four]
func (m Map[K, V]) Set(k K, v V) {
 m[k] = v
}

// Stringify the map.
//
// Example:
//
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// fmt.Println(m) // map[1:one 2:two 3:three]
func (m Map[K, V]) String() string {
 return fmt.Sprintf("%v", map[K]V(m))
}

// Returns the number of elements in the map.
//
// Example:
//
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// fmt.Println(m.Len()) // 3
func (m Map[K, V]) Len() int {
 return len(m)
}

// Returns true if the map is empty.
//
// Example:
//
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// fmt.Println(m.IsEmpty()) // false
func (m Map[K, V]) IsEmpty() bool {
 return len(m) == 0
}

// Returns true if the map contains the key k.
//
// Example:
//
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// fmt.Println(m.Contains(2)) // true
func (m Map[K, V]) Contains(k K) bool {
 _, ok := m[k]
 return ok
}

// Returns a slice of the keys in the map.
//
// Example:
//
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// fmt.Println(m.Keys()) // [1 2 3]
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
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// fmt.Println(m.Values()) // [one two three]
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
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// fmt.Println(m.Pairs()) // [[1 one] [2 two] [3 three]]
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
// m1 := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// m2 := gblink.Map[int, string]{
//     4: "four",
//     5: "five",
//     6: "six",
// }
// m3 := gblink.Map[int, string]{
//     7: "seven",
//     8: "eight",
//     9: "nine",
// }
// m := m1.Merge(m2, m3)
// fmt.Println(m) // map[1:one 2:two 3:three 4:four 5:five 6:six 7:seven 8:eight 9:nine]
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
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// m = m.Filter(func(k int, v string) bool {
//     return k > 1
// })
// fmt.Println(m) // map[2:two 3:three]
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
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// m.Delete(2)
// fmt.Println(m) // map[1:one 3:three]
func (m Map[K, V]) Delete(k K) {
 delete(m, k)
}

// Delete the key-value pairs from the map that satisfy the predicate.
//
// Example:
//
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// m.DeleteIf(func(k int, v string) bool {
//     return k > 1
// })
// fmt.Println(m) // map[1:one]
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
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// m.Each(func(k int, v string) {
//     fmt.Println(k, v)
// }) // 1 one 2 two 3 three
func (m Map[K, V]) Each(callback func(K, V)) {
 for k, v := range m {
  callback(k, v)
 }
}

// Clone the map.
//
// Example:
//
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// m2 := m.Clone()
// fmt.Println(m2) // map[1:one 2:two 3:three]
func (m Map[K, V]) Clone() Map[K, V] {
 clone := Map[K, V]{}
 for k, v := range m {
  clone[k] = v
 }
 return clone
}

// Clear the map.
//
// Example:
//
// m := gblink.Map[int, string]{
//     1: "one",
//     2: "two",
//     3: "three",
// }
// m.Clear()
// fmt.Println(m) // map[]
func (m Map[K, V]) Clear() {
 for k := range m {
  delete(m, k)
 }
}

```

### MapStringInterface

This type of map is more suitable to work with JSON data.

It will have all the methods of the `Map` type, plus some extra methods to work with JSON data.

```go
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
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// v, err := m.Get("two")
// if err != nil {
//     fmt.Println(err) // MapError: key two not found
// } else {
//     fmt.Println(v) // 2
// }
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
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// m.Set("four", 4)
// fmt.Println(m) // map[one:1 two:2 three:3 four:4]
func (m MapStringInterface) Set(k string, v interface{}) {
 m[k] = v
}

// Stringify the map.
//
// Example:
//
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// fmt.Println(m) // map[one:1 two:2 three:3]
func (m MapStringInterface) String() string {
 return fmt.Sprintf("%v", map[string]interface{}(m))
}

// Return the length of the map.
//
// Example:
//
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// fmt.Println(m.Len()) // 3
func (m MapStringInterface) Len() int {
 return len(m)
}

// Return true if the map is empty.
//
// Example:
//
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// fmt.Println(m.IsEmpty()) // false
func (m MapStringInterface) IsEmpty() bool {
 return len(m) == 0
}

// Return true if the map contains the key k.
//
// Example:
//
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// fmt.Println(m.Contains("two")) // true
func (m MapStringInterface) Contains(k string) bool {
 _, ok := m[k]
 return ok
}

// Return a slice of keys.
//
// Example:
//
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// fmt.Println(m.Keys()) // [one two three]
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
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// fmt.Println(m.Values()) // [1 2 3]
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
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// fmt.Println(m.Pairs()) // [[one 1] [two 2] [three 3]]
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
// m1 := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// m2 := gblink.MapStringInterface{
//     "four": 4,
//     "five": 5,
//     "six": 6,
// }
// m3 := gblink.MapStringInterface{
//     "seven": 7,
//     "eight": 8,
//     "nine": 9,
// }
// m4 := m1.Merge(m2, m3)
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
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// m2 := m.Filter(func(k string, v interface{}) bool {
//     return v.(int) > 1
// })
// fmt.Println(m2) // map[two:2 three:3]
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
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// m.Delete("two")
// fmt.Println(m) // map[one:1 three:3]
func (m MapStringInterface) Delete(k string) {
 delete(m, k)
}

// Delete a random key/value pair from the map that satisfies the callback
// function f.
//
// Example:
//
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// m.DeleteFunc(func(k string, v interface{}) bool {
//     return v.(int) > 1
// })
// fmt.Println(m) // map[one:1]
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
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// m.Each(func(k string, v interface{}) {
//     fmt.Println(k, v)
// })
func (m MapStringInterface) Each(f func(string, interface{})) {
 for k, v := range m {
  f(k, v)
 }
}

// Clone the map.
//
// Example:
//
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// m2 := m.Clone()
// fmt.Println(m2) // map[one:1 two:2 three:3]
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
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// m.Clear()
// fmt.Println(m) // map[]
func (m MapStringInterface) Clear() {
 for k := range m {
  delete(m, k)
 }
}

// Json stringifies the map.
//
// Example:
//
// m := gblink.MapStringInterface{
//     "one": 1,
//     "two": 2,
//     "three": 3,
// }
// s, err := m.JsonString()
// fmt.Println(s) // {"one":1,"two":2,"three":3}
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
//  m := gblink.MapStringInterface{
//      "a": 1,
//      "b": &gblink.MapStringInterface{
//          "c": 2,
//          "d": &gblink.MapStringInterface{
//              "e": 3,
//          },
//      },
//  }
//  v, err := m.GetDeep("a.b.c")
//  if err != nil {
//      fmt.Println(err) // MapError: Key "a.b.c" does not exist
//  }
//  v , err := m.GetDeep("b.d.e")
//  fmt.Println(v) // 3
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
// m := gblink.MapStringInterface{}
// m.SetDeep("a.b.c", 1)
// fmt.Println(m) // map[a:map[b:map[c:1]]]
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
// m := gblink.MapStringInterface{
//     "a": 1,
//     "b": &gblink.MapStringInterface{
//         "c": 2,
//         "d": &gblink.MapStringInterface{
//             "e": 3,
//         },
//     },
// }
// m.DeleteDeep("b.d.e")
// fmt.Println(m) // map[a:1 b:map[c:2 d:map[]]]
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
// m := gblink.MapStringInterface{
//     "a": 1,
//     "b": &gblink.MapStringInterface{
//         "c": 2,
//         "d": &gblink.MapStringInterface{
//             "e": 3,
//         },
//     },
// }
// fmt.Println(m.HasDeep("a")) // true
// fmt.Println(m.HasDeep("b.d.e")) // true
// fmt.Println(m.HasDeep("b.d.f")) // false
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
// m := gblink.MapStringInterface{
//     "a": 1,
//     "b": 2,
//     "c": 3,
// }
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
// m := gblink.MapStringInterface{
//     "a": 1,
//     "b": 2,
//     "c": 3,
// }
//
// mCleaned := m.CleanIf(func(key string, value interface{}) bool {
//     return value == 2
// })
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
// m := gblink.MapStringInterface{
//     "a": 1,
//     "b": &gblink.MapStringInterface{
//         "c": 2,
//         "d": &gblink.MapStringInterface{
//             "e": 3,
//         },
//     },
// }
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
// m := gblink.MapStringInterface{
//     "a": 1,
//     "b": &gblink.MapStringInterface{
//         "c": 2,
//         "d": &gblink.MapStringInterface{
//             "e": 3,
//         },
//     },
// }
//
// mCleaned := m.CleanDeepIf(func(key string, value interface{}) bool {
//     return value == 2
// })
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
// m1 := gblink.MapStringInterface{
//     "a": 1,
//     "b": &gblink.MapStringInterface{
//         "c": 2,
//         "d": &gblink.MapStringInterface{
//             "e": 3,
//         },
//     },
// }
//
// m2 := gblink.MapStringInterface{
//     "a": 4,
//     "b": &gblink.MapStringInterface{
//         "d": &gblink.MapStringInterface{
//             "f": 5,
//         },
//     },
// }
//
// m3 := m1.MergeDeep(m2)
// fmt.Println(m3) // map[a:4 b:map[c:2 d:map[e:3 f:5]]]
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

```

### Array (Slice, Generic)

This data structure is best suited for storing a list of items, especially primitive types like `int`, `string`, `float64`, etc.

```go
package gblink

import (
 "fmt"

 "golang.org/x/exp/constraints"
)

type ArrayError struct {
 error
}

type Array[T constraints.Ordered] []T

// Len returns the length of the array.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.Len()) // 3
func (a *Array[T]) Len() int {
 return len(*a)
}

// Append appends the given value to the array.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array) // [1 2 3]
func (a *Array[T]) Append(value T) {
 *a = append(*a, value)
}

// AppendAll appends the given values to the array.
//
// Example:
//
// var array Array[int]
// array.AppendAll(1, 2, 3)
// fmt.Println(array) // [1 2 3]
func (a *Array[T]) AppendAll(values ...T) {
 *a = append(*a, values...)
}

// AppendArray appends the given array to the array.
//
// Example:
//
// var array1 Array[int]
// array1.AppendArray(&Array[int]{1, 2, 3})
// fmt.Println(array1) // [1 2 3]
func (a *Array[T]) AppendArray(array *Array[T]) {
 *a = append(*a, *array...)
}

// AppendArrayAll appends the given arrays to the array.
//
// Example:
//
// var array1 Array[int]
// array1.AppendArrayAll(&Array[int]{1, 2, 3}, &Array[int]{4, 5, 6})
// fmt.Println(array1) // [1 2 3 4 5 6]
func (a *Array[T]) AppendArrayAll(arrays ...*Array[T]) {
 for _, array := range arrays {
  *a = append(*a, *array...)
 }
}

// Clear clears the array.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
func (a *Array[T]) Clear() {
 *a = (*a)[:0]
}

// Contains returns true if the array contains the given value.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.Contains(2)) // true
func (a *Array[T]) Contains(value T) bool {
 for _, v := range *a {
  if v == value {
   return true
  }
 }
 return false
}

// ContainsAll returns true if the array contains all the given values.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.ContainsAll(1, 2)) // true
func (a *Array[T]) ContainsAll(values ...T) bool {
 for _, value := range values {
  if !a.Contains(value) {
   return false
  }
 }
 return true
}

// Some returns true if the array contains at least one value for which the given predicate returns true.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.Some(func(value int) bool {
//  return value > 2
// })) // true
func (a *Array[T]) Some(predicate func(T) bool) bool {
 for _, v := range *a {
  if predicate(v) {
   return true
  }
 }
 return false
}

// Every returns true if the array contains all values for which the given predicate returns true.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.Every(func(value int) bool {
//  return value > 0
// })) // true
func (a *Array[T]) Every(predicate func(T) bool) bool {
 for _, v := range *a {
  if !predicate(v) {
   return false
  }
 }
 return true
}

// Filter returns a new array containing all values for which the given predicate returns true.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.Filter(func(value int) bool {
//  return value > 1
// })) // [2 3]
func (a *Array[T]) Filter(predicate func(T) bool) *Array[T] {
 var array Array[T]
 for _, v := range *a {
  if predicate(v) {
   array.Append(v)
  }
 }
 return &array
}

// Find returns the first value for which the given predicate returns true.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.Find(func(value int) bool {
//  return value > 1
// })) // 2
func (a *Array[T]) Find(predicate func(T) bool) T {
 for _, v := range *a {
  if predicate(v) {
   return v
  }
 }
 var zero T
 return zero
}

// FindIndex returns the index of the first value for which the given predicate returns true.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.FindIndex(func(value int) bool {
//  return value > 1
// })) // 1
func (a *Array[T]) FindIndex(predicate func(T) bool) int {
 for i, v := range *a {
  if predicate(v) {
   return i
  }
 }
 return -1
}

// FindLast returns the last value for which the given predicate returns true.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.FindLast(func(value int) bool {
//  return value > 1
// })) // 3
func (a *Array[T]) FindLast(predicate func(T) bool) T {
 for i := len(*a) - 1; i >= 0; i-- {
  if predicate((*a)[i]) {
   return (*a)[i]
  }
 }
 var zero T
 return zero
}

// FindLastIndex returns the index of the last value for which the given predicate returns true.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.FindLastIndex(func(value int) bool {
//  return value > 1
// })) // 2
func (a *Array[T]) FindLastIndex(predicate func(T) bool) int {
 for i := len(*a) - 1; i >= 0; i-- {
  if predicate((*a)[i]) {
   return i
  }
 }
 return -1
}

// Each calls the given function for each (index, value) pair in the array.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// array.Each(func(index int, value int) {
//  fmt.Println(index, value)
// })
func (a *Array[T]) Each(f func(int, T)) {
 for i, v := range *a {
  f(i, v)
 }
}

// EachIndex calls the given function for each index in the array.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// array.EachIndex(func(index int) {
//  fmt.Println(index)
// })
func (a *Array[T]) EachIndex(f func(int)) {
 for i := range *a {
  f(i)
 }
}

// EachValue calls the given function for each value in the array.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// array.EachValue(func(value int) {
//  fmt.Println(value)
// })
func (a *Array[T]) EachValue(f func(T)) {
 for _, v := range *a {
  f(v)
 }
}

// Map returns a new array containing the values returned by the given function for each value in the array.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.Map(func(value int) int {
//  return value + 1
// })) // [2 3 4]
func (a *Array[T]) Map(f func(T) T) *Array[T] {
 var array Array[T]
 for _, v := range *a {
  array.Append(f(v))
 }
 return &array
}

// SortBy sorts the array using the sort function (which should be a function that takes two
// values and returns -1, 0, or 1).
// -1 means the first value is less than the second value
// 0 means the first value is equal to the second value
// 1 means the first value is greater than the second value
//
// Example:
//
// var array Array[int]
// array.Append(3)
// array.Append(1)
// array.Append(2)
// array.SortBy(func(a int, b int) int {
//  if a < b {
//   return -1
//  } else if a > b {
//   return 1
//  } else {
//   return 0
//  }
// })
// fmt.Println(array) // [1 2 3]
func (a *Array[T]) SortBy(compare func(T, T) int) {
 // Using the compare function with QuickSort
 quickSort(a, 0, len(*a)-1, compare)
}

// quickSort is a helper function for SortBy
func quickSort[T constraints.Ordered](a *Array[T], lo int, hi int, compare func(T, T) int) {
 if lo < hi {
  p := partition(a, lo, hi, compare)
  quickSort(a, lo, p-1, compare)
  quickSort(a, p+1, hi, compare)
 }
}

// partition is a helper function for quickSort
func partition[T constraints.Ordered](a *Array[T], lo int, hi int, compare func(T, T) int) int {
 pivot := (*a)[hi]
 i := lo
 for j := lo; j < hi; j++ {
  if compare((*a)[j], pivot) < 0 {
   (*a)[i], (*a)[j] = (*a)[j], (*a)[i]
   i++
  }
 }
 (*a)[i], (*a)[hi] = (*a)[hi], (*a)[i]
 return i
}

func compareAsc[T constraints.Ordered](a T, b T) int {
 if a < b {
  return -1
 } else if a > b {
  return 1
 } else {
  return 0
 }
}

func compareDes[T constraints.Ordered](a T, b T) int {
 if a > b {
  return -1
 } else if a < b {
  return 1
 } else {
  return 0
 }
}

// Sort sorts the array (ascending or descending), default is ascending.
//
// Example:
//
// var array Array[int]
// array.Append(3)
// array.Append(1)
// array.Append(2)
// array.Sort()
// fmt.Println(array) // [1 2 3]
func (a *Array[T]) Sort(isAscending bool) {
 // Using QuickSort
 if isAscending {
  quickSort(a, 0, len(*a)-1, compareAsc[T])
 } else {
  quickSort(a, 0, len(*a)-1, compareDes[T])
 }
}

// Reverse reverses the array.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// array.Reverse()
// fmt.Println(array) // [3 2 1]
func (a *Array[T]) Reverse() {
 for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
  (*a)[i], (*a)[j] = (*a)[j], (*a)[i]
 }
}

// Copy returns a copy of the array.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// arrayCopy := array.Copy()
// fmt.Println(arrayCopy) // [1 2 3]
func (a *Array[T]) Copy() *Array[T] {
 var array Array[T]
 for _, v := range *a {
  array.Append(v)
 }
 return &array
}

// IndexOf returns the index of the first occurrence of the given value in the array.
// If the value is not found, -1 is returned.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// fmt.Println(array.IndexOf(2)) // 1
// fmt.Println(array.IndexOf(4)) // -1
func (a *Array[T]) IndexOf(value T) int {
 for i, v := range *a {
  if v == value {
   return i
  }
 }
 return -1
}

// Insert inserts the given value at the given index.
// If the index is out of bounds, an error is returned.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// array.Insert(1, 4)
// fmt.Println(array) // [1 4 2 3]
func (a *Array[T]) Insert(index int, value T) error {
 if index < 0 || index > len(*a) {
  return ArrayError{fmt.Errorf("ArrayError: %d index out of range", index)}
 }
 *a = append(*a, value)
 copy((*a)[index+1:], (*a)[index:])
 (*a)[index] = value
 return nil
}

// Remove removes the value at the given index.
// If the index is out of bounds, an error is returned.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// array.Remove(1)
// fmt.Println(array) // [1 3]
func (a *Array[T]) Remove(index int) error {
 if index < 0 || index >= len(*a) {
  return ArrayError{fmt.Errorf("ArrayError: %d index out of range", index)}
 }
 *a = append((*a)[:index], (*a)[index+1:]...)
 return nil
}

// RemoveRange removes the values in the range of start and end.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// array.RemoveRange(1, 2)
// fmt.Println(array) // [1 3]
func (a *Array[T]) RemoveRange(start int, end int) {
 *a = append((*a)[:start], (*a)[end:]...)
}

// Slice returns a slice of the array.
//
// Example:
//
// var array Array[int]
// array.Append(1)
// array.Append(2)
// array.Append(3)
// slice := array.Slice(1, 2)
// fmt.Println(slice) // [2]
func (a *Array[T]) Slice(start int, end int) Array[T] {
 var array Array[T]
 for i := start; i < end; i++ {
  array.Append((*a)[i])
 }
 return array
}

// InsertArray inserts the given array at the given index.
// If the index is out of bounds, an error is returned.
//
// Example:
//
// var array1 Array[int]
// array1.Append(1)
// array1.Append(2)
// array1.Append(3)
// var array2 Array[int]
// array2.Append(4)
// array2.Append(5)
// array2.Append(6)
// array1.InsertArray(1, array2)
// fmt.Println(array1) // [1 4 5 6 2 3]
func (a *Array[T]) InsertArray(index int, array Array[T]) error {
 if index < 0 || index > len(*a) {
  return ArrayError{fmt.Errorf("ArrayError: %d index out of range", index)}
 }
 *a = append(*a, array...)
 copy((*a)[index+len(array):], (*a)[index:])
 copy((*a)[index:], array)
 return nil
}

```

### Bloom Filter

Bloom Filter is a space-efficient probabilistic data structure, that is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not. Elements can be added to the set, but not removed (though this can be addressed with a counting filter). The more elements that are added to the set, the larger the probability of false positives.

Bloom Filter is implemented using a bit array and a set of hash functions. The bit array is initialized to all 0s. When an element is added to the set, the k hash functions are applied to the element, and the resulting k array indices are set to 1. To query if an element is in the set, the k hash functions are applied to the element, and the resulting k array indices are checked. If all k indices contain a 1, then the element is probably in the set. If any of the k indices contain a 0, then the element is definitely not in the set.

You can read more about Bloom Filter [here](https://en.wikipedia.org/wiki/Bloom_filter).

```go
package gblink

import (
 "fmt"
 "hash/fnv"
 "math"
)

// BloomFilter is a probabilistic data structure that can be used to test if an item is in a set.
// It is a space-efficient implementation of a set that returns false positives but never false negatives.
// The probability of a false positive can be controlled by the size of the bitset and the number of hash functions.
//
// The number of hash functions should be set to k = m/n * ln(2), where m is the size of the bitset and n is the number of items in the set.
//
// The size of the bitset should be set to m = -n * ln(p) / (ln(2))^2, where n is the number of items in the set and p is the desired probability of a false positive.
//
// More: https://en.wikipedia.org/wiki/Bloom_filter
type BloomFilter struct {
 bitset []bool // the bitset used to store the filter
 k      uint   // the number of hash functions used
}

// NewBloomFilter creates a new Bloom filter with the specified bitset size and number of hash functions.
func NewBloomFilter(m uint, k uint) *BloomFilter {
 return &BloomFilter{
  bitset: make([]bool, m),
  k:      k,
 }
}

// Add adds an item to the Bloom filter by setting the corresponding bits in the bitset.
func (bf *BloomFilter) Add(item string) {
 for i := uint(0); i < bf.k; i++ {
  hash := bf.hash(item, i)
  bf.bitset[hash] = true
 }
}

// Contains checks if an item is in the Bloom filter by checking if all the corresponding bits in the bitset are set.
func (bf *BloomFilter) Contains(item string) bool {
 for i := uint(0); i < bf.k; i++ {
  hash := bf.hash(item, i)
  if !bf.bitset[hash] {
   return false
  }
 }
 return true
}

// hash computes the hash value for an item using the FNV-1a hash function and the specified seed value.
func (bf *BloomFilter) hash(item string, seed uint) uint {
 hash := fnv.New32a()                             // create a new 32-bit FNV-1a hash object
 hash.Write([]byte(item))                         // write the item to the hash object
 hash.Write([]byte{byte(seed)})                   // write the seed value to the hash object
 return uint(hash.Sum32()) % uint(len(bf.bitset)) // compute the hash value and return it
}

// CalculateBloomFilterBitSetSize calculates the size of the bitset for a Bloom filter with the specified number of items and false positive rate.
func CalculateBloomFilterBitSetSize(numItems uint, falsePositiveRate float64) uint {
 return uint(float64(numItems) * math.Log(falsePositiveRate) / math.Pow(math.Log(2), 2))
}

// CalculateBloomFilterNumHashFunctions calculates the number of hash functions for a Bloom filter with the specified bitset size and number of items.
func CalculateBloomFilterNumHashFunctions(bitSetSize uint, numItems uint) uint {
 return uint(float64(bitSetSize) / float64(numItems) * math.Log(2))
}

// ExampleBloomFilter shows how to use a Bloom filter.
func ExampleBloomFilter() {
 // create Bloom filter
 bf := NewBloomFilter(100, 4)

 // add items
 bf.Add("foo")
 bf.Add("bar")
 bf.Add("baz")

 // check if items are in Bloom filter
 fmt.Println(bf.Contains("foo"))
 fmt.Println(bf.Contains("bar"))
 fmt.Println(bf.Contains("baz"))
 fmt.Println(bf.Contains("qux"))

 // Output:
 // true
 // true
 // true
 // false
}

```

### Stack

Stack is a data structure that is used to store a list of items. Items can be added to the top of the stack, and items can be removed from the top of the stack. The last item added to the stack is the first item that can be removed.

Stack is implemented using a slice. When an item is added to the stack, it is appended to the slice. When an item is removed from the stack, the last item in the slice is removed and returned.

The Stack rule of thumb is that the last item added to the stack is the first item that can be removed (FILO - First In Last Out)

```go
package gblink

import "errors"

type Stack[T any] []T

type StackError struct {
 error
}

// NewStack returns a new Stack.
func NewStack[T any]() *Stack[T] {
 return &Stack[T]{}
}

// Push pushes the specified value onto the stack.
//
// Push an item onto the stack.
//
// Example:
//
// s := NewStack()
// s.Push(1)
// s.Push(2)
// s.Push(3)
// fmt.Println(s) // [1 2 3]
func (s *Stack[T]) Push(v T) {
 *s = append(*s, v)
}

// Pop removes and returns the top item from the stack.
//
// Pop an item from the stack.
//
// Example:
//
// s := NewStack()
// s.Push(1)
// s.Push(2)
// s.Push(3)
// fmt.Println(s.Pop()) // 3
// fmt.Println(s.Pop()) // 2
// fmt.Println(s.Pop()) // 1
func (s *Stack[T]) Pop() (T, error) {
 if len(*s) == 0 {
  var zero T
  return zero, &StackError{errors.New("StackError: stack is empty")}
 }
 v := (*s)[len(*s)-1]
 *s = (*s)[:len(*s)-1]
 return v, nil
}

// Peek returns the top item from the stack without removing it.
//
// Peek at the top item on the stack.
//
// Example:
//
// s := NewStack()
// s.Push(1)
// s.Push(2)
// s.Push(3)
// fmt.Println(s.Peek()) // 3
func (s *Stack[T]) Peek() (T, error) {
 if len(*s) == 0 {
  var zero T
  return zero, &StackError{errors.New("StackError: stack is empty")}
 }
 return (*s)[len(*s)-1], nil
}

// Len returns the number of items in the stack.
//
// Get the number of items in the stack.
//
// Example:
//
// s := NewStack()
// s.Push(1)
// s.Push(2)
// s.Push(3)
// fmt.Println(s.Len()) // 3
func (s *Stack[T]) Len() int {
 return len(*s)
}

// IsEmpty returns true if the stack is empty.
//
// Check if the stack is empty.
//
// Example:
//
//  s := NewStack()
//  fmt.Println(s.IsEmpty()) // true
//  s.Push(1)
//  fmt.Println(s.IsEmpty()) // false
//  s.Pop()
//  fmt.Println(s.IsEmpty()) // true
func (s *Stack[T]) IsEmpty() bool {
 return len(*s) == 0
}

```

### Queue

Queue is a data structure that is used to store a list of items. Items can be added to the queue, and items can be removed from the queue. The first item added to the queue is the first item that can be removed.

Queue is implemented using a slice. When an item is added to the queue, it is appended to the slice. When an item is removed from the queue, the first item in the slice is removed and returned.

The Queue rule of thumb is that the first item added to the queue is the first item that can be removed (FIFO - First In First Out)

More about queues: [https://en.wikipedia.org/wiki/Queue_(abstract_data_type)](https://en.wikipedia.org/wiki/Queue_(abstract_data_type))

```go
package gblink

import "errors"

type Queue[T any] []T

type QueueError struct {
 error
}

// NewQueue returns a new Queue.
func NewQueue[T any]() *Queue[T] {
 return &Queue[T]{}
}

// Push pushes the specified value onto the queue.
//
// Push an item onto the queue.
//
// Example:
//
// q := NewQueue()
// q.Push(1)
// q.Push(2)
// q.Push(3)
// fmt.Println(q) // [1 2 3]
func (q *Queue[T]) Push(v T) {
 *q = append(*q, v)
}

// Pop removes and returns the first item from the queue.
//
// Pop an item from the queue.
//
// Example:
//
// q := NewQueue()
// q.Push(1)
// q.Push(2)
// q.Push(3)
// fmt.Println(q.Pop()) // 1
// fmt.Println(q.Pop()) // 2
// fmt.Println(q.Pop()) // 3
func (q *Queue[T]) Pop() (T, error) {
 if len(*q) == 0 {
  var zero T
  return zero, &QueueError{errors.New("QueueError: queue is empty")}
 }
 v := (*q)[0]
 *q = (*q)[1:]
 return v, nil
}

// Peek returns the first item from the queue without removing it.
//
// Peek at the first item on the queue.
//
// Example:
//
// q := NewQueue()
// q.Push(1)
// q.Push(2)
// q.Push(3)
// fmt.Println(q.Peek()) // 1
// fmt.Println(q.Peek()) // 1
// fmt.Println(q.Peek()) // 1
func (q *Queue[T]) Peek() (T, error) {
 if len(*q) == 0 {
  var zero T
  return zero, &QueueError{errors.New("QueueError: queue is empty")}
 }
 return (*q)[0], nil
}

// Len returns the number of items in the queue.
//
// Get the number of items in the queue.
func (q *Queue[T]) Len() int {
 return len(*q)
}

// IsEmpty returns true if the queue is empty.
//
// Check if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
 return len(*q) == 0
}

```

### Linked List

Linked List is a data structure that is used to store a list of items. Items can be added to the list, and items can be removed from the list. The first item added to the list is the first item that can be removed.

Linked List is implemented using a slice of pointers to nodes. Each node contains a value and a pointer to the next node. When an item is added to the list, a new node is created and added to the slice. When an item is removed from the list, the first node in the slice is removed and returned.

More about Linked List: [https://en.wikipedia.org/wiki/Linked_list](https://en.wikipedia.org/wiki/Linked_list)

```go
package gblink

import "fmt"

type LikedListNode[T comparable] struct {
 Value T
 Next  *LikedListNode[T]
}

type LikedList[T comparable] struct {
 Head *LikedListNode[T]
 Tail *LikedListNode[T]
}

type LikedListError struct {
 error
}

func NewLikedList[T comparable]() *LikedList[T] {
 return &LikedList[T]{}
}

// Len returns the number of elements in the list.
//
// The complexity is O(n).
//
// Example:
//
// list := NewLikedList[int]()
// list.Append(1)
// list.Append(2)
// list.Append(3)
// list.Len() // 3
func (l *LikedList[T]) Len() int {
 count := 0
 for node := l.Head; node != nil; node = node.Next {
  count++
 }
 return count
}

// Append adds a new element with the given value to the end of the list.
//
// The complexity is O(1).
//
// Example:
//
// list := NewLikedList[int]()
// list.Append(1)
// list.Append(2)
// list.Append(3)
// list.Len() // 3
func (l *LikedList[T]) Append(value T) {
 node := &LikedListNode[T]{Value: value}
 if l.Head == nil {
  l.Head = node
  l.Tail = node
  return
 }
 l.Tail.Next = node
 l.Tail = node
}

// Prepend adds a new element with the given value to the beginning of the list.
//
// The complexity is O(1).
//
// Example:
//
// list := NewLikedList[int]()
// list.Prepend(1)
// list.Prepend(2)
// list.Prepend(3)
// list.Len() // 3
// fmt.Println(list.Head.Value) // 3
func (l *LikedList[T]) Prepend(value T) {
 node := &LikedListNode[T]{Value: value}
 if l.Head == nil {
  l.Head = node
  l.Tail = node
  return
 }
 node.Next = l.Head
 l.Head = node
}

// Insert adds a new element with the given value after the n-th element of the list.
//
// The complexity is O(n).
//
// Example:
//
// list := NewLikedList[int]()
// list.Append(1)
// list.Append(2)
// list.Append(3)
// list.Insert(2, 4)
// list.Len() // 4
// fmt.Println(list.Head.Next.Next.Value) // 4
func (l *LikedList[T]) Insert(n int, value T) error {
 if n < 0 {
  return &LikedListError{fmt.Errorf("LikedListError: index out of range")}
 }
 if n == 0 {
  l.Prepend(value)
  return nil
 }
 if n == l.Len() {
  l.Append(value)
  return nil
 }
 node := l.Head
 for i := 0; i < n-1; i++ {
  node = node.Next
 }
 newNode := &LikedListNode[T]{Value: value}
 newNode.Next = node.Next
 node.Next = newNode
 return nil
}

// Remove removes the n-th element of the list and returns its value.
//
// The complexity is O(n).
//
// Example:
//
// list := NewLikedList[int]()
// list.Append(1)
// list.Append(2)
// list.Append(3)
// list.Remove(1) // 2
// list.Len() // 2
func (l *LikedList[T]) Remove(n int) (T, error) {
 if n < 0 || n >= l.Len() {
  var zero T
  return zero, &LikedListError{fmt.Errorf("LikedListError: index out of range")}
 }
 if n == 0 {
  value := l.Head.Value
  l.Head = l.Head.Next
  return value, nil
 }
 node := l.Head
 for i := 0; i < n-1; i++ {
  node = node.Next
 }
 value := node.Next.Value
 node.Next = node.Next.Next
 return value, nil
}

// Get returns the value of the n-th element of the list.
//
// The complexity is O(n).
//
// Example:
//
// list := NewLikedList[int]()
// list.Append(1)
// list.Append(2)
// list.Append(3)
// list.Get(1) // 2
func (l *LikedList[T]) Get(n int) (T, error) {
 if n < 0 || n >= l.Len() {
  var zero T
  return zero, &LikedListError{fmt.Errorf("LikedListError: index out of range")}
 }
 node := l.Head
 for i := 0; i < n; i++ {
  node = node.Next
 }
 return node.Value, nil
}

// IndexOf returns the index of the first element with the given value.
//
// The complexity is O(n).
//
// Example:
//
// list := NewLikedList[int]()
// list.Append(1)
// list.Append(2)
// list.Append(3)
// list.IndexOf(2) // 1
func (l *LikedList[T]) IndexOf(value T) int {
 node := l.Head
 for i := 0; i < l.Len(); i++ {
  if node.Value == value {
   return i
  }
  node = node.Next
 }
 return -1
}

// Contains returns true if the list contains an element with the given value.
//
// The complexity is O(n).
//
// Example:
//
// list := NewLikedList[int]()
// list.Append(1)
// list.Append(2)
// list.Append(3)
// list.Contains(2) // true
func (l *LikedList[T]) Contains(value T) bool {
 return l.IndexOf(value) != -1
}

// Clear removes all elements from the list.
//
// The complexity is O(1).
//
// Example:
//
// list := NewLikedList[int]()
// list.Append(1)
// list.Append(2)
// list.Append(3)
// list.Clear()
// list.Len() // 0
func (l *LikedList[T]) Clear() {
 l.Head = nil
 l.Tail = nil
}

```

### Hash Table

A hash table is a data structure that maps keys to values. It is used to implement associative arrays, a data structure that can map keys to values. A hash table uses a hash function to compute an index into an array of buckets or slots, from which the desired value can be found.

Hash table usually used when you need to store a large amount of data and you need to find a specific element in the data quickly.

!NOTE: to use this HashTable you need LikedList code too.

More about hash table: [https://en.wikipedia.org/wiki/Hash_table](https://en.wikipedia.org/wiki/Hash_table)

```go
package gblink

import (
 "fmt"
 "hash"
)

type HashTableError struct {
 error
}

// HashTable is a hash table implementation.
//
// The zero value for HashTable is an empty hash table ready to use.
//
// The HashTable type is not safe for concurrent use by multiple goroutines without.
type HashTable[K comparable, V comparable] struct {
 Hasher hash.Hash64
 Table  map[uint64]*LikedList[V]
}

// NewHashTable returns a new HashTable.
func NewHashTable[K comparable, V comparable](hasher hash.Hash64) *HashTable[K, V] {
 return &HashTable[K, V]{
  Table:  make(map[uint64]*LikedList[V]),
  Hasher: hasher,
 }
}

// Set sets the value for the given key.
//
// The complexity is O(1).
//
// Example:
//
// table := NewHashTable[int, string]()
// table.Set(1, "one")
// table.Set(2, "two")
// table.Set(3, "three")
// table.Set(4, "four")
// table.Set(5, "five")
// table.Set(6, "six")
func (t *HashTable[K, V]) Set(key K, value V) {
 // Hash the key.
 t.Hasher.Reset()
 t.Hasher.Write([]byte(fmt.Sprintf("%v", key)))
 hash := t.Hasher.Sum64()
 if _, ok := t.Table[hash]; !ok {
  t.Table[hash] = NewLikedList[V]()
 }
 t.Table[hash].Append(value)
}

// Get returns the value for the given key.
//
// The complexity is O(1).
//
// Example:
//
//  table := NewHashTable[int, string]()
//  table.Set(1, "one")
//  table.Set(2, "two")
//  table.Set(3, "three")
//  v , err := table.Get(2)
//  if err != nil {
//      panic(err)
//  }
//  fmt.Println(v) // two
func (t *HashTable[K, V]) Get(key K) (V, error) {
 // Hash the key.
 t.Hasher.Reset()
 t.Hasher.Write([]byte(fmt.Sprintf("%v", key)))
 hash := t.Hasher.Sum64()
 if _, ok := t.Table[hash]; !ok {
  var zero V
  return zero, &HashTableError{error: fmt.Errorf("HashTableError: key not found")}
 }
 return t.Table[hash].Head.Value, nil
}

// Len returns the number of elements in the hash table.
//
// The complexity is O(n).
//
// Example:
//
//  table := NewHashTable[int, string]()
//  table.Set(1, "one")
//  table.Set(2, "two")
//  table.Set(3, "three")
//  table.Set(4, "four")
//  table.Set(5, "five")
//     fmt.Println(table.Len()) // 5
func (t *HashTable[K, V]) Len() int {
 count := 0
 for _, list := range t.Table {
  count += list.Len()
 }
 return count
}

// Clear removes all elements from the hash table.
//
// The complexity is O(n).
//
// Example:
//
// table := NewHashTable[int, string]()
// table.Set(1, "one")
// table.Set(2, "two")
// table.Set(3, "three")
// table.Clear()
// fmt.Println(table.Len()) // 0
func (t *HashTable[K, V]) Clear() {
 t.Table = make(map[uint64]*LikedList[V])
}

// Delete removes the element with the given key from the hash table.
//
// NOTE!: This is dangerous, because it will remove all elements with the same hash.
//
// The complexity is O(1).
//
// Example:
//
// table := NewHashTable[int, string]()
// table.Set(1, "one")
// table.Set(2, "two")
// table.Set(3, "three")
// table.Delete(2)
// fmt.Println(table.Len()) // 2
func (t *HashTable[K, V]) Delete(key K) {
 // Hash the key.
 t.Hasher.Reset()
 t.Hasher.Write([]byte(fmt.Sprintf("%v", key)))
 hash := t.Hasher.Sum64()
 delete(t.Table, hash)
}

```

## Algorithms

### Rate Limiter

#### Leaky Bucket

Leaky Bucket is a token bucket algorithm that is used to limit the rate of data transfer. It is a simple algorithm that allows bursts of data, but limits the average rate of data transfer.

Leaky Bucket is implemented using a channel that is used to store water. The channel is initialized to a fixed size. When data is transferred, a water is added from the channel. If the channel is full, then no data is transferred.
Water is also leaked from the channel at a fixed rate. The leak rate is determined by the size of the channel and the time interval between leaks.

You can read more about Leaky Bucket [here](https://en.wikipedia.org/wiki/Leaky_bucket).

```go
package gblink

import (
 "fmt"
 "time"
)

// LeakyBucket simulates a bucket with a hole that leaks water at a fixed rate.
type LeakyBucket struct {
 flowRate       float64       // The rate at which water flows into the bucket.
 bucketCapacity float64       // The maximum amount of water that the bucket can hold.
 waterLevel     float64       // The current amount of water in the bucket.
 lastLeak       time.Time     // The time when the bucket was last leaked.
 flowTicker     *time.Ticker  // The ticker that controls the flow of water into the bucket.
 stopChan       chan struct{} // The channel used to stop the flow of water into the bucket.
}

// NewLeakyBucket creates a new leaky bucket with the specified flow rate and bucket capacity.
func NewLeakyBucket(flowRate float64, bucketCapacity float64) *LeakyBucket {
 return &LeakyBucket{
  flowRate:       flowRate,
  bucketCapacity: bucketCapacity,
  waterLevel:     0,
  lastLeak:       time.Now(),
  flowTicker:     time.NewTicker(time.Second), // By default, the bucket leaks water once per second.
  stopChan:       make(chan struct{}),
 }
}

// AddWater adds a specified volume of water to the bucket.
func (lb *LeakyBucket) AddWater(volume float64) bool {
 // Calculate the time since the bucket was last leaked.
 elapsed := time.Since(lb.lastLeak)

 // Calculate the amount of water that should have leaked from the bucket during this time.
 leaked := elapsed.Seconds() * lb.flowRate

 // Update the current water level by subtracting the leaked water.
 lb.waterLevel -= leaked

 // Ensure that the water level does not exceed the bucket capacity.
 if lb.waterLevel+volume > lb.bucketCapacity {
  return false // The bucket is full.
 }

 // Add the new water volume to the water level.
 lb.waterLevel += volume

 // Update the last leak time.
 lb.lastLeak = time.Now()

 return true // The water has been added to the bucket.
}

// Start starts the flow of water into the bucket.
func (lb *LeakyBucket) Start() {
 go func() {
  for {
   select {
   case <-lb.stopChan:
    lb.flowTicker.Stop() // Stop the ticker when the flow is stopped.
    return
   case <-lb.flowTicker.C:
    // Add the flow rate amount of water to the bucket.
    lb.AddWater(lb.flowRate)
   }
  }
 }()
}

// Stop stops the flow of water into the bucket.
func (lb *LeakyBucket) Stop() {
 lb.stopChan <- struct{}{}
}

// Example of usage:
// Using the leaky bucket to limit the number of requests per second.
// We want to limit the number of requests to 100 requests per second.
func ExampleLeakyBucket() {
 // Create a new leaky bucket with a flow rate of 100 requests per second and a capacity of 100 requests.
 bucket := NewLeakyBucket(100, 100)

 // Start the flow of requests into the bucket.
 bucket.Start()

 // Simulate incoming requests.
 for i := 1; i <= 200; i++ {
  if bucket.AddWater(1) {
   fmt.Printf("Request %d allowed at %s\n", i, time.Now().Format(time.StampMilli))
  } else {
   fmt.Printf("Request %d blocked at %s\n", i, time.Now().Format(time.StampMilli))
  }
  time.Sleep(10 * time.Millisecond) // Wait for 10 milliseconds between each request.
 }
}

```

#### Token Bucket

Token Bucket is a token bucket algorithm that is used to limit the rate of data transfer. It is a simple algorithm that allows bursts of data, but limits the average rate of data transfer.

Token Bucket is implemented using a channel that is used to store tokens. The channel is initialized to a fixed size. When a request invokes the algorithm, a token is removed from the channel. If the channel is empty, then no data is transferred. When a token is removed from the channel, a new token is added to the channel after a fixed time interval. The size of the channel determines the maximum burst size.

You can read more about Token Bucket [here](https://en.wikipedia.org/wiki/Token_bucket).

```go
package gblink

import (
 "fmt"
 "sync"
 "time"
)

type TokenBucket struct {
 tokens        uint64        // Current number of tokens in the bucket.
 capacity      uint64        // Maximum number of tokens that the bucket can hold.
 rate          time.Duration // Rate at which tokens are added to the bucket.
 mu            sync.Mutex    // Mutex to synchronize access to the bucket.
 lastTokenTime time.Time     // Last time a token was added to the bucket.
}

// NewTokenBucket creates a new Token Bucket with the specified capacity and refill rate.
func NewTokenBucket(capacity uint64, rate time.Duration) *TokenBucket {
 return &TokenBucket{
  tokens:        capacity,
  capacity:      capacity,
  rate:          rate,
  lastTokenTime: time.Now(),
 }
}

// TakeToken attempts to take a token from the bucket.
func (tb *TokenBucket) TakeToken() bool {
 tb.mu.Lock()
 defer tb.mu.Unlock()

 // Calculate the number of tokens that should have been added since the last token was added.
 elapsedTime := time.Since(tb.lastTokenTime)
 numTokensToAdd := uint64(elapsedTime.Nanoseconds() / tb.rate.Nanoseconds())

 // Add the calculated tokens to the bucket, up to the capacity of the bucket.
 tb.tokens += numTokensToAdd
 if tb.tokens > tb.capacity {
  tb.tokens = tb.capacity
 }

 // Attempt to take a token from the bucket.
 if tb.tokens > 0 {
  tb.tokens--
  tb.lastTokenTime = time.Now()
  return true
 }
 return false
}

// Example of a token bucket.
// Limit the rate of incoming requests to 100 requests per second.
func ExampleTokenBucket() {
 // Create a new token bucket with a capacity of 100 tokens and a fill rate of 100 tokens per second.
 tb := NewTokenBucket(100, 100)

 // Simulate 1000 incoming requests over the course of 11 seconds.
 start := time.Now()
 for i := 0; i < 1000; i++ {
  // Wait for the token bucket to allow the request to proceed.
  for !tb.TakeToken() {
   time.Sleep(time.Millisecond * 10)
  }

  // Process the request.
  processRequest()

  // Throttle the request rate to exactly 100 requests per second.
  if i%10 == 9 {
   time.Sleep(time.Second / 100)
  }
 }
 elapsed := time.Since(start)

 // Print some statistics about the simulation.
 fmt.Printf("Processed %d requests in %s (%.2f requests per second)\n", 1000, elapsed, float64(1000)/elapsed.Seconds())
}

func processRequest() {
 // Simulate some work.
 time.Sleep(time.Millisecond * 50)
}

```
