package cache

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)
var ErrTooBigText  = errors.New("ciphertext is too big")

type SecretFileCache struct {
	fileCache FileCache
	gcm cipher.AEAD
}

func (cache *SecretFileCache)Get(key string) (str string, ok bool) {
	str, ok = cache.fileCache.data[key]
	return
}

func (cache *SecretFileCache)Set(key, value string) error {
	cache.fileCache.data[key] = value
	content, err := json.Marshal(cache.fileCache.data)
	if err != nil {
		return err
	}
	return cache.put(content)
}

func (cache *SecretFileCache)pull() ([]byte, error) {
	ciphertext, err := cache.fileCache.pull()
	if err != nil {
		return nil, err
	}
	if len(ciphertext) == 0 {
		return []byte{}, nil
	}
	nonceSize := cache.gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, ErrTooBigText
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := cache.gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, err
}

func (cache *SecretFileCache)put(content []byte) error {
	fmt.Println("Hello world")
	nonce := make([]byte, cache.gcm.NonceSize())
	ciphertext := cache.gcm.Seal(nonce, nonce, content, nil)
	return cache.fileCache.put(ciphertext)
}

func NewSecretFileCache(src *os.File, key []byte) (Cache, error) {
	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(cip)
	if err != nil {
		return nil, err
	}

	c := FileCache{src: src, data: make(map[string]string)}
	cache := SecretFileCache{fileCache: c, gcm: gcm}
	data, err := cache.pull()

	if err != nil {
		return nil, err
	}
	if len(data) > 0 {

		err = json.Unmarshal(data, &cache.fileCache.data)
		if err != nil {
			return nil, err
		}
	}
	return &cache, nil
}
