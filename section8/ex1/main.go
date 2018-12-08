package main

import (
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", count)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.ListenAndServe(":8887", nil)
}

func count(resp http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	}

	// if it cannot be parsed then treat it as 0
	count, _ := strconv.Atoi(cookie.Value)

	count++

	cookie.Value = strconv.Itoa(count)

	http.SetCookie(resp, cookie)
}
