package main

import (
	"fmt"
	"net/http"
)

type MyWebserverType bool

func (m MyWebserverType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Well, hello there!")
	fmt.Fprintln(w, "Request is: %+v", r)
}

func main() {
	var k MyWebserverType
	http.ListenAndServe("localhost:8081", k)
}
