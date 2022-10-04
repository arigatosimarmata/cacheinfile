package cacheinfile

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
)

// Set writes item to cache
func Set(cache_directory string, key string, data string, expire time.Duration) error {
	key = regexp.MustCompile("[^a-zA-Z0-9_-]"+"[\\.]").ReplaceAllLiteralString(key, "")
	file := fmt.Sprintf("fcache.%s", key)
	// file := key

	fpath := filepath.Join(cache_directory, file)

	if _, err := os.Stat(cache_directory); os.IsNotExist(err) {
		if err := os.MkdirAll(cache_directory, 0755); err != nil {
			log.Fatal(err)
		}
	}

	var fmutex sync.RWMutex
	fmutex.Lock()
	defer fmutex.Unlock()
	fp, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer fp.Close()
	if _, err = fp.WriteString(data); err != nil {
		fp.Close()
		return err
	}

	fp.Close()
	return nil
}

// Get reads item from cache
func Get(cache_directory string, key string, dst string) (bool, string, error) {
	key = regexp.MustCompile("[^a-zA-Z0-9_-]").ReplaceAllLiteralString(key, "")
	pattern := filepath.Join(cache_directory, fmt.Sprintf("fcache.%s", key))
	// pattern := filepath.Join(cache_directory, key)
	files, err := filepath.Glob(pattern)
	if err != nil {
		return false, "", err
	}
	if len(files) < 1 {
		return false, "", nil
	}

	if _, err = os.Stat(files[0]); err != nil {
		return false, "", err
	}

	fp, err := os.OpenFile(files[0], os.O_RDONLY, 0400)
	if err != nil {
		return false, "", err
	}
	defer fp.Close()

	datafile, err := os.ReadFile(files[0]) // just pass the file name
	if err != nil {
		fp.Close()
		return false, "", err
	}

	data_out := string(datafile)
	fp.Close()
	return true, data_out, nil
}
