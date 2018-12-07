package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8887", http.FileServer(http.Dir("starting-files")))
}
