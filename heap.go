package heap

// CompareFunc compares a and b and returns 1, 0 or -1 if a is greater than b, a
// is equal to b, or a is smaller than b, respectively.
type CompareFunc[T any] func(a, b T) int

type item[K comparable, V any] struct {
	Key   K
	Value V
}

// Heap is a binary heap that stores items in an ascending order.
type Heap[K comparable, V any] struct {
	items  []item[K, V]
	lookup map[K]int
	cmp    CompareFunc[V]
}

// New instantiates a new heap.
func New[K comparable, V any](cmp CompareFunc[V]) *Heap[K, V] {
	return &Heap[K, V]{
		lookup: make(map[K]int),
		cmp:    cmp,
	}
}

// Len returns the number of items in the heap.
func (h *Heap[K, V]) Len() int {
	return len(h.items)
}

// Push adds a new key-value pair to the heap.
func (h *Heap[K, V]) Push(key K, value V) {
	if index, ok := h.lookup[key]; ok {
		if h.cmp(h.items[index].Value, value) == 0 {
			return
		}
		h.Remove(key)
	}

	h.items = append(h.items, item[K, V]{key, value})
	size := len(h.items)
	h.lookup[key] = size - 1
	h.siftUp(size - 1)
}

// Peek returns the minimum key-value pair.
func (h *Heap[K, V]) Peek() (key K, value V, ok bool) {
	if len(h.items) == 0 {
		return key, value, false
	}
	return h.items[0].Key, h.items[0].Value, true
}

// Pop removes the minimum key-value pair and returns it.
func (h *Heap[K, V]) Pop() (key K, value V, ok bool) {
	size := len(h.items)
	if size == 0 {
		return key, value, false
	}

	first := h.items[0]
	delete(h.lookup, first.Key)

	if size == 1 {
		h.items = h.items[:0]
	} else {
		h.items[0] = h.items[size-1]
		h.items = h.items[:size-1]
		h.lookup[h.items[0].Key] = 0
		h.siftDown(0)
	}

	return first.Key, first.Value, true
}

// Has returns true if a key exists in the heap.
func (h *Heap[K, V]) Has(key K) bool {
	_, ok := h.lookup[key]
	return ok
}

// Get returns a value given its key.
func (h *Heap[K, V]) Get(key K) (value V, ok bool) {
	if index, ok := h.lookup[key]; ok {
		return h.items[index].Value, true
	}
	return value, false
}

// Remove deletes a value given its key.
func (h *Heap[K, V]) Remove(key K) (value V, ok bool) {
	index, ok := h.lookup[key]
	if !ok {
		return value, false
	}

	if index == 0 {
		_, value, ok = h.Pop()
		return value, ok
	}

	delete(h.lookup, key)
	size := len(h.items)

	if index == size-1 {
		value = h.items[size-1].Value
		h.items = h.items[:size-1]
		return value, true
	}

	item := h.items[index]
	h.items[index] = h.items[size-1]
	h.items = h.items[:size-1]
	h.lookup[h.items[index].Key] = index

	parent := (index - 1) / 2
	if h.cmp(h.items[parent].Value, h.items[index].Value) > 0 {
		h.siftUp(index)
	} else {
		h.siftDown(index)
	}

	return item.Value, true
}

// Clear removes all items in the heap.
func (h *Heap[K, V]) Clear() {
	if len(h.items) > 0 {
		h.items = h.items[:0]
		h.lookup = make(map[K]int)
	}
}

func (h *Heap[K, V]) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.cmp(h.items[parent].Value, h.items[index].Value) <= 0 {
			break
		}

		h.items[index], h.items[parent] = h.items[parent], h.items[index]
		h.lookup[h.items[index].Key] = index
		h.lookup[h.items[parent].Key] = parent
		index = parent
	}
}

func (h *Heap[K, V]) siftDown(index int) {
	size := len(h.items)
	for index < size {
		left, right := 2*index+1, 2*index+2
		parent := index

		if left < size && h.cmp(h.items[left].Value, h.items[parent].Value) < 0 {
			parent = left
		}
		if right < size && h.cmp(h.items[right].Value, h.items[parent].Value) < 0 {
			parent = right
		}
		if parent == index {
			break
		}

		h.items[index], h.items[parent] = h.items[parent], h.items[index]
		h.lookup[h.items[index].Key] = index
		h.lookup[h.items[parent].Key] = parent
		index = parent
	}
}
