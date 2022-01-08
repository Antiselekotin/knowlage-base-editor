package cache

import (
	"os"
	"path/filepath"
	"testing"
)

var path, _ = filepath.Abs("./.cache/cache")
var dir = filepath.Dir(path)
func createTestFile() *os.File {
	err := os.Mkdir(dir ,0755)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		os.Remove(dir)
		panic(err)
	}
	return file
}

func removeTestFile(file *os.File)  {
	file.Close()
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
	err = os.Remove(dir)
	if err != nil {
		panic(err)
	}

}

func TestMultiCacheToOneFile(t *testing.T) {
	file := createTestFile()
	fCache, err := NewFileCache(file)
	if err != nil {
		panic(err)
	}
	err = fCache.Set("key", "value")
	if err != nil {
		panic(err)
	}
	sCache, err := NewFileCache(file)
	val, ok := sCache.Get("key")
	if !ok || val != "value" {
		t.Errorf("value is %v 'value' expected, ok is %v", val, ok)
	}
	removeTestFile(file)
}