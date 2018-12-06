package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8887", nil)
}

func root(response http.ResponseWriter, request *http.Request) {
	err := tpl.ExecuteTemplate(response, "tpl.gohtml", "ROOT")
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(response http.ResponseWriter, request *http.Request) {
	err := tpl.ExecuteTemplate(response, "tpl.gohtml", "BARK!!!")
	if err != nil {
		log.Fatalln(err)
	}
}

func me(response http.ResponseWriter, request *http.Request) {
	err := tpl.ExecuteTemplate(response, "tpl.gohtml", "My name is Uladzimir")
	if err != nil {
		log.Fatalln(err)
	}
}
