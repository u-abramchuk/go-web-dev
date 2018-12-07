package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8887", nil)
}

func index(resp http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(resp, nil)
}
