/*
This is how to create a simple http server with Golang!
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address =%s\n", address)
}

// Why do I have to add *(pointer) in http.Request argument? : Because type Request is a struct, hence the need to pass
// a reference explicitly.
// https://stackoverflow.com/questions/13255907/in-go-http-handlers-why-is-the-responsewriter-a-value-but-the-request-a-pointer
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8000\n")
	// Syntax: func ListenAndServe(addr string, handler Handler) error
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
