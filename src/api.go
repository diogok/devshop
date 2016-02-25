package main

import (
	"net/http"
	"encoding/json"
	"strconv"
)

func echo(r *http.Request) (interface{}, error) {
	return "Hello", nil
}

func developers_api(r *http.Request) (interface{}, error) {
	rPage := r.URL.Query().Get("page")
	page := 1
	if len(rPage) >= 1 {
		rrPage,err := strconv.Atoi(rPage)
		if err == nil {
			page = rrPage
		}
	}
	query := r.URL.Query().Get("query")
	if len(query) <1 {
		query = "*"
	}
	devs, err := developers(query,page)
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
