package main

import (
	"log"
	"net/http"
	"fmt"
	"flag"
)

func main() {
	var port = flag.Int("port",8080,"Port to bind server. Default to 8080.")
	flag.Parse()

	log.Printf("Binding at 0.0.0.0:%d", *port)

	http.Handle("/", http.FileServer(http.Dir("public")))

	routes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",*port), nil))

}
