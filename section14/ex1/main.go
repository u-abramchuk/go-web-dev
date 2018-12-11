package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type code struct {
	Code    int
	Descrip string
}

func main() {
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		http.ServeFile(resp, req, "index.html")
	})
	http.ListenAndServe(":8887", nil)
}

func encode(resp http.ResponseWriter, req *http.Request) {
	codes := []code{
		code{200, "OK"},
		code{301, "Moved Permanently"},
		code{302, "Found"},
		code{303, "See Other"},
		code{307, "Temporary Redirect"},
		code{400, "Bad Request"},
		code{401, "Unauthorized"},
		code{402, "Payment Required"},
		code{403, "Forbidden"},
		code{404, "Not Found"},
		code{405, "Method Not Allowed"},
		code{418, "Teapot"},
		code{500, "Internal Server Error"},
	}

	serialized, err := json.Marshal(codes)
	if err != nil {
		log.Fatalln(err)
	}

	resp.Header().Add("Content-Type", "application/json")
	resp.Write(serialized)
}

func decode(resp http.ResponseWriter, req *http.Request) {
	var data []code

	rcvd := `[{"Code":200,"Descrip":"OK"},{"Code":301,"Descrip":"Moved Permanently"},{"Code":302,"Descrip":"Found"},{"Code":303,"Descrip":"See Other"},{"Code":307,"Descrip":"Temporary Redirect"},{"Code":400,"Descrip":"Bad Request"},{"Code":401,"Descrip":"Unauthorized"},{"Code":402,"Descrip":"Payment Required"},{"Code":403,"Descrip":"Forbidden"},{"Code":404,"Descrip":"Not Found"},{"Code":405,"Descrip":"Method Not Allowed"},{"Code":418,"Descrip":"Teapot"},{"Code":500,"Descrip":"Internal Server Error"}]`

	json.Unmarshal([]byte(rcvd), &data)

	for _, v := range data {
		fmt.Println(v)
	}

	resp.WriteHeader(http.StatusOK)
}
