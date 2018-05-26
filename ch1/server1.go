package main

import (
	"fmt"
	"log"
	"net/http"
)

func main1() {
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %s\n", r.URL.Query().Get("cycles"))
}
