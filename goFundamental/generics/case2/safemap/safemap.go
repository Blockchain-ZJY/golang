package safemap

import (
	"fmt"
	"sync"
)

type Safemap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func NewSafeMap[K comparable, V any]() *Safemap[K, V] {
	return &Safemap[K, V]{
		data: make(map[K]V),
	}
}

func (m *Safemap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[key] = value
}

func (m *Safemap[K, V]) Get(key K) (V, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if _, ok := m.data[key]; !ok {
		return *new(V), fmt.Errorf("key not found")
	}

	return m.data[key], nil
}
