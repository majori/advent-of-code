package aoc

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func getCacheFile() *os.File {
	cachePath := os.Getenv("CACHE_PATH")
	cacheFilePath := path.Join(cachePath, "cache.json")

	file, err := os.OpenFile(cacheFilePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	return file
}

func getCacheKey(year, day int) string {
	return fmt.Sprintf("%d:%d", year, day)
}

func readCache() map[string]string {
	file := getCacheFile()

	dec := json.NewDecoder(file)

	file.Stat()

	var c map[string]string

	err := dec.Decode(&c)
	if err != nil {
		if err.Error() == "EOF" {
			return make(map[string]string)
		}

		panic(err)
	}

	return c
}

func readCacheValue(key string) string {
	cache := readCache()

	if value, ok := cache[key]; ok {
		return value
	}

	return ""
}

func writeCacheValue(key string, data string) error {
	cache := readCache()

	cache[key] = data

	file := getCacheFile()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	err := enc.Encode(cache)
	if err != nil {
		return err
	}

	return nil
}
