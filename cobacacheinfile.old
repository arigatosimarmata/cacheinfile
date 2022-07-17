package cacheinfile

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
)

const cachedir = "./cache/"

// Set writes item to cache
func Set(cache_directory string, key string, data interface{}, expire time.Duration) error {
	key = regexp.MustCompile("[^a-zA-Z0-9_-]"+"[\\.]").ReplaceAllLiteralString(key, "")
	// file := fmt.Sprintf("fcache.%s.%v", key, strconv.FormatInt(time.Now().Add(expire).Unix(), 10))
	// file := fmt.Sprintf("fcache.%s.%v", key, time.Now().Add(expire).Format("20060102"))
	file := fmt.Sprintf("fcache.%s", key)

	fpath := filepath.Join(cachedir+cache_directory, file)

	if _, err := os.Stat(cachedir + cache_directory); os.IsNotExist(err) {
		if err := os.Mkdir(cachedir+cache_directory, 0755); err != nil {
			log.Fatal(err)
		}
	}

	clean(key)

	serialized, err := serialize(data)
	if err != nil {
		return err
	}

	var fmutex sync.RWMutex
	fmutex.Lock()
	defer fmutex.Unlock()
	fp, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer fp.Close()
	if _, err = fp.Write(serialized); err != nil {
		return err
	}

	return nil
}

// Get reads item from cache
func Get(cache_directory string, key string, dst interface{}) (bool, error) {
	key = regexp.MustCompile("[^a-zA-Z0-9_-]").ReplaceAllLiteralString(key, "")
	// pattern := filepath.Join(cachedir+cache_directory, fmt.Sprintf("fcache.%s.*", key))
	pattern := filepath.Join(cachedir+cache_directory, fmt.Sprintf("fcache.%s", key))
	files, err := filepath.Glob(pattern)
	if err != nil {
		return false, err
	}
	if len(files) < 1 {
		return false, nil
	}

	if _, err = os.Stat(files[0]); err != nil {
		return false, err
	}

	fp, err := os.OpenFile(files[0], os.O_RDONLY, 0400)
	if err != nil {
		return false, err
	}
	defer fp.Close()

	var serialized []byte
	buf := make([]byte, 1024)
	for {
		var n int
		n, err = fp.Read(buf)
		serialized = append(serialized, buf[0:n]...)
		if err != nil || err == io.EOF {
			break
		}
	}

	if err = deserialize(serialized, dst); err != nil {
		return false, err
	}

	return true, nil
}

// clean removes item from cache
func clean(key string) error {
	// pattern := filepath.Join(cachedir, fmt.Sprintf("fcache.%s.*", key))
	pattern := filepath.Join(cachedir, fmt.Sprintf("fcache.%s", key))
	files, _ := filepath.Glob(pattern)
	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			os.Remove(file)
		}
	}

	return nil
}

// serialize encodes a value using binary
func serialize(src interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(src); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// deserialize decodes a value using binary
func deserialize(src []byte, dst interface{}) error {
	buf := bytes.NewReader(src)
	err := gob.NewDecoder(buf).Decode(dst)
	return err
}
