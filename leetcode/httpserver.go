package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct{}
func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

type String string
func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "String!")
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}
func (s *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	// var h Hello
	// err := http.ListenAndServe("localhost:4000", h)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	http.Handle("/", Hello{})
	http.HandleFunc("/hello", hello)
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
