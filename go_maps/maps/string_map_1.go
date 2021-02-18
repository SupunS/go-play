package maps

import "container/list"

// StringIntOrderedMap
//
type StringIntOrderedMap struct {
	pairs map[string]*StringIntPair
	list  *list.List
}

// NewStringIntOrderedMap creates a new StringIntOrderedMap.
func NewStringIntOrderedMap() *StringIntOrderedMap {
	return &StringIntOrderedMap{
		pairs: make(map[string]*StringIntPair),
		list:  list.New(),
	}
}

// Get returns the value associated with the given key.
// Returns nil if not found.
// The second return value indicates if the key is present in the map.
func (om *StringIntOrderedMap) Get(key string) (int, bool) {
	if pair, present := om.pairs[key]; present {
		return pair.Value, present
	}
	return 0, false
}

// GetPair returns the key-value pair associated with the given key.
// Returns nil if not found.
func (om *StringIntOrderedMap) GetPair(key string) *StringIntPair {
	return om.pairs[key]
}

// Set sets the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Set`.
func (om *StringIntOrderedMap) Set(key string, value int) (int, bool) {
	if pair, present := om.pairs[key]; present {
		oldValue := pair.Value
		pair.Value = value
		return oldValue, true
	}

	pair := &StringIntPair{
		Key:   key,
		Value: value,
	}
	pair.element = om.list.PushBack(pair)
	om.pairs[key] = pair

	return 0, false
}

// Delete removes the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Delete`.
func (om *StringIntOrderedMap) Delete(key string) (int, bool) {
	if pair, present := om.pairs[key]; present {
		om.list.Remove(pair.element)
		delete(om.pairs, key)
		return pair.Value, true
	}

	return 0, false
}

// Len returns the length of the ordered map.
func (om *StringIntOrderedMap) Len() int {
	return len(om.pairs)
}

// Oldest returns a pointer to the oldest pair.
func (om *StringIntOrderedMap) Oldest() *StringIntPair {
	return listElementToPair(om.list.Front())
}

// Newest returns a pointer to the newest pair.
func (om *StringIntOrderedMap) Newest() *StringIntPair {
	return listElementToPair(om.list.Back())
}

// StringIntPair
//
type StringIntPair struct {
	Key   string
	Value int

	element *list.Element
}

// Next returns a pointer to the next pair.
func (p *StringIntPair) Next() *StringIntPair {
	return listElementToPair(p.element.Next())
}

// Prev returns a pointer to the previous pair.
func (p *StringIntPair) Prev() *StringIntPair {
	return listElementToPair(p.element.Prev())
}

func listElementToPair(element *list.Element) *StringIntPair {
	if element == nil {
		return nil
	}
	return element.Value.(*StringIntPair)
}
