package main

import (
	"bufio"
	"fmt"
	"html"
	"log"
	"net"
	"strings"
)

var handlers = make(map[string]func(net.Conn))

func main() {
	listener, err := net.Listen("tcp", ":8887")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	addHandler("GET", "/", func(connection net.Conn) {
		textRespone(connection, "This is root")
	})
	addHandler("GET", "/what", func(connection net.Conn) {
		textRespone(connection, "What am I looking for?")
	})

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	defer connection.Close()

	scanner := bufio.NewScanner(connection)
	var method, path string

	index := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		if index == 0 {
			parts := strings.Fields(line)
			method = parts[0]
			path = parts[1]
		}

		index++
	}

	fmt.Println(method, path)

	callHandler(method, path, connection)
}

func getKey(method, path string) string {
	return method + "/" + path
}

func addHandler(method, path string, handler func(net.Conn)) {
	key := getKey(method, path)
	handlers[key] = handler
}

func callHandler(method, path string, connection net.Conn) {
	key := getKey(method, path)

	if handler, ok := handlers[key]; ok {
		handler(connection)
	} else {
		notFound(connection)
	}
}

func notFound(connection net.Conn) {
	fmt.Fprint(connection, "HTTP/1.1 404 Not found\r\n")
	fmt.Fprintf(connection, "Content-Length: 0\r\n")
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	fmt.Fprint(connection, "\r\n")
}

func textRespone(connection net.Conn, text string) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>` + html.EscapeString(text) + `</strong></body></html>`

	fmt.Fprint(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	fmt.Fprint(connection, "\r\n")
	fmt.Fprint(connection, body)
}
