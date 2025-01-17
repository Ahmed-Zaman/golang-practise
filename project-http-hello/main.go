package main

import (
	"net/http"
	"log"
	"fmt"
)

func GetName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, "Hello, ", name)
}

func main() {
	// Your solution goes here. Good luck!
	http.HandleFunc("/hello", GetName)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
