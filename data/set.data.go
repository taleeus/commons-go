package data

import (
	"encoding/json"
	"iter"
	"maps"
	"slices"
)

// Set is a collection of unique elements
type Set[T comparable] map[T]struct{}

// NewSet creates a new set with the given values
func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T], len(values))
	for _, val := range values {
		set.Push(val)
	}

	return set
}

// Push adds a new value to the set
func (set Set[T]) Push(val T) {
	set[val] = struct{}{}
}

// Delete removes a value from the set
func (set Set[T]) Delete(val T) {
	delete(set, val)
}

// Contains checks if the set contains the given value
func (set Set[T]) Contains(val T) bool {
	_, ok := set[val]
	return ok
}

// Values returns the set iterator
func (set Set[T]) Values() iter.Seq[T] {
	return maps.Keys(set)
}

// MarshalJSON returns the JSON encoding of the set
func (set *Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(slices.Collect(set.Values()))
}

// UnmarshalJSON parses the JSON-encoded data and stores the result in the set
func (set *Set[T]) UnmarshalJSON(data []byte) error {
	var slice []T
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	*set = NewSet(slice...)
	return nil
}
