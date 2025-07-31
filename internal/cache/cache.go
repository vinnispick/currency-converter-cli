package cache

import (
	"currency-converter-cli/internal/storage"
	"currency-converter-cli/pkg/models"
	"encoding/json"
	"time"
)

func GetCacheKey(args *models.Args) string {
	return args.From + "_" + args.To
}

func GetCacheValue(cachePath, key string) *float64 {
	cache := getCache(cachePath)
	if cache == nil {
		return nil
	}
	if data, exists := (*cache)[key]; exists {
		if data.ExpirationDate < time.Now().Format(time.RFC3339) {
			delete((*cache), key)
			updatedData, err := json.Marshal(cache)
			if err != nil {
				storage.SaveFile(cachePath, updatedData)
			}
			return nil
		}
		return &data.Currency
	}
	return nil
}

func getCache(cachePath string) *map[string]models.CacheData {
	file, err := storage.OpenFile(cachePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	value, err := storage.MarshalFile(file)
	if err != nil {
		return nil
	}
	var cache map[string]models.CacheData
	err = json.Unmarshal(value, &cache)
	if err != nil {
		return nil
	}
	return &cache
}

func SetCacheValue(cachePath, key string, value float64) {
	cache := getCache(cachePath)
	if cache == nil {
		cache = &map[string]models.CacheData{}
	}
	(*cache)[key] = models.CacheData{
		Currency:       value,
		ExpirationDate: time.Now().Add(10 * time.Minute).Format(time.RFC3339),
	}
	updatedData, err := json.Marshal(cache)
	if err != nil {
		return
	}
	err = storage.SaveFile(cachePath, updatedData)
	if err != nil {
		return
	}
}
