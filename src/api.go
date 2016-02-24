package main

import (
	"net/http"
	"encoding/json"
//	"github.com/syndtr/goleveldb/leveldb"
)

func echo(r *http.Request) (interface{}, error) {
	return "Hello", nil
}

func developers_api(r *http.Request) (interface{}, error) {
	devs, err := developers("*",0)
	return devs, err
}

func api(route string, handler func(*http.Request) (res interface{},err error)) {
	http.HandleFunc(route,func(w http.ResponseWriter, r *http.Request){
		res, err := handler(r)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(res)
		}
	})
}

func routes() {
	api("/echo",echo)
	api("/developers",developers_api)
}
