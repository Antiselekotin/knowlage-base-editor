package cache

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type FileCache struct {
	src *os.File
	data map[string]string
}

func (cache *FileCache)Get(key string) (str string, ok bool) {
	str, ok = cache.data[key]
	return
}

func (cache *FileCache)Set(key, value string) error {
	cache.data[key] = value
	content, err := json.Marshal(cache.data)
	if err != nil {
		return err
	}
	return cache.put(content)
}

func (cache *FileCache)put(content []byte) error {
	err := cache.src.Truncate(0)
	if err != nil {
		return err
	}
	_, err = cache.src.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = cache.src.Write(content)
	return err
}

func (cache *FileCache)pull() ([]byte, error) {
	cache.src.Seek(0,0)
	buf, err := ioutil.ReadAll(cache.src)
	if err != nil {
		return nil, err
	}
	return buf, err
}

func NewFileCache(src *os.File) (Cache, error) {
	cache := FileCache{src: src, data: make(map[string]string)}
	data, err := cache.pull()

	if err != nil {
		return nil, err
	}
	if len(data) > 0 {

		err = json.Unmarshal(data, &cache.data)
		if err != nil {
			return nil, err
		}
	}

	return &cache, nil
}
