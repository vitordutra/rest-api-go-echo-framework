package main

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
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

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
	<html>
  <head>
     Hi
  </head>
  <body>
     <h1>
       Login
     </h1>
  </body>
</html>`)
}

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/welcome", Welcome)
	http.ListenAndServe("localhost:8081", nil)
}
