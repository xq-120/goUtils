package pkg

import (
	"slices"
)

type SortedMap struct {
	array []string
	dict  map[string]string
}

// 为啥这里要用指针方法？

func NewSortedMap() *SortedMap {
	return &SortedMap{
		dict: make(map[string]string), // 初始化 map，避免 nil 访问问题
	}
}
func (m *SortedMap) SetObjectForKey(object string, key string) {
	if !slices.Contains(m.array, key) {
		m.array = append(m.array, key)
	}
	m.dict[key] = object
}

func (m *SortedMap) ObjectForKey(key string) (string, bool) {
	s, o := m.dict[key]
	return s, o
}

func (m *SortedMap) Contains(key string) bool {
	return slices.Contains(m.array, key)
}

func (m *SortedMap) RemoveObjectForKey(key string) {
	m.array = removeElement(m.array, key)
	delete(m.dict, key)
}

func (m *SortedMap) AllKeys() []string {
	return m.array
}

func (m *SortedMap) AllValues() []string {
	vals := make([]string, 0, len(m.array))
	for _, k := range m.array {
		if v, ok := m.dict[k]; ok {
			vals = append(vals, v)
		}
	}
	return vals
}

func (m *SortedMap) Iterator() <-chan struct {
	Key string
	Val string
} {
	ch := make(chan struct {
		Key string
		Val string
	})

	go func() {
		for _, k := range m.array {
			if v, ok := m.dict[k]; ok {
				ch <- struct {
					Key string
					Val string
				}{k, v}
			}
		}
		close(ch)
	}()

	return ch
}

func removeElement(slice []string, val string) []string {
	for i, v := range slice {
		if v == val {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
