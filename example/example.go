package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	cacheinfile "github.com/arigatosimarmata/cacheinfile"
)

func main() {
	key := "foo"
	// data := []byte("bar") // can be of any time
	data := "1;20;20220717" // can be of any time
	expire := 1 * time.Hour
	cachedirectory := "/home/lawencon/application-log/cache-dir/" + time.Now().Format("20060102")

	// caching data
	err := cacheinfile.Set(cachedirectory, key, data, expire)
	if err != nil {
		log.Fatal(err)
	}

	// reading cached data
	var dst string
	found, dat, err := cacheinfile.Get(cachedirectory, key, dst)
	if err != nil {
		log.Fatal(err)
	}
	if found {
		fmt.Println(dat) // "bar"
		sample_split := strings.Split(dat, ";")
		data1 := sample_split[0]
		data2 := sample_split[1]
		data3 := sample_split[2]

		fmt.Printf("Printf %s - %s - %s \n", data1, data2, data3)

		if data1 == "1" {
			//update data
			data1 = "12"

			data_in := data1 + ";" + data2 + ";" + data3
			err := cacheinfile.Set(cachedirectory, key, data_in, expire)
			if err != nil {
				fmt.Printf("err %s : \n", err)
			}
		}
	}
}
