package mocks

import "errors"

type MockCache struct {
	Data    map[string]float64
	FailGet bool
	FailSet bool
}

func NewMockCache() *MockCache {
	return &MockCache{
		Data: make(map[string]float64),
	}
}

func (m *MockCache) Get(key string) (*float64, error) {
	if m.FailGet {
		return nil, errors.New("failed to get from cache")
	}
	if value, exists := m.Data[key]; exists {
		return &value, nil
	}
	return nil, nil
}

func (m *MockCache) Set(key string, value float64) error {
	if m.FailSet {
		return errors.New("failed to set in cache")
	}
	m.Data[key] = value
	return nil
}
