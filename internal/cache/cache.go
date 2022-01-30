package cache

import (
	"errors"
	"github.com/krls256/knowlage-base-editor/pkg/cache"
	"os"
	"path/filepath"
)

var cacheFileName = "storage/.cache/cache"
var secretCacheFileName = "storage/.cache/secret-cache"

var MainFileCache cache.Cache
var SecretFileCache cache.Cache

var secretKey = []byte("hello worldhello worldhello worl")

func init()  {
	path, err := filepath.Abs(cacheFileName); initCheck(err)
	secretpath, err := filepath.Abs(secretCacheFileName); initCheck(err)
	dir := filepath.Dir(path)

	if _, err = os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(dir ,0755); initCheck(err)
	}
	fNormal, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666); initCheck(err)
	fSecret, err := os.OpenFile(secretpath, os.O_RDWR|os.O_CREATE, 0666); initCheck(err)
	c, err := cache.NewFileCache(fNormal); initCheck(err)
	MainFileCache = c
	c, err = cache.NewSecretFileCache(fSecret, secretKey); initCheck(err)
	SecretFileCache = c
}

func initCheck(err error) {
	if err != nil {
		panic(err)
	}
}
