package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/HandleFuncMode", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! %q", html.EscapeString(r.URL.Path))
	})

	http.Handle("/HandleMode", testHandler)
	http.ListenAndServe(":8080", nil)
}

type TestHandler struct{}

func (t TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

var testHandler = TestHandler{}
