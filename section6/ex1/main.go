package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogImage)
	http.ListenAndServe(":8887", nil)
}

func foo(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "foo ran")
}

func dog(response http.ResponseWriter, request *http.Request) {
	tpl.ExecuteTemplate(response, "dog.gohtml", nil)
}

func dogImage(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "toby.jpg")
}
