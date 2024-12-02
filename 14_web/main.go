package main

import (
	"fmt"
	"net/http" // import this package for work with http
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hello world</h1>") // also can write html in this
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "about")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server is starting ...")
	http.ListenAndServe(":3333", nil)
}