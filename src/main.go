package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
)

func main() {
	sport := os.Getenv("PORT")
	var port string
	if len(sport) > 2 {
		port = sport
	} else {
		port = "8080"
	}

	log.Printf("Binding at 0.0.0.0:%s", port)

	http.Handle("/", http.FileServer(http.Dir("public")))

	routes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port), nil))

}
