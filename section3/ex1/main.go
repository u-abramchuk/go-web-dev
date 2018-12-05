package main

import (
	"bufio"
	"fmt"
	"html"
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
			log.Println(err)
			continue
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	defer connection.Close()

	scanner := bufio.NewScanner(connection)
	var location, host string

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		values := strings.Fields(line)

		switch values[0] {
		case "GET":
			fallthrough
		case "POST":
			location = values[1]
		case "Host:":
			host = values[1]
		}
	}

	hostAndPath := host + location

	fmt.Printf("URL: %v\n", hostAndPath)

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>` + html.EscapeString(hostAndPath) + `</strong></body></html>`

	fmt.Fprint(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/html\r\n")
	fmt.Fprint(connection, "\r\n")
	fmt.Fprint(connection, body)
}
