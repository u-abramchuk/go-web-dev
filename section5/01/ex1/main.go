package main

import (
	"io"
	"net/http"
)

func root(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "ROOT")
}

func dog(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "BARK!!!")
}

func me(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "My name is Uladzimir")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8887", nil)
}
