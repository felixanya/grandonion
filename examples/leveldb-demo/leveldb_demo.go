package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("/root/test/", nil)
	if err != nil {
		fmt.Println("open leveldb error", err.Error())
		return
	}
	defer db.Close()

	if err := db.Put([]byte("say"), []byte("hello"), nil); err != nil {
		fmt.Println("put data to leveldb error", err.Error())
		return
	}

	data, err := db.Get([]byte("say"), nil)
	fmt.Println("say:", string(data))
}
