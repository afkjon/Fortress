package main

import (
	"fmt"
	"log"
	"net/http"
)

var port = ":8080"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! from %s\n", r.URL.Path)
}

func main() {

	log.Println("Starting my service")

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
