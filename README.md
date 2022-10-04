# go cacheinfile

> Fast arbitrary data caching to tmp files in Go

## Install

```bash
go get -u github.com/arigatosimarmata/cacheinfile
```

## Getting started

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/arigatosimarmata/cacheinfile"
)

func main() {
	key := "foo"
	data := []byte("bar") // can be of any time
	expire := 1 * time.Hour
	cachedirectory := time.Now().Format("20060102")//set directory date or anything

	// caching data
	err := cacheinfile.Set(cachedirectory, key, data, expire)
	if err != nil {
		log.Fatal(err)
	}

	// reading cached data
	var dst []byte
	found, err := cacheinfile.Get(cachedirectory, key, &dst)
	if err != nil {
		log.Fatal(err)
	}
	if found {
		fmt.Println(string(dst)) // "bar"
	}
}
```