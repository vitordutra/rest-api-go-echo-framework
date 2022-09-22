package main

import (
	"fmt"
	"net/http"
)

type MyWebserverType bool

func (m MyWebserverType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
	<html>
  <head>
     Hi
  </head>
  <body>
     <h1>
        My name is João Vitor Dutra
     </h1>
  </body>
</html>`)
}

func main() {
	var k MyWebserverType
	http.ListenAndServe("localhost:8081", k)
}
