package maps

// StringIntOrderedMap
//
type StringIntOrderedFastMap struct {
	pairs   map[string]int
	entries [] Entry
}

// NewStringIntOrderedMap creates a new StringIntOrderedMap.
func NewStringIntOrderedFastMap() *StringIntOrderedFastMap {
	return &StringIntOrderedFastMap{
		pairs:   make(map[string]int),
		entries: []Entry{},
	}
}

// Get returns the value associated with the given key.
// Returns nil if not found.
// The second return value indicates if the key is present in the map.
func (om *StringIntOrderedFastMap) Get(key string) (int, bool) {
	if index, present := om.pairs[key]; present {
		return om.entries[index].Value, present
	}

	return 0, false
}

//// GetPair returns the key-value pair associated with the given key.
//// Returns nil if not found.
//func (om *StringIntOrderedFastMap) GetPair(key string) *StringIntPair {
//	return om.pairs[key]
//}

// Set sets the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Set`.
func (om *StringIntOrderedFastMap) Set(key string, value int) (int, bool) {
	if index, present := om.pairs[key]; present {
		oldValue := om.entries[index].Value
		om.entries[index].Value = value
		return oldValue, true
	}

	pair := Entry{
		Key:   key,
		Value: value,
	}
	om.entries = append(om.entries, pair)
	om.pairs[key] = len(om.entries) - 1
	return 0, false
}

// Delete removes the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Delete`.
//func (om *StringIntOrderedFastMap) Delete(key string) (int, bool) {
//	if pair, present := om.pairs[key]; present {
//		om.entries.Remove(pair.element)
//		delete(om.pairs, key)
//		return pair.Value, true
//	}
//
//	return 0, false
//}

// Len returns the length of the ordered map.
func (om *StringIntOrderedFastMap) Len() int {
	return len(om.pairs)
}

//// Oldest returns a pointer to the oldest pair.
//func (om *StringIntOrderedFastMap) Oldest() *StringIntPair {
//	return listElementToPair2(om.entries.Front())
//}
//
//// Newest returns a pointer to the newest pair.
//func (om *StringIntOrderedFastMap) Newest() *StringIntPair {
//	return listElementToPair2(om.entries.Back())
//}

// StringIntPair
//
type Entry struct {
	Key   string
	Value int
}

//// Next returns a pointer to the next pair.
//func (p *StringIntPair) Next() *StringIntPair {
//	return listElementToPair2(p.element.Next())
//}
//
//// Prev returns a pointer to the previous pair.
//func (p *StringIntPair) Prev() *StringIntPair {
//	return listElementToPair2(p.element.Prev())
//}
//
//func listElementToPair2(element *entries.Element) *StringIntPair {
//	if element == nil {
//		return nil
//	}
//	return element.Value.(*StringIntPair)
//}
