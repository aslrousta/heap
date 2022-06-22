package heap_test

import (
	mrand "math/rand"
	"testing"

	"github.com/aslrousta/heap"
	"github.com/aslrousta/rand"

	. "github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	records := make(map[string]int)
	for i := 0; i < 1000; i++ {
		records[rand.MustString(8, rand.All)] = mrand.Int()
	}

	t.Run("Push", func(t *testing.T) {
		h := heap.New[string](func(a, b int) int {
			if a > b {
				return 1
			}
			if a < b {
				return -1
			}
			return 0
		})

		for k, v := range records {
			h.Push(k, v)
		}

		Equal(t, len(records), h.Len())
	})

	t.Run("Pop", func(t *testing.T) {
		h := heap.New[string](func(a, b int) int {
			if a > b {
				return 1
			}
			if a < b {
				return -1
			}
			return 0
		})

		for k, v := range records {
			h.Push(k, v)
		}

		shifted := 0
		for h.Len() > 0 {
			key, value, ok := h.Pop()
			if True(t, ok) {
				Equal(t, records[key], value)
			}
			shifted++
		}

		Equal(t, len(records), shifted)
	})

	t.Run("Remove", func(t *testing.T) {
		h := heap.New[string](func(a, b int) int {
			if a > b {
				return 1
			}
			if a < b {
				return -1
			}
			return 0
		})

		for k, v := range records {
			h.Push(k, v)
		}

		for key := range records {
			value, ok := h.Remove(key)
			if True(t, ok) {
				Equal(t, records[key], value)
			}
		}

		Equal(t, 0, h.Len())
	})
}
