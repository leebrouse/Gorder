package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// log
	log.Println("Listen order service on port 8082")
	//mux new
	mux := http.NewServeMux()

	// router default
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 	print request url
		log.Printf("%v", r.RequestURI)
		fmt.Fprintln(w, "<h1>Welcome to my Project: Gorder</h1>")
	})

	// router ping
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		//print request url
		log.Printf("%v", r.RequestURI)
		fmt.Fprintln(w, "pong")
	})
	// listen order service on port 8082
	if err := http.ListenAndServe("localhost:8082", mux); err != nil {
		log.Fatal(err)
	}
}
