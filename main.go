package main

import (
	"io"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world.")
}

func main() {
	http.HandleFunc("/", sayHello)
	panic(http.ListenAndServe(":8090", nil))
}
