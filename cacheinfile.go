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
	fpath := filepath.Join(cache_directory, file)

	if _, err := os.Stat(cache_directory); os.IsNotExist(err) {
		if err := os.MkdirAll(cache_directory, 0755); err != nil {
			log.Fatal("ERROR ", err)
		}
	}

	filepath := cache_directory + "/" + file
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fp, err := os.Create(filepath)
		if err != nil {
			fp.Close()
			log.Fatal("ERROR CREATE FILE : ", err)
		}
		fp.Close()
	}

	var fmutex sync.RWMutex
	fmutex.Lock()
	defer fmutex.Unlock()

	if err := os.WriteFile(fpath, []byte(data), 0644); err != nil {
		log.Fatal("ERROR WRITE FILE : ", err)
	}

	return nil
}

// Get reads item from cache
func Get(cache_directory string, key string, dst string) (bool, string, error) {
	key = regexp.MustCompile("[^a-zA-Z0-9_-]").ReplaceAllLiteralString(key, "")
	pattern := filepath.Join(cache_directory, fmt.Sprintf("fcache.%s", key))
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

	datafile, err := os.ReadFile(files[0]) // just pass the file name
	if err != nil {
		return false, "", err
	}

	data_out := string(datafile)
	return true, data_out, nil
}
