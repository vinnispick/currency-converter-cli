package cache

import (
	"currency-converter-cli/internal/storage"
	"currency-converter-cli/pkg/models"
	"encoding/json"
	"time"
)

type FileCache struct {
	File string
}

func NewFileCache(file string) *FileCache {
	return &FileCache{File: file}
}

func (fn *FileCache) Get(key string) (*float64, error) {
	cache, err := readCache(fn.File)
	if err != nil {
		return nil, err
	}
	if cache == nil {
		return nil, nil
	}
	if data, exists := (*cache)[key]; exists {
		if data.ExpirationDate < time.Now().Format(time.RFC3339) {
			delete((*cache), key)
			err = writeCache(fn.File, cache)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
		return &data.Currency, nil
	}
	return nil, nil
}

func (fn *FileCache) Set(key string, value float64) error {
	cache, err := readCache(fn.File)
	if err != nil {
		return err
	}
	if cache == nil {
		cache = &map[string]models.CacheData{}
	}
	(*cache)[key] = models.CacheData{
		Currency:       value,
		ExpirationDate: time.Now().Add(10 * time.Minute).Format(time.RFC3339),
	}
	err = writeCache(fn.File, cache)
	if err != nil {
		return err
	}

	return nil
}

func readCache(cachePath string) (*map[string]models.CacheData, error) {
	file, err := storage.OpenFile(cachePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	value, err := storage.MarshalFile(file)
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, nil
	}
	var cache map[string]models.CacheData
	err = json.Unmarshal(value, &cache)
	if err != nil {
		return nil, err
	}
	return &cache, nil
}

func writeCache(cachePath string, cache *map[string]models.CacheData) error {
	if cache == nil {
		return nil
	}
	data, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	err = storage.SaveFile(cachePath, data)
	if err != nil {
		return err
	}
	return nil
}
