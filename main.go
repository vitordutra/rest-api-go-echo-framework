package main

import (
	"fmt"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	http.ListenAndServe("localhost:8081", http.HandlerFunc(ServeHTTP))
}
