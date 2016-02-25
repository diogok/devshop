package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"os"
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

var db *leveldb.DB

func conn() *leveldb.DB {
	if db != nil {
		return db
	}
	path := "../data"
	rpath := os.Getenv("DATA_DIR")
	if len(rpath) > 1 {
		path = rpath
	}
	idb, err := leveldb.OpenFile(path, nil)
	if err != nil {
		panic(err)
	}
	db = idb
	return idb
}

func AddToCart(dev Developer) {
	db := conn()
	var blob bytes.Buffer
	enc := gob.NewEncoder(&blob)
	_ = enc.Encode(dev)
	key := fmt.Sprintf("cart-%s",dev.Login)
	bkey := bytes.NewBufferString(key)
	err := db.Put(bkey.Bytes(),blob.Bytes(),nil)
	if err != nil {
		panic(err)
	}
}

func DelFromCart(login string) {
	db := conn()
	key := fmt.Sprintf("cart-%s",login)
	bkey := bytes.NewBufferString(key)
	_ = db.Delete(bkey.Bytes(),nil)
}

func ListCart() Developers {
	db := conn()
	devs := Developers{}
	iter := db.NewIterator(util.BytesPrefix([]byte("cart-")), nil)
	for iter.Next() {
		var dev Developer
		val := iter.Value()
		blob := bytes.NewReader(val)
		dec := gob.NewDecoder(blob)
		_ = dec.Decode(&dev)
		devs = append(devs,dev)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		panic(err)
	}
	return devs
}

func ClearCart() {
	db := conn()
	var keys []string
	iter := db.NewIterator(util.BytesPrefix([]byte("cart-")), nil)
	for iter.Next() {
		key := string(iter.Key())
		keys = append(keys,key)
	}
	iter.Release()
	for _, key := range keys {
		bkey := bytes.NewBufferString(key)
		db.Delete(bkey.Bytes(),nil)
	}
}

func Purchase() {
	db := conn()
	cart := ListCart()
	ClearCart()

	timestamp := time.Now().Unix()
	key := fmt.Sprintf("purchase-%d",timestamp)
	bkey := bytes.NewBufferString(key)

	var blob bytes.Buffer
	enc := gob.NewEncoder(&blob)
	_ = enc.Encode(cart)

	err := db.Put(bkey.Bytes(),blob.Bytes(),nil)
	if err != nil {
		panic(err)
	}
}

func ListPurchases() []Developers {
	db := conn()
	purchases := []Developers{}
	iter := db.NewIterator(util.BytesPrefix([]byte("purchase-")), nil)
	for iter.Next() {
		var devs Developers
		val := iter.Value()
		blob := bytes.NewReader(val)
		dec := gob.NewDecoder(blob)
		_ = dec.Decode(&devs)
		purchases = append(purchases,devs)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		panic(err)
	}
	return purchases
}

func ClearPurchases() {
	db := conn()
	var keys []string
	iter := db.NewIterator(util.BytesPrefix([]byte("purchase-")), nil)
	for iter.Next() {
		key := string(iter.Key())
		keys = append(keys,key)
	}
	iter.Release()
	for _, key := range keys {
		bkey := bytes.NewBufferString(key)
		db.Delete(bkey.Bytes(),nil)
	}
}
