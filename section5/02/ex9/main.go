package main

import (
	"bufio"
	"fmt"
	"html"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8887")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go serve(connection)
	}
}

func serve(connection net.Conn) {
	defer connection.Close()

	scanner := bufio.NewScanner(connection)

	i := 0
	var method, path string

	for scanner.Scan() {
		ln := scanner.Text()

		if i == 0 {
			parts := strings.Fields(ln)
			method = parts[0]
			path = parts[1]

			fmt.Println(method, path)
		}

		i++

		if ln == "" {
			break
		}
	}

	switch {
	case method == "GET" && path == "/":
		respond(connection, "ROOT")
	case method == "GET" && path == "/apply":
		respond(connection, "Show me what you've got")
	case method == "POST" && path == "/apply":
		respond(connection, "WASTED!")
	default:
		respond(connection, "Not found")
	}
}

func respond(connection net.Conn, text string) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Document</title>
	</head>
	<body>
		<h1>` + html.EscapeString(text) + `</h1>
	</body>
	</html>`

	io.WriteString(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	io.WriteString(connection, "\r\n")
	io.WriteString(connection, body)
}
