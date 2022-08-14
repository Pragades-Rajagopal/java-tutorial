package main

import (
	"fmt"
	"net/http"
	"log"
)

func handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("app is running in port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
