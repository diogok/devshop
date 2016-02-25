package main

import (
	"net/http"
	"encoding/json"
	"strconv"
	"errors"
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

func cart_api(r *http.Request) (interface{}, error) {
	if r.Method == "GET" {
		return ListCart(), nil
	} else if r.Method == "POST" {
		login := r.FormValue("login")
		cost,err0 := strconv.Atoi(r.FormValue("cost"))
		id,err1 := strconv.Atoi(r.FormValue("id"))
		if len(login) < 1 {
			return nil, errors.New("Bad request: login")
		}
		if err0 != nil {
			return nil, err0
		}
		if err1 != nil {
			return nil, err1
		}
		dev := Developer{Login: login, ID: id, Cost: cost}
		AddToCart(dev)
		return ListCart(), nil
	} else if(r.Method == "DELETE") {
		login := r.URL.Query().Get("login")
		if len(login) < 1 {
			return nil, errors.New("Bad request: login")
		}
		DelFromCart(login)
		return ListCart(), nil
	} else {
		return nil, errors.New("Bad request")
	}
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
	api("/cart",cart_api)
}
