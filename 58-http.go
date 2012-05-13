package main

import (
    "net/http"
    "fmt"
)

type String string

func (s String) ServeHTTP (w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, s)
}
type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
func (s Struct) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}
func main() {
    // your http.Handle calls here
    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.ListenAndServe("localhost:4000", nil)
}
