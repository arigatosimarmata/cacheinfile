package cacheinfile

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

const cachedir = "./cache/"

func Set(key string, data interface{}, expire time.Duration) error {
	key = regexp.MustCompile("[^a-zA-Z0-9_-]").ReplaceAllLiteralString(key, "")
	file := fmt.Sprintf("fcache.%s.%v", key, strconv.FormatInt(time.Now().Add(expire).Unix(), 10))
	fp := filepath.Join(cachedir, file)

	fmt.Println(fp)
	cleancache(key)

	return nil
}

func cleancache(key string) error {
	pattern := filepath.Join(cachedir, fmt.Sprintf("fcache.%s.*", key))
	files, _ := filepath.Glob(pattern)
	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			os.Remove(file)
		}
	}

	return nil
}

func serialize(src interface{}) ([]byte, error) {
	buff := new(bytes.Buffer)
	if err := gob.NewEncoder(buff).Encode(src); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func deserialize(src []byte, dst interface{}) error {
	buff := bytes.NewReader(src)
	err := gob.NewDecoder(buff).Decode(dst)
	return err
}
